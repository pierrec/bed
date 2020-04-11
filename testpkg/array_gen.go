package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *Array) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "RCRDRERFRGRHRIRJRKRLRPRQRY"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	{
		_s := &self.Int
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_int(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Int8
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_int8(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Int16
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_int16(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Int32
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_int32(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Int64
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_int64(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Uint
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_uint(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	err = serializer.Write_bytea(w, (self.Uint8)[:])
	if err != nil {
		return
	}

	{
		_s := &self.Uint16
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_uint16(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Uint32
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_uint32(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Uint64
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_uint64(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Complex64
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_complex64(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.Complex128
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_complex128(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &self.String
		for _i := 0; _i < len(_s); _i++ {
			err = serializer.Write_string(w, _b, _s[_i])
			if err != nil {
				return
			}
		}
	}
	return
}

func (self *Array) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "RCRDRERFRGRHRIRJRKRLRPRQRY"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _complex128 complex128
	var _complex64 complex64
	var _int int
	var _int16 int16
	var _int32 int32
	var _int64 int64
	var _int8 int8
	var _string string
	var _uint uint
	var _uint16 uint16
	var _uint32 uint32
	var _uint64 uint64

	{
		_s := &self.Int
		for _i := 0; _i < len(_s); _i++ {
			_int, err = serializer.Read_int(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int
		}
	}

	{
		_s := &self.Int8
		for _i := 0; _i < len(_s); _i++ {
			_int8, err = serializer.Read_int8(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int8
		}
	}

	{
		_s := &self.Int16
		for _i := 0; _i < len(_s); _i++ {
			_int16, err = serializer.Read_int16(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int16
		}
	}

	{
		_s := &self.Int32
		for _i := 0; _i < len(_s); _i++ {
			_int32, err = serializer.Read_int32(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int32
		}
	}

	{
		_s := &self.Int64
		for _i := 0; _i < len(_s); _i++ {
			_int64, err = serializer.Read_int64(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _int64
		}
	}

	{
		_s := &self.Uint
		for _i := 0; _i < len(_s); _i++ {
			_uint, err = serializer.Read_uint(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint
		}
	}

	err = serializer.Read_bytea(r, (self.Uint8)[:])
	if err != nil {
		return
	}

	{
		_s := &self.Uint16
		for _i := 0; _i < len(_s); _i++ {
			_uint16, err = serializer.Read_uint16(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint16
		}
	}

	{
		_s := &self.Uint32
		for _i := 0; _i < len(_s); _i++ {
			_uint32, err = serializer.Read_uint32(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint32
		}
	}

	{
		_s := &self.Uint64
		for _i := 0; _i < len(_s); _i++ {
			_uint64, err = serializer.Read_uint64(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _uint64
		}
	}

	{
		_s := &self.Complex64
		for _i := 0; _i < len(_s); _i++ {
			_complex64, err = serializer.Read_complex64(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _complex64
		}
	}

	{
		_s := &self.Complex128
		for _i := 0; _i < len(_s); _i++ {
			_complex128, err = serializer.Read_complex128(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _complex128
		}
	}

	{
		_s := &self.String
		for _i := 0; _i < len(_s); _i++ {
			_string, err = serializer.Read_string(r, _b)
			if err != nil {
				return
			}
			_s[_i] = _string
		}
	}

	return
}

func (self *ArrayPtr) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "WRCWRDWREWRFWRGWRHWRIWRJWRKWRLWRPWRQWRY"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	err = serializer.Write_bool(w, _b, self.Int == nil)
	if err != nil {
		return
	}
	if self.Int != nil {
		{
			_s := &*self.Int
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Int8
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Int16
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Int32
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Int64
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Uint
			for _i := 0; _i < len(_s); _i++ {
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
		err = serializer.Write_bytea(w, (*self.Uint8)[:])
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
			_s := &*self.Uint16
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Uint32
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Uint64
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Complex64
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.Complex128
			for _i := 0; _i < len(_s); _i++ {
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
			_s := &*self.String
			for _i := 0; _i < len(_s); _i++ {
				err = serializer.Write_string(w, _b, _s[_i])
				if err != nil {
					return
				}
			}
		}
	}

	return
}

func (self *ArrayPtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "WRCWRDWREWRFWRGWRHWRIWRJWRKWRLWRPWRQWRY"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _complex128 complex128
	var _complex64 complex64
	var _int int
	var _int16 int16
	var _int32 int32
	var _int64 int64
	var _int8 int8
	var _string string
	var _uint uint
	var _uint16 uint16
	var _uint32 uint32
	var _uint64 uint64

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Int = nil
	} else {
		*self.Int = [4]int{}

		{
			_s := &*self.Int
			for _i := 0; _i < len(_s); _i++ {
				_int, err = serializer.Read_int(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Int8 = nil
	} else {
		*self.Int8 = [4]int8{}

		{
			_s := &*self.Int8
			for _i := 0; _i < len(_s); _i++ {
				_int8, err = serializer.Read_int8(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int8
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Int16 = nil
	} else {
		*self.Int16 = [4]int16{}

		{
			_s := &*self.Int16
			for _i := 0; _i < len(_s); _i++ {
				_int16, err = serializer.Read_int16(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int16
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Int32 = nil
	} else {
		*self.Int32 = [4]int32{}

		{
			_s := &*self.Int32
			for _i := 0; _i < len(_s); _i++ {
				_int32, err = serializer.Read_int32(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int32
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Int64 = nil
	} else {
		*self.Int64 = [4]int64{}

		{
			_s := &*self.Int64
			for _i := 0; _i < len(_s); _i++ {
				_int64, err = serializer.Read_int64(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _int64
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Uint = nil
	} else {
		*self.Uint = [4]uint{}

		{
			_s := &*self.Uint
			for _i := 0; _i < len(_s); _i++ {
				_uint, err = serializer.Read_uint(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Uint8 = nil
	} else {

		err = serializer.Read_bytea(r, (*self.Uint8)[:])
		if err != nil {
			return
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Uint16 = nil
	} else {
		*self.Uint16 = [4]uint16{}

		{
			_s := &*self.Uint16
			for _i := 0; _i < len(_s); _i++ {
				_uint16, err = serializer.Read_uint16(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint16
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Uint32 = nil
	} else {
		*self.Uint32 = [4]uint32{}

		{
			_s := &*self.Uint32
			for _i := 0; _i < len(_s); _i++ {
				_uint32, err = serializer.Read_uint32(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint32
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Uint64 = nil
	} else {
		*self.Uint64 = [4]uint64{}

		{
			_s := &*self.Uint64
			for _i := 0; _i < len(_s); _i++ {
				_uint64, err = serializer.Read_uint64(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _uint64
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Complex64 = nil
	} else {
		*self.Complex64 = [4]complex64{}

		{
			_s := &*self.Complex64
			for _i := 0; _i < len(_s); _i++ {
				_complex64, err = serializer.Read_complex64(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _complex64
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.Complex128 = nil
	} else {
		*self.Complex128 = [4]complex128{}

		{
			_s := &*self.Complex128
			for _i := 0; _i < len(_s); _i++ {
				_complex128, err = serializer.Read_complex128(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _complex128
			}
		}
	}

	if isNil, _e := serializer.Read_bool(r, _b); _e != nil {
		return _e
	} else if isNil {
		self.String = nil
	} else {
		*self.String = [4]string{}

		{
			_s := &*self.String
			for _i := 0; _i < len(_s); _i++ {
				_string, err = serializer.Read_string(r, _b)
				if err != nil {
					return
				}
				_s[_i] = _string
			}
		}
	}

	return
}
