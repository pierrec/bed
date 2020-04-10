package testpkg

import "io"

func (self *CompositeOnly) MarshalBinaryTo(w io.Writer) (err error) {
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
