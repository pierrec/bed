package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *CompositeOnly) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "ZZZZ"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	err = self.Basic.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	err = self.Slice.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	err = self.Array.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	err = self.Map.MarshalBinaryTo(w)
	if err != nil {
		return
	}

	return
}

func (self *CompositeOnly) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "ZZZZ"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	err = self.Basic.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = self.Slice.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = self.Array.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = self.Map.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	return
}
