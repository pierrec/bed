// Code generated by "stringer -type=isKind -linecomment"; DO NOT EDIT.

package serializer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[isSlice-1]
	_ = x[isArray-2]
	_ = x[isByteArray-3]
	_ = x[isStruct-4]
	_ = x[isAnonStruct-5]
	_ = x[isMap-6]
	_ = x[isMapStruct-7]
	_ = x[isPointer-8]
}

const _isKind_name = "SliceArrayByteArrayStructAnonStructMapMapStructPointer"

var _isKind_index = [...]uint8{0, 5, 10, 19, 25, 35, 38, 47, 54}

func (i isKind) String() string {
	i -= 1
	if i >= isKind(len(_isKind_index)-1) {
		return "isKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _isKind_name[_isKind_index[i]:_isKind_index[i+1]]
}