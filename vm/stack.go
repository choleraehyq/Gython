package vm

type IntStack struct {
	topElem *intElem
	length  int
}

type intElem struct {
	value int
	next  *intElem
}

func NewIntStack() *IntStack {
	return new(IntStack)
}

// Return the IntStack's length
func (s *IntStack) size() int {
	return s.length
}

// Push a new element onto the IntStack
func (s *IntStack) push(value int) {
	s.topElem = &intElem{
		value: value,
		next:  s.topElem,
	}
	s.length++
}

// Remove the top element from the IntStack and return it's value
// If the IntStack is empty, return -1
func (s *IntStack) pop() int {
	if s.length > 0 {
		var value int
		value, s.topElem = s.topElem.value, s.topElem.next
		s.length--
		return value
	}
	return -1
}

func (s *IntStack) top() int {
	return s.topElem.value
}
