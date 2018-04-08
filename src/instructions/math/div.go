package math

import (
	"jvmgo/src/rtda"
	"jvmgo/src/instructions/base"
)

type IDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }
type DDIV struct{ base.NoOperandsInstruction }

func (self DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

func (self IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 / v2
	stack.PushInt(result)
}

func (self FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

func (self LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 / v2
	stack.PushLong(result)
}

