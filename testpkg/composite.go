package testpkg

type CompositeOnly struct {
	Basic Basic
	Slice Slice
	Array Array
	Map   Map
}

type Composite struct {
	Bytes []byte
	Basic Basic
	Slice Slice
	Array Array
}
