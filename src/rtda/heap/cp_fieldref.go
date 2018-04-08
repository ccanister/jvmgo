package heap

import "jvmgo/src/classfile"

type FieldRef struct {
	MemberRef
	field	*Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)

	return ref
}


func (self *FieldRef) ResolveField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if !field.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(class *Class, name string, descriptor string) *Field {
	for _,field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range class.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}

	return nil
}


