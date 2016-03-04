package vm

type PyNone struct{}

func NewPyNone() *PyNone {
	return &PyNone{}
}

func (this *PyNone) getType() *PyType {
	return PyTypes[PyNoneType]
}

func (this *PyNone) toString() string {
	return "None"
}
