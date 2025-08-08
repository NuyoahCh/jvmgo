package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/

// SourceFileAttribute 结构体
type SourceFileAttribute struct {
	cp              ConstantPool // 常量池引用
	sourceFileIndex uint16       // 源文件索引
}

// readInfo 从 ClassReader 中读取源文件属性信息
func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

// FileName 返回源文件的名称
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
