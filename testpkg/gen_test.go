package testpkg

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pierrec/serializer"
)

func TestGen(t *testing.T) {
	var (
		_s       = func(v ...int) []int { return v }
		one, two = 1, 2
		basic    = &Basic{
			Int:        -100,
			Int8:       -80,
			Int16:      -1600,
			Int32:      -3200,
			Int64:      -6400,
			Uint:       100,
			Uint8:      80,
			Uint16:     1600,
			Uint32:     3200,
			Uint64:     6400,
			Complex64:  64,
			Complex128: 128,
			String:     "hi",
		}
		slice = &Slice{
			Int:        []int{-100, -200, -300, -400},
			Int8:       []int8{-81, -82, -83, -84},
			Int16:      []int16{-1601, -1602, -1603, -1604},
			Int32:      []int32{-3201, -3202, -3203, -3204},
			Int64:      []int64{-6401, -6402, -6403, -6404},
			Uint:       []uint{100, 200, 300, 400},
			Uint8:      []uint8{81, 82, 83, 84},
			Uint16:     []uint16{1601, 1602, 1603, 1604},
			Uint32:     []uint32{3201, 3202, 3203, 3204},
			Uint64:     []uint64{6401, 6402, 6403, 6404},
			Complex64:  []complex64{1, 2, 3, 4},
			Complex128: []complex128{11, 22, 33, 44},
			String:     []string{"one", "two", "three", "four"},
		}
		array = &Array{
			Int:        [4]int{-100, -200, -300, -400},
			Int8:       [4]int8{-81, -82, -83, -84},
			Int16:      [4]int16{-1601, -1602, -1603, -1604},
			Int32:      [4]int32{-3201, -3202, -3203, -3204},
			Int64:      [4]int64{-6401, -6402, -6403, -6404},
			Uint:       [4]uint{100, 200, 300, 400},
			Uint8:      [4]uint8{81, 82, 83, 84},
			Uint16:     [4]uint16{1601, 1602, 1603, 1604},
			Uint32:     [4]uint32{3201, 3202, 3203, 3204},
			Uint64:     [4]uint64{6401, 6402, 6403, 6404},
			Complex64:  [4]complex64{1, 2, 3, 4},
			Complex128: [4]complex128{11, 22, 33, 44},
			String:     [4]string{"one", "two", "three", "four"},
		}
		mapp = &Map{
			StringInt:  map[string]int{"a": 1, "b": 2},
			StringInts: map[string][]int{"a": _s(1, 11), "b": _s(2, 22)},
			IntPtrInt:  map[*int]int{&one: 11, &two: 22},
		}
		cmpUintPointers = cmpopts.SortMaps(func(x, y *uint) bool { return *x < *y })
		cmpIntPointers  = cmpopts.SortMaps(func(x, y *int) bool { return *x < *y })
	)
	for _, tc := range []interface{}{
		basic, slice, array,
		&CompositeOnly{
			Basic: *basic,
			Slice: *slice,
			Array: *array,
			Map:   *mapp,
		},
		&Composite{
			Bytes: []byte("I am a slice of bytes"),
			Basic: *basic,
			Slice: *slice,
			Array: *array,
		},
	} {
		label := fmt.Sprintf("%T", tc)
		t.Run(label, func(t *testing.T) {
			from := tc.(serializer.Interface)
			var buf bytes.Buffer
			if err := from.MarshalBinaryTo(&buf); err != nil {
				t.Fatal(err)
			}
			typ := reflect.TypeOf(tc).Elem()
			into := reflect.New(typ).Interface().(serializer.Interface)
			if err := into.UnmarshalBinaryFrom(&buf); err != nil {
				t.Fatal(err)
			}
			if got, want := into, from; !cmp.Equal(got, want, cmpIntPointers, cmpUintPointers) {
				t.Fatalf("diff\n%v", cmp.Diff(got, want))
			}
		})
	}
}
