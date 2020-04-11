package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *Composite) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "IZZZ"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	err = serializer.Write_bytes(w, _b, self.Bytes)
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

	return
}

func (self *Composite) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "IZZZ"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _bytes []byte

	_bytes, err = serializer.Read_bytes(r, _b)
	if err != nil {
		return
	}
	self.Bytes = _bytes

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

	return
}
