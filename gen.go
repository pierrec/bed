package serializer

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

// Interface is the interface added to types processed by Gen.
type Interface interface {
	MarshalBinaryTo(io.Writer) error
	UnmarshalBinaryFrom(io.Reader) error
}

var _Interface = reflect.TypeOf([]Interface(nil)).Elem()

// Gen generates the MarshalBinaryTo and UnmarshalBinaryFrom methods for
// the provided types out to the given Writer.
//
// It does *not* generate the package header or imports.
func Gen(out io.Writer, data ...interface{}) error {
	const receiver = "self"
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

const pkgName = "serializer"

// genRecord keeps track of the struct elements being serialized.
// Slices are encoded as: <slice length><item0>...
// Structs are encoded in their fields order.
type genRecord struct {
	Is      uint8
	Ident   string // target identifier
	Kind    string // target kind (only fixed size kinds)
	Name    string // target type name
	Include []genRecord
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
			{Ident: ident, Kind: kind.String(), Name: typ.String()},
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
				{Ident: ident, Kind: "bytes", Name: "bytes"},
			}
		} else {
			records = []genRecord{
				{Is: isSlice, Ident: ident, Kind: "[]" + k, Include: subrecords},
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
				{Is: isByteArray, Ident: ident, Kind: "bytea", Name: "bytea"},
			}
		} else {
			records = []genRecord{
				{Is: isArray, Ident: ident, Kind: k, Include: subrecords},
			}
		}
		deps = append(deps, subdeps...)
	//case reflect.Map:
	//	value := reflect.New(typ.Elem()).Elem()
	//	subrecords, err := walkDataType(ident, value.Interface())
	//	if err != nil {
	//		return nil, err
	//	}
	//	k := subrecords[0].Kind
	//	records = []genRecord{
	//		{Is: isMap, Ident: ident, Kind: k, Include: subrecords}, // value
	//		{Kind: typ.Key().Kind().String()},                       // key
	//	}
	case reflect.Struct:
		value := reflect.ValueOf(data)
		if len(p) > 0 && typ.Name() != "" {
			// Named struct type: add to the the list of dependents to get the marshal methods
			// if it does not already implement the methods.
			records = append(records, genRecord{Is: isStruct, Ident: ident, Kind: typ.String()})
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

func genMarshalBinTo(w io.Writer, records []genRecord, receiver string, data interface{}) error {
	const (
		head = `
func (%rcv% *%type%) MarshalBinaryTo(w io.Writer) (err error) {`
		call = `
%tab%err = %pkg%.Write_%kind%(w, _b, %conv%); if err != nil { return }
`
		slice = `
%tab%{
	%tab%_s := %idlevel%
	%tab%_n = len(_s)
	%tab%err = %pkg%.Write_int(w, _b, _n); if err != nil { return }
	%tab%for _i := 0; _i < _n; _i++ {%include%	%tab%}
%tab%}`
		array = `
%tab%{
	%tab%_s := %idlevel%
	%tab%for _i := 0; _i < len(_s); _i++ {%include%	%tab%}
%tab%}`
		bytearray = `
%tab%err = %pkg%.Write_bytea(w, %conv%[:]); if err != nil { return }
`
		structt = `
%tab%err = %id%.MarshalBinaryTo(w); if err != nil { return }
`
		mapp = `
%tab%{
	%tab%_s := %idlevel%
	%tab%err = %pkg%.Write_int(w, _b, len(_s)); if err != nil { return }
	%tab%for _i := range _s {%include%	%tab%}
%tab%}`
		tail = `
	return
}
`
	)
	m := map[string]string{
		"pkg":  pkgName,
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

// genHeader writes the method header with pre declared variables if required.
func genHeader(w io.Writer, records []genRecord, withDecl bool, head string, data map[string]string) error {
	const decl = `	var %var% %kind%
`
	vars := make(map[string]string)
	genHeaderNext(w, records, withDecl, vars)

	if err := templateExec(w, head, data); err != nil {
		return err
	}
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

func genHeaderNext(w io.Writer, records []genRecord, withDecl bool, vars map[string]string) {
	// Prepare all variables.
	for _, rec := range records {
		if rec.Is == isSlice {
			vars["_n"] = "int"
			genHeaderNext(w, rec.Include, withDecl, vars)
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
		default:
			continue
		}
		// Add a buffer as we have a Read or Write call.
		vars["__buf"] = "[16]byte"
		vars["_b"] = "= __buf[:]"
	}
}

func genBody(level int, w io.Writer, records []genRecord, tmpls genConfig, data map[string]string, conv convFunc) error {
	defer func(t string) { data["tab"] = t }(data["tab"])
	data["tab"] += "\t"
	tab := data["tab"]
	inctab := tab + "\t"
	var include strings.Builder
	for _, rec := range records {
		if rec.Is > 0 {
			include.Reset()
			data["tab"] = inctab
			conv := conv
			if level == 0 {
				conv = wrapConv(conv)
			}
			if err := genBody(level+1, &include, rec.Include, tmpls, data, conv); err != nil {
				return err
			}
			if level == 0 {
				data["idlevel"] = rec.Ident
			} else {
				data["idlevel"] = "_s[_i]"
			}
			data["tab"] = tab
			data["include"] = include.String()
		}
		var s string
		switch rec.Is {
		case isSlice:
			s = tmpls.Slice
		case isArray:
			s = tmpls.Array
		case isByteArray:
			s = tmpls.ByteArray
		case isStruct:
			s = tmpls.Struct
		case isMap:
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

func wrapConv(conv convFunc) convFunc {
	return func(rec genRecord) (string, string) {
		const _s = "_s[_i]"
		v, c := conv(rec)
		return _s, strings.ReplaceAll(c, v, _s)
	}
}

func genUnmarshalBinFrom(w io.Writer, records []genRecord, receiver string, data interface{}) error {
	const (
		head = `
func (%rcv% *%type%) UnmarshalBinaryFrom(r io.Reader) (err error) {`
		call = `
%tab%_%kind%, err = %pkg%.Read_%kind%(r, _b); if err != nil { return }
%tab%%value% = %conv%
`
		slice = `
%tab%_n, err = %pkg%.Read_int(r, _b); if err != nil { return }
%tab%if c := cap(%id%); _n > c || c - _n > c/8 { %id% = make(%kind%, _n) } else { %id% = %id%[:_n] }
%tab%{
	%tab%_s := %idlevel%
	%tab%for _i := 0; _i < _n; _i++ {%include%	%tab%}
%tab%}
`
		array = `
%tab%{
	%tab%_s := %idlevel%
	%tab%for _i := 0; _i < len(_s); _i++ {%include%	%tab%}
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
%tab%%id% = make(map[%kindkey%]%kind%, _n)
%tab%{
	%tab%_s := %idlevel%
	%tab%for _mi := 0; _mi < _n; _i++ {%include%	%tab%}
%tab%}
`
		tail = `
	return
}
`
	)
	m := map[string]string{
		"pkg":  pkgName,
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
