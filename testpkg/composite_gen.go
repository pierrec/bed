package testpkg

import "io"
import "github.com/pierrec/serializer"

func (self *Composite) MarshalBinaryTo(w io.Writer) (err error) {
	var __buf [16]byte
	var _b = __buf[:]

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
	var __buf [16]byte
	var _b = __buf[:]
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
