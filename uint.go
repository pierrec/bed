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
	num                 int // number of non zero bytes
	a, b, c, d, e, f, g int // shifts
}

// packUint64 packs x into buf and returns the number of bytes used.
// buf must be at least 9 bytes long.
func packUint64(buf []byte, x uint64) int {
	switch x {
	case 0:
		buf[0] = 0
		return 1
	case ^uint64(0):
		buf[0] = 0xF
		binary.BigEndian.PutUint64(buf[1:], x)
		return 9
	}
	left := bits.LeadingZeros64(x)
	var bitmap uint8
	b := buf[1:1]
	for x > 0 {
		bitmap <<= 1
		if x := byte(x); x > 0 {
			bitmap |= 1
			b = append(b, x)
		}
		x >>= 8
	}
	buf[0] = bitmap << (left / 8)

	return len(b) + 1
}

func packUint64To(w io.Writer, buf []byte, x uint64) error {
	n := packUint64(buf, x)
	_, err := w.Write(buf[:n])
	return err
}

// unpackUint64 unpacks buf and returns the value.
func unpackUint64(bitmap byte, buf []byte) (x uint64) {
	entry := unpack64Table[bitmap]
	switch entry.num {
	case 1:
		x = uint64(buf[0]) << entry.a
	case 2:
		x = uint64(buf[0])<<entry.a | uint64(buf[1])<<entry.b
	case 3:
		x = uint64(buf[0])<<entry.a | uint64(buf[1])<<entry.b | uint64(buf[2])<<entry.c
	case 4:
		x = uint64(buf[0])<<entry.a | uint64(buf[1])<<entry.b | uint64(buf[2])<<entry.c | uint64(buf[3])<<entry.d
	case 5:
		x = uint64(buf[0])<<entry.a | uint64(buf[1])<<entry.b | uint64(buf[2])<<entry.c | uint64(buf[3])<<entry.d |
			uint64(buf[4])<<entry.e
	case 6:
		x = uint64(buf[0])<<entry.a | uint64(buf[1])<<entry.b | uint64(buf[2])<<entry.c | uint64(buf[3])<<entry.d |
			uint64(buf[4])<<entry.e | uint64(buf[5])<<entry.f
	case 7:
		x = uint64(buf[0])<<entry.a | uint64(buf[1])<<entry.b | uint64(buf[2])<<entry.c | uint64(buf[3])<<entry.d |
			uint64(buf[4])<<entry.e | uint64(buf[5])<<entry.f | uint64(buf[6])<<entry.g
	case 8:
		x = binary.LittleEndian.Uint64(buf)
	}
	return
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
	_ = buf[:5]

	var bitmap uint8
	b := buf[1:1]
	var c byte
	for x := x; x > 0; x >>= 4 {
		bitmap <<= 1
		if q := byte(x & 0xF); q > 0 {
			bitmap |= 1
			if c == 0 {
				c = q
			} else {
				b = append(b, c|(q<<4))
				c = 0
			}
		}
	}
	buf[0] = bits.Reverse8(bitmap)

	return len(b) + 1
}

func packUint32To(w io.Writer, buf []byte, x uint32) error {
	n := packUint32(buf, x)
	_, err := w.Write(buf[:n])
	return err
}
