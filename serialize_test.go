package serializer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWalkData(t *testing.T) {
	for _, tc := range []interface{}{
		struct {
			S string
			D []int32
			B []byte
		}{S: "hello"},
	} {
		t.Run("", func(t *testing.T) {
			records, err := walkData("", tc)
			if err != nil {
				t.Fatal(err)
			}
			buf := new(bytes.Buffer)
			if err := genMarshalBinTo(buf, records, "s", tc); err != nil {
				t.Fatal(err)
			}
			fmt.Println(">>>", buf.String())
			buf.Reset()
			if err := genUnmarshalBinFrom(buf, records, "s", tc); err != nil {
				t.Fatal(err)
			}
			fmt.Println(">>>", buf.String())
		})
	}
}
