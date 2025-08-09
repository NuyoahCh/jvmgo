package rtda

// Thread 代表一个线程
type Thread struct {
	pc    int
	stack *Stack // 栈帧栈
}

// NewThread 创建一个新的线程实例
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), // 初始化栈，假设最大栈深度为 1024
	}
}

// PC 获取当前线程的程序计数器
func (self *Thread) PC() int {
	return self.pc // getter
}

// SetPC 设置当前线程的程序计数器
func (self *Thread) SetPC(pc int) {
	self.pc = pc // setter
}

// PushFrame 将一个栈帧压入线程的栈中
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame) // 将栈帧压入栈中
}

// PopFrame 弹出线程栈顶的栈帧，并返回该栈帧
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop() // 弹出栈顶的栈帧
}

// CurrentFrame 获取当前线程栈顶的栈帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top() // 返回栈顶的栈帧
}
