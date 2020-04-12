package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *Slice) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "XCXDXEXFXGXHXIXJXKXLXPXQXYXVYXC"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	var _n int

	{
		_s := self.Int
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
	{
		_s := self.Int8
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_int8(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Int16
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_int16(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Int32
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_int32(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Int64
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_int64(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Uint
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_uint(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	err = serializer.Write_bytes(w, _b, self.Uint8)
	if err != nil {
		return
	}

	{
		_s := self.Uint16
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_uint16(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Uint32
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_uint32(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Uint64
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_uint64(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Complex64
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_complex64(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Complex128
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_complex128(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.String
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			err = serializer.Write_string(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := self.Maps
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _i := 0; _i < _n; _i++ {
			{
				_s := _s[_i]
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
		}
	}
	return
}

func (self *Slice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "XCXDXEXFXGXHXIXJXKXLXPXQXYXVYXC"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _bytes []byte
	var _complex128 complex128
	var _complex64 complex64
	var _int int
	var _int16 int16
	var _int32 int32
	var _int64 int64
	var _int8 int8
	var _n int
	var _string string
	var _uint uint
	var _uint16 uint16
	var _uint32 uint32
	var _uint64 uint64

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Int); _n > c || c-_n > c/8 {
		self.Int = make([]int, _n)
	} else {
		self.Int = (self.Int)[:_n]
	}
	if _n > 0 {
		_s := self.Int
		for _i := 0; _i < _n; _i++ {
			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Int8); _n > c || c-_n > c/8 {
		self.Int8 = make([]int8, _n)
	} else {
		self.Int8 = (self.Int8)[:_n]
	}
	if _n > 0 {
		_s := self.Int8
		for _i := 0; _i < _n; _i++ {
			_int8, err = serializer.Read_int8(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int8
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Int16); _n > c || c-_n > c/8 {
		self.Int16 = make([]int16, _n)
	} else {
		self.Int16 = (self.Int16)[:_n]
	}
	if _n > 0 {
		_s := self.Int16
		for _i := 0; _i < _n; _i++ {
			_int16, err = serializer.Read_int16(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int16
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Int32); _n > c || c-_n > c/8 {
		self.Int32 = make([]int32, _n)
	} else {
		self.Int32 = (self.Int32)[:_n]
	}
	if _n > 0 {
		_s := self.Int32
		for _i := 0; _i < _n; _i++ {
			_int32, err = serializer.Read_int32(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int32
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Int64); _n > c || c-_n > c/8 {
		self.Int64 = make([]int64, _n)
	} else {
		self.Int64 = (self.Int64)[:_n]
	}
	if _n > 0 {
		_s := self.Int64
		for _i := 0; _i < _n; _i++ {
			_int64, err = serializer.Read_int64(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int64
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Uint); _n > c || c-_n > c/8 {
		self.Uint = make([]uint, _n)
	} else {
		self.Uint = (self.Uint)[:_n]
	}
	if _n > 0 {
		_s := self.Uint
		for _i := 0; _i < _n; _i++ {
			_uint, err = serializer.Read_uint(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint
		}
	}

	_bytes, err = serializer.Read_bytes(r, _b)
	if err != nil {
		return
	}
	self.Uint8 = _bytes

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Uint16); _n > c || c-_n > c/8 {
		self.Uint16 = make([]uint16, _n)
	} else {
		self.Uint16 = (self.Uint16)[:_n]
	}
	if _n > 0 {
		_s := self.Uint16
		for _i := 0; _i < _n; _i++ {
			_uint16, err = serializer.Read_uint16(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint16
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Uint32); _n > c || c-_n > c/8 {
		self.Uint32 = make([]uint32, _n)
	} else {
		self.Uint32 = (self.Uint32)[:_n]
	}
	if _n > 0 {
		_s := self.Uint32
		for _i := 0; _i < _n; _i++ {
			_uint32, err = serializer.Read_uint32(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint32
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Uint64); _n > c || c-_n > c/8 {
		self.Uint64 = make([]uint64, _n)
	} else {
		self.Uint64 = (self.Uint64)[:_n]
	}
	if _n > 0 {
		_s := self.Uint64
		for _i := 0; _i < _n; _i++ {
			_uint64, err = serializer.Read_uint64(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint64
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Complex64); _n > c || c-_n > c/8 {
		self.Complex64 = make([]complex64, _n)
	} else {
		self.Complex64 = (self.Complex64)[:_n]
	}
	if _n > 0 {
		_s := self.Complex64
		for _i := 0; _i < _n; _i++ {
			_complex64, err = serializer.Read_complex64(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _complex64
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Complex128); _n > c || c-_n > c/8 {
		self.Complex128 = make([]complex128, _n)
	} else {
		self.Complex128 = (self.Complex128)[:_n]
	}
	if _n > 0 {
		_s := self.Complex128
		for _i := 0; _i < _n; _i++ {
			_complex128, err = serializer.Read_complex128(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _complex128
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.String); _n > c || c-_n > c/8 {
		self.String = make([]string, _n)
	} else {
		self.String = (self.String)[:_n]
	}
	if _n > 0 {
		_s := self.String
		for _i := 0; _i < _n; _i++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _string
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if c := cap(self.Maps); _n > c || c-_n > c/8 {
		self.Maps = make([]map[string][]int, _n)
	} else {
		self.Maps = (self.Maps)[:_n]
	}
	if _n > 0 {
		_s := self.Maps
		for _i := 0; _i < _n; _i++ {
			_n, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			if _n > 0 {
				_s[_i] = make(map[string][]int, _n)
				_s := _s[_i]
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
						_s[_i] = (_s[_i])[:_n]
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
				_s[_i] = nil
			}
		}
	}

	return
}

func (self *SlicePtr) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "WXCWXDWXEWXFWXGWXHWXIWXJWXKWXLWXPWXQWXYWXVYXC"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	var _n int

	err = serializer.Write_bool(w, _b, self.Int == nil)
	if err != nil {
		return
	}
	if self.Int != nil {
		{
			_s := *self.Int
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

	err = serializer.Write_bool(w, _b, self.Int8 == nil)
	if err != nil {
		return
	}
	if self.Int8 != nil {
		{
			_s := *self.Int8
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_int8(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Int16 == nil)
	if err != nil {
		return
	}
	if self.Int16 != nil {
		{
			_s := *self.Int16
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_int16(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Int32 == nil)
	if err != nil {
		return
	}
	if self.Int32 != nil {
		{
			_s := *self.Int32
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_int32(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Int64 == nil)
	if err != nil {
		return
	}
	if self.Int64 != nil {
		{
			_s := *self.Int64
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_int64(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint == nil)
	if err != nil {
		return
	}
	if self.Uint != nil {
		{
			_s := *self.Uint
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_uint(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint8 == nil)
	if err != nil {
		return
	}
	if self.Uint8 != nil {
		err = serializer.Write_bytes(w, _b, *self.Uint8)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint16 == nil)
	if err != nil {
		return
	}
	if self.Uint16 != nil {
		{
			_s := *self.Uint16
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_uint16(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint32 == nil)
	if err != nil {
		return
	}
	if self.Uint32 != nil {
		{
			_s := *self.Uint32
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_uint32(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint64 == nil)
	if err != nil {
		return
	}
	if self.Uint64 != nil {
		{
			_s := *self.Uint64
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_uint64(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Complex64 == nil)
	if err != nil {
		return
	}
	if self.Complex64 != nil {
		{
			_s := *self.Complex64
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_complex64(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Complex128 == nil)
	if err != nil {
		return
	}
	if self.Complex128 != nil {
		{
			_s := *self.Complex128
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_complex128(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.String == nil)
	if err != nil {
		return
	}
	if self.String != nil {
		{
			_s := *self.String
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				err = serializer.Write_string(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, self.Maps == nil)
	if err != nil {
		return
	}
	if self.Maps != nil {
		{
			_s := *self.Maps
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _i := 0; _i < _n; _i++ {
				{
					_s := _s[_i]
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
			}
		}
	}

	return
}

func (self *SlicePtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "WXCWXDWXEWXFWXGWXHWXIWXJWXKWXLWXPWXQWXYWXVYXC"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _bool bool
	var _bytes []byte
	var _complex128 complex128
	var _complex64 complex64
	var _int int
	var _int16 int16
	var _int32 int32
	var _int64 int64
	var _int8 int8
	var _n int
	var _string string
	var _uint uint
	var _uint16 uint16
	var _uint32 uint32
	var _uint64 uint64

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Int = nil
	} else {
		*self.Int = []int{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Int); _n > c || c-_n > c/8 {
			*self.Int = make([]int, _n)
		} else {
			*self.Int = (*self.Int)[:_n]
		}
		if _n > 0 {
			_s := *self.Int
			for _i := 0; _i < _n; _i++ {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Int8 = nil
	} else {
		*self.Int8 = []int8{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Int8); _n > c || c-_n > c/8 {
			*self.Int8 = make([]int8, _n)
		} else {
			*self.Int8 = (*self.Int8)[:_n]
		}
		if _n > 0 {
			_s := *self.Int8
			for _i := 0; _i < _n; _i++ {
				_int8, err = serializer.Read_int8(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int8
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Int16 = nil
	} else {
		*self.Int16 = []int16{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Int16); _n > c || c-_n > c/8 {
			*self.Int16 = make([]int16, _n)
		} else {
			*self.Int16 = (*self.Int16)[:_n]
		}
		if _n > 0 {
			_s := *self.Int16
			for _i := 0; _i < _n; _i++ {
				_int16, err = serializer.Read_int16(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int16
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Int32 = nil
	} else {
		*self.Int32 = []int32{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Int32); _n > c || c-_n > c/8 {
			*self.Int32 = make([]int32, _n)
		} else {
			*self.Int32 = (*self.Int32)[:_n]
		}
		if _n > 0 {
			_s := *self.Int32
			for _i := 0; _i < _n; _i++ {
				_int32, err = serializer.Read_int32(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int32
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Int64 = nil
	} else {
		*self.Int64 = []int64{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Int64); _n > c || c-_n > c/8 {
			*self.Int64 = make([]int64, _n)
		} else {
			*self.Int64 = (*self.Int64)[:_n]
		}
		if _n > 0 {
			_s := *self.Int64
			for _i := 0; _i < _n; _i++ {
				_int64, err = serializer.Read_int64(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int64
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Uint = nil
	} else {
		*self.Uint = []uint{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Uint); _n > c || c-_n > c/8 {
			*self.Uint = make([]uint, _n)
		} else {
			*self.Uint = (*self.Uint)[:_n]
		}
		if _n > 0 {
			_s := *self.Uint
			for _i := 0; _i < _n; _i++ {
				_uint, err = serializer.Read_uint(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Uint8 = nil
	} else {
		_bytes, err = serializer.Read_bytes(r, _b)
		if err != nil {
			return
		}
		*self.Uint8 = _bytes
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Uint16 = nil
	} else {
		*self.Uint16 = []uint16{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Uint16); _n > c || c-_n > c/8 {
			*self.Uint16 = make([]uint16, _n)
		} else {
			*self.Uint16 = (*self.Uint16)[:_n]
		}
		if _n > 0 {
			_s := *self.Uint16
			for _i := 0; _i < _n; _i++ {
				_uint16, err = serializer.Read_uint16(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint16
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Uint32 = nil
	} else {
		*self.Uint32 = []uint32{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Uint32); _n > c || c-_n > c/8 {
			*self.Uint32 = make([]uint32, _n)
		} else {
			*self.Uint32 = (*self.Uint32)[:_n]
		}
		if _n > 0 {
			_s := *self.Uint32
			for _i := 0; _i < _n; _i++ {
				_uint32, err = serializer.Read_uint32(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint32
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Uint64 = nil
	} else {
		*self.Uint64 = []uint64{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Uint64); _n > c || c-_n > c/8 {
			*self.Uint64 = make([]uint64, _n)
		} else {
			*self.Uint64 = (*self.Uint64)[:_n]
		}
		if _n > 0 {
			_s := *self.Uint64
			for _i := 0; _i < _n; _i++ {
				_uint64, err = serializer.Read_uint64(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint64
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Complex64 = nil
	} else {
		*self.Complex64 = []complex64{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Complex64); _n > c || c-_n > c/8 {
			*self.Complex64 = make([]complex64, _n)
		} else {
			*self.Complex64 = (*self.Complex64)[:_n]
		}
		if _n > 0 {
			_s := *self.Complex64
			for _i := 0; _i < _n; _i++ {
				_complex64, err = serializer.Read_complex64(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _complex64
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Complex128 = nil
	} else {
		*self.Complex128 = []complex128{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Complex128); _n > c || c-_n > c/8 {
			*self.Complex128 = make([]complex128, _n)
		} else {
			*self.Complex128 = (*self.Complex128)[:_n]
		}
		if _n > 0 {
			_s := *self.Complex128
			for _i := 0; _i < _n; _i++ {
				_complex128, err = serializer.Read_complex128(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _complex128
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.String = nil
	} else {
		*self.String = []string{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.String); _n > c || c-_n > c/8 {
			*self.String = make([]string, _n)
		} else {
			*self.String = (*self.String)[:_n]
		}
		if _n > 0 {
			_s := *self.String
			for _i := 0; _i < _n; _i++ {
				_string, err = serializer.Read_string(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _string
			}
		}
	}

	if _bool, err = serializer.Read_bool(r, _b); err != nil {
		return
	}
	if _bool {
		self.Maps = nil
	} else {
		*self.Maps = []map[string][]int{}
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if c := cap(*self.Maps); _n > c || c-_n > c/8 {
			*self.Maps = make([]map[string][]int, _n)
		} else {
			*self.Maps = (*self.Maps)[:_n]
		}
		if _n > 0 {
			_s := *self.Maps
			for _i := 0; _i < _n; _i++ {
				_n, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				if _n > 0 {
					_s[_i] = make(map[string][]int, _n)
					_s := _s[_i]
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
							_s[_i] = (_s[_i])[:_n]
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
					_s[_i] = nil
				}
			}
		}
	}

	return
}
