package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/

// ConstantNameAndTypeInfo 表示常量池中的名称和类型信息
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

// readInfo 从 ClassReader 中读取名称和类型信息
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
