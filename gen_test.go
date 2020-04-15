package serializer_test

import (
	"os"
	"path"
	"strings"
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
		{"basic_gen.go", _s(testpkg.Basic{}, testpkg.BasicPtr{}, testpkg.BasicEmbed{}, testpkg.BasicAnon{})},
		{"slice_gen.go", _s(testpkg.Slice{}, testpkg.SlicePtr{}, testpkg.SliceAnon{})},
		{"array_gen.go", _s(testpkg.Array{}, testpkg.ArrayPtr{}, testpkg.ArrayAnon{})},
		{"map_gen.go", _s(testpkg.Map{})},
		{"time_gen.go", _s(testpkg.Time{}, testpkg.TimePtr{}, testpkg.TimeSlice{}, testpkg.TimeArray{}, testpkg.TimeMap{})},
		{"big_gen.go", _s(testpkg.Big{}, testpkg.BigPtr{}, testpkg.BigPtrSlice{}, testpkg.BigSlice{}, testpkg.BigArray{}, testpkg.BigMap{})},
	} {
		label := strings.TrimSuffix(tc.out, "_gen.go")
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
