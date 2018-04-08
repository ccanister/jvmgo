package base

import (
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
)

func InvokeMethod(invokeParam *rtda.Frame, method *heap.Method) {
	thread := invokeParam.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argsSlot := int(method.ArgSlotCount())
	if argsSlot > 0 {
		for i := argsSlot - 1; i >= 0; i-- {
			slot := invokeParam.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
