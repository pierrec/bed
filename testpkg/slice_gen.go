package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _SliceLayout = "XCXDXEXFXGXHXIXJXKXLXPXQXYXVYXC"

func (s *Slice) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _SliceLayout)
	if err != nil {
		return
	}

	var _n int

	{
		_s := s.Int
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
	{
		_s := s.Int8
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_int8(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Int16
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_int16(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Int32
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_int32(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Int64
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_int64(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Uint
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_uint(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Uint8
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_uint8(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Uint16
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_uint16(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Uint32
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_uint32(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Uint64
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_uint64(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Complex64
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_complex64(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Complex128
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_complex128(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.String
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_string(w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := s.Maps
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			{
				_s := _s[_k]
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
		}
	}
	return
}

func (s *Slice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _SliceLayout)
	if err != nil {
		return
	}

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
	var _uint8 uint8

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Int); _n > _c || _c-_n > _c/8 {
		s.Int = make([]int, _n)
	} else {
		s.Int = (s.Int)[:_n]
	}
	if _n > 0 {
		_s := s.Int
		for _k := 0; _k < _n; _k++ {
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
	if _c := cap(s.Int8); _n > _c || _c-_n > _c/8 {
		s.Int8 = make([]int8, _n)
	} else {
		s.Int8 = (s.Int8)[:_n]
	}
	if _n > 0 {
		_s := s.Int8
		for _k := 0; _k < _n; _k++ {
			_int8, err = serializer.Read_int8(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int8
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Int16); _n > _c || _c-_n > _c/8 {
		s.Int16 = make([]int16, _n)
	} else {
		s.Int16 = (s.Int16)[:_n]
	}
	if _n > 0 {
		_s := s.Int16
		for _k := 0; _k < _n; _k++ {
			_int16, err = serializer.Read_int16(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int16
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Int32); _n > _c || _c-_n > _c/8 {
		s.Int32 = make([]int32, _n)
	} else {
		s.Int32 = (s.Int32)[:_n]
	}
	if _n > 0 {
		_s := s.Int32
		for _k := 0; _k < _n; _k++ {
			_int32, err = serializer.Read_int32(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int32
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Int64); _n > _c || _c-_n > _c/8 {
		s.Int64 = make([]int64, _n)
	} else {
		s.Int64 = (s.Int64)[:_n]
	}
	if _n > 0 {
		_s := s.Int64
		for _k := 0; _k < _n; _k++ {
			_int64, err = serializer.Read_int64(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int64
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Uint); _n > _c || _c-_n > _c/8 {
		s.Uint = make([]uint, _n)
	} else {
		s.Uint = (s.Uint)[:_n]
	}
	if _n > 0 {
		_s := s.Uint
		for _k := 0; _k < _n; _k++ {
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
	if _c := cap(s.Uint8); _n > _c || _c-_n > _c/8 {
		s.Uint8 = make([]uint8, _n)
	} else {
		s.Uint8 = (s.Uint8)[:_n]
	}
	if _n > 0 {
		_s := s.Uint8
		for _k := 0; _k < _n; _k++ {
			_uint8, err = serializer.Read_uint8(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint8
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Uint16); _n > _c || _c-_n > _c/8 {
		s.Uint16 = make([]uint16, _n)
	} else {
		s.Uint16 = (s.Uint16)[:_n]
	}
	if _n > 0 {
		_s := s.Uint16
		for _k := 0; _k < _n; _k++ {
			_uint16, err = serializer.Read_uint16(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint16
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Uint32); _n > _c || _c-_n > _c/8 {
		s.Uint32 = make([]uint32, _n)
	} else {
		s.Uint32 = (s.Uint32)[:_n]
	}
	if _n > 0 {
		_s := s.Uint32
		for _k := 0; _k < _n; _k++ {
			_uint32, err = serializer.Read_uint32(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint32
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Uint64); _n > _c || _c-_n > _c/8 {
		s.Uint64 = make([]uint64, _n)
	} else {
		s.Uint64 = (s.Uint64)[:_n]
	}
	if _n > 0 {
		_s := s.Uint64
		for _k := 0; _k < _n; _k++ {
			_uint64, err = serializer.Read_uint64(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint64
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Complex64); _n > _c || _c-_n > _c/8 {
		s.Complex64 = make([]complex64, _n)
	} else {
		s.Complex64 = (s.Complex64)[:_n]
	}
	if _n > 0 {
		_s := s.Complex64
		for _k := 0; _k < _n; _k++ {
			_complex64, err = serializer.Read_complex64(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _complex64
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Complex128); _n > _c || _c-_n > _c/8 {
		s.Complex128 = make([]complex128, _n)
	} else {
		s.Complex128 = (s.Complex128)[:_n]
	}
	if _n > 0 {
		_s := s.Complex128
		for _k := 0; _k < _n; _k++ {
			_complex128, err = serializer.Read_complex128(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _complex128
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.String); _n > _c || _c-_n > _c/8 {
		s.String = make([]string, _n)
	} else {
		s.String = (s.String)[:_n]
	}
	if _n > 0 {
		_s := s.String
		for _k := 0; _k < _n; _k++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_s[_k] = _string
		}
	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.Maps); _n > _c || _c-_n > _c/8 {
		s.Maps = make([]map[string][]int, _n)
	} else {
		s.Maps = (s.Maps)[:_n]
	}
	if _n > 0 {
		_s := s.Maps
		for _k := 0; _k < _n; _k++ {
			_n, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			if _n == 0 {
				_s[_k] = nil
			} else {
				_s[_k] = make(map[string][]int, _n)
				_s := _s[_k]
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
					if _c := cap(_s[_k]); _n > _c || _c-_n > _c/8 {
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
		}
	}

	return
}

const _SlicePtrLayout = "WXCWXDWXEWXFWXGWXHWXIWXJWXKWXLWXPWXQWXYWXVYXC"

func (s *SlicePtr) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _SlicePtrLayout)
	if err != nil {
		return
	}

	var _n int

	err = serializer.Write_bool(w, _b, s.Int == nil)
	if err != nil {
		return
	}
	if s.Int != nil {
		{
			_s := *s.Int
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

	err = serializer.Write_bool(w, _b, s.Int8 == nil)
	if err != nil {
		return
	}
	if s.Int8 != nil {
		{
			_s := *s.Int8
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_int8(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Int16 == nil)
	if err != nil {
		return
	}
	if s.Int16 != nil {
		{
			_s := *s.Int16
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_int16(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Int32 == nil)
	if err != nil {
		return
	}
	if s.Int32 != nil {
		{
			_s := *s.Int32
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_int32(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Int64 == nil)
	if err != nil {
		return
	}
	if s.Int64 != nil {
		{
			_s := *s.Int64
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_int64(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Uint == nil)
	if err != nil {
		return
	}
	if s.Uint != nil {
		{
			_s := *s.Uint
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_uint(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Uint8 == nil)
	if err != nil {
		return
	}
	if s.Uint8 != nil {
		{
			_s := *s.Uint8
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_uint8(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Uint16 == nil)
	if err != nil {
		return
	}
	if s.Uint16 != nil {
		{
			_s := *s.Uint16
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_uint16(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Uint32 == nil)
	if err != nil {
		return
	}
	if s.Uint32 != nil {
		{
			_s := *s.Uint32
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_uint32(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Uint64 == nil)
	if err != nil {
		return
	}
	if s.Uint64 != nil {
		{
			_s := *s.Uint64
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_uint64(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Complex64 == nil)
	if err != nil {
		return
	}
	if s.Complex64 != nil {
		{
			_s := *s.Complex64
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_complex64(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Complex128 == nil)
	if err != nil {
		return
	}
	if s.Complex128 != nil {
		{
			_s := *s.Complex128
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_complex128(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.String == nil)
	if err != nil {
		return
	}
	if s.String != nil {
		{
			_s := *s.String
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				err = serializer.Write_string(w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(w, _b, s.Maps == nil)
	if err != nil {
		return
	}
	if s.Maps != nil {
		{
			_s := *s.Maps
			_n = len(_s)
			err = serializer.Write_int(w, _b, _n)
			if err != nil {
				return
			}
			for _k := 0; _k < _n; _k++ {
				{
					_s := _s[_k]
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
			}
		}
	}

	return
}

func (s *SlicePtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _SlicePtrLayout)
	if err != nil {
		return
	}

	var _bool bool
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
	var _uint8 uint8

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Int = nil
	} else {
		s.Int = new([]int)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Int); _n > _c || _c-_n > _c/8 {
			*s.Int = make([]int, _n)
		} else {
			*s.Int = (*s.Int)[:_n]
		}
		if _n > 0 {
			_s := *s.Int
			for _k := 0; _k < _n; _k++ {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Int8 = nil
	} else {
		s.Int8 = new([]int8)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Int8); _n > _c || _c-_n > _c/8 {
			*s.Int8 = make([]int8, _n)
		} else {
			*s.Int8 = (*s.Int8)[:_n]
		}
		if _n > 0 {
			_s := *s.Int8
			for _k := 0; _k < _n; _k++ {
				_int8, err = serializer.Read_int8(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int8
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Int16 = nil
	} else {
		s.Int16 = new([]int16)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Int16); _n > _c || _c-_n > _c/8 {
			*s.Int16 = make([]int16, _n)
		} else {
			*s.Int16 = (*s.Int16)[:_n]
		}
		if _n > 0 {
			_s := *s.Int16
			for _k := 0; _k < _n; _k++ {
				_int16, err = serializer.Read_int16(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int16
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Int32 = nil
	} else {
		s.Int32 = new([]int32)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Int32); _n > _c || _c-_n > _c/8 {
			*s.Int32 = make([]int32, _n)
		} else {
			*s.Int32 = (*s.Int32)[:_n]
		}
		if _n > 0 {
			_s := *s.Int32
			for _k := 0; _k < _n; _k++ {
				_int32, err = serializer.Read_int32(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int32
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Int64 = nil
	} else {
		s.Int64 = new([]int64)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Int64); _n > _c || _c-_n > _c/8 {
			*s.Int64 = make([]int64, _n)
		} else {
			*s.Int64 = (*s.Int64)[:_n]
		}
		if _n > 0 {
			_s := *s.Int64
			for _k := 0; _k < _n; _k++ {
				_int64, err = serializer.Read_int64(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int64
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Uint = nil
	} else {
		s.Uint = new([]uint)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Uint); _n > _c || _c-_n > _c/8 {
			*s.Uint = make([]uint, _n)
		} else {
			*s.Uint = (*s.Uint)[:_n]
		}
		if _n > 0 {
			_s := *s.Uint
			for _k := 0; _k < _n; _k++ {
				_uint, err = serializer.Read_uint(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Uint8 = nil
	} else {
		s.Uint8 = new([]uint8)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Uint8); _n > _c || _c-_n > _c/8 {
			*s.Uint8 = make([]uint8, _n)
		} else {
			*s.Uint8 = (*s.Uint8)[:_n]
		}
		if _n > 0 {
			_s := *s.Uint8
			for _k := 0; _k < _n; _k++ {
				_uint8, err = serializer.Read_uint8(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint8
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Uint16 = nil
	} else {
		s.Uint16 = new([]uint16)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Uint16); _n > _c || _c-_n > _c/8 {
			*s.Uint16 = make([]uint16, _n)
		} else {
			*s.Uint16 = (*s.Uint16)[:_n]
		}
		if _n > 0 {
			_s := *s.Uint16
			for _k := 0; _k < _n; _k++ {
				_uint16, err = serializer.Read_uint16(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint16
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Uint32 = nil
	} else {
		s.Uint32 = new([]uint32)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Uint32); _n > _c || _c-_n > _c/8 {
			*s.Uint32 = make([]uint32, _n)
		} else {
			*s.Uint32 = (*s.Uint32)[:_n]
		}
		if _n > 0 {
			_s := *s.Uint32
			for _k := 0; _k < _n; _k++ {
				_uint32, err = serializer.Read_uint32(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint32
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Uint64 = nil
	} else {
		s.Uint64 = new([]uint64)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Uint64); _n > _c || _c-_n > _c/8 {
			*s.Uint64 = make([]uint64, _n)
		} else {
			*s.Uint64 = (*s.Uint64)[:_n]
		}
		if _n > 0 {
			_s := *s.Uint64
			for _k := 0; _k < _n; _k++ {
				_uint64, err = serializer.Read_uint64(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint64
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Complex64 = nil
	} else {
		s.Complex64 = new([]complex64)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Complex64); _n > _c || _c-_n > _c/8 {
			*s.Complex64 = make([]complex64, _n)
		} else {
			*s.Complex64 = (*s.Complex64)[:_n]
		}
		if _n > 0 {
			_s := *s.Complex64
			for _k := 0; _k < _n; _k++ {
				_complex64, err = serializer.Read_complex64(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _complex64
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Complex128 = nil
	} else {
		s.Complex128 = new([]complex128)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Complex128); _n > _c || _c-_n > _c/8 {
			*s.Complex128 = make([]complex128, _n)
		} else {
			*s.Complex128 = (*s.Complex128)[:_n]
		}
		if _n > 0 {
			_s := *s.Complex128
			for _k := 0; _k < _n; _k++ {
				_complex128, err = serializer.Read_complex128(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _complex128
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.String = nil
	} else {
		s.String = new([]string)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.String); _n > _c || _c-_n > _c/8 {
			*s.String = make([]string, _n)
		} else {
			*s.String = (*s.String)[:_n]
		}
		if _n > 0 {
			_s := *s.String
			for _k := 0; _k < _n; _k++ {
				_string, err = serializer.Read_string(r, _b)
				if err != nil {
					return
				}
				_s[_k] = _string
			}
		}
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		s.Maps = nil
	} else {
		s.Maps = new([]map[string][]int)
		_n, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		if _c := cap(*s.Maps); _n > _c || _c-_n > _c/8 {
			*s.Maps = make([]map[string][]int, _n)
		} else {
			*s.Maps = (*s.Maps)[:_n]
		}
		if _n > 0 {
			_s := *s.Maps
			for _k := 0; _k < _n; _k++ {
				_n, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				if _n == 0 {
					_s[_k] = nil
				} else {
					_s[_k] = make(map[string][]int, _n)
					_s := _s[_k]
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
						if _c := cap(_s[_k]); _n > _c || _c-_n > _c/8 {
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
			}
		}
	}

	return
}

const _SliceAnonLayout = "ZCYXZHK"

func (s *SliceAnon) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _SliceAnonLayout)
	if err != nil {
		return
	}

	var _n int

	{
		_s := &s.Anon

		err = serializer.Write_int(w, _b, _s.Int)
		if err != nil {
			return
		}

		err = serializer.Write_string(w, _b, _s.String)
		if err != nil {
			return
		}

	}

	{
		_s := s.AnonSlice
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			{
				_s := &_s[_k]

				err = serializer.Write_uint(w, _b, _s.Uint)
				if err != nil {
					return
				}

				err = serializer.Write_uint32(w, _b, _s.Uint32)
				if err != nil {
					return
				}

			}
		}
	}
	return
}

func (s *SliceAnon) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _SliceAnonLayout)
	if err != nil {
		return
	}

	var _int int
	var _n int
	var _string string
	var _uint uint
	var _uint32 uint32

	{
		_s := &s.Anon

		_int, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		_s.Int = _int

		_string, err = serializer.Read_string(r, _b)
		if err != nil {
			return
		}
		_s.String = _string

	}

	_n, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	if _c := cap(s.AnonSlice); _n > _c || _c-_n > _c/8 {
		s.AnonSlice = make([]struct {
			Uint   uint
			Uint32 uint32
		}, _n)
	} else {
		s.AnonSlice = (s.AnonSlice)[:_n]
	}
	if _n > 0 {
		_s := s.AnonSlice
		for _k := 0; _k < _n; _k++ {
			{
				_s := &_s[_k]

				_uint, err = serializer.Read_uint(r, _b)
				if err != nil {
					return
				}
				_s.Uint = _uint

				_uint32, err = serializer.Read_uint32(r, _b)
				if err != nil {
					return
				}
				_s.Uint32 = _uint32

			}
		}
	}

	return
}
