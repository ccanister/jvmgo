package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isSubInstanceOf(t)
	}
}


func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass;c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}

	return false
}

func (self *Class) isSubInstanceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if iface == superInterface || superInterface.isSubInstanceOf(iface){
			return true
		}
	}

	return false
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, superInterface := range c.interfaces {
			if superInterface == iface || superInterface.isSubInstanceOf(iface) {
				return true
			}
		}
	}
	return false
}