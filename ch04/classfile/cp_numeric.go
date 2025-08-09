package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/

// ConstantIntegerInfo 表示常量池中的整数类型
type ConstantIntegerInfo struct {
	val int32
}

// readInfo 从 ClassReader 中读取整数信息
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// Value 返回整数值
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/*
	CONSTANT_Float_info {
	    u1 tag;
	    u4 bytes;
	}
*/

// ConstantFloatInfo 表示常量池中的浮点数类型
type ConstantFloatInfo struct {
	val float32
}

// readInfo 从 ClassReader 中读取浮点数信息
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

// Value 返回浮点数值
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
	CONSTANT_Long_info {
	    u1 tag;
	    u4 high_bytes;
	    u4 low_bytes;
	}
*/

// ConstantLongInfo 表示常量池中的长整型
type ConstantLongInfo struct {
	val int64
}

// readInfo 从 ClassReader 中读取长整型信息
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

// Value 返回长整型值
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/*
	CONSTANT_Double_info {
	    u1 tag;
	    u4 high_bytes;
	    u4 low_bytes;
	}
*/

// ConstantDoubleInfo 表示常量池中的双精度浮点数类型
type ConstantDoubleInfo struct {
	val float64
}

// readInfo 从 ClassReader 中读取双精度浮点数信息
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

// Value 返回双精度浮点数值
func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
