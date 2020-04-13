package serializer

import (
	"encoding/binary"
	"io"
	"math"
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

func Read_bytes(r io.Reader, buf []byte) ([]byte, error) {
	n, err := Read_int(r, buf)
	if err != nil || n == 0 {
		return nil, err
	}
	buf = make([]byte, n)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func Read_bytea(r io.Reader, buf []byte) error {
	_, err := io.ReadFull(r, buf)
	return err
}

func Read_time(r io.Reader, buf []byte) (t time.Time, err error) {
	_ = buf[:12]
	if _, err = io.ReadFull(r, buf[:2]); err != nil {
		return
	}
	year := binary.LittleEndian.Uint16(buf)
	if year == 0 {
		return
	}
	if _, err = io.ReadFull(r, buf[:10]); err != nil {
		return
	}

	month, day := buf[0], buf[1]
	hour, min, sec := buf[2], buf[3], buf[4]
	offset := buf[5]

	ns := binary.LittleEndian.Uint32(buf[6:])

	loc := time.FixedZone("", int(offset)*(60*60))

	t = time.Date(int(year), time.Month(month), int(day), int(hour), int(min), int(sec), int(ns), loc)
	return
}
