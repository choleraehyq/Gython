package vm

type PyBool struct {
	val  bool
	Dict map[string]func(this *PyBool, args []PyObject) PyObject
}

func NewPyBool(val bool) *PyBool {
	dict := make(map[string]func(this *PyBool, args []PyObject) PyObject)
	dict["__type__"] = func(this *PyBool, args []PyObject) PyObject {
		return this.getType()
	}
	dict["__bool__"] = func(this *PyBool, args []PyObject) PyObject {
		return this
	}
	dict["__str__"] = func(this *PyBool, args []PyObject) PyObject {
		return NewPyStr(this.toString())
	}
	return &PyBool{
		val:  val,
		Dict: dict,
	}
}

func (*PyBool) getType() *PyType {
	return PyTypes[PyBoolType]
}

func (this *PyBool) toString() string {
	if this.val {
		return "True"
	} else {
		return "False"
	}
}

func (this *PyBool) getVal() bool {
	return this.val
}
