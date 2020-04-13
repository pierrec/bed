package testpkg

import (
	"io"
	"time"

	"github.com/pierrec/serializer"
)

var _ time.Time

const _CompositeOnlyLayout = "ZZZZ"

func (c *CompositeOnly) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _CompositeOnlyLayout)
	if err != nil {
		return
	}

	err = c.Basic.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	err = c.Slice.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	err = c.Array.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	err = c.Map.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	return
}

func (c *CompositeOnly) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _CompositeOnlyLayout)
	if err != nil {
		return
	}

	err = c.Basic.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = c.Slice.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = c.Array.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = c.Map.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	return
}
