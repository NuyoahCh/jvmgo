package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

// ConstantMethodHandleInfo represents a method handle in the constant pool.
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

// readInfo reads the method handle information from the ClassReader.
func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

// ConstantMethodTypeInfo represents a method type in the constant pool.
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

// readInfo reads the method type information from the ClassReader.
func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

// ConstantInvokeDynamicInfo represents an invoke dynamic constant in the constant pool.
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

// readInfo reads the invoke dynamic information from the ClassReader.
func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
