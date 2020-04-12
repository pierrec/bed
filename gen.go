package serializer

import (
	"fmt"
	"io"
	"path"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

type _error string

func (e _error) Error() string { return string(e) }

const ErrInvalidData _error = "invalid data layout"

var (
	localpkgPath = reflect.TypeOf(ErrInvalidData).PkgPath()
	localpkgName = path.Base(localpkgPath)
)

// Interface is the interface added to types processed by Gen.
type Interface interface {
	MarshalBinaryTo(io.Writer) error
	UnmarshalBinaryFrom(io.Reader) error
}

var _Interface = reflect.TypeOf([]Interface(nil)).Elem()

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
	const imports = `package %pkgname%

import (
	"io"

	"%pkgpath%"
)

`
	tdata := map[string]string{
		"pkgname": config.PkgName,
		"pkgpath": localpkgPath,
	}
	if err := templateExec(out, imports, tdata); err != nil {
		return err
	}

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
		stripLocalPkgName(records, config.PkgName+".")
		err = genMarshalBinTo(out, records, receiver, item)
		if err != nil {
			return err
		}
		err = genUnmarshalBinFrom(out, records, receiver, item)
		if err != nil {
			return err
		}
	}
	return nil
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

// stripLocalPkgName removes the package name if it is equal to the one where the methods are created.
func stripLocalPkgName(records []genRecord, name string) {
	for i, rec := range records {
		records[i].Kind = strings.ReplaceAll(rec.Kind, name, "")
		stripLocalPkgName(rec.Include, name)
		stripLocalPkgName(rec.Key, name)
	}
}

// genRecord keeps track of the struct elements being serialized.
// Slices are encoded as: <slice length><item0>...
// Structs are encoded in their fields order.
type genRecord struct {
	Is      uint8
	RKind   reflect.Kind
	Ident   string      // target identifier
	Kind    string      // target kind (only fixed size kinds)
	Name    string      // target type name
	Include []genRecord // slice, array or map value
	Key     []genRecord // map key
}

type convFunc func(genRecord) (id, value, conv string)

const (
	_ uint8 = iota
	isSlice
	isArray
	isByteArray
	isStruct
	isMap
	isMapStruct
	isPointer
)

