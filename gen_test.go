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
	config := serializer.Config{PkgName: "testpkg", Receiver: "self"}
	type tcase struct {
		out  string
		data interface{}
	}
	for _, tc := range []tcase{
		{"basic_gen.go", testpkg.Basic{}},
		{"slice_gen.go", testpkg.Slice{}},
		{"array_gen.go", testpkg.Array{}},
		{"compositeonly_gen.go", testpkg.CompositeOnly{}},
		{"composite_gen.go", testpkg.Composite{}},
		{"map_gen.go", testpkg.Map{}},
	} {
		label := fmt.Sprintf("%T", tc.data)
		t.Run(label, func(t *testing.T) {
			out, err := os.Create(path.Join("testpkg", tc.out))
			if err != nil {
				t.Fatal(err)
			}
			defer out.Close()

			if err := serializer.Gen(out, config, tc.data); err != nil {
				t.Fatal(err)
			}
		})
	}
}
