package testpkg

import (
	"io"
	"time"

	"github.com/pierrec/serializer"
)

const _MiscLayout = "ZZ"

func (m *Misc) MarshalBinaryTo(w io.Writer) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(w, _b, _MiscLayout)
	if err != nil {
		return
	}

	err = serializer.Write_time(w, _b, m.Time)
	if err != nil {
		return
	}

	err = serializer.Write_time(w, _b, time.Time(m.MyTime))
	if err != nil {
		return
	}

	return
}

func (m *Misc) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(r, _b, _MiscLayout)
	if err != nil {
		return
	}

	var _time time.Time

	_time, err = serializer.Read_time(r, _b)
	if err != nil {
		return
	}
	m.Time = _time

	_time, err = serializer.Read_time(r, _b)
	if err != nil {
		return
	}
	m.MyTime = MyTime(_time)

	return
}
