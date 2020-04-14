package testpkg

import (
	"io"
	"math/big"

	"github.com/pierrec/serializer"
)

const _BigLayout = "ZZZ"

func (b *Big) MarshalBinaryTo(w io.Writer) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(w, _b, _BigLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)

	err = serializer.Write_bigfloat(w, _b, _bb, b.Float)
	if err != nil {
		return
	}

	err = serializer.Write_bigint(w, _b, _bb, b.Int)
	if err != nil {
		return
	}

	err = serializer.Write_bigrat(w, _b, _bb, b.Rat)
	if err != nil {
		return
	}

	return
}

func (b *Big) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(r, _b, _BigLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat

	_bigfloat, err = serializer.Read_bigfloat(r, _b, _bb)
	if err != nil {
		return
	}
	b.Float = _bigfloat

	_bigint, err = serializer.Read_bigint(r, _b, _bb)
	if err != nil {
		return
	}
	b.Int = _bigint

	_bigrat, err = serializer.Read_bigrat(r, _b, _bb)
	if err != nil {
		return
	}
	b.Rat = _bigrat

	return
}

const _BigPtrLayout = "WZWZWZ"

func (b *BigPtr) MarshalBinaryTo(w io.Writer) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(w, _b, _BigPtrLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)

	err = serializer.Write_bool(w, _b, b.Float == nil)
	if err != nil {
		return
	}
	if b.Float != nil {
		err = serializer.Write_bigfloat(w, _b, _bb, *b.Float)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Int == nil)
	if err != nil {
		return
	}
	if b.Int != nil {
		err = serializer.Write_bigint(w, _b, _bb, *b.Int)
		if err != nil {
			return
		}
	}

	err = serializer.Write_bool(w, _b, b.Rat == nil)
	if err != nil {
		return
	}
	if b.Rat != nil {
		err = serializer.Write_bigrat(w, _b, _bb, *b.Rat)
		if err != nil {
			return
		}
	}

	return
}

func (b *BigPtr) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(r, _b, _BigPtrLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat
	var _bool bool

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Float = nil
	} else {
		b.Float = new(big.Float)
		_bigfloat, err = serializer.Read_bigfloat(r, _b, _bb)
		if err != nil {
			return
		}
		*b.Float = _bigfloat
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Int = nil
	} else {
		b.Int = new(big.Int)
		_bigint, err = serializer.Read_bigint(r, _b, _bb)
		if err != nil {
			return
		}
		*b.Int = _bigint
	}

	_bool, err = serializer.Read_bool(r, _b)
	if err != nil {
		return
	}
	if _bool {
		b.Rat = nil
	} else {
		b.Rat = new(big.Rat)
		_bigrat, err = serializer.Read_bigrat(r, _b, _bb)
		if err != nil {
			return
		}
		*b.Rat = _bigrat
	}

	return
}

const _BigPtrSliceLayout = "XWZXWZXWZ"

func (b *BigPtrSlice) MarshalBinaryTo(w io.Writer) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(w, _b, _BigPtrSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _n int

	{
		_s := b.Float
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_bool(w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_bigfloat(w, _b, _bb, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := b.Int
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_bool(w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_bigint(w, _b, _bb, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	{
		_s := b.Rat
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_bool(w, _b, _s[_k] == nil)
			if err != nil {
				return
			}
			if _s[_k] != nil {
				err = serializer.Write_bigrat(w, _b, _bb, *_s[_k])
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (b *BigPtrSlice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(r, _b, _BigPtrSliceLayout)
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

	_n, err = serializer.Read_int(r, _b)
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
		for _k := 0; _k < _n; _k++ {
			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(big.Float)
				_bigfloat, err = serializer.Read_bigfloat(r, _b, _bb)
				if err != nil {
					return
				}
				*_s[_k] = _bigfloat
			}
		}
	}

	_n, err = serializer.Read_int(r, _b)
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
		for _k := 0; _k < _n; _k++ {
			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(big.Int)
				_bigint, err = serializer.Read_bigint(r, _b, _bb)
				if err != nil {
					return
				}
				*_s[_k] = _bigint
			}
		}
	}

	_n, err = serializer.Read_int(r, _b)
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
		for _k := 0; _k < _n; _k++ {
			_bool, err = serializer.Read_bool(r, _b)
			if err != nil {
				return
			}
			if _bool {
				_s[_k] = nil
			} else {
				_s[_k] = new(big.Rat)
				_bigrat, err = serializer.Read_bigrat(r, _b, _bb)
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
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Write_layout(w, _b, _BigSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _n int

	{
		_s := b.Float
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_bigfloat(w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := b.Int
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_bigint(w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	{
		_s := b.Rat
		_n = len(_s)
		err = serializer.Write_int(w, _b, _n)
		if err != nil {
			return
		}
		for _k := 0; _k < _n; _k++ {
			err = serializer.Write_bigrat(w, _b, _bb, _s[_k])
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *BigSlice) UnmarshalBinaryFrom(r io.Reader) (err error) {
	_b := serializer.Buffers.Get()
	defer serializer.Buffers.Put(_b)
	err = serializer.Read_layout(r, _b, _BigSliceLayout)
	if err != nil {
		return
	}

	var _bb = serializer.BigBuffers.Get()
	defer serializer.BigBuffers.Put(_bb)
	var _bigfloat big.Float
	var _bigint big.Int
	var _bigrat big.Rat
	var _n int

	_n, err = serializer.Read_int(r, _b)
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
		for _k := 0; _k < _n; _k++ {
			_bigfloat, err = serializer.Read_bigfloat(r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigfloat
		}
	}

	_n, err = serializer.Read_int(r, _b)
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
		for _k := 0; _k < _n; _k++ {
			_bigint, err = serializer.Read_bigint(r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigint
		}
	}

	_n, err = serializer.Read_int(r, _b)
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
		for _k := 0; _k < _n; _k++ {
			_bigrat, err = serializer.Read_bigrat(r, _b, _bb)
			if err != nil {
				return
			}
			_s[_k] = _bigrat
		}
	}

	return
}
