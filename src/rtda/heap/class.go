package heap

import (
	"jvmgo/src/classfile"
	"strings"
)

type Class struct {
	accessFlags		uint16
	name			string
	superName		string
	interfaceNames		[]string
	constantPool		*ConstantPool
	fields			[]*Field
	methods			[]*Method
	loader			*ClassLoader
	superClass		*Class
	interfaces		[]*Class
	instanceSlotCount	uint
	staticSlotCount		uint
	staticVars		Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.methods = NewMethods(class, cf.Methods())
	class.fields = newFields(class, cf.Fields())

	return class
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) IsInterface() bool {
	return  0 != self.accessFlags & ACC_INTERFACES
}

func (self *Class) IsAbstract() bool {
	return  0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

func (self *Class) NewObject() *Object {
	return NewObject(self)
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Class) GetMainMethod() *Method {
	return self.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.Name() == name && method.Descriptor() == descriptor {
			return method
		}
	}

	return nil
}