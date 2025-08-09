package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// CodeAttribute Code 变长属性结构体，表示方法的字节码和异常处理信息
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

// readInfo 从 ClassReader 中读取 Code 属性信息
func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

// MaxStack 操作数栈的最大深度
func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

// MaxLocals 局部变量表的最大大小
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

// Code 返回方法的字节码
func (self *CodeAttribute) Code() []byte {
	return self.code
}

// ExceptionTable 返回异常处理表
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

// ExceptionTableEntry 异常处理表
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// readExceptionTable 读取异常处理表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

// StartPc 开始位置
func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}

// EndPc 结束位置
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}

// HandlerPc 异常处理程序的位置
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}

// CatchType 捕获的异常类型索引
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
