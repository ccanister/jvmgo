package heap

type Object struct {
	class 	*Class
	fields 	Slots
}

func NewObject(class *Class) *Object {
	return &Object{
		class:		class,
		fields:		NewSlots(class.instanceSlotCount),
	}
}

func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) Class() *Class {
	return self.class
}
