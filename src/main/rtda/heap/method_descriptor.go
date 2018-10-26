package heap

type MethodDescriptor struct {
	paramterTypes []string
	returnType    string
}

func (self *MethodDescriptor) addParameterType(t string) {
	pLen := len(self.paramterTypes)
	if pLen == cap(self.paramterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, self.paramterTypes)
		self.paramterTypes = s
	}
	self.paramterTypes = append(self.paramterTypes, t)
}
