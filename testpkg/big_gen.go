package testpkg

import (
	"github.com/pierrec/packer/iobyte"
	"github.com/pierrec/serializer"
	readwrite "github.com/pierrec/serializer/packed"
	"io"
	"math/big"
)

const _BigLayout = "ZZZ"

func (b *Big) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Write_layout(_w, _b, _BigLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)

	err = readwrite.Write_bigfloat(_w, _b, _bb, b.Float)
	if err != nil {
		return
	}

	err = readwrite.Write_bigint(_w, _b, _bb, b.Int)
	if err != nil {
		return
	}

	err = readwrite.Write_bigrat(_w, _b, _bb, b.Rat)
	if err != nil {
		return
	}

	return
}

func (b *Big) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Read_layout(_r, _b, _BigLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat

	_bigfloat, err = readwrite.Read_bigfloat(_r, _b, _bb)
	if err != nil {
		return
	}
	b.Float = _bigfloat

	_bigint, err = readwrite.Read_bigint(_r, _b, _bb)
	if err != nil {
		return
	}
	b.Int = _bigint

	_bigrat, err = readwrite.Read_bigrat(_r, _b, _bb)
	if err != nil {
		return
	}
	b.Rat = _bigrat

	return
}

const _BigPtrLayout = "WZWZWZ"

func (b *BigPtr) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Write_layout(_w, _b, _BigPtrLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)

	err = readwrite.Write_bool(_w, _b, b.Float == nil)
	if err != nil {
		return
	}
	if b.Float != nil {
		err = readwrite.Write_bigfloat(_w, _b, _bb, *b.Float)
		if err != nil {
			return
		}
	}

	err = readwrite.Write_bool(_w, _b, b.Int == nil)
	if err != nil {
		return
	}
	if b.Int != nil {
		err = readwrite.Write_bigint(_w, _b, _bb, *b.Int)
		if err != nil {
			return
		}
	}

	err = readwrite.Write_bool(_w, _b, b.Rat == nil)
	if err != nil {
		return
	}
	if b.Rat != nil {
		err = readwrite.Write_bigrat(_w, _b, _bb, *b.Rat)
		if err != nil {
			return
		}
	}

	return
}

func (b *BigPtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Read_layout(_r, _b, _BigPtrLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat
	var _bool bool

	_bool, err = readwrite.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Float = nil
	} else {
		b.Float = new(big.Float)
		_bigfloat, err = readwrite.Read_bigfloat(_r, _b, _bb)
		if err != nil {
			return
		}
		*b.Float = _bigfloat
	}

	_bool, err = readwrite.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int = nil
	} else {
		b.Int = new(big.Int)
		_bigint, err = readwrite.Read_bigint(_r, _b, _bb)
		if err != nil {
			return
		}
		*b.Int = _bigint
	}

	_bool, err = readwrite.Read_bool(_r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Rat = nil
	} else {
		b.Rat = new(big.Rat)
		_bigrat, err = readwrite.Read_bigrat(_r, _b, _bb)
		if err != nil {
			return
		}
		*b.Rat = _bigrat
	}

	return
}

const _BigPtrSliceLayout = "XWZXWZXWZ"

