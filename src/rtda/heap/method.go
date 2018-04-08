package heap

import (
	"jvmgo/src/classfile"
	"fmt"
)

type Method struct {
	ClassMember
	maxLocals	uint
	maxStack	uint
	code		[]byte
	argSlotCount	uint
}

func (self *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = uint(codeAttr.MaxStack())
		self.maxLocals = uint(codeAttr.MaxLocals())
		self.code = codeAttr.Code()
	}
}

func NewMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
		methods[i].calcArgSlotCount()
	}

	return methods
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount ++
		fmt.Println(paramType)
		//TODO judge L or D
	}
	if !self.IsStatic() {
		self.argSlotCount ++
	}
}