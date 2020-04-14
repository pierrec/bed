package testpkg

import (
	"io"
	"time"

	"github.com/pierrec/serializer"
)

const _MiscLayout = "ZZ"

func (m *Misc) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _MiscLayout)
	if err != nil {
		return
	}

	err = serializer.Write_time(_w, _b, m.Time)
	if err != nil {
		return
	}

	err = serializer.Write_time(_w, _b, time.Time(m.MyTime))
	if err != nil {
		return
	}

	return
}

func (m *Misc) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _MiscLayout)
	if err != nil {
		return
	}

	var _time time.Time

	_time, err = serializer.Read_time(_r, _b)
	if err != nil {
		return
	}
	m.Time = _time

	_time, err = serializer.Read_time(_r, _b)
	if err != nil {
		return
	}
	m.MyTime = MyTime(_time)

	return
}
