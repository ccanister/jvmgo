package comparisons

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type IF_ACMPEQ struct{ base.BranchInstruction }

type IF_ACMPNE struct{ base.BranchInstruction }

func _skipRef(frame *rtda.Frame, offset int, instruction string) {
	stack := frame.OperandStack()
	v1 := stack.PopRef()
	v2 := stack.PopRef()
	flag := false
	switch instruction {
	case "==":flag = v1 == v2
	case "!=":flag = v1 != v2
	}
	if flag {
		base.Branch(frame, offset)
	}
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, "==")
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, "!=")
}