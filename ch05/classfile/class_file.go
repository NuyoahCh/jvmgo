package classfile

import "fmt"

/*
	ClassFile {
	    u4             magic;
	    u2             minor_version;
	    u2             major_version;
	    u2             constant_pool_count;
	    cp_info        constant_pool[constant_pool_count-1];
	    u2             access_flags;
	    u2             this_class;
	    u2             super_class;
	    u2             interfaces_count;
	    u2             interfaces[interfaces_count];
	    u2             fields_count;
	    field_info     fields[fields_count];
	    u2             methods_count;
	    method_info    methods[methods_count];
	    u2             attributes_count;
	    attribute_info attributes[attributes_count];
	}
*/
type ClassFile struct {
	//magic      uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse 把 []byte 解析成 ClassFile 结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	// 使用 defer 和 recover 来捕获 panic
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// 创建 ClassReader 实例并读取 classData
	cr := &ClassReader{classData}
	// 创建 ClassFile 实例并调用 read 方法
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// read 方法从 ClassReader 中读取 class 文件的各个部分
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)                              // 检查魔数
	self.readAndCheckVersion(reader)                            // 检查版本号
	self.constantPool = readConstantPool(reader)                // 读取常量池
	self.accessFlags = reader.readUint16()                      // 读取访问标志
	self.thisClass = reader.readUint16()                        // 读取当前类索引
	self.superClass = reader.readUint16()                       // 读取父类索引
	self.interfaces = reader.readUint16s()                      // 读取接口索引
	self.fields = readMembers(reader, self.constantPool)        // 读取字段信息
	self.methods = readMembers(reader, self.constantPool)       // 读取方法信息
	self.attributes = readAttributes(reader, self.constantPool) // 读取属性信息
}

// readAndCheckMagic 检查 class 文件的魔数是否正确
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	// 读取魔数，应该是 0xCAFEBABE
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// readAndCheckVersion 检查 class 文件的版本号是否符合要求
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16() // 读取次版本号
	self.majorVersion = reader.readUint16() // 读取主版本号
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

// MinorVersion 返回次版本号
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

// MajorVersion 返回主版本号
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// ConstantPool 返回常量池
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// AccessFlags 返回访问标志
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// Fields 返回字段信息
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// Methods 返回方法信息
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// ClassName 返回当前类的名称
func (self *ClassFile) ClassName() string {
	// 获取常量池中 thisClass 的类名
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName 返回父类的名称
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

// InterfaceNames 返回接口的名称列表
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
