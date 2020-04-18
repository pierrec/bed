package testpkg

import (
	"github.com/pierrec/packer/iobyte"
	"io"

	"github.com/pierrec/serializer"
)

const _MapLayout = "VCCVBBVYCVYXCVWHHVWCCVCWCVWCWCVZCYC"

func (m *Map) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
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
		err = serializer.Write_len(_w, _b, len(_s))
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
		_s := m.BoolBool
		err = serializer.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(_w, _b, _k)
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
		_s := m.StringInt
		err = serializer.Write_len(_w, _b, len(_s))
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
		err = serializer.Write_len(_w, _b, len(_s))
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
				err = serializer.Write_len(_w, _b, _n)
				if err != nil {
					return
				}
				for _k, _kn := 0, _n; _k < _kn; _k++ {
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
		err = serializer.Write_len(_w, _b, len(_s))
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
		err = serializer.Write_len(_w, _b, len(_s))
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
		err = serializer.Write_len(_w, _b, len(_s))
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
		err = serializer.Write_len(_w, _b, len(_s))
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
		_s := m.AnonInt
		err = serializer.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			{
				_s := &_k

				err = serializer.Write_int(_w, _b, _s.Int)
				if err != nil {
					return
				}

				err = serializer.Write_string(_w, _b, _s.String)
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
	return
}

func (m *Map) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.Empty = nil
	} else {
		m.Empty = make(map[int]int, _n)
		_s := m.Empty
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.BoolBool = nil
	} else {
		m.BoolBool = make(map[bool]bool, _n)
		_s := m.BoolBool
		var _k bool
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_bool, err = serializer.Read_bool(_r, _b)
			if err != nil {
				return
			}
			_k = _bool

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
		m.StringInt = nil
	} else {
		m.StringInt = make(map[string]int, _n)
		_s := m.StringInt
		var _k string
		for _j, _jn := 0, _n; _j < _jn; _j++ {
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.StringInts = nil
	} else {
		m.StringInts = make(map[string][]int, _n)
		_s := m.StringInts
		var _k string
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_string, err = serializer.Read_string(_r, _b)
			if err != nil {
				return
			}
			_k = _string

			_n, err = serializer.Read_len(_r)
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
				for _k, _kn := 0, _n; _k < _kn; _k++ {
					_int, err = serializer.Read_int(_r, _b)
					if err != nil {
						return
					}
					_s[_k] = _int
				}
			}
		}
	}

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.UintPtrUint = nil
	} else {
		m.UintPtrUint = make(map[*uint]uint, _n)
		_s := m.UintPtrUint
		var _k *uint
		for _j, _jn := 0, _n; _j < _jn; _j++ {
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntPtrInt = nil
	} else {
		m.IntPtrInt = make(map[*int]int, _n)
		_s := m.IntPtrInt
		var _k *int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntIntPtr = nil
	} else {
		m.IntIntPtr = make(map[int]*int, _n)
		_s := m.IntIntPtr
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.IntPtrIntPtr = nil
	} else {
		m.IntPtrIntPtr = make(map[*int]*int, _n)
		_s := m.IntPtrIntPtr
		var _k *int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
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

	_n, err = serializer.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		m.AnonInt = nil
	} else {
		m.AnonInt = make(map[struct {
			Int    int
			String string
		}]int, _n)
		_s := m.AnonInt
		var _k struct {
			Int    int
			String string
		}
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			{
				_s := &_k

				_int, err = serializer.Read_int(_r, _b)
				if err != nil {
					return
				}
				_s.Int = _int

				_string, err = serializer.Read_string(_r, _b)
				if err != nil {
					return
				}
				_s.String = _string

			}

			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	return
}
