package vm

type PyList struct {
	data []PyObject
	Dict map[string]func(this *PyList, args []PyObject) PyObject
}

func NewPyList(items []PyObject) *PyList {
	dict := make(map[string]func(this *PyList, args []PyObject) PyObject)
	dict["__len__"] = func(this *PyList, args []PyObject) PyObject {
		return NewPyInt(len(this.data))
	}
	dict["__getitem__"] = func(this *PyList, args []PyObject) PyObject {
		return this.data[args[0].(*PyInt).getVal()]
	}
	dict["__setitem__"] = func(this *PyList, args []PyObject) PyObject {
		this.data[args[0].(*PyInt).getVal()] = args[1]
		return NewPyNone()
	}
	// TODO
	// __iter__
	dict["__str__"] = func(this *PyList, args []PyObject) PyObject {
		return NewPyStr(this.toString())
	}
	dict["__type__"] = func(this *PyList, args []PyObject) PyObject {
		return this.getType()
	}
	return &PyList{
		data: items,
		Dict: dict,
	}
}

func (this *PyList) getType() *PyType {
	return PyTypes[PyListType]
}

func (this *PyList) toString() string {
	ret := "["
	ret += this.data[0].toString()
	if len(this.data) == 1 {
		ret += "]"
		return ret
	}
	for i := 1; i < len(this.data); i++ {
		ret += ", "
		ret += this.data[i].toString()
	}
	ret += ")"
	return ret
}
