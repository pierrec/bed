package serializer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead_bool(t *testing.T) {
	for _, tc := range []bool{false, true} {
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			var buf [16]byte
			b := buf[:]
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
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			var buf [16]byte
			b := buf[:]
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
		label := fmt.Sprintf("%v", tc)
		t.Run(label, func(t *testing.T) {
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
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
			var buf [16]byte
			b := buf[:]
			rw := new(bytes.Buffer)

			err := Write_bytes(rw, b, []byte(tc))
			if err != nil {
				t.Error(err)
			}
			v, err := Read_bytes(rw, b)
			if err != nil {
				t.Error(err)
			}
			if got, want := string(v), tc; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}
