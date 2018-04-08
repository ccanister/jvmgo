package conversions

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type D2I struct{ base.NoOperandsInstruction }
type D2F struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := float64(d)
	stack.PushDouble(l)
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := int64(d)
	stack.PushLong(f)
}

