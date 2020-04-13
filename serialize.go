package serializer

import (
	"encoding/binary"
	"io"
	"math"
	"math/bits"
	"time"
)

func Write_layout(w io.Writer, buf []byte, layout string) error {
	return Write_string(w, buf, layout)
}

func Write_bool(w io.Writer, buf []byte, v bool) error {
	if v {
		buf[0] = 1
	} else {
		buf[0] = 0
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
	u := bits.Reverse32(math.Float32bits(v))
	return packUint64To(w, buf, uint64(u))
}

func Write_float64(w io.Writer, buf []byte, v float64) error {
	u := bits.Reverse64(math.Float64bits(v))
	return packUint64To(w, buf, u)
}

func Write_complex64(w io.Writer, buf []byte, v complex64) error {
	if err := Write_float32(w, buf, real(v)); err != nil {
		return err
	}
	return Write_float32(w, buf, imag(v))
}

func Write_complex128(w io.Writer, buf []byte, v complex128) error {
	if err := Write_float64(w, buf, real(v)); err != nil {
		return err
	}
	return Write_float64(w, buf, imag(v))
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

func Write_time(w io.Writer, buf []byte, t time.Time) error {
	if t.IsZero() {
		_, err := w.Write([]byte{0, 0})
		return err
	}
	// item<size in bits>
	// year<16> month<4> day<5> hour<5> minute<6> second<6> nanosecond<32> TZoffset<5>
	// item<size in bytes>
	// year<2> month<1> day<1> hour<1> minute<1> second<1> nanosecond<4> TZoffset<1>
	buf = buf[:12]
	b := buf
	year, month, day := t.Date()
	binary.LittleEndian.PutUint16(b, uint16(year))
	b = b[2:]

	hour, min, sec := t.Clock()
	_, offset := t.Zone()
	offset /= 60 * 60 // offset in hours
	b[0] = uint8(month)
	b[1] = uint8(day)
	b[2] = uint8(hour)
	b[3] = uint8(min)
	b[4] = uint8(sec)
	b[5] = uint8(offset)
	b = b[6:]

	ns := t.Nanosecond()
	binary.LittleEndian.PutUint32(b, uint32(ns))
	b = b[4:]

	_, err := w.Write(buf)
	return err
}
