package testpkg

import (
	"io"

	"github.com/pierrec/serializer"
)

const _ArrayLayout = "R4CR4DR4ER4FR4GR4HR4R4JR4KR4LR4PR4QR4Y"

func (a *Array) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _ArrayLayout)
	if err != nil {
		return
	}

	{
		_s := &a.Int
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_int(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Int8
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_int8(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Int16
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_int16(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Int32
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_int32(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Int64
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_int64(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Uint
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_uint(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	err = serializer.Write_bytea(_w, (a.Uint8)[:])
	if err != nil {
		return
	}

	{
		_s := &a.Uint16
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_uint16(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Uint32
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_uint32(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Uint64
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_uint64(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Complex64
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_complex64(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.Complex128
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_complex128(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &a.String
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = serializer.Write_string(_w, _b, _s[_k])
			if err != nil {
				return
			}
		}
	}
	return
}

func (a *Array) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _ArrayLayout)
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
	var _string string
	var _uint uint
	var _uint16 uint16
	var _uint32 uint32
	var _uint64 uint64

	{
		_s := &a.Int
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_int, err = serializer.Read_int(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int
		}
	}

	{
		_s := &a.Int8
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_int8, err = serializer.Read_int8(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int8
		}
	}

	{
		_s := &a.Int16
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_int16, err = serializer.Read_int16(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int16
		}
	}

	{
		_s := &a.Int32
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_int32, err = serializer.Read_int32(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int32
		}
	}

	{
		_s := &a.Int64
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_int64, err = serializer.Read_int64(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _int64
		}
	}

	{
		_s := &a.Uint
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_uint, err = serializer.Read_uint(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint
		}
	}

	err = serializer.Read_bytea(_r, (a.Uint8)[:])
	if err != nil {
		return
	}

	{
		_s := &a.Uint16
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_uint16, err = serializer.Read_uint16(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint16
		}
	}

	{
		_s := &a.Uint32
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_uint32, err = serializer.Read_uint32(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint32
		}
	}

	{
		_s := &a.Uint64
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_uint64, err = serializer.Read_uint64(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _uint64
		}
	}

	{
		_s := &a.Complex64
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_complex64, err = serializer.Read_complex64(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _complex64
		}
	}

	{
		_s := &a.Complex128
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_complex128, err = serializer.Read_complex128(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _complex128
		}
	}

	{
		_s := &a.String
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_string, err = serializer.Read_string(_r, _b)
			if err != nil {
				return
			}
			_s[_k] = _string
		}
	}

	return
}

const _ArrayPtrLayout = "WR4CWR4DWR4EWR4FWR4GWR4HWR4WR4JWR4KWR4LWR4PWR4QWR4Y"

func (a *ArrayPtr) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _ArrayPtrLayout)
	if err != nil {
		return
	}

	err = serializer.Write_bool(_w, _b, a.Int == nil)
	if err != nil {
		return
	}
	if a.Int != nil {
		{
			_s := &*a.Int
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_int(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Int8 == nil)
	if err != nil {
		return
	}
	if a.Int8 != nil {
		{
			_s := &*a.Int8
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_int8(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Int16 == nil)
	if err != nil {
		return
	}
	if a.Int16 != nil {
		{
			_s := &*a.Int16
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_int16(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Int32 == nil)
	if err != nil {
		return
	}
	if a.Int32 != nil {
		{
			_s := &*a.Int32
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_int32(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Int64 == nil)
	if err != nil {
		return
	}
	if a.Int64 != nil {
		{
			_s := &*a.Int64
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_int64(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Uint == nil)
	if err != nil {
		return
	}
	if a.Uint != nil {
		{
			_s := &*a.Uint
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_uint(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Uint8 == nil)
	if err != nil {
		return
	}
	if a.Uint8 != nil {
		err = serializer.Write_bytea(_w, (*a.Uint8)[:])
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(_w, _b, a.Uint16 == nil)
	if err != nil {
		return
	}
	if a.Uint16 != nil {
		{
			_s := &*a.Uint16
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_uint16(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Uint32 == nil)
	if err != nil {
		return
	}
	if a.Uint32 != nil {
		{
			_s := &*a.Uint32
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_uint32(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Uint64 == nil)
	if err != nil {
		return
	}
	if a.Uint64 != nil {
		{
			_s := &*a.Uint64
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_uint64(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Complex64 == nil)
	if err != nil {
		return
	}
	if a.Complex64 != nil {
		{
			_s := &*a.Complex64
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_complex64(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.Complex128 == nil)
	if err != nil {
		return
	}
	if a.Complex128 != nil {
		{
			_s := &*a.Complex128
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_complex128(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	err = serializer.Write_bool(_w, _b, a.String == nil)
	if err != nil {
		return
	}
	if a.String != nil {
		{
			_s := &*a.String
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				err = serializer.Write_string(_w, _b, _s[_k])
				if err != nil {
					return
				}
			}
		}
	}

	return
}

func (a *ArrayPtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _ArrayPtrLayout)
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

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Int = nil
	} else {
		a.Int = new([4]int)
		{
			_s := &*a.Int
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_int, err = serializer.Read_int(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Int8 = nil
	} else {
		a.Int8 = new([4]int8)
		{
			_s := &*a.Int8
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_int8, err = serializer.Read_int8(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int8
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Int16 = nil
	} else {
		a.Int16 = new([4]int16)
		{
			_s := &*a.Int16
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_int16, err = serializer.Read_int16(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int16
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Int32 = nil
	} else {
		a.Int32 = new([4]int32)
		{
			_s := &*a.Int32
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_int32, err = serializer.Read_int32(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int32
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Int64 = nil
	} else {
		a.Int64 = new([4]int64)
		{
			_s := &*a.Int64
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_int64, err = serializer.Read_int64(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _int64
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Uint = nil
	} else {
		a.Uint = new([4]uint)
		{
			_s := &*a.Uint
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_uint, err = serializer.Read_uint(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Uint8 = nil
	} else {
		a.Uint8 = new([4]uint8)
		err = serializer.Read_bytea(_r, (*a.Uint8)[:])
		if err != nil {
			return
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Uint16 = nil
	} else {
		a.Uint16 = new([4]uint16)
		{
			_s := &*a.Uint16
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_uint16, err = serializer.Read_uint16(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint16
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Uint32 = nil
	} else {
		a.Uint32 = new([4]uint32)
		{
			_s := &*a.Uint32
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_uint32, err = serializer.Read_uint32(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint32
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Uint64 = nil
	} else {
		a.Uint64 = new([4]uint64)
		{
			_s := &*a.Uint64
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_uint64, err = serializer.Read_uint64(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _uint64
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Complex64 = nil
	} else {
		a.Complex64 = new([4]complex64)
		{
			_s := &*a.Complex64
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_complex64, err = serializer.Read_complex64(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _complex64
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.Complex128 = nil
	} else {
		a.Complex128 = new([4]complex128)
		{
			_s := &*a.Complex128
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_complex128, err = serializer.Read_complex128(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _complex128
			}
		}
	}

	_bool, err = serializer.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		a.String = nil
	} else {
		a.String = new([4]string)
		{
			_s := &*a.String
			for _k, _kn := 0, len(_s); _k < _kn; _k++ {
				_string, err = serializer.Read_string(_r, _b)
				if err != nil {
					return
				}
				_s[_k] = _string
			}
		}
	}

	return
}

const _ArrayAnonLayout = "R4ZCYR4XZHK"

func (a *ArrayAnon) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := serializer.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(_w, _b, _ArrayAnonLayout)
	if err != nil {
		return
	}

	var _n int

	{
		_s := &a.Anon
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			{
				_s := &_s[_k]

				err = serializer.Write_int(_w, _b, _s.Int)
				if err != nil {
					return
				}

				err = serializer.Write_string(_w, _b, _s.String)
				if err != nil {
					return
				}

			}
		}
	}
	{
		_s := &a.AnonSlice
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			{
				_s := _s[_k]
				_n = len(_s)
				err = serializer.Write_len(_w, _b, _n)
				if err != nil {
					return
				}
				for _k, _kn := 0, _n; _k < _kn; _k++ {
					{
						_s := &_s[_k]

						err = serializer.Write_uint(_w, _b, _s.Uint)
						if err != nil {
							return
						}

						err = serializer.Write_uint32(_w, _b, _s.Uint32)
						if err != nil {
							return
						}

					}
				}
			}
		}
	}
	return
}

func (a *ArrayAnon) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := serializer.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(_r, _b, _ArrayAnonLayout)
	if err != nil {
		return
	}

	var _int int
	var _n int
	var _string string
	var _uint uint
	var _uint32 uint32

	{
		_s := &a.Anon
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			{
				_s := &_s[_k]

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
		}
	}

	{
		_s := &a.AnonSlice
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_n, err = serializer.Read_len(_r)
			if err != nil {
				return
			}
			if _c := cap(_s[_k]); _n > _c || _c-_n > _c/8 {
				_s[_k] = make([]struct {
					Uint   uint
					Uint32 uint32
				}, _n)
			} else {
				_s[_k] = (_s[_k])[:_n]
			}
			if _n > 0 {
				_s := _s[_k]
				for _k, _kn := 0, _n; _k < _kn; _k++ {
					{
						_s := &_s[_k]

						_uint, err = serializer.Read_uint(_r, _b)
						if err != nil {
							return
						}
						_s.Uint = _uint

						_uint32, err = serializer.Read_uint32(_r, _b)
						if err != nil {
							return
						}
						_s.Uint32 = _uint32

					}
				}
			}
		}
	}

	return
}
