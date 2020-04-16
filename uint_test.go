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
			t.Logf("packed size for %d = %d", tc, n)
			x := unpackUint64(buf[0], buf[1:n])
			if got, want := x, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}
