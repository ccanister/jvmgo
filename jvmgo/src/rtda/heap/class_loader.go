package heap

import (
	"jvmgo/src/classpath"
	"jvmgo/src/classfile"
	"fmt"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class
}

func NewCLassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:                cp,
		classMap:        make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}

	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)	//链接，
	fmt.Printf("[Loader %s from %s]\n", name, entry)
	
	return class
}
func link(class *Class) {
	verify(class)	//验证
	prepare(class)	//初始化
	
}
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}
func allocAndInitStaticVars(class *Class) {
	class.staticVars = NewSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstantValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.descriptor {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String":
			panic("todo")
		}
	}
}
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId ++
			if field.isLongOrDouble() {
				slotId ++
			}
		}
	}

	class.staticSlotCount = slotId
}
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId ++
			if field.isLongOrDouble() {
				slotId ++
			}
		}
	}

	class.instanceSlotCount = slotId
}
func verify(class *Class) {
	
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}

	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)	//加载父类
	resolveSuperInterfaces(class)	//加载继承的接口
	self.classMap[class.name] = class
	
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.classFormatException")
	}
	
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superName)
	}
}

func resolveSuperInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces =  make([]*Class, interfaceCount)
		for index, name := range class.interfaceNames {
			class.interfaces[index] = class.loader.LoadClass(name)
		}
	}
}
