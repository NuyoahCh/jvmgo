package classfile

import "encoding/binary"

// ClassReader 读取数据
type ClassReader struct {
	data []byte
}

// readUint8 读取一个字节，返回 uint8 类型
func (self *ClassReader) readUint8() uint8 { // u1
	val := self.data[0]
	self.data = self.data[1:] // 移除已读取的字节
	return val
}

// readUint16 读取一个字节，返回 uint16 类型
func (self *ClassReader) readUint16() uint16 { // u2
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:] // 移除已读取的字节
	return val
}

// readUint32 读取一个字节，返回 uint32 类型
func (self *ClassReader) readUint32() uint32 { // u4
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:] // 移除已读取的字节
	return val
}

// readUint64 读取一个字节，返回 uint64 类型
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:] // 移除已读取的字节
	return val
}

// readUint16s 读取一个字节，返回 uint16 类型
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() // 读取 u2，获取长度
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16() // 读取 u2，填充到切片中
	}
	return s
}

// readBytes 读取指定长度的字节，返回 []byte 类型
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]    // 获取指定长度的字节
	self.data = self.data[n:] // 移除已读取的字节
	return bytes
}
