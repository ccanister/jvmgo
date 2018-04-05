package comparisons

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type IF_ICMPEQ struct{ base.BranchInstruction }

type IF_ICMPNE struct{ base.BranchInstruction }

type IF_ICMPLT struct{ base.BranchInstruction }

type IF_ICMPLE struct{ base.BranchInstruction }

type IF_ICMPGT struct{ base.BranchInstruction }

type IF_ICMPGE struct{ base.BranchInstruction }

func _skipInt(frame *rtda.Frame, offset int, instruction string) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	flag := false
	switch instruction {
	case "==":flag = v1 == v2
	case "<":flag = v1 < v2
	case "<=":flag = v1 <= v2
	case ">":flag = v1 > v2
	case ">=":flag = v1 >= v2
	case "!=":flag = v1 != v2
	}
	if flag {
		base.Branch(frame, offset)
	}
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, "==")
}
func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, "!=")
}
func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, "<")
}
func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, "<=")
}
func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, ">")
}
func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	_skipInt(frame, self.Offset, ">=")
}
