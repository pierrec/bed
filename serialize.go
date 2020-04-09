package serializer

import (
	"encoding/binary"
	"io"
	"math"
)

func Write_bool(w io.Writer, buf []byte, v bool) error {
	if v {
		buf[0] = 1
	}
	_, err := w.Write(buf[:1])
	return err
}

func Write_int(w io.Writer, buf []byte, v int) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_int8(w io.Writer, buf []byte, v int8) error {
	buf[0] = byte(v)
	_, err := w.Write(buf[:1])
	return err
}

func Write_int16(w io.Writer, buf []byte, v int16) error {
	binary.LittleEndian.PutUint16(buf, uint16(v))
	_, err := w.Write(buf[:2])
	return err
}

func Write_int32(w io.Writer, buf []byte, v int32) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_int64(w io.Writer, buf []byte, v int64) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_uint(w io.Writer, buf []byte, v uint) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_uint8(w io.Writer, buf []byte, v uint8) error {
	buf[0] = v
	_, err := w.Write(buf[:1])
	return err
}

func Write_uint16(w io.Writer, buf []byte, v uint16) error {
	binary.LittleEndian.PutUint16(buf, v)
	_, err := w.Write(buf[:2])
	return err
}

func Write_uint32(w io.Writer, buf []byte, v uint32) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_uint64(w io.Writer, buf []byte, v uint64) error {
	return packUint64To(w, buf, v)
}

func Write_float32(w io.Writer, buf []byte, v float32) error {
	u := math.Float32bits(v)
	return packUint64To(w, buf, uint64(u))
}

func Write_float64(w io.Writer, buf []byte, v float64) error {
	u := math.Float64bits(v)
	return packUint64To(w, buf, u)
}

func Write_complex64(w io.Writer, buf []byte, v complex64) error {
	re := uint64(math.Float32bits(real(v)))
	im := uint64(math.Float32bits(imag(v)))
	return packUint64To(w, buf, re<<32|im)
}

func Write_complex128(w io.Writer, buf []byte, v complex128) error {
	re := math.Float64bits(real(v))
	if err := packUint64To(w, buf, re); err != nil {
		return err
	}
	im := math.Float64bits(imag(v))
	return packUint64To(w, buf, im)
}

func Write_string(w io.Writer, buf []byte, v string) error {
	if err := Write_int(w, buf, len(v)); err != nil {
		return err
	}
	_, err := w.Write([]byte(v))
	return err
}

func Write_bytes(w io.Writer, buf []byte, v []byte) error {
	if err := Write_int(w, buf, len(v)); err != nil {
		return err
	}
	_, err := w.Write(v)
	return err
}

func Write_bytea(w io.Writer, v []byte) error {
	_, err := w.Write(v)
	return err
}
