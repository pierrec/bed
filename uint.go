package serializer

import (
	"io"
	"math/bits"
)

// Pack and unpack integers:
//  - first byte contains a bitmap of the non zero bytes found in the integer
//  - following bytes are the integer non zero bytes

// packUint64 packs x into buf and returns the number of bytes used.
// buf must be at least 9 bytes long.
func packUint64(buf []byte, x uint64) int {
	_ = buf[:9]

	if x == 0 {
		buf[0] = 0
		return 1
	}
	x = bits.ReverseBytes64(x)
	var bitmap uint8
	b := buf[1:1]
	for i := 0; i < 8; i++ {
		bitmap <<= 1
		if x&0xFF > 0 {
			bitmap |= 1
			b = append(b, byte(x))
		}
		x >>= 8
	}
	buf[0] = bits.Reverse8(bitmap)

	return len(b) + 1
}

func packUint64To(w io.Writer, buf []byte, x uint64) error {
	n := packUint64(buf, x)
	_, err := w.Write(buf[:n])
	return err
}

// unpackUint64 unpacks buf and returns the value.
func unpackUint64(bitmap byte, buf []byte) (x uint64) {
	if bitmap == 0 {
		return
	}
	for i := 0; i < 8; i++ {
		x <<= 8
		if bitmap&1 > 0 {
			x |= uint64(buf[0])
			buf = buf[1:]
		}
		bitmap >>= 1
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
