package testpkg

import (
	"io"
	"time"

	"github.com/pierrec/serializer"
)

const _TimeLayout = "ZZ"

func (t *Time) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _TimeLayout)
	if err != nil {
		return
	}

	err = serializer.Write_time(_w, _b, t.Time)
	if err != nil {
		return
	}

	err = serializer.Write_time(_w, _b, time.Time(t.MyTime))
	if err != nil {
		return
	}

	return
}

func (t *Time) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _TimeLayout)
	if err != nil {
		return
	}

	var _time time.Time

	_time, err = serializer.Read_time(_r, _b)
	if err != nil {
		return
	}
	t.Time = _time

	_time, err = serializer.Read_time(_r, _b)
	if err != nil {
		return
	}
	t.MyTime = MyTime(_time)

	return
}

const _TimePtrLayout = "WZWZ"

func (t *TimePtr) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _TimePtrLayout)
	if err != nil {
		return
	}

	err = serializer.Write_bool(_w, _b, t.Time == nil)
	if err != nil {
		return
	}
	if t.Time != nil {
		err = serializer.Write_time(_w, _b, *t.Time)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(_w, _b, t.MyTime == nil)
	if err != nil {
		return
	}
	if t.MyTime != nil {
		err = serializer.Write_time(_w, _b, time.Time(*t.MyTime))
		if err != nil {
			return
		}
	}

	return
}

func (t *TimePtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _TimePtrLayout)
	if err != nil {
		return
	}

	var _bool bool
	var _time time.Time

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		t.Time = nil
	} else {
		t.Time = new(time.Time)
		_time, err = serializer.Read_time(_r, _b)
		if err != nil {
			return
		}
		*t.Time = _time
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		t.MyTime = nil
	} else {
		t.MyTime = new(MyTime)
		_time, err = serializer.Read_time(_r, _b)
		if err != nil {
			return
		}
		*t.MyTime = MyTime(_time)
	}

	return
}

const _TimeSliceLayout = "XZXZ"

func (t *TimeSlice) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _TimeSliceLayout)
	if err != nil {
		return
	}

	var _n int

	{
		_s := t.Time
		_n = len(_s)
		err = serializer.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = serializer.Write_time(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := t.MyTime
		_n = len(_s)
		err = serializer.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = serializer.Write_time(_w, _b, time.Time(_s[_k]))
			if err != nil {
				return
			}
		}
	}
	return
}

func (t *TimeSlice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _TimeSliceLayout)
	if err != nil {
		return
	}

	var _n int
	var _time time.Time

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(t.Time); _n > _c || _c-_n > _c/8 {
		t.Time = make([]time.Time, _n)
	} else {
		t.Time = (t.Time)[:_n]
	}
	if _n > 0 {
		_s := t.Time
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _time
		}
	}

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(t.MyTime); _n > _c || _c-_n > _c/8 {
		t.MyTime = make([]MyTime, _n)
	} else {
		t.MyTime = (t.MyTime)[:_n]
	}
	if _n > 0 {
		_s := t.MyTime
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = MyTime(_time)
		}
	}

	return
}

const _TimeArrayLayout = "R4ZR4Z"

func (t *TimeArray) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _TimeArrayLayout)
	if err != nil {
		return
	}

	{
		_s := &t.Time
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_time(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &t.MyTime
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_time(_w, _b, time.Time(_s[_k]))
			if err != nil {
				return
			}
		}
	}
	return
}

func (t *TimeArray) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _TimeArrayLayout)
	if err != nil {
		return
	}

	var _time time.Time

	{
		_s := &t.Time
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _time
		}
	}

	{
		_s := &t.MyTime
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = MyTime(_time)
		}
	}

	return
}

const _TimeMapLayout = "VZBVZBVCZVCZ"

func (t *TimeMap) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _TimeMapLayout)
	if err != nil {
		return
	}

	{
		_s := t.TimeKey
		err = serializer.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_time(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_bool(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := t.MyTimeKey
		err = serializer.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_time(_w, _b, time.Time(_k))
			if err != nil {
				return
			}

			err = serializer.Write_bool(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := t.Time
		err = serializer.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_time(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := t.MyTime
		err = serializer.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_time(_w, _b, time.Time(_s[_k]))
			if err != nil {
				return
			}
		}
	}
	return
}

func (t *TimeMap) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _TimeMapLayout)
	if err != nil {
		return
	}

	var _bool bool
	var _int int
	var _n int
	var _time time.Time

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		t.TimeKey = nil
	} else {
		t.TimeKey = make(map[time.Time]bool, _n)
		_s := t.TimeKey
		var _k time.Time
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_k = _time

			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _bool
		}
	}

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		t.MyTimeKey = nil
	} else {
		t.MyTimeKey = make(map[MyTime]bool, _n)
		_s := t.MyTimeKey
		var _k MyTime
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_k = MyTime(_time)

			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _bool
		}
	}

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		t.Time = nil
	} else {
		t.Time = make(map[int]time.Time, _n)
		_s := t.Time
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _time
		}
	}

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		t.MyTime = nil
	} else {
		t.MyTime = make(map[int]MyTime, _n)
		_s := t.MyTime
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_time, err = serializer.Read_time(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = MyTime(_time)
		}
	}

	return
}
