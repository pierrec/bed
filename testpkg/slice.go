package testpkg

type Slice struct {
	Bool       []bool
	Int        []int
	Int8       []int8
	Int16      []int16
	Int32      []int32
	Int64      []int64
	Uint       []uint
	Uint8      []uint8
	Uint16     []uint16
	Uint32     []uint32
	Uint64     []uint64
	Complex64  []complex64
	Complex128 []complex128
	String     []string
	Maps       []map[string][]int
}

type SlicePtr struct {
	Bool       *[]bool
	Int        *[]int
	Int8       *[]int8
	Int16      *[]int16
	Int32      *[]int32
	Int64      *[]int64
	Uint       *[]uint
	Uint8      *[]uint8
	Uint16     *[]uint16
	Uint32     *[]uint32
	Uint64     *[]uint64
	Complex64  *[]complex64
	Complex128 *[]complex128
	String     *[]string
	Maps       *[]map[string][]int
}

type SliceAnon struct {
	Anon struct {
		Int    int
		String string
	}
	AnonSlice []struct {
		Uint   uint
		Uint32 uint32
	}
}
