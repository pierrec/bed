package testpkg

import "io"
import "github.com/pierrec/serializer"

func (self *Array) MarshalBinaryTo(w io.Writer) (err error) {
	var __buf [16]byte
	var _b = __buf[:]

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
	err = serializer.Write_bytea(w, self.Uint8[:])
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
	var __buf [16]byte
	var _b = __buf[:]
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

	err = serializer.Read_bytea(r, self.Uint8[:])
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
