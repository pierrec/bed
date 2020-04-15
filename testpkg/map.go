package testpkg

type Map struct {
	Empty        map[int]int
	BoolBool     map[bool]bool
	StringInt    map[string]int
	StringInts   map[string][]int
	UintPtrUint  map[*uint]uint
	IntPtrInt    map[*int]int
	IntIntPtr    map[int]*int
	IntPtrIntPtr map[*int]*int
	AnonInt      map[struct {
		Int    int
		String string
	}]int
}
