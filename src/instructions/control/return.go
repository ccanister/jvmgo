package control

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type RETURN struct { base.NoOperandsInstruction }
type ARETURN struct { base.NoOperandsInstruction }
type DRETURN struct { base.NoOperandsInstruction }
type FRETURN struct { base.NoOperandsInstruction }
type IRETURN struct { base.NoOperandsInstruction }
type LRETURN struct { base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func get_frame(frame *rtda.Frame) (*rtda.Frame, *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	return currentFrame, invokerFrame
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	currentFrame, invokerFrame := get_frame(frame)
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	currentFrame, invokerFrame := get_frame(frame)
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	currentFrame, invokerFrame := get_frame(frame)
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	currentFrame, invokerFrame := get_frame(frame)
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	currentFrame, invokerFrame := get_frame(frame)
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}



