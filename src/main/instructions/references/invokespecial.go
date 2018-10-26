package references

import (
	"main/instructions/base"
	"main/rtda"
	"main/rtda/heap"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	//调用方法的类
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	//调用的方法引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	//声明定义方法的类
	resolvedClass := methodRef.ResolvedClass()
	//定义的方法
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//调用方法的对象引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	//被调用的方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass,
			methodRef.Name(), methodRef.Descriptor())
	}
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	//执行方法
	base.InvokeMethod(frame, methodToBeInvoked)
}
