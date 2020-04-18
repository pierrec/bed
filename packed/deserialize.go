package packed

import (
	"io"
	"math"
	"math/big"
	"math/bits"
	"strings"
	"time"

	"github.com/pierrec/bed"
	"github.com/pierrec/bed/raw"
	"github.com/pierrec/packer"
	"github.com/pierrec/packer/iobyte"
)

func Read_layout(r iobyte.ByteReader, buf []byte, layout string) error {
	s, err := Read_string(r, buf)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(s, layout) {
		return bed.ErrInvalidData
	}
	return nil
}

func Read_bool(r iobyte.ByteReader, _ []byte) (bool, error) {
	return raw.Read_bool(r, nil)
}

func Read_len(r iobyte.ByteReader) (int, error) {
	return raw.Read_len(r)
}

func Read_int(r iobyte.ByteReader, buf []byte) (int, error) {
	v, err := packer.UnpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func Read_int8(r iobyte.ByteReader, _ []byte) (int8, error) {
	u, err := r.ReadByte()
	return int8(u), err
}

func Read_int16(r iobyte.ByteReader, buf []byte) (int16, error) {
	return raw.Read_int16(r, buf)
}

func Read_int32(r iobyte.ByteReader, buf []byte) (int32, error) {
	v, err := packer.UnpackUint32From(r, buf)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func Read_int64(r iobyte.ByteReader, buf []byte) (int64, error) {
	v, err := packer.UnpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func Read_uint(r iobyte.ByteReader, buf []byte) (uint, error) {
	v, err := packer.UnpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func Read_uint8(r iobyte.ByteReader, _ []byte) (uint8, error) {
	return r.ReadByte()
}

func Read_uint16(r iobyte.ByteReader, buf []byte) (uint16, error) {
	return raw.Read_uint16(r, buf)
}

func Read_uint32(r iobyte.ByteReader, buf []byte) (uint32, error) {
	return packer.UnpackUint32From(r, buf)
}

func Read_uint64(r iobyte.ByteReader, buf []byte) (uint64, error) {
	return packer.UnpackUint64From(r, buf)
}

func Read_float32(r iobyte.ByteReader, buf []byte) (float32, error) {
	v, err := packer.UnpackUint32From(r, buf)
	if err != nil {
		return 0, err
	}
	u := bits.Reverse32(v)
	return math.Float32frombits(u), err
}

func Read_float64(r iobyte.ByteReader, buf []byte) (float64, error) {
	v, err := packer.UnpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	v = bits.Reverse64(v)
	return math.Float64frombits(v), err
}

func Read_complex64(r iobyte.ByteReader, buf []byte) (complex64, error) {
	re, err := Read_float32(r, buf)
	if err != nil {
		return 0, err
	}
	im, err := Read_float32(r, buf)
	return complex(re, im), err
}

func Read_complex128(r iobyte.ByteReader, buf []byte) (complex128, error) {
	re, err := Read_float64(r, buf)
	if err != nil {
		return 0, err
	}
	im, err := Read_float64(r, buf)
	return complex(re, im), err
}

func Read_string(r iobyte.ByteReader, buf []byte) (string, error) {
	n, err := Read_int(r, buf)
	if err != nil || n == 0 {
		return "", err
	}
	if n > cap(buf) {
		buf = make([]byte, n)
	} else {
		buf = buf[:n]
	}
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func Read_bytes(r iobyte.ByteReader, buf, out []byte) ([]byte, error) {
	n, err := Read_int(r, buf)
	if err != nil || n == 0 {
		return nil, err
	}
	if len(out) < n {
		out = make([]byte, n)
	} else {
		out = out[:n]
	}
	_, err = io.ReadFull(r, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func Read_bytea(r iobyte.ByteReader, buf []byte) error {
	_, err := io.ReadFull(r, buf)
	return err
}

func Read_time(r iobyte.ByteReader, buf []byte) (t time.Time, err error) {
	return raw.Read_time(r, buf)
}

func Read_bigfloat(r iobyte.ByteReader, buf, bigbuf []byte) (b big.Float, err error) {
	return raw.Read_bigfloat(r, buf, bigbuf)
}

func Read_bigint(r iobyte.ByteReader, buf, bigbuf []byte) (b big.Int, err error) {
	return raw.Read_bigint(r, buf, bigbuf)
}

func Read_bigrat(r iobyte.ByteReader, buf, bigbuf []byte) (b big.Rat, err error) {
	return raw.Read_bigrat(r, buf, bigbuf)
}