func walkDataType(p []string, ident string, data interface{}) ([]genRecord, []interface{}, error) {
	var records []genRecord
	var deps []interface{}
	typ := reflect.TypeOf(data)
	switch kind := typ.Kind(); kind {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		records = []genRecord{
			{RKind: kind, Ident: ident, Kind: kind.String(), Name: typ.String()},
		}
	case reflect.Slice:
		etyp := typ.Elem()
		value := reflect.New(etyp).Elem()
		subrecords, subdeps, err := walkDataType(append(p, etyp.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		records = []genRecord{
			{RKind: kind, Is: isSlice, Ident: ident, Kind: typ.String(), Include: subrecords},
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
		value := reflect.ValueOf(data)
		if len(p) > 0 && typ.Name() != "" {
			// Named struct type: add to the the list of dependents to get the marshal methods
			// if it does not already implement the methods.
			records = append(records, genRecord{RKind: kind, Is: isStruct, Ident: ident, Kind: typ.String()})
			if !reflect.New(typ).Type().Implements(_Interface) {
				deps = append(deps, value.Interface())
			}
			break
		}
		n := typ.NumField()
		for i := 0; i < n; i++ {
			sf := typ.Field(i)
			if sf.Name == "" {
				continue
			}
			ident := fmt.Sprintf("%s.%s", ident, sf.Name)
			field := value.Field(i)
			if !field.CanInterface() {
				continue
			}
			data := field.Interface()
			s, d, err := walkDataType(append(p, sf.Name), ident, data)
			if err != nil {
				return records, deps, err
			}
			records = append(records, s...)
			deps = append(deps, d...)
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
	Call      string
	Slice     string
	Array     string
	ByteArray string
	Struct    string
	Map       string
	MapStruct string
	Pointer   string
}

// genHeader writes the method header with pre declared variables if required.
func genHeader(w io.Writer, records []genRecord, withDecl bool, head string, data map[string]string) error {
	const decl = `	var %var% %kind%
`
	vars := make(map[string]string)
	genHeaderNext(records, withDecl, vars)

	data["layout"] = strings.Join(genCheck(records, nil), "")
	if err := templateExec(w, head, data); err != nil {
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

func genCheck(records []genRecord, s []string) []string {
	// Cache layout strings.
	m := map[reflect.Kind]string{
		reflect.Slice: string('A' + reflect.Slice),
		reflect.Uint8: string('A' + reflect.Uint8),
	}
	for _, rec := range records {
		if _, ok := m[rec.RKind]; !ok {
			m[rec.RKind] = string('A' + rec.RKind)
		}
		s = append(s, m[rec.RKind])
		switch rec.Is {
		case isSlice, isPointer:
			s = genCheck(rec.Include, s)
		case isArray, isByteArray:
			size := rec.Kind[1:strings.Index(rec.Kind, "]")]
			s = append(s, size)
			s = genCheck(rec.Include, s)
		case isMap:
			s = genCheck(rec.Key, s)
			s = genCheck(rec.Include, s)
		}
	}
	return s
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
			genHeaderNext(rec.Include, withDecl, vars)
		case isMap:
			reg("_n", "int")
			genHeaderNext(rec.Key, withDecl, vars)
			genHeaderNext(rec.Include, withDecl, vars)
		case isArray:
			genHeaderNext(rec.Include, withDecl, vars)
		case isPointer:
			reg("_bool", "bool")
			genHeaderNext(rec.Include, withDecl, vars)
		}
		switch kind := rec.Kind; kind {
		case "bool",
			"int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64",
			"float32", "float64",
			"complex64", "complex128",
			"string":
			reg("_"+kind, kind)
		case "[]byte":
			reg("_bytes", kind)
		}
	}
}

func genBody(level int, w io.Writer, records []genRecord, tmpls genConfig, data map[string]string, conv convFunc) error {
	defer func(t string) { data["tab"] = t }(data["tab"])
	data["tab"] += "\t"
	tab := data["tab"]
	inctab := tab + "\t"
	var include strings.Builder

	wconv := sliceConv(conv)
	kconv := keyConv(conv)
	vconv := valueConv(conv)
	pconv := pointerConv(conv)
	doinc := func(incname string, records []genRecord, conv convFunc) error {
		include.Reset()
		data["tab"] = inctab
		if err := genBody(level+1, &include, records, tmpls, data, conv); err != nil {
			return err
		}
		data["tab"] = tab
		data[incname] = include.String()
		return nil
	}
	doalloc := func(rec genRecord) error {
		var alloc string
		if rec.Is == isMap {
			alloc = `
%tab%%idlevel% = make(%kind%)`
		} else {
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
			if err := doinc("include", rec.Include, wconv); err != nil {
				return err
			}
			s = tmpls.Slice
		case isArray:
			if err := doinc("include", rec.Include, wconv); err != nil {
				return err
			}
			s = tmpls.Array
		case isByteArray:
			s = tmpls.ByteArray
		case isMapStruct:
			s = tmpls.MapStruct
		case isStruct:
			s = tmpls.Struct
		case isMap:
			if err := doinc("includekey", rec.Key, kconv); err != nil {
				return err
			}
			if err := doinc("include", rec.Include, vconv); err != nil {
				return err
			}
			data["kindkey"] = rec.Key[0].Kind
			s = tmpls.Map
		case isPointer:
			if err := doinc("include", rec.Include, pconv); err != nil {
				return err
			}
			if err := doalloc(rec.Include[0]); err != nil {
				return err
			}
			s = tmpls.Pointer
		default:
			s = tmpls.Call
		}
		data["id"] = rec.Ident
		data["kind"] = rec.Kind
		data["idlevel"], data["value"], data["conv"] = conv(rec)
		if err := templateExec(w, s, data); err != nil {
			return err
		}
		data["kindkey"] = ""
		data["alloc"] = ""
	}
	return nil
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

func genMarshalBinTo(w io.Writer, records []genRecord, receiver string, data interface{}) error {
	const (
		head = `
const _%type%Layout = "%layout%"

func (%rcv% *%type%) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = %pkg%.Write_layout(w, _b, _%type%Layout); if err != nil { return }
`
		call = `
%tab%err = %pkg%.Write_%kind%(w, _b, %conv%); if err != nil { return }
`
		slice = `
%tab%{
%tab%	_s := %value%
%tab%	_n = len(_s)
%tab%	err = %pkg%.Write_int(w, _b, _n); if err != nil { return }
%tab%	for _k := 0; _k < _n; _k++ {%include%	%tab%}
%tab%}`
		array = `
%tab%{
%tab%	_s := &%value%
%tab%	for _k := 0; _k < len(_s); _k++ {%include%	%tab%}
%tab%}`
		bytearray = `
%tab%err = %pkg%.Write_bytea(w, (%conv%)[:]); if err != nil { return }
`
		structt = `
%tab%err = %idlevel%.MarshalBinaryTo(w); if err != nil { return }
`
		mapstruct = `
%tab%{
%tab%	_struct := %idlevel%
%tab%	err = _struct.MarshalBinaryTo(w); if err != nil { return }
%tab%}
`
		mapp = `
%tab%{
%tab%	_s := %value%
%tab%	err = %pkg%.Write_int(w, _b, len(_s)); if err != nil { return }
%tab%	for _k := range _s {%includekey%%include%	%tab%}
%tab%}`
		pointer = `
%tab%err = %pkg%.Write_bool(w, _b, %idlevel% == nil); if err != nil { return }
%tab%if %idlevel% != nil {%include%	%tab%}
`
		tail = `
	return
}
`
	)
	m := map[string]string{
		"pkg":  localpkgName,
		"rcv":  receiver,
		"type": reflect.TypeOf(data).Name(),
	}
	if err := genHeader(w, records, false, head, m); err != nil {
		return err
	}
	conv := func(rec genRecord) (a, b, c string) {
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
	}
	err := genBody(0, w, records,
		genConfig{
			Call:      call,
			Slice:     slice,
			Array:     array,
			ByteArray: bytearray,
			Struct:    structt,
			Map:       mapp,
			MapStruct: mapstruct,
			Pointer:   pointer,
		},
		m,
		conv)
	if err == nil {
		_, err = w.Write([]byte(tail))
	}
	return err
}

func genUnmarshalBinFrom(w io.Writer, records []genRecord, receiver string, data interface{}) error {
	const (
		head = `
func (%rcv% *%type%) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = %pkg%.Read_layout(r, _b, _%type%Layout); if err != nil { return }
`
		call = `
%tab%_%kind%, err = %pkg%.Read_%kind%(r, _b); if err != nil { return }
%tab%%value% = %conv%
`
		slice = `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if _c := cap(%value%); _n > _c || _c - _n > _c/8 { %value% = make(%kind%, _n) } else { %value% = (%value%)[:_n] }
%tab%if _n > 0 {
%tab%	_s := %value%
%tab%	for _k := 0; _k < _n; _k++ {%include%	%tab%}
%tab%}
`
		array = `
%tab%{
%tab%	_s := &%value%
%tab%	for _k := 0; _k < len(_s); _k++ {%include%	%tab%}
%tab%}
`
		bytearray = `
%tab%err = %pkg%.Read_bytea(r, (%value%)[:]); if err != nil { return }
`
		structt = `
%tab%err = %idlevel%.UnmarshalBinaryFrom(r); if err != nil { return }
`
		mapstruct = `
%tab%{
%tab%	_struct := %idlevel%
%tab%	err = _struct.UnmarshalBinaryFrom(r); if err != nil { return }
%tab%	%idlevel% = _struct
%tab%}
`
		mapp = `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if _n == 0 {  %idlevel% = nil } else {
%tab%	%idlevel% = make(%kind%, _n)
%tab%	_s := %idlevel%
%tab%	var _k %kindkey%
%tab%	for _j := 0; _j < _n; _j++ {%includekey%%include%	%tab%}
%tab%}
`
		pointer = `
%tab%_bool, err = %pkg%.Read_bool(r, _b); if err != nil { return }
%tab%if _bool { %idlevel% = nil } else {%alloc%%include%%tab%}
`
		tail = `
	return
}
`
	)
	m := map[string]string{
		"pkg":  localpkgName,
		"rcv":  receiver,
		"type": reflect.TypeOf(data).Name(),
	}
	if err := genHeader(w, records, true, head, m); err != nil {
		return err
	}
	conv := func(rec genRecord) (a, b, c string) {
		id := rec.Ident
		val := rec.Ident
		kind := rec.Kind
		conv := "_" + kind
		if rec.Is == isPointer {
			val = "*" + val
		}
		if rec.Name == kind {
			return id, val, conv
		}
		return id, val, fmt.Sprintf("%s(%s)", rec.Name, conv)
	}
	err := genBody(0, w, records,
		genConfig{
			Call:      call,
			Slice:     slice,
			Array:     array,
			ByteArray: bytearray,
			Struct:    structt,
			MapStruct: mapstruct,
			Map:       mapp,
			Pointer:   pointer,
		},
		m,
		conv)
	if err == nil {
		_, err = fmt.Fprintf(w, tail)
	}
	return err
}
