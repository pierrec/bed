package testpkg

type CompositeOnly struct {
	Basic Basic
	Slice Slice
	Array Array
}

type Composite struct {
	Bytes []byte
	Basic Basic
	Slice Slice
	Array Array
}
