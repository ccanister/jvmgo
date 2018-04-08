package heap

import "jvmgo/src/classfile"

type ClassMember struct {
	accessFlags	uint16
	name		string
	descriptor	string
	class 		*Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}


func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *ClassMember) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *ClassMember) IsNative() bool {
	return 0 != self.accessFlags & ACC_NATIVE
}

func (self *ClassMember) IsAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) || d.GetPackageName() == c.GetPackageName()
	}

	if !self.IsPrivate() {
		return d == c || d.GetPackageName() == c.GetPackageName()
	}
	return d == c
}


func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}