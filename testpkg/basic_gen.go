package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

func (self *Basic) MarshalBinaryTo(w io.Writer) (err error) {
	var buf [16]byte
	b := buf[:]

	err = serializer.Write_int(w, b, self.Int)
	if err != nil {
		return
	}

	err = serializer.Write_int8(w, b, self.Int8)
	if err != nil {
		return
	}

	err = serializer.Write_int16(w, b, self.Int16)
	if err != nil {
		return
	}

	err = serializer.Write_int32(w, b, self.Int32)
	if err != nil {
		return
	}

	err = serializer.Write_int64(w, b, self.Int64)
	if err != nil {
		return
	}

	err = serializer.Write_uint(w, b, self.Uint)
	if err != nil {
		return
	}

	err = serializer.Write_uint8(w, b, self.Uint8)
	if err != nil {
		return
	}

	err = serializer.Write_uint16(w, b, self.Uint16)
	if err != nil {
		return
	}

	err = serializer.Write_uint32(w, b, self.Uint32)
	if err != nil {
		return
	}

	err = serializer.Write_uint64(w, b, self.Uint64)
	if err != nil {
		return
	}

	err = serializer.Write_complex64(w, b, self.Complex64)
	if err != nil {
		return
	}

	err = serializer.Write_complex128(w, b, self.Complex128)
	if err != nil {
		return
	}

	err = serializer.Write_string(w, b, self.String)
	if err != nil {
		return
	}

	return
}

func (self *Basic) UnmarshalBinaryFrom(r io.Reader) (err error) {
	var buf [16]byte
	b := buf[:]
	var _int16 int16
	var _uint uint
	var _uint16 uint16
	var _complex64 complex64
	var _int8 int8
	var _int32 int32
	var _int64 int64
	var _uint8 uint8
	_ = _uint8
	var _uint32 uint32
	var _uint64 uint64
	var _complex128 complex128
	var _string string
	var _int int

	_int, err = serializer.Read_int(r, b)
	if err != nil {
		return
	}
	self.Int = _int

	_int8, err = serializer.Read_int8(r, b)
	if err != nil {
		return
	}
	self.Int8 = _int8

	_int16, err = serializer.Read_int16(r, b)
	if err != nil {
		return
	}
	self.Int16 = _int16

	_int32, err = serializer.Read_int32(r, b)
	if err != nil {
		return
	}
	self.Int32 = _int32

	_int64, err = serializer.Read_int64(r, b)
	if err != nil {
		return
	}
	self.Int64 = _int64

	_uint, err = serializer.Read_uint(r, b)
	if err != nil {
		return
	}
	self.Uint = _uint

	_uint8, err = serializer.Read_uint8(r, b)
	if err != nil {
		return
	}
	self.Uint8 = _uint8

	_uint16, err = serializer.Read_uint16(r, b)
	if err != nil {
		return
	}
	self.Uint16 = _uint16

	_uint32, err = serializer.Read_uint32(r, b)
	if err != nil {
		return
	}
	self.Uint32 = _uint32

	_uint64, err = serializer.Read_uint64(r, b)
	if err != nil {
		return
	}
	self.Uint64 = _uint64

	_complex64, err = serializer.Read_complex64(r, b)
	if err != nil {
		return
	}
	self.Complex64 = _complex64

	_complex128, err = serializer.Read_complex128(r, b)
	if err != nil {
		return
	}
	self.Complex128 = _complex128

	_string, err = serializer.Read_string(r, b)
	if err != nil {
		return
	}
	self.String = _string

	return
}
