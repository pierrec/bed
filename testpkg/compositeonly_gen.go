package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _CompositeOnlyLayout = "ZZZZ"

func (c *CompositeOnly) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _CompositeOnlyLayout)
	if err != nil {
		return
	}

	err = c.Basic.MarshalBinaryTo(_w)
	if err != nil {
		return
	}

	err = c.Slice.MarshalBinaryTo(_w)
	if err != nil {
		return
	}

	err = c.Array.MarshalBinaryTo(_w)
	if err != nil {
		return
	}

	err = c.Map.MarshalBinaryTo(_w)
	if err != nil {
		return
	}

	return
}

func (c *CompositeOnly) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _CompositeOnlyLayout)
	if err != nil {
		return
	}

	err = c.Basic.UnmarshalBinaryFrom(_r)
	if err != nil {
		return
	}

	err = c.Slice.UnmarshalBinaryFrom(_r)
	if err != nil {
		return
	}

	err = c.Array.UnmarshalBinaryFrom(_r)
	if err != nil {
		return
	}

	err = c.Map.UnmarshalBinaryFrom(_r)
	if err != nil {
		return
	}

	return
}
