package serializer

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=isKind -linecomment

type _error string

func (e _error) Error() string { return string(e) }

const (
	ErrMissingPackageName _error = "missing package name"
	ErrInvalidData        _error = "invalid data layout"
)

var (
	localpkgPath = reflect.TypeOf(ErrInvalidData).PkgPath()
	localpkgName = path.Base(localpkgPath)
)

// Interface defines the methods added to types processed by Gen.
type Interface interface {
	MarshalBinaryTo(io.Writer) error
	UnmarshalBinaryFrom(io.Reader) error
}

var (
	_Interface = reflect.TypeOf([]Interface(nil)).Elem()
	_Time      = reflect.TypeOf(time.Time{})
)

// Config defines the elements used to generate the code.
//
// If the Receiver is not set, it uses the lowered first letter of the type name.
type Config struct {
	PkgName  string
	Receiver string
}

// Gen generates the MarshalBinaryTo and UnmarshalBinaryFrom methods for
// the provided types out to the given Writer.
func Gen(out io.Writer, config Config, data ...interface{}) error {
	// imports contains all the packages that are or may be required, which
	// are referenced via _ to avoid import errors.
	const imports = `package %pkgname%

import (
	%imports%

	"%pkgpath%"
)
`
	if config.PkgName == "" {
		return ErrMissingPackageName
	}

	tdata := map[string]string{
		"pkgname": config.PkgName,
		"pkgpath": localpkgPath,
		"pkg":     localpkgName,
	}
	imps := map[string]bool{
		"io": true,
	}

	// Generate the methods code into buf first so that proper imports can be determined.
	codeBuf := new(bytes.Buffer)
	for i := 0; i < len(data); i++ {
		item := data[i]
		receiver := config.Receiver
		if receiver == "" {
			tname := reflect.TypeOf(item).Name()
			for i, c := range tname {
				if c != '_' && !unicode.IsDigit(c) {
					receiver = strings.ToLower(tname[i : i+1])
					break
				}
			}
		}
		records, deps, err := walkDataType(nil, receiver, item)
		if err != nil {
			return err
		}
		for _, dep := range deps {
			if !hasType(reflect.TypeOf(dep).Name(), data) {
				data = append(data, dep)
			}
		}

		processRecords(records, config.PkgName+".", imps)
		tdata["type"] = reflect.TypeOf(item).Name()
		tdata["rcv"] = receiver

		for _, c := range []genConfig{marshalBinaryTo, unmarshalBinaryFrom} {
			err = c.genHeader(codeBuf, records, tdata)
			if err != nil {
				return err
			}
			err = c.genBody(0, codeBuf, records, tdata, c.Conv)
			if err != nil {
				return err
			}
			err = c.genTail(codeBuf, tdata)
			if err != nil {
				return err
			}
		}
	}

	// Figure out the imports.
	importsInUse := make([]string, 0, len(imps))
	for imp := range imps {
		importsInUse = append(importsInUse, strconv.Quote(imp))
	}
	sort.Strings(importsInUse)
	tdata["imports"] = strings.Join(importsInUse, "\n\t")

	if err := templateExec(out, imports, tdata); err != nil {
		return err
	}

	_, err := out.Write(codeBuf.Bytes())
	return err
}

// hasType returns whether or not data contains a struct type with the given name.
func hasType(name string, data []interface{}) bool {
	for _, item := range data {
		if name == reflect.TypeOf(item).Name() {
			return true
		}
	}
	return false
}

// processRecords removes the package name if it is equal to the one where the methods are created.
// It also records imports into the imps map.
func processRecords(records []genRecord, name string, imps map[string]bool) {
	for i, rec := range records {
		switch rec.FuncKind {
		case "time":
			imps["time"] = true
		}
		records[i].Kind = strings.ReplaceAll(rec.Kind, name, "")
		records[i].Name = strings.ReplaceAll(rec.Name, name, "")
		processRecords(rec.Include, name, imps)
		processRecords(rec.Key, name, imps)
	}
}

// genRecord keeps track of the struct elements being serialized.
// Slices are encoded as: <slice length><item0>...
// Structs are encoded in their fields order.
type genRecord struct {
	Is       isKind
	RKind    reflect.Kind
	Ident    string      // target identifier
	Kind     string      // target kind (only fixed size kinds)
	Name     string      // target type name
	FuncKind string      // if set, used for Write_ and Read_ instead of Kind
	Include  []genRecord // slice, array or map value
	Key      []genRecord // map key
}

