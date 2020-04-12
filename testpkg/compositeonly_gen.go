package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

const _CompositeOnlyLayout = "ZZZZ"

func (c *CompositeOnly) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _CompositeOnlyLayout)
	if err != nil {
		return
	}

	{
		_s := c.Basic
		err = _s.MarshalBinaryTo(w)
		if err != nil {
			return
		}
	}

	{
		_s := c.Slice
		err = _s.MarshalBinaryTo(w)
		if err != nil {
			return
		}
	}

	{
		_s := c.Array
		err = _s.MarshalBinaryTo(w)
		if err != nil {
			return
		}
	}

	{
		_s := c.Map
		err = _s.MarshalBinaryTo(w)
		if err != nil {
			return
		}
	}

	return
}

func (c *CompositeOnly) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _CompositeOnlyLayout) {
		return serializer.ErrInvalidData
	}

	{
		_s := c.Basic
		err = _s.UnmarshalBinaryFrom(r)
		if err != nil {
			return
		}
	}

	{
		_s := c.Slice
		err = _s.UnmarshalBinaryFrom(r)
		if err != nil {
			return
		}
	}

	{
		_s := c.Array
		err = _s.UnmarshalBinaryFrom(r)
		if err != nil {
			return
		}
	}

	{
		_s := c.Map
		err = _s.UnmarshalBinaryFrom(r)
		if err != nil {
			return
		}
	}

	return
}
