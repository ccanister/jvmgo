package math

import (
	"jvmgo/src/instructions/base"
	"jvmgo/src/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self IINC) Execute(frame *rtda.Frame) {
	localVar := frame.LocalVars()
	val := localVar.GetInt(self.Index)
	val += self.Const
	localVar.SetInt(self.Index, val)
}
