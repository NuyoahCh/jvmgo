package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/

// ConstantValueAttribute 定长属性，表示一个常量值
type ConstantValueAttribute struct {
	constantValueIndex uint16 // 常量值索引
}

// readInfo 从 ClassReader 中读取常量值属性信息
func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

// ConstantValueIndex 返回常量值索引
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
