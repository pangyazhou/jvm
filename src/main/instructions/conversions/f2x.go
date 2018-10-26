package conversions

import (
	"main/instructions/base"
	"main/rtda"
)

type F2D struct {base.NoOperandsInstruction}
func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := float64(d)
	stack.PushDouble(i)
}


type F2I struct {base.NoOperandsInstruction}
func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int32(d)
	stack.PushInt(i)
}


type F2L struct {base.NoOperandsInstruction}
func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int64(d)
	stack.PushLong(i)
}

