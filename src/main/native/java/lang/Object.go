package lang

import (
	"main/native"
	"main/rtda"
)

func init() {
	native.Register("java/lang/Object",
		"getClass",
		"()Ljava/lang/Class",
		getClass)
}

// public final native Class<?> getClass();
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
