package references

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
	"jvmgo/src/rtda/heap"
)

type PUT_FIELD struct{
	base.Index16Instruction
}

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()
	class := field.Class()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassCHangeError")
	}

	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<client>" {
			panic("java.lang.IlleagalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()


	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := get_ref(stack)
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := get_ref(stack)
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := get_ref(stack)
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := get_ref(stack)
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := get_ref(stack)
		ref.Fields().SetRef(slotId, val)
	}
}

func get_ref(stack *rtda.OperandStack) *heap.Object {
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	return ref
}


