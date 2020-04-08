package serializer

import (
	"encoding/binary"
	"io"
	"math/bits"
)

// Pack and unpack integers:
//  - first byte contains the number of non zero bytes of the original integer
//  - first byte packs the last byte if it fits
//  - following bytes are the integer bytes up to the first non leading zero

func packSize(n int) uint8 {
	if n == 0 {
		return 1
	}
	return uint8(n-1) << 5
}

func unpackSize(c byte) int {
	return 1 + int(c)>>5
}

func packIsSmall(c byte) bool {
	return c < 1<<5
}

func unpackSmall(c byte) (byte, bool) {
	r := c & (1<<5 - 1)
	return r, r > 0
}

// packUint64 packs x into buf and returns the number of bytes used.
// buf must be at least 9 bytes long.
func packUint64(buf []byte, x uint64) int {
	_ = buf[:9]
	buf[0] = 0
	b := buf[1:]
	binary.LittleEndian.PutUint64(b, x)

	n := 1
	if size := 8 - bits.LeadingZeros64(x)/8; size > 0 {
		n = size
		if last := uint8(x >> (8 * (n - 1))); packIsSmall(last) {
			// Pack the last byte into the header.
			buf[0] = last
		}
	}
	buf[0] |= packSize(n)
	return n + 1
}

func packUint64To(w io.Writer, buf []byte, x uint64) error {
	n := packUint64(buf, x)
	_, err := w.Write(buf[:n])
	return err
}

// unpackUint64 unpacks buf and returns the value.
func unpackUint64(buf []byte) (x uint64) {
	h := buf[0]
	b := buf[1:]
	size := unpackSize(h)
	if last, ok := unpackSmall(h); ok {
		size--
		x |= uint64(last) << (8 * size)
	}
	for i := 0; i < size; i++ {
		x |= uint64(b[i]) << (8 * i)
	}
	return
}

func unpackUint64From(r io.Reader, buf []byte) (uint64, error) {
	if br, ok := r.(io.ByteReader); ok {
		b, err := br.ReadByte()
		if err != nil {
			return 0, err
		}
		buf[0] = b
	} else if _, err := io.ReadFull(r, buf[:1]); err != nil {
		return 0, err
	}
	h := buf[0]
	b := buf[1:]
	size := unpackSize(h)
	if _, err := io.ReadFull(r, b[:size]); err != nil {
		return 0, err
	}
	var x uint64
	if last, ok := unpackSmall(h); ok {
		size--
		x |= uint64(last) << (8 * size)
	}
	for i := 0; i < size; i++ {
		x |= uint64(b[i]) << (8 * i)
	}
	return x, nil
}
