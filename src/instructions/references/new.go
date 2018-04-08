package references

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if class.IsAbstract() || class.IsInterface() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
