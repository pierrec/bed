package testpkg

type Map struct {
	Empty        map[int]int
	StringInt    map[string]int
	StringInts   map[string][]int
	UintPtrUint  map[*uint]uint
	IntPtrInt    map[*int]int
	IntIntPtr    map[int]*int
	IntPtrIntPtr map[*int]*int
	IntStruct    map[int]Basic
}
