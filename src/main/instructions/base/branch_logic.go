package base

import "main/rtda"

func Branch(frame *rtda.Frame, offset int){
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPc(nextPC)
}
