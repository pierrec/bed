package serializer_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/pierrec/serializer"
	"github.com/pierrec/serializer/testpkg"
)

func TestGen(t *testing.T) {
	_s := func(v ...interface{}) []interface{} { return v }
	config := serializer.Config{PkgName: "testpkg", Receiver: ""}
	type tcase struct {
		out  string
		data []interface{}
	}
	for _, tc := range []tcase{
		//{"basic_gen.go", _s(testpkg.Basic{}, testpkg.BasicPtr{}, testpkg.BasicEmbed{}, testpkg.BasicAnon{})},
		//{"slice_gen.go", _s(testpkg.Slice{}, testpkg.SlicePtr{}, testpkg.SliceAnon{})},
		//{"array_gen.go", _s(testpkg.Array{}, testpkg.ArrayPtr{})},
		//{"map_gen.go", _s(testpkg.Map{})},
		//{"compositeonly_gen.go", _s(testpkg.CompositeOnly{})},
		//{"composite_gen.go", _s(testpkg.Composite{})},
		{"misc_gen.go", _s(testpkg.Misc{})},
	} {
		label := fmt.Sprintf("%T", tc.data)
		t.Run(label, func(t *testing.T) {
			out, err := os.Create(path.Join("testpkg", tc.out))
			if err != nil {
				t.Fatal(err)
			}
			defer out.Close()

			if err := serializer.Gen(out, config, tc.data...); err != nil {
				t.Fatal(err)
			}
		})
	}
}
