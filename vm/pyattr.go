package vm

type PyAttr struct {
	self *PyObject
	name string
	Dict map[string]func(this *PyAttr, args []PyObject) PyObject
}

func NewPyAttr(self PyObject, attrName string) *PyAttr {
	dict := make(map[string]func(this *PyAttr, args []PyObject) PyObject)
	dict["__call__"] = func(this *PyAttr, args []PyObject) PyObject {
		return CallAttr(*(this.self), this.name, args)
	}
	return &PyAttr{
		self: &self,
		name: attrName,
		Dict: dict,
	}
}

func (this *PyAttr) getType() *PyType {
	return PyTypes[PyBuiltInType]
}

func (this *PyAttr) toString() string {
	return "<method " + this.name + " of " + (*(this.self)).getType().toString() + "object>"
}
