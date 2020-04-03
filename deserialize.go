package serializer

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"
)

type ByteReader struct {
	n uint8
	io.Reader
}

func (r *ByteReader) ReadByte() (byte, error) {
	var buf [1]byte
	n, err := r.Reader.Read(buf[:])
	if n > 0 {
		r.n++
	}
	return buf[0], err
}

func (s *Accumulator) readvarint(r *ByteReader, buf []byte) (uint64, error) {
	err := s.read(r, buf[:1])
	if err != nil {
		return 0, err
	}
	if v := buf[0]; v&0x8 > 0 {
		return uint64(v), nil
	}
	panic("TODO")
	return 0, nil
}

func (s *Accumulator) read(r *ByteReader, buf []byte) error {
	n, err := io.ReadFull(r, buf)
	*s += Accumulator(n)
	return err
}

func (s *Accumulator) Read_int(r *ByteReader, buf []byte) (int, error) {
	v, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}
	*s += Accumulator(r.n)
	return int(v), nil
}

func (s *Accumulator) Read_int8(r *ByteReader, buf []byte) (int8, error) {
	err := s.read(r, buf[:1])
	return int8(buf[0]), err
}

func (s *Accumulator) Read_int16(r *ByteReader, buf []byte) (int16, error) {
	err := s.read(r, buf[:2])
	v := binary.LittleEndian.Uint16(buf)
	return int16(v), err
}

func (s *Accumulator) Read_int32(r *ByteReader, buf []byte) (int32, error) {
	err := s.read(r, buf[:4])
	v := binary.LittleEndian.Uint32(buf)
	return int32(v), err
}

func (s *Accumulator) Read_int64(r *ByteReader, buf []byte) (int64, error) {
	err := s.read(r, buf[:8])
	v := binary.LittleEndian.Uint64(buf)
	return int64(v), err
}

func (s *Accumulator) Read_uint(r *ByteReader, buf []byte) (uint, error) {
	v, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}
	*s += Accumulator(r.n)
	return uint(v), nil
}

func (s *Accumulator) Read_uint8(r *ByteReader, buf []byte) (uint8, error) {
	err := s.read(r, buf[:1])
	return buf[0], err
}

func (s *Accumulator) Read_uint16(r *ByteReader, buf []byte) (uint16, error) {
	err := s.read(r, buf[:2])
	v := binary.LittleEndian.Uint16(buf)
	return v, err
}

func (s *Accumulator) Read_uint32(r *ByteReader, buf []byte) (uint32, error) {
	err := s.read(r, buf[:4])
	v := binary.LittleEndian.Uint32(buf)
	return v, err
}

func (s *Accumulator) Read_uint64(r *ByteReader, buf []byte) (uint64, error) {
	err := s.read(r, buf[:8])
	v := binary.LittleEndian.Uint64(buf)
	return v, err
}

func (s *Accumulator) Read_float32(r *ByteReader, buf []byte) (float32, error) {
	_, err := r.Read(buf[:4])
	v := binary.LittleEndian.Uint32(buf)
	return math.Float32frombits(v), err
}

func (s *Accumulator) Read_float64(r *ByteReader, buf []byte) (float64, error) {
	err := s.read(r, buf[:4])
	v := binary.LittleEndian.Uint64(buf)
	return math.Float64frombits(v), err
}

func (s *Accumulator) Read_Complex64(r *ByteReader, buf []byte) (complex64, error) {
	err := s.read(r, buf[:8])
	re := math.Float32frombits(binary.LittleEndian.Uint32(buf))
	im := math.Float32frombits(binary.LittleEndian.Uint32(buf[4:]))
	return complex(re, im), err
}

func (s *Accumulator) Read_Complex128(r *ByteReader, buf []byte) (complex128, error) {
	err := s.read(r, buf[:16])
	re := math.Float64frombits(binary.LittleEndian.Uint64(buf))
	im := math.Float64frombits(binary.LittleEndian.Uint64(buf[8:]))
	return complex(re, im), err
}

func (s *Accumulator) Read_string(r *ByteReader, buf []byte) (string, error) {
	n, err := s.Read_int(r, buf)
	if err != nil || n == 0 {
		return "", err
	}
	if n > cap(buf) {
		buf = make([]byte, n)
	} else {
		buf = buf[:n]
	}
	err = s.read(r, buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (s *Accumulator) Read_bytes(r *ByteReader, buf []byte) ([]byte, error) {
	n, err := s.Read_int(r, buf)
	if err != nil || n == 0 {
		return nil, err
	}
	buf = make([]byte, n)
	err = s.read(r, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func genUnmarshalBinFrom(w io.Writer, records []Record, name string, data interface{}) error {
	const (
		head = `
func (%s *%T) UnmarshalBinaryFrom(r io.Reader) (n int64, err error) {
	var acc serializer.Accumulator
	defer func() { n = int64(acc) }()
	var buf [16]byte
	b := buf[:]
	var in int
	_ = in
	br := &serializer.ByteReader{Reader: r}
`
		call = `
	%s, err = acc.Read_%s(br, b); if err != nil { return }
`
		direct = `
	%s, err = acc.Read_bytes(br); if err != nil { return }
`
		loop = `
	in, err = acc.Read_int(br, b); if err != nil { return }
	if c := cap(%s); in > c || c - in > c/8 { %s = make([]%s, in) } else { %s = %s[:in] }
	for i := 0; i < in; i++ {
		%s[i], err = acc.Read_%s(br, b); if err != nil { return }
	}
`
		tail = `
	return
}`
	)

	var err error
	if _, err = fmt.Fprintf(w, head, name, data); err != nil {
		return err
	}
	for _, rec := range records {
		id := name + rec.Ident
		switch {
		case rec.IsSlice && rec.Kind == reflect.Uint8:
			_, err = fmt.Fprintf(w, direct, id)
		case rec.IsSlice && rec.Kind != reflect.Uint8:
			_, err = fmt.Fprintf(w, loop, id, id, rec.Kind, id, id, id, rec.Kind)
		default:
			_, err = fmt.Fprintf(w, call, id, rec.Kind)
		}
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(w, tail)
	return err
}
