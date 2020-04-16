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
		0xF0F0F0F0,
		0xF0F0F0F0F0F0F0F0,
		math.Float64bits(math.Pi),
		math.Float64bits(math.Phi),
		math.Float64bits(math.E),
	} {
		label := fmt.Sprintf("%d", tc)
		t.Run(label, func(t *testing.T) {
			buf := make([]byte, 16)
			n := packUint64(buf, tc)
			t.Logf("packed size for %d = %d", tc, n)
			x := unpackUint64(buf[0], buf[1:])
			if got, want := x, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

var benchInt int
var benchX uint64

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

func BenchmarkPackUint64_zebra2(b *testing.B) {
	benchmarkPackUint64(b, 0xF0F0F0F0F0F0F0F0)
}

func benchmarkUnpackUint64(b *testing.B, x uint64) {
	buf := make([]byte, 16)
	packUint64(buf, x)
	bitmap := buf[0]
	buf = buf[1:]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchX = unpackUint64(bitmap, buf)
	}
}

func BenchmarkUnpackUint64_E(b *testing.B) {
	benchmarkUnpackUint64(b, math.Float64bits(math.E))
}

func BenchmarkUnpackUint64_1024(b *testing.B) {
	benchmarkUnpackUint64(b, 1024)
}

func BenchmarkUnpackUint64_zebra2(b *testing.B) {
	benchmarkUnpackUint64(b, 0xF0F0)
}

func BenchmarkUnpackUint64_zebra4(b *testing.B) {
	benchmarkUnpackUint64(b, 0xF0F0F0F0)
}

func BenchmarkUnpackUint64_zebra6(b *testing.B) {
	benchmarkUnpackUint64(b, 0xF0F0F0F0F0F0F0)
}

func BenchmarkUnpackUint64_zebra8(b *testing.B) {
	benchmarkUnpackUint64(b, 0xF0F0F0F0F0F0F0F0)
}
