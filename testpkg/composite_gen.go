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

	var _n int

	{
		_s := c.Bytes
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_uint8(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
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

	var _n int
	var _uint8 uint8

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(c.Bytes); _n > _c || _c-_n > _c/8 {
		c.Bytes = make([]uint8, _n)
	} else {
		c.Bytes = (c.Bytes)[:_n]
	}
	if _n > 0 {
		_s := c.Bytes
		for _k := 0; _k < _n; _k++ {
			_uint8, err = serializer.Read_uint8(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint8
		}
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

	return
}
