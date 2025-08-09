package stack

import (
	"github.com/NuyoahCh/jvmgo/ch05/instructions/base"
	"github.com/NuyoahCh/jvmgo/ch05/rtda"
)

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]

	|
	V

[...][c][b]
*/

// Execute 执行指令
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]

	|  |
	V  V

[...][c]
*/

// Execute 执行指令
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
