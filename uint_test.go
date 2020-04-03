package serializer

import (
	"fmt"
	"testing"
)

func TestPackUint(t *testing.T) {
	for _, tc := range []uint64{
		0,
		1,
		127,
		128,
		256,
		1 << 10,
		1 << 20,
		1<<20 | 1<<10,
		1 << 63,
	} {
		label := fmt.Sprintf("%d", tc)
		t.Run(label, func(t *testing.T) {
			buf := make([]byte, 16)
			n := packUint64(buf, tc)
			x := unpackUint64(buf[:n])
			if got, want := x, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}
