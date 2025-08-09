package rtda

type Stack struct {
	maxSize uint   // 栈的最大深度
	size    uint   // 当前栈的深度
	_top    *Frame // 栈顶的栈帧
}

// newStack 创建一个新的栈实例
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// push 将一个栈帧压入栈中
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top // 设置新栈帧的下一个栈帧为当前栈顶
		self.size++             // 增加栈的深度
	}
}

// pop 弹出栈顶的栈帧，并返回该栈帧
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top      // 获取当前栈顶的栈帧
	self._top = top.lower // 将栈顶指针指向下一个栈
	top.lower = nil       // 清除下一个栈帧的引用
	self.size--           // 减少栈的深度
	return top            // 返回被弹出的栈帧
}

// top 返回当前栈顶的栈帧
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
