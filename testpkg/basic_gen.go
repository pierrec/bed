package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _BasicLayout = "BCDEFGHIJKLPQY"

func (b *Basic) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _BasicLayout)
	if err != nil {
		return
	}

	err = serializer.Write_bool(w, _b, b.Bool)
	if err != nil {
		return
	}

	err = serializer.Write_int(w, _b, b.Int)
	if err != nil {
		return
	}

	err = serializer.Write_int8(w, _b, b.Int8)
	if err != nil {
		return
	}

	err = serializer.Write_int16(w, _b, b.Int16)
	if err != nil {
		return
	}

	err = serializer.Write_int32(w, _b, b.Int32)
	if err != nil {
		return
	}

	err = serializer.Write_int64(w, _b, b.Int64)
	if err != nil {
		return
	}

	err = serializer.Write_uint(w, _b, b.Uint)
	if err != nil {
		return
	}

	err = serializer.Write_uint8(w, _b, b.Uint8)
	if err != nil {
		return
	}

	err = serializer.Write_uint16(w, _b, b.Uint16)
	if err != nil {
		return
	}

	err = serializer.Write_uint32(w, _b, b.Uint32)
	if err != nil {
		return
	}

	err = serializer.Write_uint64(w, _b, b.Uint64)
	if err != nil {
		return
	}

	err = serializer.Write_complex64(w, _b, b.Complex64)
	if err != nil {
		return
	}

	err = serializer.Write_complex128(w, _b, b.Complex128)
	if err != nil {
		return
	}

	err = serializer.Write_string(w, _b, b.String)
	if err != nil {
		return
	}

	return
}

func (b *Basic) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _BasicLayout)
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
	b.Bool = _bool

	_int, err = serializer.Read_int(r, _b)
	if err != nil {
		return
	}
	b.Int = _int

	_int8, err = serializer.Read_int8(r, _b)
	if err != nil {
		return
	}
	b.Int8 = _int8

	_int16, err = serializer.Read_int16(r, _b)
	if err != nil {
		return
	}
	b.Int16 = _int16

	_int32, err = serializer.Read_int32(r, _b)
	if err != nil {
		return
	}
	b.Int32 = _int32

	_int64, err = serializer.Read_int64(r, _b)
	if err != nil {
		return
	}
	b.Int64 = _int64

	_uint, err = serializer.Read_uint(r, _b)
	if err != nil {
		return
	}
	b.Uint = _uint

	_uint8, err = serializer.Read_uint8(r, _b)
	if err != nil {
		return
	}
	b.Uint8 = _uint8

	_uint16, err = serializer.Read_uint16(r, _b)
	if err != nil {
		return
	}
	b.Uint16 = _uint16

	_uint32, err = serializer.Read_uint32(r, _b)
	if err != nil {
		return
	}
	b.Uint32 = _uint32

	_uint64, err = serializer.Read_uint64(r, _b)
	if err != nil {
		return
	}
	b.Uint64 = _uint64

	_complex64, err = serializer.Read_complex64(r, _b)
	if err != nil {
		return
	}
	b.Complex64 = _complex64

	_complex128, err = serializer.Read_complex128(r, _b)
	if err != nil {
		return
	}
	b.Complex128 = _complex128

	_string, err = serializer.Read_string(r, _b)
	if err != nil {
		return
	}
	b.String = _string

	return
}

const _BasicPtrLayout = "WBWCWDWEWFWGWHWIWJWKWLWPWQWY"

