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

type BigPtrSlice struct {
	Float []*big.Float
	Int   []*big.Int
	Rat   []*big.Rat
}
