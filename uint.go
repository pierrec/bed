package serializer

import (
	"encoding/binary"
	"io"
	"math/bits"
)

//go:generate go run uintgen.go

// Pack and unpack integers:
//  - first byte contains a bitmap of the non zero bytes found in the integer
//  - following bytes are the integer non zero bytes

type unpack64Entry struct {
	num                 uint8 // number of non zero bytes
	a, b, c, d, e, f, g uint8 // shifts
}

// packUint64 packs x into buf and returns the number of bytes used.
// buf must be at least 9 bytes long.
func packUint64(buf []byte, x uint64) int {
	_ = buf[8]
	if x == 0 {
		buf[0] = 0
		return 1
	}
	const (
		shift = 8
		max   = 1 << shift
	)
	var (
		bitmap uint8
		i      int
	)

	if x := byte(x); x > 0 {
		bitmap = 1 << 7
		i = 1
		buf[1] = x
	}
	if x < max {
		buf[0] = bitmap
		return 2
	}
	x >>= shift

	if x := byte(x); x > 0 {
		bitmap |= 1 << 6
		i++
		buf[i] = x
	}
	if x < max {
		buf[0] = bitmap
		return i + 1
	}
	x >>= shift

	if x := byte(x); x > 0 {
		bitmap |= 1 << 5
		i++
		buf[i] = x
	}
	if x < max {
		buf[0] = bitmap
		return i + 1
	}
	x >>= shift

	if x := byte(x); x > 0 {
		bitmap |= 1 << 4
		i++
		buf[i] = x
	}
	if x < max {
		buf[0] = bitmap
		return i + 1
	}
	x >>= shift

	if x := byte(x); x > 0 {
		bitmap |= 1 << 3
		i++
		buf[i] = x
	}
	if x < max {
		buf[0] = bitmap
		return i + 1
	}
	x >>= shift

	if x := byte(x); x > 0 {
		bitmap |= 1 << 2
		i++
		buf[i] = x
	}
	if x < max {
		buf[0] = bitmap
		return i + 1
	}
	x >>= shift

	if x := byte(x); x > 0 {
		bitmap |= 1 << 1
		i++
		buf[i] = x
	}
	if x < max {
		buf[0] = bitmap
		return i + 1
	}
	buf[i+1] = byte(x >> shift)
	buf[0] = bitmap | 1
	return i + 1
}

func packUint64To(w io.Writer, buf []byte, x uint64) error {
	n := packUint64(buf, x)
	_, err := w.Write(buf[:n])
	return err
}

// unpackUint64 unpacks buf and returns the value.
func unpackUint64(bitmap byte, buf []byte) uint64 {
	switch bitmap {
	case 0:
		return 0
	case 0xFF:
		return binary.LittleEndian.Uint64(buf)
	}
	entry := unpack64Table[bitmap-1]
	a, b, c, d, e, f, g := entry.a, entry.b, entry.c, entry.d, entry.e, entry.f, entry.g
	switch entry.num {
	case 1:
		return uint64(buf[0]) << a
	case 2:
		return uint64(buf[0])<<a | uint64(buf[1])<<b
	case 3:
		return uint64(buf[0])<<a | uint64(buf[1])<<b | uint64(buf[2])<<c
	case 4:
		return uint64(buf[0])<<a | uint64(buf[1])<<b | uint64(buf[2])<<c | uint64(buf[3])<<d
	case 5:
		return uint64(buf[0])<<a | uint64(buf[1])<<b | uint64(buf[2])<<c | uint64(buf[3])<<d |
			uint64(buf[4])<<e
	case 6:
		return uint64(buf[0])<<a | uint64(buf[1])<<b | uint64(buf[2])<<c | uint64(buf[3])<<d |
			uint64(buf[4])<<e | uint64(buf[5])<<f
	}
	return uint64(buf[0])<<a | uint64(buf[1])<<b | uint64(buf[2])<<c | uint64(buf[3])<<d |
		uint64(buf[4])<<e | uint64(buf[5])<<f | uint64(buf[6])<<g
}

func unpackUint64From(r ByteReader, buf []byte) (uint64, error) {
	bitmap, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	if bitmap == 0 {
		return 0, nil
	}
	n := bits.OnesCount8(bitmap)
	if _, err := io.ReadFull(r, buf[:n]); err != nil {
		return 0, err
	}
	return unpackUint64(bitmap, buf), nil
}