func (b *BasicPtr) MarshalBinaryTo(w io.Writer) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Write_layout(w, _b, _BasicPtrLayout)
	if err != nil {
		return
	}

	err = serializer.Write_bool(w, _b, b.Bool == nil)
	if err != nil {
		return
	}
	if b.Bool != nil {
		err = serializer.Write_bool(w, _b, *b.Bool)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Int == nil)
	if err != nil {
		return
	}
	if b.Int != nil {
		err = serializer.Write_int(w, _b, *b.Int)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Int8 == nil)
	if err != nil {
		return
	}
	if b.Int8 != nil {
		err = serializer.Write_int8(w, _b, *b.Int8)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Int16 == nil)
	if err != nil {
		return
	}
	if b.Int16 != nil {
		err = serializer.Write_int16(w, _b, *b.Int16)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Int32 == nil)
	if err != nil {
		return
	}
	if b.Int32 != nil {
		err = serializer.Write_int32(w, _b, *b.Int32)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Int64 == nil)
	if err != nil {
		return
	}
	if b.Int64 != nil {
		err = serializer.Write_int64(w, _b, *b.Int64)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Uint == nil)
	if err != nil {
		return
	}
	if b.Uint != nil {
		err = serializer.Write_uint(w, _b, *b.Uint)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Uint8 == nil)
	if err != nil {
		return
	}
	if b.Uint8 != nil {
		err = serializer.Write_uint8(w, _b, *b.Uint8)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Uint16 == nil)
	if err != nil {
		return
	}
	if b.Uint16 != nil {
		err = serializer.Write_uint16(w, _b, *b.Uint16)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Uint32 == nil)
	if err != nil {
		return
	}
	if b.Uint32 != nil {
		err = serializer.Write_uint32(w, _b, *b.Uint32)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Uint64 == nil)
	if err != nil {
		return
	}
	if b.Uint64 != nil {
		err = serializer.Write_uint64(w, _b, *b.Uint64)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Complex64 == nil)
	if err != nil {
		return
	}
	if b.Complex64 != nil {
		err = serializer.Write_complex64(w, _b, *b.Complex64)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Complex128 == nil)
	if err != nil {
		return
	}
	if b.Complex128 != nil {
		err = serializer.Write_complex128(w, _b, *b.Complex128)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.String == nil)
	if err != nil {
		return
	}
	if b.String != nil {
		err = serializer.Write_string(w, _b, *b.String)
		if err != nil {
			return
		}
	}

	return
}

func (b *BasicPtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var _buf [16]byte
	_b := _buf[:]
	err = serializer.Read_layout(r, _b, _BasicPtrLayout)
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
		b.Bool = nil
	} else {
		b.Bool = new(bool)
		_bool, err = serializer.Read_bool(r, _b)
		if err != nil {
			return
		}
		*b.Bool = _bool
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int = nil
	} else {
		b.Int = new(int)
		_int, err = serializer.Read_int(r, _b)
		if err != nil {
			return
		}
		*b.Int = _int
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int8 = nil
	} else {
		b.Int8 = new(int8)
		_int8, err = serializer.Read_int8(r, _b)
		if err != nil {
			return
		}
		*b.Int8 = _int8
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int16 = nil
	} else {
		b.Int16 = new(int16)
		_int16, err = serializer.Read_int16(r, _b)
		if err != nil {
			return
		}
		*b.Int16 = _int16
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int32 = nil
	} else {
		b.Int32 = new(int32)
		_int32, err = serializer.Read_int32(r, _b)
		if err != nil {
			return
		}
		*b.Int32 = _int32
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int64 = nil
	} else {
		b.Int64 = new(int64)
		_int64, err = serializer.Read_int64(r, _b)
		if err != nil {
			return
		}
		*b.Int64 = _int64
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Uint = nil
	} else {
		b.Uint = new(uint)
		_uint, err = serializer.Read_uint(r, _b)
		if err != nil {
			return
		}
		*b.Uint = _uint
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Uint8 = nil
	} else {
		b.Uint8 = new(uint8)
		_uint8, err = serializer.Read_uint8(r, _b)
		if err != nil {
			return
		}
		*b.Uint8 = _uint8
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Uint16 = nil
	} else {
		b.Uint16 = new(uint16)
		_uint16, err = serializer.Read_uint16(r, _b)
		if err != nil {
			return
		}
		*b.Uint16 = _uint16
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Uint32 = nil
	} else {
		b.Uint32 = new(uint32)
		_uint32, err = serializer.Read_uint32(r, _b)
		if err != nil {
			return
		}
		*b.Uint32 = _uint32
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Uint64 = nil
	} else {
		b.Uint64 = new(uint64)
		_uint64, err = serializer.Read_uint64(r, _b)
		if err != nil {
			return
		}
		*b.Uint64 = _uint64
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Complex64 = nil
	} else {
		b.Complex64 = new(complex64)
		_complex64, err = serializer.Read_complex64(r, _b)
		if err != nil {
			return
		}
		*b.Complex64 = _complex64
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Complex128 = nil
	} else {
		b.Complex128 = new(complex128)
		_complex128, err = serializer.Read_complex128(r, _b)
		if err != nil {
			return
		}
		*b.Complex128 = _complex128
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.String = nil
	} else {
		b.String = new(string)
		_string, err = serializer.Read_string(r, _b)
		if err != nil {
			return
		}
		*b.String = _string
	}

	return
}
