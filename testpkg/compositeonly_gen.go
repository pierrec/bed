package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (c *CompositeOnly) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "ZZZZ"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
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
	const _check = "ZZZZ"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
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
