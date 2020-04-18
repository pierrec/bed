package packed

import (
	"encoding/binary"
	"io"
	"math"
	"math/big"
	"math/bits"
	"time"

	"github.com/pierrec/packer"
	"github.com/pierrec/packer/iobyte"
	"github.com/pierrec/serializer/raw"
)

func Write_layout(w iobyte.ByteWriter, buf []byte, layout string) error {
	return Write_string(w, buf, layout)
}

func Write_bool(w iobyte.ByteWriter, _ []byte, v bool) error {
	return raw.Write_bool(w, nil, v)
}

// Write_Len is used when the value is most likely `small` (less than 1<<16).
func Write_len(w iobyte.ByteWriter, buf []byte, v int) error {
	return raw.Write_len(w, buf, v)
}

func Write_int(w iobyte.ByteWriter, buf []byte, v int) error {
	return packer.PackUint64To(w, buf, uint64(v))
}

func Write_int8(w iobyte.ByteWriter, _ []byte, v int8) error {
	return w.WriteByte(byte(v))
}

func Write_int16(w iobyte.ByteWriter, buf []byte, v int16) error {
	binary.LittleEndian.PutUint16(buf, uint16(v))
	_, err := w.Write(buf[:2])
	return err
}

func Write_int32(w iobyte.ByteWriter, buf []byte, v int32) error {
	return packer.PackUint64To(w, buf, uint64(v))
}

func Write_int64(w iobyte.ByteWriter, buf []byte, v int64) error {
	return packer.PackUint64To(w, buf, uint64(v))
}

func Write_uint(w iobyte.ByteWriter, buf []byte, v uint) error {
	return packer.PackUint64To(w, buf, uint64(v))
}

func Write_uint8(w iobyte.ByteWriter, _ []byte, v uint8) error {
	return w.WriteByte(v)
}

func Write_uint16(w iobyte.ByteWriter, buf []byte, v uint16) error {
	binary.LittleEndian.PutUint16(buf, v)
	_, err := w.Write(buf[:2])
	return err
}

func Write_uint32(w iobyte.ByteWriter, buf []byte, v uint32) error {
	return packer.PackUint64To(w, buf, uint64(v))
}

func Write_uint64(w iobyte.ByteWriter, buf []byte, v uint64) error {
	return packer.PackUint64To(w, buf, v)
}

func Write_float32(w iobyte.ByteWriter, buf []byte, v float32) error {
	u := bits.Reverse32(math.Float32bits(v))
	return packer.PackUint64To(w, buf, uint64(u))
}

func Write_float64(w iobyte.ByteWriter, buf []byte, v float64) error {
	u := bits.Reverse64(math.Float64bits(v))
	return packer.PackUint64To(w, buf, u)
}

func Write_complex64(w iobyte.ByteWriter, buf []byte, v complex64) error {
	if err := Write_float32(w, buf, real(v)); err != nil {
		return err
	}
	return Write_float32(w, buf, imag(v))
}

func Write_complex128(w iobyte.ByteWriter, buf []byte, v complex128) error {
	if err := Write_float64(w, buf, real(v)); err != nil {
		return err
	}
	return Write_float64(w, buf, imag(v))
}

func Write_string(w iobyte.ByteWriter, buf []byte, v string) error {
	return raw.Write_string(w, buf, v)
}

func Write_bytes(w iobyte.ByteWriter, buf []byte, v []byte) error {
	return raw.Write_bytes(w, buf, v)
}

func Write_bytea(w io.Writer, v []byte) error {
	return raw.Write_bytea(w, v)
}

func Write_time(w iobyte.ByteWriter, buf []byte, t time.Time) error {
	return raw.Write_time(w, buf, t)
}

func Write_bigfloat(w iobyte.ByteWriter, buf, bigbuf []byte, v big.Float) error {
	return raw.Write_bigfloat(w, buf, bigbuf, v)
}

func Write_bigint(w iobyte.ByteWriter, buf, bigbuf []byte, v big.Int) error {
	return raw.Write_bigint(w, buf, bigbuf, v)
}

func Write_bigrat(w iobyte.ByteWriter, buf, bigbuf []byte, v big.Rat) error {
	return raw.Write_bigrat(w, buf, bigbuf, v)
}
