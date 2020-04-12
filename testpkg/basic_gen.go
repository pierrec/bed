package testpkg

import (
	"io"
	"strings"

	"github.com/pierrec/serializer"
)

func (self *Basic) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "BCDEFGHIJKLPQY"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	err = serializer.Write_bool(w, _b, self.Bool)
	if err != nil {
		return
	}

	err = serializer.Write_int(w, _b, self.Int)
	if err != nil {
		return
	}

	err = serializer.Write_int8(w, _b, self.Int8)
	if err != nil {
		return
	}

	err = serializer.Write_int16(w, _b, self.Int16)
	if err != nil {
		return
	}

	err = serializer.Write_int32(w, _b, self.Int32)
	if err != nil {
		return
	}

	err = serializer.Write_int64(w, _b, self.Int64)
	if err != nil {
		return
	}

	err = serializer.Write_uint(w, _b, self.Uint)
	if err != nil {
		return
	}

	err = serializer.Write_uint8(w, _b, self.Uint8)
	if err != nil {
		return
	}

	err = serializer.Write_uint16(w, _b, self.Uint16)
	if err != nil {
		return
	}

	err = serializer.Write_uint32(w, _b, self.Uint32)
	if err != nil {
		return
	}

	err = serializer.Write_uint64(w, _b, self.Uint64)
	if err != nil {
		return
	}

	err = serializer.Write_complex64(w, _b, self.Complex64)
	if err != nil {
		return
	}

	err = serializer.Write_complex128(w, _b, self.Complex128)
	if err != nil {
		return
	}

	err = serializer.Write_string(w, _b, self.String)
	if err != nil {
		return
	}

	return
}

func (self *Basic) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "BCDEFGHIJKLPQY"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _bool bool
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
	var _uint8 uint8

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	self.Bool = _bool

	_int, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	self.Int = _int

	_int8, err = serializer.Read_int8(r, _b)
	if err != nil {
		return
	}
	self.Int8 = _int8

	_int16, err = serializer.Read_int16(r, _b)
	if err != nil {
		return
	}
	self.Int16 = _int16

	_int32, err = serializer.Read_int32(r, _b)
	if err != nil {
		return
	}
	self.Int32 = _int32

	_int64, err = serializer.Read_int64(r, _b)
	if err != nil {
		return
	}
	self.Int64 = _int64

	_uint, err = serializer.Read_uint(r, _b)
	if err != nil {
		return
	}
	self.Uint = _uint

	_uint8, err = serializer.Read_uint8(r, _b)
	if err != nil {
		return
	}
	self.Uint8 = _uint8

	_uint16, err = serializer.Read_uint16(r, _b)
	if err != nil {
		return
	}
	self.Uint16 = _uint16

	_uint32, err = serializer.Read_uint32(r, _b)
	if err != nil {
		return
	}
	self.Uint32 = _uint32

	_uint64, err = serializer.Read_uint64(r, _b)
	if err != nil {
		return
	}
	self.Uint64 = _uint64

	_complex64, err = serializer.Read_complex64(r, _b)
	if err != nil {
		return
	}
	self.Complex64 = _complex64

	_complex128, err = serializer.Read_complex128(r, _b)
	if err != nil {
		return
	}
	self.Complex128 = _complex128

	_string, err = serializer.Read_string(r, _b)
	if err != nil {
		return
	}
	self.String = _string

	return
}

