package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

// ConstantClassInfo 表示常量池中的类或接口信息
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

// readInfo 从 ClassReader 中读取类信息
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

// Name 返回类或接口的名称
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
