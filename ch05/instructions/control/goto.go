package control

import (
	"github.com/NuyoahCh/jvmgo/ch05/instructions/base"
	"github.com/NuyoahCh/jvmgo/ch05/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
