package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _CompositeLayout = "XZZZ"

func (c *Composite) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _CompositeLayout)
	if err != nil {
		return
	}

	err = serializer.Write_bytes(_w, _b, c.Bytes)
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

	return
}

func (c *Composite) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _CompositeLayout)
	if err != nil {
		return
	}

	c.Bytes, err = serializer.Read_bytes(_r, _b, nil)
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

	return
}
