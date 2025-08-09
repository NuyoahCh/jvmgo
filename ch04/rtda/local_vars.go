package rtda

import "math"

// LocalVars 代表局部变量表，存储方法调用时的局部变量
type LocalVars []Slot

// newLocalVars 创建一个新的局部变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// SetInt 设置指定索引的局部变量为一个整数值
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

// GetInt 获取指定索引的局部变量的整数值
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// SetFloat 设置指定索引的局部变量为一个浮点数值
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

// GetFloat 获取指定索引的局部变量的浮点数值
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// SetLong 设置指定索引的局部变量为一个长整型值
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

// GetLong 获取指定索引的局部变量的长整型值
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// SetDouble 设置指定索引的局部变量为一个双精度浮点数值
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

// GetDouble 获取指定索引的局部变量的双精度浮点数值
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef 设置指定索引的局部变量为一个对象引用
func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

// GetRef 获取指定索引的局部变量的对象引用
func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}
