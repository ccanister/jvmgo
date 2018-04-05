package heap

import "jvmgo/src/classfile"

type Field struct {
	ClassMember
	slotId 			uint
	constantValueIndex	uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}

	return fields
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "D" || self.descriptor == "J"
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstantValueIndex() uint {
	return self.constantValueIndex
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constantValueIndex = uint(valAttr.ConstantValueIndex())
	}
}