package vm

type PyStr struct {
	val  string
	Dict map[string]func(this *PyStr, args []PyObject) PyObject
}

func NewPyStr(s string) *PyStr {
	dict := make(map[string]func(this *PyStr, args []PyObject) PyObject)
	dict["__type__"] = func(this *PyStr, args []PyObject) PyObject {
		return this.getType()
	}
	dict["__len__"] = func(this *PyStr, args []PyObject) PyObject {
		return NewPyInt(len(this.val))
	}
	dict["__str__"] = func(this *PyStr, args []PyObject) PyObject {
		return this
	}
	dict["__add__"] = func(this *PyStr, args []PyObject) PyObject {
		return NewPyStr(this.val + args[0].(*PyStr).val)
	}
	// TODO
	// __float__
	// __int__
	// __bool__
	// __eq__
	// __getitem__
	// __iter__
	return &PyStr{
		val:  s,
		Dict: dict,
	}
}

func (this *PyStr) getType() *PyType {
	return PyTypes[PyStrType]
}

func (this *PyStr) toString() string {
	return this.val
}