func (self *BasicPtr) MarshalBinaryTo(w io.Writer) (err error) {
	const _check = "WBWCWDWEWFWGWHWIWJWKWLWPWQWY"
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_string(w, _b, _check)
	if err != nil {
		return
	}

	err = serializer.Write_bool(w, _b, self.Bool == nil)
	if err != nil {
		return
	}
	if self.Bool != nil {
		err = serializer.Write_bool(w, _b, *self.Bool)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Int == nil)
	if err != nil {
		return
	}
	if self.Int != nil {
		err = serializer.Write_int(w, _b, *self.Int)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Int8 == nil)
	if err != nil {
		return
	}
	if self.Int8 != nil {
		err = serializer.Write_int8(w, _b, *self.Int8)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Int16 == nil)
	if err != nil {
		return
	}
	if self.Int16 != nil {
		err = serializer.Write_int16(w, _b, *self.Int16)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Int32 == nil)
	if err != nil {
		return
	}
	if self.Int32 != nil {
		err = serializer.Write_int32(w, _b, *self.Int32)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Int64 == nil)
	if err != nil {
		return
	}
	if self.Int64 != nil {
		err = serializer.Write_int64(w, _b, *self.Int64)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint == nil)
	if err != nil {
		return
	}
	if self.Uint != nil {
		err = serializer.Write_uint(w, _b, *self.Uint)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint8 == nil)
	if err != nil {
		return
	}
	if self.Uint8 != nil {
		err = serializer.Write_uint8(w, _b, *self.Uint8)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint16 == nil)
	if err != nil {
		return
	}
	if self.Uint16 != nil {
		err = serializer.Write_uint16(w, _b, *self.Uint16)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint32 == nil)
	if err != nil {
		return
	}
	if self.Uint32 != nil {
		err = serializer.Write_uint32(w, _b, *self.Uint32)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Uint64 == nil)
	if err != nil {
		return
	}
	if self.Uint64 != nil {
		err = serializer.Write_uint64(w, _b, *self.Uint64)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Complex64 == nil)
	if err != nil {
		return
	}
	if self.Complex64 != nil {
		err = serializer.Write_complex64(w, _b, *self.Complex64)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.Complex128 == nil)
	if err != nil {
		return
	}
	if self.Complex128 != nil {
		err = serializer.Write_complex128(w, _b, *self.Complex128)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, self.String == nil)
	if err != nil {
		return
	}
	if self.String != nil {
		err = serializer.Write_string(w, _b, *self.String)
		if err != nil {
			return
		}
	}

	return
}

func (self *BasicPtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	const _check = "WBWCWDWEWFWGWHWIWJWKWLWPWQWY"
	var _buf [16]byte
	_b := _buf[:]
	if s, err := serializer.Read_string(r, _b); err != nil {
		return err
	} else if !strings.HasPrefix(s, _check) {
		return serializer.ErrInvalidData
	}

	var _bool bool
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
	var _uint8 uint8

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Bool = nil
	} else {
		_bool, err = serializer.Read_bool(r, _b)
		if err != nil {
			return
		}
		*self.Bool = _bool
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Int = nil
	} else {
		_int, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		*self.Int = _int
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Int8 = nil
	} else {
		_int8, err = serializer.Read_int8(r, _b)
		if err != nil {
			return
		}
		*self.Int8 = _int8
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Int16 = nil
	} else {
		_int16, err = serializer.Read_int16(r, _b)
		if err != nil {
			return
		}
		*self.Int16 = _int16
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Int32 = nil
	} else {
		_int32, err = serializer.Read_int32(r, _b)
		if err != nil {
			return
		}
		*self.Int32 = _int32
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Int64 = nil
	} else {
		_int64, err = serializer.Read_int64(r, _b)
		if err != nil {
			return
		}
		*self.Int64 = _int64
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Uint = nil
	} else {
		_uint, err = serializer.Read_uint(r, _b)
		if err != nil {
			return
		}
		*self.Uint = _uint
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Uint8 = nil
	} else {
		_uint8, err = serializer.Read_uint8(r, _b)
		if err != nil {
			return
		}
		*self.Uint8 = _uint8
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Uint16 = nil
	} else {
		_uint16, err = serializer.Read_uint16(r, _b)
		if err != nil {
			return
		}
		*self.Uint16 = _uint16
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Uint32 = nil
	} else {
		_uint32, err = serializer.Read_uint32(r, _b)
		if err != nil {
			return
		}
		*self.Uint32 = _uint32
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Uint64 = nil
	} else {
		_uint64, err = serializer.Read_uint64(r, _b)
		if err != nil {
			return
		}
		*self.Uint64 = _uint64
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Complex64 = nil
	} else {
		_complex64, err = serializer.Read_complex64(r, _b)
		if err != nil {
			return
		}
		*self.Complex64 = _complex64
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.Complex128 = nil
	} else {
		_complex128, err = serializer.Read_complex128(r, _b)
		if err != nil {
			return
		}
		*self.Complex128 = _complex128
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		self.String = nil
	} else {
		_string, err = serializer.Read_string(r, _b)
		if err != nil {
			return
		}
		*self.String = _string
	}

	return
}
