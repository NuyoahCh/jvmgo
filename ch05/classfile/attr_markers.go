package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

// DeprecatedAttribute is a marker attribute that indicates that the class, method, or
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

// SyntheticAttribute is a marker attribute that indicates that the class, method, or
type SyntheticAttribute struct {
	MarkerAttribute
}

// MarkerAttribute is a base type for marker attributes that do not have additional information.
type MarkerAttribute struct{}

// readInfo is a no-op for marker attributes since they do not contain additional information.
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
