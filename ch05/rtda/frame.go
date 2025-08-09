package rtda

type Frame struct {
	lower        *Frame        // 下一个栈帧
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 所属线程
	nextPC       int           // 下一条指令的地址
}

// newFrame 创建一个新的栈帧
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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
