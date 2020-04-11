package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *Map) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "VCCVYCVYXC"
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
		for _i := range _s {
			err = serializer.Write_int(w, _b, _i)
			if err != nil {
				return
			}

			err = serializer.Write_int(w, _b, _s[_i])
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
		for _i := range _s {
			err = serializer.Write_string(w, _b, _i)
			if err != nil {
				return
			}

			err = serializer.Write_int(w, _b, _s[_i])
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
		for _i := range _s {
			err = serializer.Write_string(w, _b, _i)
			if err != nil {
				return
			}

			{
				_s := _s[_i]
				_n = len(_s)
				err = serializer.Write_int(w, _b, _n)
				if err != nil {
					return
				}
				for _i := 0; _i < _n; _i++ {
					err = serializer.Write_int(w, _b, _s[_i])
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}

func (self *Map) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "VCCVYCVYXC"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _int int
	var _n int
	var _string string

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n > 0 {
		self.Empty = make(map[int]int, _n)
		_s := self.Empty
		var _i int
		for _j := 0; _j < _n; _j++ {
			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_i = _int

			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int
		}
	} else {
		self.Empty = nil
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n > 0 {
		self.StringInt = make(map[string]int, _n)
		_s := self.StringInt
		var _i string
		for _j := 0; _j < _n; _j++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_i = _string

			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int
		}
	} else {
		self.StringInt = nil
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _n > 0 {
		self.StringInts = make(map[string][]int, _n)
		_s := self.StringInts
		var _i string
		for _j := 0; _j < _n; _j++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_i = _string

			_n, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			if c := cap(_s[_i]); _n > c || c-_n > c/8 {
				_s[_i] = make([]int, _n)
			} else {
				_s[_i] = _s[_i][:_n]
			}
			if _n > 0 {
				_s := _s[_i]
				for _i := 0; _i < _n; _i++ {
					_int, err = serializer.Read_int(r, _b)
					if err != nil {
						return
					}
					_s[_i] = _int
				}
			}
		}
	} else {
		self.StringInts = nil
	}

	return
}
