package rtda

type Frame struct {
	lower        *Frame        // 下一个栈帧
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	// todo
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),   // 创建局部变量表
		operandStack: newOperandStack(maxStack), // 创建操作数栈
	}
}

// LocalVars getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

// OperandStack getters
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
