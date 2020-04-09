package testpkg

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pierrec/serializer"
)

func TestGen(t *testing.T) {
	_if := func(v ...interface{}) []interface{} { return v }
	type tcase struct {
		out  string
		data []interface{}
	}
	for _, tc := range []tcase{
		{"basic_gen.go", _if(Basic{})},
		{"slice_gen.go", _if(Slice{})},
		{"array_gen.go", _if(Array{})},
		{"composite_gen.go", _if(CompositeOnly{}, Composite{})},
	} {
		label := fmt.Sprintf("%T", tc.data)
		t.Run(label, func(t *testing.T) {

			typ := reflect.TypeOf(tc.data[0])
			from := reflect.New(typ).Interface().(serializer.Interface)
			var buf bytes.Buffer
			if err := from.MarshalBinaryTo(&buf); err != nil {
				t.Fatal(err)
			}
			into := reflect.New(typ).Interface().(serializer.Interface)
			if err := into.UnmarshalBinaryFrom(&buf); err != nil {
				t.Fatal(err)
			}
			if got, want := into, from; !cmp.Equal(got, want) {
				t.Fatalf("diff\n%v", cmp.Diff(got, want))
			}
		})
	}
}
