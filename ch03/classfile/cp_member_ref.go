package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

// ConstantFieldrefInfo 表示常量池中的字段引用
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

// ConstantMethodrefInfo 表示常量池中的方法引用
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

// ConstantInterfaceMethodrefInfo 表示常量池中的接口方法引用
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

// ConstantMemberrefInfo 是一个抽象类型，包含了类索引和名称与类型索引
type ConstantMemberrefInfo struct {
	cp               ConstantPool // 常量池引用
	classIndex       uint16       // 类索引
	nameAndTypeIndex uint16       // 名称和类型索引
}

// readInfo 从 ClassReader 中读取成员引用信息
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

// ClassName 返回类或接口的名称
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

// NameAndDescriptor 返回名称和描述符
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
