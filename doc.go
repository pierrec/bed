// Package bed provides an easy way to serialize and deserialize any data at runtime or via generated methods.
//
// It does not handle versions and is only safe to deserialize data when structs have been augmented with already
// serialized fields not modified.
package bed

//TODO guard against panics on invalid inputs
//TODO benchmarks
//TODO version support?
//TODO pack slice items in batches?
