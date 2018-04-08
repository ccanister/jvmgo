package references

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	/*cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V":fmt.Printf("%v\n", stack.PopInt())
		case "(B)V":fmt.Printf("%v\n", stack.PopInt())
		case "(S)V":fmt.Printf("%v\n", stack.PopInt())
		case "(I)V":fmt.Printf("%v\n", stack.PopInt())
		case "(J)V":fmt.Printf("%v\n", stack.PopLong())
		case "(F)V":fmt.Printf("%v\n", stack.PopFloat())
		case "(D)V":fmt.Printf("%v\n", stack.PopDouble())
		default:
			panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}*/

	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncpmpatibaleClassChangeErroe")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		//hack??
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPinterEception")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)

}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":fmt.Printf("%v\n", stack.PopInt())
	case "(B)V":fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":fmt.Printf("%v\n", stack.PopInt())
	case "(J)V":fmt.Printf("%v\n", stack.PopLong())
	case "(F)V":fmt.Printf("%v\n", stack.PopFloat())
	case "(D)V":fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
