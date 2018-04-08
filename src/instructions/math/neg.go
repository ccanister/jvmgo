package math

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type INEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }
type DNEG struct{ base.NoOperandsInstruction }

func (self DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	result := -v1
	stack.PushDouble(result)
}

func (self INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	result := -v1
	stack.PushInt(result)
}

func (self FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	result := -v1
	stack.PushFloat(result)
}

func (self LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	result := -v1
	stack.PushLong(result)
}
