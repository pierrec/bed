package bed

import (
	"bytes"
	"testing"
)

func TestTemplateExec(t *testing.T) {
	type _m map[string]string
	type tcase struct {
		src  string
		data _m
		res  string
	}
	for _, tc := range []tcase{
		{"%%", _m{}, "%%"},
		{"%key%", _m{"key": "hello"}, "hello"},
		{"key%key%", _m{"key": "hello"}, "keyhello"},
		{"%key%", _m{"k": "hello"}, "%key%"},
		{"hello %key%!", _m{"key": "world"}, "hello world!"},
		{"%greet% %who%!", _m{"greet": "hello", "who": "world"}, "hello world!"},
		{"%greet% %who%? %who%!", _m{"greet": "hello", "who": "world"}, "hello world? world!"},
	} {
		t.Run(tc.src, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := templateExec(buf, tc.src, tc.data)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := buf.String(), tc.res; got != want {
				t.Fatalf("got %q; want %q", got, want)
			}
		})
	}
}
