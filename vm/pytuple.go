package vm

type PyTuple struct {
	data []PyObject
	Dict map[string]func(this *PyTuple, args []PyObject) PyObject
}

func NewPyTuple(items []PyObject) *PyTuple {
	dict := make(map[string]func(this *PyTuple, args []PyObject) PyObject)
	dict["__type__"] = func(this *PyTuple, args []PyObject) PyObject {
		return this.getType()
	}
	dict["__len__"] = func(this *PyTuple, args []PyObject) PyObject {
		return NewPyInt(len(this.data))
	}
	dict["__getitem__"] = func(this *PyTuple, args []PyObject) PyObject {
		return this.getItem(args[0].(*PyInt).getVal())
	}
	dict["__str__"] = func(this *PyTuple, args []PyObject) PyObject {
		return NewPyStr(this.toString())
	}
	return &PyTuple{
		data: items,
		Dict: dict,
	}
}

func (this *PyTuple) getType() *PyType {
	return PyTypes[PyTupleType]
}

func (this *PyTuple) toString() string {
	ret := "("
	ret += this.data[0].toString()
	if len(this.data) == 1 {
		ret += ")"
		return ret
	}
	for i := 1; i < len(this.data); i++ {
		ret += ", "
		ret += this.data[i].toString()
	}
	ret += ")"
	return ret
}

func (this *PyTuple) getItem(i int) PyObject {
	return this.data[i]
}

func (this *PyTuple) size() int {
	return len(this.data)
}
