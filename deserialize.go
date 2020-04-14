package serializer

import (
	"encoding/binary"
	"io"
	"math"
	"math/big"
	"math/bits"
	"strings"
	"time"
)

func Read_layout(r io.Reader, buf []byte, layout string) error {
	s, err := Read_string(r, buf)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(s, layout) {
		return ErrInvalidData
	}
	return nil
}

func Read_bool(r io.Reader, buf []byte) (bool, error) {
	if _, err := io.ReadFull(r, buf[:1]); err != nil {
		return false, err
	}
	return buf[0] != 0, nil
}

func Read_int(r io.Reader, buf []byte) (int, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func Read_int8(r io.Reader, buf []byte) (int8, error) {
	if _, err := io.ReadFull(r, buf[:1]); err != nil {
		return 0, err
	}
	return int8(buf[0]), nil
}

func Read_int16(r io.Reader, buf []byte) (int16, error) {
	if _, err := io.ReadFull(r, buf[:2]); err != nil {
		return 0, err
	}
	v := binary.LittleEndian.Uint16(buf)
	return int16(v), nil
}

func Read_int32(r io.Reader, buf []byte) (int32, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func Read_int64(r io.Reader, buf []byte) (int64, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func Read_uint(r io.Reader, buf []byte) (uint, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func Read_uint8(r io.Reader, buf []byte) (uint8, error) {
	if _, err := io.ReadFull(r, buf[:1]); err != nil {
		return 0, err
	}
	return buf[0], nil
}

func Read_uint16(r io.Reader, buf []byte) (uint16, error) {
	if _, err := io.ReadFull(r, buf[:2]); err != nil {
		return 0, err
	}
	v := binary.LittleEndian.Uint16(buf)
	return v, nil
}

func Read_uint32(r io.Reader, buf []byte) (uint32, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func Read_uint64(r io.Reader, buf []byte) (uint64, error) {
	return unpackUint64From(r, buf)
}

func Read_float32(r io.Reader, buf []byte) (float32, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	u := bits.Reverse32(uint32(v))
	return math.Float32frombits(u), err
}

func Read_float64(r io.Reader, buf []byte) (float64, error) {
	v, err := unpackUint64From(r, buf)
	if err != nil {
		return 0, err
	}
	v = bits.Reverse64(v)
	return math.Float64frombits(v), err
}

func Read_complex64(r io.Reader, buf []byte) (complex64, error) {
	re, err := Read_float32(r, buf)
	if err != nil {
		return 0, err
	}
	im, err := Read_float32(r, buf)
	return complex(re, im), err
}

func Read_complex128(r io.Reader, buf []byte) (complex128, error) {
	re, err := Read_float64(r, buf)
	if err != nil {
		return 0, err
	}
	im, err := Read_float64(r, buf)
	return complex(re, im), err
}

func Read_string(r io.Reader, buf []byte) (string, error) {
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

func Read_bytes(r io.Reader, buf, out []byte) ([]byte, error) {
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

func Read_bytea(r io.Reader, buf []byte) error {
	_, err := io.ReadFull(r, buf)
	return err
}

func Read_time(r io.Reader, buf []byte) (t time.Time, err error) {
	_ = buf[:10]
	if _, err = io.ReadFull(r, buf[:2]); err != nil {
		return
	}
	year := binary.LittleEndian.Uint16(buf)
	if year == 0 {
		return
	}
	if _, err = io.ReadFull(r, buf[:4]); err != nil {
		return
	}

	u := binary.LittleEndian.Uint32(buf)

	var ns int
	if u&1 > 0 {
		if _, err = io.ReadFull(r, buf[4:8]); err != nil {
			return
		}
		ns = int(binary.LittleEndian.Uint32(buf[4:]))
	}

	const (
		fiveMask = 1<<5 - 1
		sixMask  = 1<<6 - 1
	)

	u >>= 1
	offset := u & fiveMask
	u >>= 5
	sec := u & sixMask
	u >>= 6
	min := u & sixMask
	u >>= 6
	hour := u & fiveMask
	u >>= 5
	day := u & fiveMask
	u >>= 5
	month := u

	loc := time.FixedZone("", int(offset)*(60*60))

	t = time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), ns, loc)

	return
}

func Read_bigfloat(r io.Reader, buf, bigbuf []byte) (b big.Float, err error) {
	bigbuf, err = Read_bytes(r, buf, bigbuf)
	if err == nil {
		err = b.UnmarshalText(bigbuf)
	}
	return
}

func Read_bigint(r io.Reader, buf, bigbuf []byte) (b big.Int, err error) {
	sign, err := Read_uint8(r, buf)
	if err != nil || sign == 1 {
		return
	}
	bigbuf, err = Read_bytes(r, buf, bigbuf)
	if err != nil {
		return
	}
	b.SetBytes(bigbuf)
	if sign == 0 {
		b.Neg(&b)
	}
	return
}

func Read_bigrat(r io.Reader, buf, bigbuf []byte) (b big.Rat, err error) {
	num, err := Read_bigint(r, buf, bigbuf)
	if err != nil {
		return
	}
	denom, err := Read_bigint(r, buf, bigbuf)
	if err != nil {
		return
	}
	b.SetFrac(&num, &denom)
	return
}
