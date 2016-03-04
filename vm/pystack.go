package vm

type PyStack struct {
	topElem *elem
	length  int
}

type elem struct {
	value PyObject
	next  *elem
}

func NewPyStack() *PyStack {
	return new(PyStack)
}

// Return the stack's length
func (s *PyStack) size() int {
	return s.length
}

// Push a new element onto the stack
func (s *PyStack) push(value PyObject) {
	s.topElem = &elem{
		value: value,
		next:  s.topElem,
	}
	s.length++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *PyStack) pop() PyObject {
	if s.length > 0 {
		var value PyObject
		value, s.topElem = s.topElem.value, s.topElem.next
		s.length--
		return value
	}
	return nil
}

func (s *PyStack) top() PyObject {
	return s.topElem.value
}
