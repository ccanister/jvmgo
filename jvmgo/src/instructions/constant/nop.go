package constant

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type NOP struct{ base.BranchInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {

}