// packUint32 packs x into buf and returns the number of bytes used.
// buf must be at least 5 bytes long.
func packUint32(buf []byte, x uint32) int {
	_ = buf[5]
	if x == 0 {
		buf[0] = 0
		return 1
	}
	const (
		shift = 4
		max   = 1 << shift
		mask  = max - 1
	)
	var (
		acc    uint32
		i      int
		bitmap uint8
	)

	if x := x & mask; x > 0 {
		bitmap = 1 << 7
		acc = x
		i = 1
	}
	if x < max {
		buf[0] = bitmap
		buf[1] = byte(acc)
		return 2
	}
	x >>= shift

	if x := x & mask; x > 0 {
		bitmap |= 1 << 6
		acc |= x << (i * 4)
		i++
	}
	if x < max {
		buf[0] = bitmap
		buf[1] = byte(acc)
		return 2
	}
	x >>= shift

	if x := x & mask; x > 0 {
		bitmap |= 1 << 5
		acc |= x << (i * 4)
		i++
	}
	if x < max {
		buf[0] = bitmap
		buf[1] = byte(acc)
		buf[2] = byte(acc >> 8)
		return 3
	}
	x >>= shift

	if x := x & mask; x > 0 {
		bitmap |= 1 << 4
		acc |= x << (i * 4)
		i++
	}
	if x < max {
		buf[0] = bitmap
		buf[1] = byte(acc)
		buf[2] = byte(acc >> 8)
		return 3
	}
	x >>= shift

	if x := x & mask; x > 0 {
		bitmap |= 1 << 3
		acc |= x << (i * 4)
		i++
	}
	if x < max {
		buf[0] = bitmap
		buf[1] = byte(acc)
		buf[2] = byte(acc >> 8)
		buf[3] = byte(acc >> 16)
		return 4
	}
	x >>= shift

	if x := x & mask; x > 0 {
		bitmap |= 1 << 2
		acc |= x << (i * 4)
		i++
	}
	if x < max {
		buf[0] = bitmap
		buf[1] = byte(acc)
		buf[2] = byte(acc >> 8)
		buf[3] = byte(acc >> 16)
		return 4
	}
	x >>= shift

	if x := x & mask; x > 0 {
		bitmap |= 1 << 1
		acc |= x << (i * 4)
		i++
	}
	x >>= shift
	if x := x & mask; x > 0 {
		bitmap |= 1
		acc |= x << (i * 4)
	}

	buf[0] = bitmap
	buf[1] = byte(acc)
	buf[2] = byte(acc >> 8)
	buf[3] = byte(acc >> 16)
	buf[4] = byte(acc >> 24)
	return 5
}

func packUint32To(w io.Writer, buf []byte, x uint32) error {
	n := packUint32(buf, x)
	_, err := w.Write(buf[:n])
	return err
}

// unpackUint32 unpacks buf and returns the value.
func unpackUint32(bitmap byte, buf []byte) uint32 {
	switch bitmap {
	case 0:
		return 0
	case 255:
		return binary.LittleEndian.Uint32(buf)
	}
	entry := unpack64Table[bitmap-1]
	a, b, c, d, e, f, g := entry.a/2, entry.b/2, entry.c/2, entry.d/2, entry.e/2, entry.f/2, entry.g/2
	switch entry.num {
	case 1:
		return uint32(buf[0]) << a
	case 2:
		return uint32(buf[0]&0xF)<<a | uint32(buf[0]>>4)<<b
	case 3:
		return uint32(buf[0]&0xF)<<a | uint32(buf[0]>>4)<<b | uint32(buf[1])<<c
	case 4:
		return uint32(buf[0]&0xF)<<a | uint32(buf[0]>>4)<<b | uint32(buf[1]&0xF)<<c | uint32(buf[1]>>4)<<d
	case 5:
		return uint32(buf[0]&0xF)<<a | uint32(buf[0]>>4)<<b | uint32(buf[1]&0xF)<<c | uint32(buf[1]>>4)<<d |
			uint32(buf[2])<<e
	case 6:
		return uint32(buf[0]&0xF)<<a | uint32(buf[0]>>4)<<b | uint32(buf[1]&0xF)<<c | uint32(buf[1]>>4)<<d |
			uint32(buf[2]&0xF)<<e | uint32(buf[2]>>4)<<f
	}
	return uint32(buf[0]&0xF)<<a | uint32(buf[0]>>4)<<b | uint32(buf[1]&0xF)<<c | uint32(buf[1]>>4)<<d |
		uint32(buf[2]&0xF)<<e | uint32(buf[2]>>4)<<f | uint32(buf[3])<<g
}

func unpackUint32From(r ByteReader, buf []byte) (uint32, error) {
	bitmap, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	if bitmap == 0 {
		return 0, nil
	}
	n := bits.OnesCount8(bitmap)
	if _, err := io.ReadFull(r, buf[:n]); err != nil {
		return 0, err
	}
	return unpackUint32(bitmap, buf), nil
}
