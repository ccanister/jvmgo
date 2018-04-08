package references

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncpmpatibaleClassChangeErroe")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPinterEception")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncpmpatibaleClassChangeErroe")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstarctMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.illegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
