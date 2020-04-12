// Package serializer provides an easy way to serialize and deserialize any data at runtime or via generated methods.
//
// It does not handle versions and is only safe to deserialize data when structs have been augmented with already
// serialized fields not modified.
package serializer

//TODO allow unnamed embedded structs
//TODO io.Reader to io.Reader + io.ByteReader
//TODO benchmarks
//TODO version support?
//TODO pack slice items in batches?