type convFunc func(genRecord) (id, value, conv string)

type isKind uint8

const (
	_            isKind = iota // None
	isSlice                    // Slice
	isByteSlice                // Bytes
	isArray                    // Array
	isByteArray                // ByteArray
	isStruct                   // Struct
	isAnonStruct               // AnonStruct
	isMap                      // Map
	isMapStruct                // MapStruct
	isPointer                  // Pointer
)

func walkDataType(p []string, ident string, data interface{}) ([]genRecord, []interface{}, error) {
	var records []genRecord
	var deps []interface{}
	typ := reflect.TypeOf(data)
	//switch kind := typ.Kind(); kind {
	kind := typ.Kind()
	switch kind {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		records = []genRecord{
			{RKind: kind, Ident: ident, Kind: kind.String(), Name: typ.String(), FuncKind: kind.String()},
		}
	case reflect.Slice:
		etyp := typ.Elem()
		value := reflect.New(etyp).Elem()
		subrecords, subdeps, err := walkDataType(append(p, etyp.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		if typ.String() == "[]uint8" {
			records = []genRecord{
				{RKind: kind, Is: isByteSlice, Ident: ident, Kind: "[]byte", Name: "[]byte"}}
		} else {
			records = []genRecord{
				{RKind: kind, Is: isSlice, Ident: ident, Kind: typ.String(), Include: subrecords},
			}
		}
		deps = append(deps, subdeps...)
	case reflect.Array:
		etyp := typ.Elem()
		value := reflect.New(etyp).Elem()
		subrecords, subdeps, err := walkDataType(append(p, etyp.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		if k := subrecords[0].Kind; len(subrecords) == 1 && k == "uint8" {
			// Special case: [...]byte
			records = []genRecord{
				{RKind: kind, Is: isByteArray, Ident: ident, Kind: typ.String(), Name: typ.String()},
			}
		} else {
			records = []genRecord{
				{RKind: kind, Is: isArray, Ident: ident, Kind: typ.String(), Include: subrecords},
			}
		}
		deps = append(deps, subdeps...)
	case reflect.Map:
		value := reflect.New(typ.Elem()).Elem()
		subrecords, subdeps, err := walkDataType(append(p, typ.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		key := reflect.New(typ.Key()).Elem()
		keyrecords, keydeps, err := walkDataType(append(p, typ.String()), ident, key.Interface())
		if err != nil {
			return nil, nil, err
		}
		if subrecords[0].Is == isStruct {
			subrecords[0].Is = isMapStruct
		}
		records = []genRecord{
			{RKind: kind, Is: isMap, Ident: ident, Kind: typ.String(), Include: subrecords, Key: keyrecords}}
		deps = append(deps, subdeps...)
		deps = append(deps, keydeps...)
	case reflect.Struct:
		// Special structs.
		if _Time.ConvertibleTo(typ) {
			records = []genRecord{
				{RKind: kind, Ident: ident, Kind: "time.Time", Name: typ.String(), FuncKind: "time"}}
			break
		}

		isTopStruct := len(p) == 0
		value := reflect.ValueOf(data)
		if name := typ.Name(); !isTopStruct && name != "" {
			// Named struct type: add to the the list of dependents to get the marshal methods
			// if it does not already implement the methods.
			records = []genRecord{
				{RKind: kind, Is: isStruct, Ident: ident, Kind: typ.String()}}
			if !reflect.New(typ).Type().Implements(_Interface) {
				deps = append(deps, value.Interface())
			}
			break
		}
		n := typ.NumField()
		subrecords := make([]genRecord, 0, n)
		for i := 0; i < n; i++ {
			sf := typ.Field(i)
			if sf.Name == "" {
				continue
			}
			id := sf.Name
			if isTopStruct {
				id = fmt.Sprintf("%s.%s", ident, sf.Name)
			}
			field := value.Field(i)
			if !field.CanInterface() {
				continue
			}
			data := field.Interface()
			s, d, err := walkDataType(append(p, sf.Name), id, data)
			if err != nil {
				return records, deps, err
			}
			subrecords = append(subrecords, s...)
			deps = append(deps, d...)
		}
		if isTopStruct {
			// Top level struct, use its fields' content.
			records = subrecords
		} else {
			records = []genRecord{
				{RKind: kind, Is: isAnonStruct, Ident: ident, Kind: typ.String(), Include: subrecords}}
		}
	case reflect.Ptr:
		etyp := typ.Elem()
		value := reflect.New(etyp).Elem()
		subrecords, subdeps, err := walkDataType(append(p, etyp.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		records = []genRecord{
			{RKind: kind, Is: isPointer, Ident: ident, Kind: typ.String(), Include: subrecords},
		}
		deps = append(deps, subdeps...)
	default:
		path := strings.Join(p, ".")
		err := fmt.Errorf("binary.Write: %s: unsupported type %T", path, data)
		return nil, nil, err
	}
	return records, deps, nil
}

type genConfig struct {
	WithDecl   bool
	Head       string
	Tail       string
	Call       string
	Slice      string
	ByteSlice  string
	AnonStruct string
	Array      string
	ByteArray  string
	Struct     string
	Map        string
	MapStruct  string
	Pointer    string

	Conv convFunc
}

var (
	marshalBinaryTo = genConfig{
		WithDecl: false,
		Head: `
const _%type%Layout = "%layout%"

func (%rcv% *%type%) MarshalBinaryTo(w io.Writer) (err error) {
	_b := %pkg%.Buffers.Get()
	defer %pkg%.Buffers.Put(_b)
	err = %pkg%.Write_layout(w, _b, _%type%Layout); if err != nil { return }
`,
		Call: `
%tab%err = %pkg%.Write_%funckind%(w, _b, %conv%); if err != nil { return }
`,
		Slice: `
%tab%{
%tab%	_s := %value%
%tab%	_n = len(_s)
%tab%	err = %pkg%.Write_int(w, _b, _n); if err != nil { return }
%tab%	for _k := 0; _k < _n; _k++ {%include%	%tab%}
%tab%}`,
		ByteSlice: `
%tab%err = %pkg%.Write_bytes(w, _b, %conv%); if err != nil { return }
`,
		Array: `
%tab%{
%tab%	_s := &%value%
%tab%	for _k := 0; _k < len(_s); _k++ {%include%	%tab%}
%tab%}`,
		ByteArray: `
%tab%err = %pkg%.Write_bytea(w, (%conv%)[:]); if err != nil { return }
`,
		Struct: `
%tab%err = %idlevel%.MarshalBinaryTo(w); if err != nil { return }
`,
		AnonStruct: `
%tab%{
%tab%	_s := &%value%
%tab%	%include%
%tab%}
`,
		MapStruct: `
%tab%{
%tab%	_struct := %idlevel%
%tab%	err = _struct.MarshalBinaryTo(w); if err != nil { return }
%tab%}
`,
		Map: `
%tab%{
%tab%	_s := %value%
%tab%	err = %pkg%.Write_int(w, _b, len(_s)); if err != nil { return }
%tab%	for _k := range _s {%includekey%%include%	%tab%}
%tab%}`,
		Pointer: `
%tab%err = %pkg%.Write_bool(w, _b, %idlevel% == nil); if err != nil { return }
%tab%if %idlevel% != nil {%include%	%tab%}
`,
		Tail: `
	return
}
`,
		Conv: func(rec genRecord) (a, b, c string) {
			id := rec.Ident
			val := rec.Ident
			kind := rec.Kind
			if rec.Is == isPointer {
				val = "*" + val
			}
			if rec.Name == kind {
				return id, val, val
			}
			return id, val, fmt.Sprintf("%s(%s)", kind, val)
		},
	}
	unmarshalBinaryFrom = genConfig{
		WithDecl: true,
		Head: `
func (%rcv% *%type%) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_b := %pkg%.Buffers.Get()
	defer %pkg%.Buffers.Put(_b)
	err = %pkg%.Read_layout(r, _b, _%type%Layout); if err != nil { return }
`,
		Call: `
%tab%_%funckind%, err = %pkg%.Read_%funckind%(r, _b); if err != nil { return }
%tab%%value% = %conv%
`,
		Slice: `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if _c := cap(%value%); _n > _c || _c - _n > _c/8 { %value% = make(%kind%, _n) } else { %value% = (%value%)[:_n] }
%tab%if _n > 0 {
%tab%	_s := %value%
%tab%	for _k := 0; _k < _n; _k++ {%include%	%tab%}
%tab%}
`,
		ByteSlice: `
%tab%%value%, err = %pkg%.Read_bytes(r, _b, nil); if err != nil { return }
`,
		Array: `
%tab%{
%tab%	_s := &%value%
%tab%	for _k := 0; _k < len(_s); _k++ {%include%	%tab%}
%tab%}
`,
		ByteArray: `
%tab%err = %pkg%.Read_bytea(r, (%value%)[:]); if err != nil { return }
`,
		Struct: `
%tab%err = %idlevel%.UnmarshalBinaryFrom(r); if err != nil { return }
`,
		AnonStruct: `
%tab%{
%tab%	_s := &%value%
%tab%	%include%
%tab%}
`,
		MapStruct: `
%tab%{
%tab%	_struct := %idlevel%
%tab%	err = _struct.UnmarshalBinaryFrom(r); if err != nil { return }
%tab%	%idlevel% = _struct
%tab%}
`,
		Map: `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if _n == 0 {  %idlevel% = nil } else {
%tab%	%idlevel% = make(%kind%, _n)
%tab%	_s := %idlevel%
%tab%	var _k %kindkey%
%tab%	for _j := 0; _j < _n; _j++ {%includekey%%include%	%tab%}
%tab%}
`,
		Pointer: `
%tab%_bool, err = %pkg%.Read_bool(r, _b); if err != nil { return }
%tab%if _bool { %idlevel% = nil } else {%alloc%%include%%tab%}
`,
		Tail: `
	return
}
`,
		Conv: func(rec genRecord) (a, b, c string) {
			id := rec.Ident
			val := rec.Ident
			kind := rec.Kind
			conv := "_" + rec.FuncKind
			if rec.Is == isPointer {
				val = "*" + val
			}
			if rec.Name == kind {
				return id, val, conv
			}
			return id, val, fmt.Sprintf("%s(%s)", rec.Name, conv)
		},
	}
)

// genHeader writes the method header with pre declared variables if required.
func (c genConfig) genHeader(w io.Writer, records []genRecord, data map[string]string) error {
	const decl = `	var %var% %kind%
`
	vars := make(map[string]string)
	genHeaderNext(records, c.WithDecl, vars)

	layouts := genCheck(records, nil, nil)
	data["layout"] = strings.Join(layouts, "")
	if err := templateExec(w, c.Head, data); err != nil {
		return err
	}
	delete(data, "layout")
	if len(vars) > 0 {
		// Separate variable declarations from the rest of the function body.
		if _, err := w.Write([]byte{'\n'}); err != nil {
			return err
		}
	}

	type kv struct{ k, v string }
	sortedVars := make([]kv, 0, len(vars))
	for v, kind := range vars {
		sortedVars = append(sortedVars, kv{v, kind})
	}
	sort.Slice(sortedVars, func(i, j int) bool { return sortedVars[i].k < sortedVars[j].k })
	for _, kv := range sortedVars {
		data["var"] = kv.k
		data["kind"] = kv.v
		if err := templateExec(w, decl, data); err != nil {
			return err
		}
	}
	delete(data, "var")
	delete(data, "kind")

	return nil
}

func genCheck(records []genRecord, cache map[reflect.Kind]string, ls []string) []string {
	if cache == nil {
		// Cache layout strings.
		cache = map[reflect.Kind]string{
			reflect.Slice: string('A' + reflect.Slice),
			reflect.Uint8: string('A' + reflect.Uint8),
		}
	}
	for _, rec := range records {
		if _, ok := cache[rec.RKind]; !ok {
			cache[rec.RKind] = string('A' + rec.RKind)
		}
		ls = append(ls, cache[rec.RKind])
		switch rec.Is {
		case isArray, isByteArray:
			size := rec.Kind[1:strings.Index(rec.Kind, "]")]
			ls = append(ls, size)
		}
		ls = genCheck(rec.Key, cache, ls)
		ls = genCheck(rec.Include, cache, ls)
	}
	return ls
}

func genHeaderNext(records []genRecord, withDecl bool, vars map[string]string) {
	reg := func(k, v string) {
		if withDecl {
			vars[k] = v
		}
	}
	// Prepare all variables.
	for _, rec := range records {
		switch rec.Is {
		case isSlice:
			vars["_n"] = "int"
		case isMap:
			reg("_n", "int")
		case isPointer:
			reg("_bool", "bool")
		}

		genHeaderNext(rec.Key, withDecl, vars)
		genHeaderNext(rec.Include, withDecl, vars)

		if f := rec.FuncKind; f != "" {
			reg("_"+f, rec.Kind)
		}
	}
}

func (c *genConfig) genBody(level int, w io.Writer, records []genRecord, data map[string]string, conv convFunc) error {
	defer func(t string) { data["tab"] = t }(data["tab"])
	data["tab"] += "\t"
	tab := data["tab"]
	inctab := tab + "\t"
	var include strings.Builder

	sconv := sliceConv(conv)
	kconv := keyConv(conv)
	vconv := valueConv(conv)
	pconv := pointerConv(conv)
	aconv := anonConv(conv)
	doinc := func(incname string, records []genRecord, conv convFunc) error {
		include.Reset()
		data["tab"] = inctab
		if err := c.genBody(level+1, &include, records, data, conv); err != nil {
			return err
		}
		data["tab"] = tab
		data[incname] = include.String()
		return nil
	}
	doalloc := func(rec genRecord) error {
		var alloc string
		switch rec.Is {
		case isByteSlice:
			// No allocation required as Read_bytes does it.
			return nil
		case isMap:
			alloc = `
%tab%%idlevel% = make(%kind%)`
		default:
			alloc = `
%tab%%idlevel% = new(%kind%)`
		}
		var include strings.Builder
		data["tab"] = inctab
		if err := templateExec(&include, alloc, data); err != nil {
			return err
		}
		data["tab"] = tab
		data["alloc"] = include.String()
		return nil
	}

	for _, rec := range records {
		var s string
		switch rec.Is {
		case isSlice:
			if err := doinc("include", rec.Include, sconv); err != nil {
				return err
			}
			s = c.Slice
		case isByteSlice:
			s = c.ByteSlice
		case isArray:
			if err := doinc("include", rec.Include, sconv); err != nil {
				return err
			}
			s = c.Array
		case isByteArray:
			s = c.ByteArray
		case isMapStruct:
			s = c.MapStruct
		case isStruct:
			s = c.Struct
		case isAnonStruct:
			if err := doinc("include", rec.Include, aconv); err != nil {
				return err
			}
			s = c.AnonStruct
		case isMap:
			if err := doinc("includekey", rec.Key, kconv); err != nil {
				return err
			}
			if err := doinc("include", rec.Include, vconv); err != nil {
				return err
			}
			data["kindkey"] = rec.Key[0].Kind
			s = c.Map
		case isPointer:
			if err := doinc("include", rec.Include, pconv); err != nil {
				return err
			}
			if err := doalloc(rec.Include[0]); err != nil {
				return err
			}
			s = c.Pointer
		default:
			s = c.Call
		}
		data["id"] = rec.Ident
		data["kind"] = rec.Kind
		data["funckind"] = rec.FuncKind
		data["idlevel"], data["value"], data["conv"] = conv(rec)
		if err := templateExec(w, s, data); err != nil {
			return err
		}
		data["kindkey"] = ""
		data["alloc"] = ""
	}
	return nil
}

func (c *genConfig) genTail(w io.Writer, data map[string]string) error {
	return templateExec(w, c.Tail, data)
}

func sliceConv(conv convFunc) convFunc {
	return func(rec genRecord) (a, b, c string) {
		const _s = "_s[_k]"
		_, v, c := conv(rec)
		return _s, _s, strings.ReplaceAll(c, v, _s)
	}
}

func keyConv(conv convFunc) convFunc {
	return func(rec genRecord) (a, b, c string) {
		const _s = "_k"
		_, v, c := conv(rec)
		return _s, _s, strings.ReplaceAll(c, v, _s)
	}
}

func valueConv(conv convFunc) convFunc {
	return func(rec genRecord) (a, b, c string) {
		const _s = "_s[_k]"
		_, v, c := conv(rec)
		return _s, _s, strings.ReplaceAll(c, v, _s)
	}
}

func pointerConv(conv convFunc) convFunc {
	return func(rec genRecord) (a, b, c string) {
		i, v, c := conv(rec)
		_s := "*" + v
		return i, _s, strings.ReplaceAll(c, v, _s)
	}
}

func anonConv(conv convFunc) convFunc {
	return func(rec genRecord) (a, b, c string) {
		_, v, c := conv(rec)
		_s := "_s." + rec.Ident
		return _s, _s, strings.ReplaceAll(c, v, _s)
	}
}
