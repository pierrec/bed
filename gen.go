package serializer

import (
	"fmt"
	"io"
	"path"
	"reflect"
	"sort"
	"strings"
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
	"strings"

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

	receiver := config.Receiver
	for i := 0; i < len(data); i++ {
		item := data[i]
		records, deps, err := walkDataType(nil, receiver, item)
		if err != nil {
			return err
		}
		for _, dep := range deps {
			if !hasType(reflect.TypeOf(dep).Name(), data) {
				data = append(data, dep)
			}
		}
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

type convFunc func(genRecord) (value, conv string)

const (
	_ uint8 = iota
	isSlice
	isArray
	isByteArray
	isStruct
	isMap
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
		typ = typ.Elem()
		value := reflect.New(typ).Elem()
		subrecords, subdeps, err := walkDataType(append(p, typ.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		if k := subrecords[0].Kind; len(subrecords) == 1 && k == "uint8" {
			// Special case: []byte
			records = []genRecord{
				{RKind: kind, Ident: ident, Kind: "bytes", Name: "bytes"},
			}
		} else {
			records = []genRecord{
				{RKind: kind, Is: isSlice, Ident: ident, Kind: "[]" + k, Include: subrecords},
			}
		}
		deps = append(deps, subdeps...)
	case reflect.Array:
		typ = typ.Elem()
		value := reflect.New(typ).Elem()
		subrecords, subdeps, err := walkDataType(append(p, typ.String()), ident, value.Interface())
		if err != nil {
			return nil, nil, err
		}
		if k := subrecords[0].Kind; len(subrecords) == 1 && k == "uint8" {
			// Special case: [...]byte
			records = []genRecord{
				{RKind: kind, Is: isByteArray, Ident: ident, Kind: "bytea", Name: "bytea"},
			}
		} else {
			records = []genRecord{
				{RKind: kind, Is: isArray, Ident: ident, Kind: k, Include: subrecords},
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
		// Append the key definition to the value one.
		records = []genRecord{
			{RKind: kind, Is: isMap, Ident: ident, Kind: subrecords[0].Kind, Include: subrecords, Key: keyrecords}}
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
			if sf.Anonymous {
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
}

// genHeader writes the method header with pre declared variables if required.
func genHeader(w io.Writer, records []genRecord, withDecl bool, head string, data map[string]string) error {
	const decl = `	var %var% %kind%
`
	vars := make(map[string]string)
	genHeaderNext(records, withDecl, vars)

	data["check"] = strings.Join(genCheck(records, nil), "")
	if err := templateExec(w, head, data); err != nil {
		return err
	}
	delete(data, "check")
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
	for _, rec := range records {
		if rec.Kind == "bytes" {
			s = append(s, string('A'+reflect.Uint8))
			continue
		}
		s = append(s, string('A'+rec.RKind))
		switch rec.Is {
		case isSlice, isArray, isByteArray:
			s = genCheck(rec.Include, s)
		case isMap:
			s = genCheck(rec.Key, s)
			s = genCheck(rec.Include, s)
		}
	}
	return s
}

func genHeaderNext(records []genRecord, withDecl bool, vars map[string]string) {
	// Prepare all variables.
	for _, rec := range records {
		switch rec.Is {
		case isSlice:
			vars["_n"] = "int"
			genHeaderNext(rec.Include, withDecl, vars)
		case isMap:
			if withDecl {
				vars["_n"] = "int"
			}
			genHeaderNext(rec.Key, withDecl, vars)
			genHeaderNext(rec.Include, withDecl, vars)
		}
		switch kind := rec.Kind; kind {
		case "bool",
			"int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64",
			"float32", "float64",
			"complex64", "complex128",
			"string":
			if withDecl {
				vars["_"+kind] = kind
			}
		case "bytes":
			if withDecl {
				vars["_bytes"] = "[]byte"
			}
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
	doinc := func(ident, incname string, records []genRecord, conv convFunc) error {
		include.Reset()
		data["tab"] = inctab
		if err := genBody(level+1, &include, records, tmpls, data, conv); err != nil {
			return err
		}
		if level == 0 {
			data["idlevel"] = ident
		} else {
			data["idlevel"] = "_s[_i]"
		}
		data["tab"] = tab
		data[incname] = include.String()
		return nil
	}

	for _, rec := range records {
		var s string
		switch rec.Is {
		case isSlice:
			if err := doinc(rec.Ident, "include", rec.Include, wconv); err != nil {
				return err
			}
			s = tmpls.Slice
		case isArray:
			if err := doinc(rec.Ident, "include", rec.Include, wconv); err != nil {
				return err
			}
			s = tmpls.Array
		case isByteArray:
			s = tmpls.ByteArray
		case isStruct:
			s = tmpls.Struct
		case isMap:
			if err := doinc(rec.Ident, "includekey", rec.Key, kconv); err != nil {
				return err
			}
			if err := doinc(rec.Ident, "include", rec.Include, vconv); err != nil {
				return err
			}
			data["kindkey"] = rec.Key[0].Kind
			s = tmpls.Map
		default:
			s = tmpls.Call
		}
		data["id"] = rec.Ident
		data["kind"] = rec.Kind
		data["value"], data["conv"] = conv(rec)
		if err := templateExec(w, s, data); err != nil {
			return err
		}
	}
	return nil
}

func sliceConv(conv convFunc) convFunc {
	return func(rec genRecord) (string, string) {
		const _s = "_s[_i]"
		v, c := conv(rec)
		return _s, strings.ReplaceAll(c, v, _s)
	}
}

func keyConv(conv convFunc) convFunc {
	return func(rec genRecord) (string, string) {
		const _s = "_i"
		v, c := conv(rec)
		return _s, strings.ReplaceAll(c, v, _s)
	}
}

func valueConv(conv convFunc) convFunc {
	return func(rec genRecord) (string, string) {
		const _s = "_s[_i]"
		v, c := conv(rec)
		return _s, strings.ReplaceAll(c, v, _s)
	}
}

func genMarshalBinTo(w io.Writer, records []genRecord, receiver string, data interface{}) error {
	const (
		head = `
func (%rcv% *%type%) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "%check%"
	var _buf [16]byte
	_b := _buf[:]
	err = %pkg%.Write_string(w, _b, _check); if err != nil { return }
`
		call = `
%tab%err = %pkg%.Write_%kind%(w, _b, %conv%); if err != nil { return }
`
		slice = `
%tab%{
%tab%	_s := %idlevel%
%tab%	_n = len(_s)
%tab%	err = %pkg%.Write_int(w, _b, _n); if err != nil { return }
%tab%	for _i := 0; _i < _n; _i++ {%include%	%tab%}
%tab%}`
		array = `
%tab%{
%tab%	_s := &%idlevel%
%tab%	for _i := 0; _i < len(_s); _i++ {%include%	%tab%}
%tab%}`
		bytearray = `
%tab%err = %pkg%.Write_bytea(w, %conv%[:]); if err != nil { return }
`
		structt = `
%tab%err = %id%.MarshalBinaryTo(w); if err != nil { return }
`
		mapp = `
%tab%{
%tab%	_s := %idlevel%
%tab%	err = %pkg%.Write_int(w, _b, len(_s)); if err != nil { return }
%tab%	for _i := range _s {%includekey%%include%	%tab%}
%tab%}`
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
	conv := func(rec genRecord) (string, string) {
		val := rec.Ident
		kind := rec.Kind
		if rec.Name == kind {
			return val, val
		}
		return val, fmt.Sprintf("%s(%s)", kind, val)
	}
	err := genBody(0, w, records,
		genConfig{
			Call:      call,
			Slice:     slice,
			Array:     array,
			ByteArray: bytearray,
			Struct:    structt,
			Map:       mapp,
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
	const _check = "%check%"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := %pkg%.Read_string(r, _b); err != nil { return err
	} else if !strings.HasPrefix(s, _check) { return %pkg%.ErrInvalidData }
`
		call = `
%tab%_%kind%, err = %pkg%.Read_%kind%(r, _b); if err != nil { return }
%tab%%value% = %conv%
`
		slice = `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if c := cap(%idlevel%); _n > c || c - _n > c/8 { %idlevel% = make(%kind%, _n) } else { %idlevel% = %idlevel%[:_n] }
%tab%if _n > 0 {
%tab%	_s := %idlevel%
%tab%	for _i := 0; _i < _n; _i++ {%include%	%tab%}
%tab%}
`
		array = `
%tab%{
%tab%	_s := &%idlevel%
%tab%	for _i := 0; _i < len(_s); _i++ {%include%	%tab%}
%tab%}
`
		bytearray = `
%tab%err = %pkg%.Read_bytea(r, %value%[:]); if err != nil { return }
`
		structt = `
%tab%err = %id%.UnmarshalBinaryFrom(r); if err != nil { return }
`
		mapp = `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if _n > 0 {
%tab%	%idlevel% = make(map[%kindkey%]%kind%, _n)
%tab%	_s := %idlevel%
%tab%	var _i %kindkey%
%tab%	for _j := 0; _j < _n; _j++ {%includekey%%include%	%tab%}
%tab%} else { %idlevel% = nil }
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
	conv := func(rec genRecord) (string, string) {
		kind := rec.Kind
		conv := "_" + kind
		if rec.Name == kind {
			return rec.Ident, conv
		}
		return rec.Ident, fmt.Sprintf("%s(%s)", rec.Name, conv)
	}
	err := genBody(0, w, records,
		genConfig{
			Call:      call,
			Slice:     slice,
			Array:     array,
			ByteArray: bytearray,
			Struct:    structt,
			Map:       mapp,
		},
		m,
		conv)
	if err == nil {
		_, err = fmt.Fprintf(w, tail)
	}
	return err
}
