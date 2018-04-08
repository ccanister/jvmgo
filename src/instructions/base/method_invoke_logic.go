package base

import (
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
	"fmt"
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

	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native methodL %v.%v%v\n", method.Class().Name(),
					method.Name(), method.Descriptor()))
		}
	}
}
