package base

type BytecodeReader struct {
	pc	int
	code	[]byte
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) Reset(pc int, code []byte) {
	self.pc = pc
	self.code = code
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	low := self.ReadUint8()
	high := self.ReadUint8()
	return uint16(low) << 8 | uint16(high)
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadUint32() uint32 {
	low := self.ReadUint16()
	high := self.ReadUint16()
	return uint32(low) << 16 | uint32(high)
}

func (self *BytecodeReader) ReadInt32() int32 {
	return int32(self.ReadUint32())
}

func (self *BytecodeReader) ReadInt32s(count int32) []int32 {
	bytes := make([]int32, count)
	for index := range bytes {
		bytes[index] = self.ReadInt32()
	}
	return bytes
}

/**
读取填充，保证PC是4的倍数
 */
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4!=0 {
		self.ReadUint8()
	}
}

