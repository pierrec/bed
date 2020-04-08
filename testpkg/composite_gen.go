package testpkg

import "io"

func (self *Composite) MarshalBinaryTo(w io.Writer) (err error) {
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
