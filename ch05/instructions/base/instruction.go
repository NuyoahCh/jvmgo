package base

import "github.com/NuyoahCh/jvmgo/ch05/rtda"

// Instruction 指令操作接口
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 无操作数指令
type NoOperandsInstruction struct {
	// empty
}

// FetchOperands 从字节码读取操作数
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction 分支指令，跳转指令
type BranchInstruction struct {
	Offset int
}

// FetchOperands 从字节码读取分支偏移量
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction 和 Index16Instruction 用于处理索引指令，存储和加载类指令需要根据索引存取局部变量表
type Index8Instruction struct {
	Index uint
}

// FetchOperands 从字节码读取8位索引
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction 用于处理16位索引指令
type Index16Instruction struct {
	Index uint
}

// FetchOperands 从字节码读取16位索引
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
