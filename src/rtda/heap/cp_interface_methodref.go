package heap

import "jvmgo/src/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method		*Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethoderInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)

	return ref
}

func (self *InterfaceMethodRef) ResolvedInterfacesMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

func (self *InterfaceMethodRef) resolveInterfaceMethodRef() *Method {
	d := self.cp.class
	c := self.ResolvedClass()

	if !c.IsInterface() {
		panic("java.lang.IncompatibleCLassChangeError")
	}

	method := lookInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	return method
}

func lookInterfaceMethod(iface *Class, name string, descriptor string) *Method {
	for _,method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInterfaces(iface.interfaces, name, descriptor)
}