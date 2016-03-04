package vm

type PyTypeId int

const (
	PyTypeType PyTypeId = iota
	PyNoneType
	PyBoolType
	PyIntType
	PyFloatType
	PyStrType
	PyFunctionType
	PyBuiltInType
	PyRangeTypeId
	PyListType
	PyCodeType
	PyTupleType
	// TODO
	// PyRangeIteratorType
	// PyListIteratorType
	// PyStrIteratorType
	// PyTupleIteratorType
	// PyExceptionTypeId
)

type PyType struct {
	typeString string
	index      PyTypeId
}

// Init in init.go
var PyTypes map[PyTypeId]*PyType

func NewPyType(typeString string, id PyTypeId) *PyType {
	return &PyType{
		typeString: typeString,
		index:      id,
	}
}

func (this *PyType) toString() string {
	return this.typeString
}

func (this *PyType) getType() *PyType {
	return PyTypes[PyTypeType]
}

func (this *PyType) typeId() PyTypeId {
	return this.index
}
