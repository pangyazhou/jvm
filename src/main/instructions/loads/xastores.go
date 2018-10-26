package loads

import (
	"main/instructions/base"
	"main/rtda"
)

type AASTORE struct{ base.NoOperandsInstruction }
type BASTORE struct{ base.NoOperandsInstruction }
type CASTORE struct{ base.NoOperandsInstruction }
type DASTORE struct{ base.NoOperandsInstruction }
type FASTORE struct{ base.NoOperandsInstruction }
type IASTORE struct{ base.NoOperandsInstruction }
type LASTORE struct{ base.NoOperandsInstruction }
type SASTORE struct{ base.NoOperandsInstruction }

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	refs := ref.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
}

func (self *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	bytes := ref.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(val)
}
func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	chars := ref.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(val)
}
func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	doubles := ref.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = val
}
func (self *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	floats := ref.Floats()
	checkIndex(len(floats), index)
	floats[index] = val
}
func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	ints := ref.Ints()
	checkIndex(len(ints), index)
	ints[index] = val
}
func (self *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	longs := ref.Longs()
	checkIndex(len(longs), index)
	longs[index] = val
}
func (self *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	ref := stack.PopRef()
	checkNotNil(ref)
	shorts := ref.Shorts()
	checkIndex(len(shorts), index)
	shorts[index] = int16(val)
}
