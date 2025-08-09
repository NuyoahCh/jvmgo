package base

// BytecodeReader 用于读取字节码
type BytecodeReader struct {
	code []byte // bytecodes
	pc   int
}

// ReSet 重置 BytecodeReader 的状态
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

// PC 返回当前的程序计数器（pc）
func (self *BytecodeReader) PC() int {
	return self.pc
}

// ReadInt8 读取一个有符号的8位整数
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

// ReadUint8 读取一个无符号的8位整数
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

// ReadInt16 读取一个有符号的16位整数
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

// ReadUint16 读取一个无符号的16位整数
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

// ReadInt32 读取一个有符号的32位整数
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// ReadInt32s 读取一个有符号的32位整数
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// SkipPadding 跳过字节码中的填充字节
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
