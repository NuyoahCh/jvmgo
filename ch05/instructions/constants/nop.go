package constants

import (
	"github.com/NuyoahCh/jvmgo/ch05/instructions/base"
	"github.com/NuyoahCh/jvmgo/ch05/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

// Execute 执行 NOP 指令
func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
