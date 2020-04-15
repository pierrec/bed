package testpkg

import "math/big"

type Big struct {
	Float big.Float
	Int   big.Int
	Rat   big.Rat
}

type BigPtr struct {
	Float *big.Float
	Int   *big.Int
	Rat   *big.Rat
}

type BigSlice struct {
	Float []big.Float
	Int   []big.Int
	Rat   []big.Rat
}

type BigArray struct {
	Float [4]big.Float
	Int   [4]big.Int
	Rat   [4]big.Rat
}

type BigMap struct {
	Float map[int]big.Float
	Int   map[int]big.Int
	Rat   map[int]big.Rat
}

type BigPtrSlice struct {
	Float []*big.Float
	Int   []*big.Int
	Rat   []*big.Rat
}
