package serializer

import (
	"fmt"
	"math"
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
		math.Float64bits(math.Pi),
		math.Float64bits(math.Phi),
		math.Float64bits(math.E),
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

var benchInt int

func benchmarkPackUint64(b *testing.B, x uint64) {
	buf := make([]byte, 16)
	for i := 0; i < b.N; i++ {
		benchInt = packUint64(buf, x)
	}
}

func BenchmarkPackUint64_E(b *testing.B) {
	benchmarkPackUint64(b, math.Float64bits(math.E))
}

func BenchmarkPackUint64_1024(b *testing.B) {
	benchmarkPackUint64(b, 1024)
}

func BenchmarkPackUint64_zebra(b *testing.B) {
	benchmarkPackUint64(b, 0xF0F0F0F0)
}
