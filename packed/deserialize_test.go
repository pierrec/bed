package packed

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pierrec/bed"
)

func TestRead_bool(t *testing.T) {
	for _, tc := range []bool{false, true} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_bool(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_bool(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestRead_int(t *testing.T) {
	for _, tc := range []int{0, 1, 10, 128, 256, 1014, 1 << 10, 1 << 20} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_int(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_int(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_int8(t *testing.T) {
	for _, tc := range []int8{0, 1, 10, 127, -1, -10, -127} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_int8(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_int8(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_int16(t *testing.T) {
	for _, tc := range []int16{0, 1, 10, 127, 1 << 10, -1, -10, -127, -1 << 10} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_int16(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_int16(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_int32(t *testing.T) {
	for _, tc := range []int32{0, 1, 10, 127, 1 << 10, -1, -10, -127, -1 << 10} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_int32(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_int32(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_int64(t *testing.T) {
	for _, tc := range []int64{0, 1, 10, 127, 1 << 10, -1, -10, -127, -1 << 10} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_int64(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_int64(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_uint(t *testing.T) {
	for _, tc := range []uint{0, 1, 10, 128, 256, 1014, 1 << 10, 1 << 20} {
		label := fmt.Sprintf("%x", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_uint(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_uint(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_uint8(t *testing.T) {
	for _, tc := range []uint8{0, 1, 10, 127, 255} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_uint8(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_uint8(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_uint16(t *testing.T) {
	for _, tc := range []uint16{0, 1, 10, 127, 1 << 10} {
		label := fmt.Sprintf("%x", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_uint16(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_uint16(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_uint32(t *testing.T) {
	for _, tc := range []uint32{0, 1, 10, 127, 1 << 10, 1 << 20, 1 << 30} {
		label := fmt.Sprintf("%x", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_uint32(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_uint32(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_uint64(t *testing.T) {
	for _, tc := range []uint64{0, 1, 10, 127, 1 << 10, 1 << 20, 1 << 30, 1 << 60} {
		label := fmt.Sprintf("%x", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_uint64(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_uint64(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %d; want %d", got, want)
			}
		})
	}
}

func TestRead_float32(t *testing.T) {
	for _, tc := range []float32{0, 1, 10, 127, 1 << 10, 1 << 20, 1 << 30} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_float32(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_float32(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestRead_float64(t *testing.T) {
	for _, tc := range []float64{0, 1, 10, 127, 1 << 10, 1 << 20, 1 << 30} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_float64(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_float64(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %f; want %f", got, want)
			}
		})
	}
}

func TestRead_complex64(t *testing.T) {
	for _, tc := range []complex64{0, 1, 10, 127, 1 << 10, 1 << 20, 1 << 30} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_complex64(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_complex64(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; !cmp.Equal(got, want) {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestRead_complex128(t *testing.T) {
	for _, tc := range []complex128{0, 1, 10, 127, 1 << 10, 1 << 20, 1 << 30} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_complex128(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_complex128(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; !cmp.Equal(got, want) {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestRead_string(t *testing.T) {
	for _, tc := range []string{"0", "1", "10", "127", "abbbbbbbbbbbbbb", "abcdefghijklmnopqrstuvwxyz"} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_string(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_string(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

func TestRead_bytes(t *testing.T) {
	for _, tc := range []string{"0", "1", "10", "127", "abbbbbbbbbbbbbb", "abcdefghijklmnopqrstuvwxyz"} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_bytes(rw, b, []byte(tc))
			if err != nil {
				t.Error(err)
			}
			v, err := Read_bytes(rw, b, nil)
			if err != nil {
				t.Error(err)
			}
			if got, want := string(v), tc; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

func TestRead_time(t *testing.T) {
	for _, tc := range []time.Time{{}, time.Now(), time.Now().Local(), time.Now().UTC()} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			rw := new(bytes.Buffer)

			err := Write_time(rw, b, tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_time(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; !got.Equal(want) {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestRead_bigfloat(t *testing.T) {
	for _, tc := range []*big.Float{
		big.NewFloat(0),
		big.NewFloat(1.0),
		big.NewFloat(math.Pi),
		big.NewFloat(math.Phi),
	} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			bb := bed.BigBuffers.Get()
			defer bed.Buffers.Put(bb)
			rw := new(bytes.Buffer)

			err := Write_bigfloat(rw, b, bb, *tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_bigfloat(rw, b, bb)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got.Cmp(want) != 0 {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestRead_bigint(t *testing.T) {
	for _, tc := range []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(123),
		big.NewInt(1 << 10),
		big.NewInt(1 << 20),
		big.NewInt(1 << 30),
		big.NewInt(-1),
		big.NewInt(-123),
		big.NewInt(-1 << 10),
		big.NewInt(-1 << 20),
		big.NewInt(-1 << 30),
	} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			bb := bed.BigBuffers.Get()
			defer bed.Buffers.Put(bb)
			rw := new(bytes.Buffer)

			err := Write_bigint(rw, b, bb, *tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_bigint(rw, b, bb)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got.Cmp(want) != 0 {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}

func TestRead_bigrat(t *testing.T) {
	for _, tc := range []*big.Rat{
		big.NewRat(0, 1),
		big.NewRat(1, 2),
		big.NewRat(123, 456),
		big.NewRat(1<<10, 42),
		big.NewRat(1<<20, 42),
		big.NewRat(1<<30, 42),
		big.NewRat(-1, 2),
		big.NewRat(-123, 456),
		big.NewRat(-1<<10, 42),
		big.NewRat(-1<<20, 42),
		big.NewRat(-1<<30, 42),
	} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			b := bed.Buffers.Get()
			defer bed.Buffers.Put(b)
			bb := bed.BigBuffers.Get()
			defer bed.Buffers.Put(bb)
			rw := new(bytes.Buffer)

			err := Write_bigrat(rw, b, bb, *tc)
			if err != nil {
				t.Error(err)
			}
			v, err := Read_bigrat(rw, b, bb)
			if err != nil {
				t.Error(err)
			}
			if got, want := v, tc; got.Cmp(want) != 0 {
				t.Errorf("got %v; want %v", got, want)
			}
		})
	}
}
