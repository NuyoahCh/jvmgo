package rtda

import "math"

// OperandStack 操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

// newOperandStack 创建一个新的操作数栈实例
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// PushInt 将一个整数值压入操作数栈
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

// PopInt 弹出操作数栈顶的整数值
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

// PushFloat 将一个浮点数值压入操作数栈
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

// PopFloat 弹出操作数栈顶的浮点数值
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

// PushLong 将一个长整型值压入操作数栈
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

// PopLong 弹出操作数栈顶的长整型值
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

// PushDouble 将一个双精度浮点数值压入操作数栈
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

// PopDouble 弹出操作数栈顶的双精度浮点数值
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

// PushRef 将一个对象引用压入操作数栈
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

// PopRef 弹出操作数栈顶的对象引用
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}