func (b *BigPtrSlice) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Write_layout(_w, _b, _BigPtrSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _n int

	{
		_s := b.Float
		_n = len(_s)
		err = readwrite.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = readwrite.Write_bool(_w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = readwrite.Write_bigfloat(_w, _b, _bb, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := b.Int
		_n = len(_s)
		err = readwrite.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = readwrite.Write_bool(_w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = readwrite.Write_bigint(_w, _b, _bb, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := b.Rat
		_n = len(_s)
		err = readwrite.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = readwrite.Write_bool(_w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = readwrite.Write_bigrat(_w, _b, _bb, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (b *BigPtrSlice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Read_layout(_r, _b, _BigPtrSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat
	var _bool bool
	var _n int

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(b.Float); _n > _c || _c-_n > _c/8 {
		b.Float = make([]*big.Float, _n)
	} else {
		b.Float = (b.Float)[:_n]
	}
	if _n > 0 {
		_s := b.Float
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_bool, err = readwrite.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(big.Float)
				_bigfloat, err = readwrite.Read_bigfloat(_r, _b, _bb)
				if err != nil {
					return
				}
				*_s[_k] = _bigfloat
			}
		}
	}

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(b.Int); _n > _c || _c-_n > _c/8 {
		b.Int = make([]*big.Int, _n)
	} else {
		b.Int = (b.Int)[:_n]
	}
	if _n > 0 {
		_s := b.Int
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_bool, err = readwrite.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(big.Int)
				_bigint, err = readwrite.Read_bigint(_r, _b, _bb)
				if err != nil {
					return
				}
				*_s[_k] = _bigint
			}
		}
	}

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(b.Rat); _n > _c || _c-_n > _c/8 {
		b.Rat = make([]*big.Rat, _n)
	} else {
		b.Rat = (b.Rat)[:_n]
	}
	if _n > 0 {
		_s := b.Rat
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_bool, err = readwrite.Read_bool(_r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(big.Rat)
				_bigrat, err = readwrite.Read_bigrat(_r, _b, _bb)
				if err != nil {
					return
				}
				*_s[_k] = _bigrat
			}
		}
	}

	return
}

const _BigSliceLayout = "XZXZXZ"

func (b *BigSlice) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Write_layout(_w, _b, _BigSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _n int

	{
		_s := b.Float
		_n = len(_s)
		err = readwrite.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = readwrite.Write_bigfloat(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := b.Int
		_n = len(_s)
		err = readwrite.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = readwrite.Write_bigint(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := b.Rat
		_n = len(_s)
		err = readwrite.Write_len(_w, _b, _n)
		if err != nil {
			return
		}
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			err = readwrite.Write_bigrat(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *BigSlice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Read_layout(_r, _b, _BigSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat
	var _n int

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(b.Float); _n > _c || _c-_n > _c/8 {
		b.Float = make([]big.Float, _n)
	} else {
		b.Float = (b.Float)[:_n]
	}
	if _n > 0 {
		_s := b.Float
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_bigfloat, err = readwrite.Read_bigfloat(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigfloat
		}
	}

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(b.Int); _n > _c || _c-_n > _c/8 {
		b.Int = make([]big.Int, _n)
	} else {
		b.Int = (b.Int)[:_n]
	}
	if _n > 0 {
		_s := b.Int
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_bigint, err = readwrite.Read_bigint(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigint
		}
	}

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _c := cap(b.Rat); _n > _c || _c-_n > _c/8 {
		b.Rat = make([]big.Rat, _n)
	} else {
		b.Rat = (b.Rat)[:_n]
	}
	if _n > 0 {
		_s := b.Rat
		for _k, _kn := 0, _n; _k < _kn; _k++ {
			_bigrat, err = readwrite.Read_bigrat(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigrat
		}
	}

	return
}

const _BigArrayLayout = "R4ZR4ZR4Z"

func (b *BigArray) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Write_layout(_w, _b, _BigArrayLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)

	{
		_s := &b.Float
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = readwrite.Write_bigfloat(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &b.Int
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = readwrite.Write_bigint(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := &b.Rat
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			err = readwrite.Write_bigrat(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *BigArray) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Read_layout(_r, _b, _BigArrayLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat

	{
		_s := &b.Float
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_bigfloat, err = readwrite.Read_bigfloat(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigfloat
		}
	}

	{
		_s := &b.Int
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_bigint, err = readwrite.Read_bigint(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigint
		}
	}

	{
		_s := &b.Rat
		for _k, _kn := 0, len(_s); _k < _kn; _k++ {
			_bigrat, err = readwrite.Read_bigrat(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigrat
		}
	}

	return
}

const _BigMapLayout = "VCZVCZVCZ"

func (b *BigMap) MarshalBinaryTo(w io.Writer) (err error) {
	_w, _done := iobyte.NewWriter(w)
	defer _done(&err)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Write_layout(_w, _b, _BigMapLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)

	{
		_s := b.Float
		err = readwrite.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = readwrite.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = readwrite.Write_bigfloat(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := b.Int
		err = readwrite.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = readwrite.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = readwrite.Write_bigint(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := b.Rat
		err = readwrite.Write_len(_w, _b, len(_s))
		if err != nil {
			return
		}
		for _k := range _s {
			err = readwrite.Write_int(_w, _b, _k)
			if err != nil {
				return
			}

			err = readwrite.Write_bigrat(_w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *BigMap) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_r := iobyte.NewReader(r)
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = readwrite.Read_layout(_r, _b, _BigMapLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat
	var _int int
	var _n int

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		b.Float = nil
	} else {
		b.Float = make(map[int]big.Float, _n)
		_s := b.Float
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_int, err = readwrite.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_bigfloat, err = readwrite.Read_bigfloat(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigfloat
		}
	}

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		b.Int = nil
	} else {
		b.Int = make(map[int]big.Int, _n)
		_s := b.Int
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_int, err = readwrite.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_bigint, err = readwrite.Read_bigint(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigint
		}
	}

	_n, err = readwrite.Read_len(_r)
	if err != nil {
		return
	}
	if _n == 0 {
		b.Rat = nil
	} else {
		b.Rat = make(map[int]big.Rat, _n)
		_s := b.Rat
		var _k int
		for _j, _jn := 0, _n; _j < _jn; _j++ {
			_int, err = readwrite.Read_int(_r, _b)
			if err != nil {
				return
			}
			_k = _int

			_bigrat, err = readwrite.Read_bigrat(_r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigrat
		}
	}

	return
}
