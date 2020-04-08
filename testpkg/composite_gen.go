package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

func (self *Composite) MarshalBinaryTo(w io.Writer) (err error) {
	var buf [16]byte
	b := buf[:]
	var _n int
	_ = _n

	err = serializer.Write_int(w, b, len(self.Bytes))
	if err != nil {
		return
	}
	_, err = w.Write(self.Bytes)
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

	return
}

func (self *Composite) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var buf [16]byte
	b := buf[:]
	var _uint8 uint8
	_ = _uint8
	var _n int
	_ = _n

	_n, err = serializer.Read_int(r, b)
	if err != nil {
		return
	}
	if c := cap(self.Bytes); _n > c || c-_n > c/8 {
		self.Bytes = make([]byte, _n)
	} else {
		self.Bytes = self.Bytes[:_n]
	}
	_, err = io.ReadFull(r, self.Bytes)
	if err != nil {
		return
	}

	err = self.Basic.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	err = self.Slice.UnmarshalBinaryFrom(r)
	if err != nil {
		return
	}

	return
}
