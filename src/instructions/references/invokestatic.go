package references

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
)

type INVOKE_STATIC struct {base.Index16Instruction}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveMethod := methodRef.ResolvedMethod()

	if !resolveMethod.IsStatic() {
		panic("java.lang.IncpmpatibaleClassChangeErroe")
	}
	base.InvokeMethod(frame, resolveMethod)
}
