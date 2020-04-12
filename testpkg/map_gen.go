package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *Map) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "VCCVYCVYXCVWHHVWCCVCWCVWCWC"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	var _n int

	{
		_s := self.Empty
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_int(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.StringInt
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_string(w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_int(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.StringInts
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_string(w, _b, _k)
			if err != nil {
				return
			}

			{
				_s := _s[_k]
				_n = len(_s)
				err = serializer.Write_int(w, _b, _n)
				if err != nil {
					return
				}
				for _k := 0; _k < _n; _k++ {
					err = serializer.Write_int(w, _b, _s[_k])
					if err != nil {
						return
					}
				}
			}
		}
	}
	{
		_s := self.UintPtrUint
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(w, _b, _k == nil)
			if err != nil {
				return
			}
			if _k != nil {
				err = serializer.Write_uint(w, _b, *_k)
				if err != nil {
					return
				}
			}

			err = serializer.Write_uint(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.IntPtrInt
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(w, _b, _k == nil)
			if err != nil {
				return
			}
			if _k != nil {
				err = serializer.Write_int(w, _b, *_k)
				if err != nil {
					return
				}
			}

			err = serializer.Write_int(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.IntIntPtr
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_int(w, _b, _k)
			if err != nil {
				return
			}

			err = serializer.Write_bool(w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_int(w, _b, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := self.IntPtrIntPtr
		err = serializer.Write_int(w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = serializer.Write_bool(w, _b, _k == nil)
			if err != nil {
				return
			}
			if _k != nil {
				err = serializer.Write_int(w, _b, *_k)
				if err != nil {
					return
				}
			}

			err = serializer.Write_bool(w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_int(w, _b, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (self *Map) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "VCCVYCVYXCVWHHVWCCVCWCVWCWC"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _bool bool
	var _int int
	var _n int
	var _string string
	var _uint uint

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.Empty = nil
	} else {
		self.Empty = make(map[int]int, _n)
		_s := self.Empty
		var _k int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_k = _int

			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.StringInt = nil
	} else {
		self.StringInt = make(map[string]int, _n)
		_s := self.StringInt
		var _k string
		for _j := 0; _j < _n; _j++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_k = _string

			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.StringInts = nil
	} else {
		self.StringInts = make(map[string][]int, _n)
		_s := self.StringInts
		var _k string
		for _j := 0; _j < _n; _j++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_k = _string

			_n, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			if c := cap(_s[_k]); _n > c || c-_n > c/8 {
				_s[_k] = make([]int, _n)
			} else {
				_s[_k] = (_s[_k])[:_n]
			}
			if _n > 0 {
				_s := _s[_k]
				for _k := 0; _k < _n; _k++ {
					_int, err = serializer.Read_int(r, _b)
					if err != nil {
						return
					}
					_s[_k] = _int
				}
			}
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.UintPtrUint = nil
	} else {
		self.UintPtrUint = make(map[*uint]uint, _n)
		_s := self.UintPtrUint
		var _k *uint
		for _j := 0; _j < _n; _j++ {
			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_k = nil
			} else {
				_uint, err = serializer.Read_uint(r, _b)
				if err != nil {
					return
				}
				*_k = _uint
			}

			_uint, err = serializer.Read_uint(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.IntPtrInt = nil
	} else {
		self.IntPtrInt = make(map[*int]int, _n)
		_s := self.IntPtrInt
		var _k *int
		for _j := 0; _j < _n; _j++ {
			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_k = nil
			} else {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				*_k = _int
			}

			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.IntIntPtr = nil
	} else {
		self.IntIntPtr = make(map[int]*int, _n)
		_s := self.IntIntPtr
		var _k int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_k = _int

			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				*_s[_k] = _int
			}
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n == 0 {
		self.IntPtrIntPtr = nil
	} else {
		self.IntPtrIntPtr = make(map[*int]*int, _n)
		_s := self.IntPtrIntPtr
		var _k *int
		for _j := 0; _j < _n; _j++ {
			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_k = nil
			} else {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				*_k = _int
			}

			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				*_s[_k] = _int
			}
		}
	}

	return
}
