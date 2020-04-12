package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _CompositeLayout = "XIZZZ"

func (c *Composite) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _CompositeLayout)
	if err != nil {
		return
	}

	err = serializer.Write_bytes(w, _b, c.Bytes)
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

	return
}

func (c *Composite) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _CompositeLayout)
	if err != nil {
		return
	}

	var _bytes []byte

	_bytes, err = serializer.Read_bytes(r, _b)
	if err != nil {
		return
	}
	c.Bytes = _bytes

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

	return
}
