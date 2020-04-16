package serializer

import (
	"encoding/binary"
	"io"
	"math/bits"
)

// Pack and unpack integers:
//  - first byte contains a bitmap of the non zero bytes found in the integer
//  - following bytes are the integer non zero bytes

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
	left := bits.LeadingZeros8(bitmap)
	for i := bits.OnesCount8(bitmap); i > 0; {
		x <<= 8
		if bitmap&1 > 0 {
			i--
			x |= uint64(buf[i])
		}
		bitmap >>= 1
	}
	x <<= 8 * left
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
