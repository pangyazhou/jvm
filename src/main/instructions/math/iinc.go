package math

import (
	"main/instructions/base"
	"main/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader)  {
	self.Index = uint(uint(reader.ReadUint8()))
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(uint(self.Index))
	val += self.Const
	localVars.SetInt(uint(self.Index), val)
}