package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/

// ConstantStringInfo 表示常量池中的字符串类型
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// readInfo 从 ClassReader 中读取字符串信息
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// StringIndex 返回字符串索引
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
