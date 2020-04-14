package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _MapLayout = "VCCVYCVYXCVWHHVWCCVCWCVWCWCVCZVCWZ"

func (m *Map) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _MapLayout)
	if err != nil {
		return
	}

	var _n int

	{
		_s := m.Empty
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_int(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := m.StringInt
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_string(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_int(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := m.StringInts
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_string(_w, _b, _k)
			if err != nil {
				return
			}

			{
				_s := _s[_k]
				_n = len(_s)
				err = serializer.Write_int(_w, _b, _n)
				if err != nil {
					return
				}
				for _k := 0; _k < _n; _k++ {
					err = serializer.Write_int(_w, _b, _s[_k])
					if err != nil {
						return
					}
				}
			}
		}
	}
	{
		_s := m.UintPtrUint
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(_w, _b, _k == nil)
			if err != nil {
				return
			}
			if _k != nil {
				err = serializer.Write_uint(_w, _b, *_k)
				if err != nil {
					return
				}
			}

			err = serializer.Write_uint(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := m.IntPtrInt
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(_w, _b, _k == nil)
			if err != nil {
				return
			}
			if _k != nil {
				err = serializer.Write_int(_w, _b, *_k)
				if err != nil {
					return
				}
			}

			err = serializer.Write_int(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := m.IntIntPtr
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_bool(_w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_int(_w, _b, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := m.IntPtrIntPtr
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(_w, _b, _k == nil)
			if err != nil {
				return
			}
			if _k != nil {
				err = serializer.Write_int(_w, _b, *_k)
				if err != nil {
					return
				}
			}

			err = serializer.Write_bool(_w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_int(_w, _b, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := m.IntStruct
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			{
				_struct := _s[_k]
				err = _struct.MarshalBinaryTo(_w)
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := m.IntStructPtr
		err = serializer.Write_int(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_bool(_w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = _s[_k].MarshalBinaryTo(_w)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (m *Map) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _MapLayout)
	if err != nil {
		return
	}

	var _bool bool
	var _int int
	var _n int
	var _string string
	var _uint uint

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.Empty = nil
	} else {
		m.Empty = make(map[int]int, _n)
		_s := m.Empty
		var _k int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.StringInt = nil
	} else {
		m.StringInt = make(map[string]int, _n)
		_s := m.StringInt
		var _k string
		for _j := 0; _j < _n; _j++ {
			_string, err = serializer.Read_string(_r, _b)
			if err != nil {
				return
			}
			_k = _string

			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.StringInts = nil
	} else {
		m.StringInts = make(map[string][]int, _n)
		_s := m.StringInts
		var _k string
		for _j := 0; _j < _n; _j++ {
			_string, err = serializer.Read_string(_r, _b)
			if err != nil {
				return
			}
			_k = _string

			_n, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			if _c := cap(_s[_k]); _n > _c || _c-_n > _c/8 {
				_s[_k] = make([]int, _n)
			} else {
				_s[_k] = (_s[_k])[:_n]
			}
			if _n > 0 {
				_s := _s[_k]
				for _k := 0; _k < _n; _k++ {
					_int, err = serializer.Read_int(_r, _b)
					if err != nil {
						return
					}
					_s[_k] = _int
				}
			}
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.UintPtrUint = nil
	} else {
		m.UintPtrUint = make(map[*uint]uint, _n)
		_s := m.UintPtrUint
		var _k *uint
		for _j := 0; _j < _n; _j++ {
			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_k = nil
			} else {
				_k = new(uint)
				_uint, err = serializer.Read_uint(_r, _b)
				if err != nil {
					return
				}
				*_k = _uint
			}

			_uint, err = serializer.Read_uint(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntPtrInt = nil
	} else {
		m.IntPtrInt = make(map[*int]int, _n)
		_s := m.IntPtrInt
		var _k *int
		for _j := 0; _j < _n; _j++ {
			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_k = nil
			} else {
				_k = new(int)
				_int, err = serializer.Read_int(_r, _b)
				if err != nil {
					return
				}
				*_k = _int
			}

			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntIntPtr = nil
	} else {
		m.IntIntPtr = make(map[int]*int, _n)
		_s := m.IntIntPtr
		var _k int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(int)
				_int, err = serializer.Read_int(_r, _b)
				if err != nil {
					return
				}
				*_s[_k] = _int
			}
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntPtrIntPtr = nil
	} else {
		m.IntPtrIntPtr = make(map[*int]*int, _n)
		_s := m.IntPtrIntPtr
		var _k *int
		for _j := 0; _j < _n; _j++ {
			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_k = nil
			} else {
				_k = new(int)
				_int, err = serializer.Read_int(_r, _b)
				if err != nil {
					return
				}
				*_k = _int
			}

			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(int)
				_int, err = serializer.Read_int(_r, _b)
				if err != nil {
					return
				}
				*_s[_k] = _int
			}
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntStruct = nil
	} else {
		m.IntStruct = make(map[int]Basic, _n)
		_s := m.IntStruct
		var _k int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			{
				_struct := _s[_k]
				err = _struct.UnmarshalBinaryFrom(_r)
				if err != nil {
					return
				}
				_s[_k] = _struct
			}
		}
	}

	_n, err = serializer.Read_int(_r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntStructPtr = nil
	} else {
		m.IntStructPtr = make(map[int]*Basic, _n)
		_s := m.IntStructPtr
		var _k int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(Basic)
				err = _s[_k].UnmarshalBinaryFrom(_r)
				if err != nil {
					return
				}
			}
		}
	}

	return
}
