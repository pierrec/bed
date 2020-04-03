// Package serializer provides an easy to use way of serializing and deserializing any data at runtime.
package serializer

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"
)

// Record keeps track of the struct elements being serialized.
// Slices are encoded as: <slice length><item0>...
// Structs are encoded in their fields order.
type Record struct {
	IsSlice bool
	Ident   string       // target identifier
	Kind    reflect.Kind // target kind (only fixed size kinds)
}

type Accumulator int64

func (s *Accumulator) writevarint(w io.Writer, buf []byte, v uint64) error {
	if v < 128 {
		buf[0] = 0x8 | byte(v)
		return s.Write_bytes(w, buf[:1])
	}
	n := binary.PutUvarint(buf[1:], v)
	buf[0] = byte(n)
	return s.Write_bytes(w, buf[:n+1])
}

func (s *Accumulator) Write_bool(w io.Writer, buf []byte, v bool) error {
	if v {
		buf[0] = 1
	}
	return s.Write_bytes(w, buf[:1])
}

func (s *Accumulator) Write_int(w io.Writer, buf []byte, v int) error {
	return s.writevarint(w, buf, uint64(v))
}

func (s *Accumulator) Write_int8(w io.Writer, buf []byte, v int8) error {
	buf[0] = byte(v)
	return s.Write_bytes(w, buf[:1])
}

func (s *Accumulator) Write_int16(w io.Writer, buf []byte, v int16) error {
	binary.LittleEndian.PutUint16(buf, uint16(v))
	return s.Write_bytes(w, buf[:2])
}

func (s *Accumulator) Write_int32(w io.Writer, buf []byte, v int32) error {
	binary.LittleEndian.PutUint32(buf, uint32(v))
	return s.Write_bytes(w, buf[:4])
}

func (s *Accumulator) Write_int64(w io.Writer, buf []byte, v int64) error {
	binary.LittleEndian.PutUint64(buf, uint64(v))
	return s.Write_bytes(w, buf[:8])
}

func (s *Accumulator) Write_uint(w io.Writer, buf []byte, v uint) error {
	return s.writevarint(w, buf, uint64(v))
}

func (s *Accumulator) Write_uint8(w io.Writer, buf []byte, v uint8) error {
	buf[0] = v
	return s.Write_bytes(w, buf[:1])
}

func (s *Accumulator) Write_uint16(w io.Writer, buf []byte, v uint16) error {
	binary.LittleEndian.PutUint16(buf, v)
	return s.Write_bytes(w, buf[:2])
}

func (s *Accumulator) Write_uint32(w io.Writer, buf []byte, v uint32) error {
	binary.LittleEndian.PutUint32(buf, v)
	return s.Write_bytes(w, buf[:4])
}

func (s *Accumulator) Write_uint64(w io.Writer, buf []byte, v uint64) error {
	binary.LittleEndian.PutUint64(buf, v)
	return s.Write_bytes(w, buf[:8])
}

func (s *Accumulator) Write_float32(w io.Writer, buf []byte, v float32) error {
	binary.LittleEndian.PutUint32(buf, math.Float32bits(v))
	return s.Write_bytes(w, buf[:4])
}

func (s *Accumulator) Write_float64(w io.Writer, buf []byte, v float64) error {
	binary.LittleEndian.PutUint64(buf, math.Float64bits(v))
	return s.Write_bytes(w, buf[:8])
}

func (s *Accumulator) Write_complex64(w io.Writer, buf []byte, v complex64) error {
	binary.LittleEndian.PutUint32(buf, math.Float32bits(real(v)))
	binary.LittleEndian.PutUint32(buf[4:], math.Float32bits(imag(v)))
	return s.Write_bytes(w, buf[:8])
}

func (s *Accumulator) Write_complex128(w io.Writer, buf []byte, v complex128) error {
	binary.LittleEndian.PutUint64(buf, math.Float64bits(real(v)))
	binary.LittleEndian.PutUint64(buf[8:], math.Float64bits(imag(v)))
	return s.Write_bytes(w, buf[:16])
}

func (s *Accumulator) Write_string(w io.Writer, buf []byte, v string) error {
	if err := s.Write_int(w, buf, len(v)); err != nil {
		return err
	}
	return s.Write_bytes(w, []byte(v))
}

func (s *Accumulator) Write_bytes(w io.Writer, b []byte) error {
	n, err := w.Write(b)
	*s += Accumulator(n)
	return err
}

func walkData(ident string, data interface{}) ([]Record, error) {
	var records []Record
	typ := reflect.TypeOf(data)
	switch kind := typ.Kind(); kind {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		records = []Record{
			{Ident: ident, Kind: kind},
		}
	case reflect.Slice, reflect.Array:
		elemKind := typ.Elem().Kind()
		records = []Record{
			{IsSlice: true, Ident: ident, Kind: elemKind},
		}
	case reflect.Struct:
		value := reflect.ValueOf(data)
		n := typ.NumField()
		for i := 0; i < n; i++ {
			sf := typ.Field(i)
			if sf.Anonymous {
				continue
			}
			ident := fmt.Sprintf("%s.%s", ident, sf.Name)
			data := value.Field(i).Interface()
			s, err := walkData(ident, data)
			if err != nil {
				return records, err
			}
			records = append(records, s...)
		}
	default:
		err := fmt.Errorf("binary.Write: unsupported type %T", data)
		return nil, err
	}
	return records, nil
}

func genMarshalBinTo(w io.Writer, records []Record, name string, data interface{}) error {
	const (
		head = `
func (%rcv% *%type%) MarshalBinaryTo(w io.Writer) (n int64, err error) {
	var acc serializer.Accumulator
	defer func() { n = int64(acc) }()
	var buf [16]byte
	b := buf[:]
	var in int
	_ = in
`
		call = `
	err = acc.Write_%kind%(w, b, %id%); if err != nil { return }
`
		direct = `
	err = acc.Write_bytes(w, %id%); if err != nil { return }
`
		loop = `
	in = len(%id%)
	err = acc.Write_int(w, b, in); if err != nil { return }
	for i := 0; i < in; i++ {
		err = acc.Write_%kind%(w, b, %id%[i]); if err != nil { return }
	}
`
		tail = `
	return
}`
	)

	m := map[string]interface{}{
		"rcv":  name,
		"type": fmt.Sprintf("%T", data),
	}
	if err := templateExec(w, head, m); err != nil {
		return err
	}
	for _, rec := range records {
		m["kind"] = rec.Kind
		m["id"] = name + rec.Ident
		s := call
		switch {
		case rec.IsSlice && rec.Kind == reflect.Uint8:
			s = direct
		case rec.IsSlice && rec.Kind != reflect.Uint8:
			s = loop
		}
		if err := templateExec(w, s, m); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(w, tail)
	return err
}
