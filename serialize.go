package serializer

import (
	"encoding/binary"
	"io"
	"math"
	"math/big"
	"math/bits"
	"time"
)

func Write_layout(w ByteWriter, buf []byte, layout string) error {
	return Write_string(w, buf, layout)
}

const (
	_false = 0
	_true  = 1
)

func Write_bool(w ByteWriter, _ []byte, v bool) error {
	if v {
		return w.WriteByte(_true)
	}
	return w.WriteByte(_false)
}

func Write_int(w ByteWriter, buf []byte, v int) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_int8(w ByteWriter, _ []byte, v int8) error {
	return w.WriteByte(byte(v))
}

func Write_int16(w ByteWriter, buf []byte, v int16) error {
	binary.LittleEndian.PutUint16(buf, uint16(v))
	_, err := w.Write(buf[:2])
	return err
}

func Write_int32(w ByteWriter, buf []byte, v int32) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_int64(w ByteWriter, buf []byte, v int64) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_uint(w ByteWriter, buf []byte, v uint) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_uint8(w ByteWriter, _ []byte, v uint8) error {
	return w.WriteByte(v)
}

func Write_uint16(w ByteWriter, buf []byte, v uint16) error {
	binary.LittleEndian.PutUint16(buf, v)
	_, err := w.Write(buf[:2])
	return err
}

func Write_uint32(w ByteWriter, buf []byte, v uint32) error {
	return packUint64To(w, buf, uint64(v))
}

func Write_uint64(w ByteWriter, buf []byte, v uint64) error {
	return packUint64To(w, buf, v)
}

func Write_float32(w ByteWriter, buf []byte, v float32) error {
	u := bits.Reverse32(math.Float32bits(v))
	return packUint64To(w, buf, uint64(u))
}

func Write_float64(w ByteWriter, buf []byte, v float64) error {
	u := bits.Reverse64(math.Float64bits(v))
	return packUint64To(w, buf, u)
}

func Write_complex64(w ByteWriter, buf []byte, v complex64) error {
	if err := Write_float32(w, buf, real(v)); err != nil {
		return err
	}
	return Write_float32(w, buf, imag(v))
}

func Write_complex128(w ByteWriter, buf []byte, v complex128) error {
	if err := Write_float64(w, buf, real(v)); err != nil {
		return err
	}
	return Write_float64(w, buf, imag(v))
}

func Write_string(w ByteWriter, buf []byte, v string) error {
	if err := Write_int(w, buf, len(v)); err != nil {
		return err
	}
	//TODO
	if w, ok := w.(io.StringWriter); ok {
		_, err := w.WriteString(v)
		return err
	}
	_, err := w.Write([]byte(v))
	return err
}

func Write_bytes(w ByteWriter, buf []byte, v []byte) error {
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

func Write_time(w ByteWriter, buf []byte, t time.Time) error {
	if t.IsZero() {
		_, err := w.Write([]byte{0, 0})
		return err
	}
	// item<size in bits>
	// month=[1,12] day=[1,31] hour[0,23] minute,second=[0,59] offset=[0,23] nsec=[0,999999999]
	// year<16> month<4> day<5> hour<5> minute<6> second<6> TZoffset<5> hasNanosecond<1> nanosecond<32>
	// = 48 or 80 bits
	year, month, day := t.Date()
	binary.LittleEndian.PutUint16(buf, uint16(year))

	hour, min, sec := t.Clock()
	ns := t.Nanosecond()
	_, offset := t.Zone()
	offset /= 60 * 60 // offset in hours

	u := uint32(month)
	u = u<<5 | uint32(day)
	u = u<<5 | uint32(hour)
	u = u<<6 | uint32(min)
	u = u<<6 | uint32(sec)
	u = u<<5 | uint32(offset)
	u <<= 1
	if ns == 0 {
		buf = buf[:2+4]
	} else {
		u |= 1
		binary.LittleEndian.PutUint32(buf[2+4:], uint32(ns))
		buf = buf[:2+4+4]
	}
	binary.LittleEndian.PutUint32(buf[2:], u)
	_, err := w.Write(buf)
	return err
}

func Write_bigfloat(w ByteWriter, buf, bigbuf []byte, v big.Float) error {
	prec := int(v.MinPrec())
	bigbuf = v.Append(bigbuf[:0], 'g', prec)
	return Write_bytes(w, buf, bigbuf)
}

func Write_bigint(w ByteWriter, buf, bigbuf []byte, v big.Int) error {
	sign := v.Sign() + 1
	if err := Write_uint8(w, buf, uint8(sign)); err != nil {
		return err
	}
	if sign == 1 {
		// v == 0
		return nil
	}
	return Write_bytes(w, buf, v.Bytes())
}

func Write_bigrat(w ByteWriter, buf, bigbuf []byte, v big.Rat) error {
	if err := Write_bigint(w, buf, bigbuf, *v.Num()); err != nil {
		return err
	}
	return Write_bigint(w, buf, bigbuf, *v.Denom())
}
