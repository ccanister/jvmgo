package references

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	//frame.OperandStack().PopRef()
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncpmpatibaleClassChangeErroe")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPinterEception")
	}

	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
	   resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
	   ref.Class() != currentClass &&
	   !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSubClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
