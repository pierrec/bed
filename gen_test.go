package serializer_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"testing"

	"github.com/pierrec/serializer"
	"github.com/pierrec/serializer/testpkg"
)

func TestGen(t *testing.T) {
	const (
		pkgheader = `package testpkg

import "io"
`
		pkgserial = pkgheader + `import "github.com/pierrec/serializer"
`
	)
	type tcase struct {
		out    string
		header string
		data   interface{}
	}
	for _, tc := range []tcase{
		{"basic_gen.go", pkgserial, testpkg.Basic{}},
		{"slice_gen.go", pkgserial, testpkg.Slice{}},
		{"array_gen.go", pkgserial, testpkg.Array{}},
		{"composite_gen.go", pkgheader, testpkg.Composite{}},
	} {
		label := fmt.Sprintf("%T", tc.data)
		t.Run(label, func(t *testing.T) {
			out, err := os.Create(path.Join("testpkg", tc.out))
			if err != nil {
				t.Fatal(err)
			}
			defer out.Close()

			if _, err := io.WriteString(out, tc.header); err != nil {
				log.Fatal(err)
			}

			if err := serializer.Gen(out, tc.data); err != nil {
				t.Fatal(err)
			}
		})
	}
}
