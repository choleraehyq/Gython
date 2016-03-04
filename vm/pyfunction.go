package vm

type PyFunction struct {
	code     *PyCode
	globals  map[string]PyObject
	cellvars map[string]PyObject
	Dict     map[string]func(this *PyFunction, args []PyObject) PyObject
}

func NewPyFunction(code *PyCode, globals map[string]PyObject, env PyObject) *PyFunction {
	dict := make(map[string]func(this *PyFunction, args []PyObject) PyObject)
	dict["__call__"] = func(this *PyFunction, args []PyObject) PyObject {
		frame := NewPyFrame(code, args, this.globals, code.getConstants(), this.cellvars)
		result := frame.execute()
		return result
	}

	// Just for closure
	cellvars := make(map[string]PyObject)
	if env != nil {
		for i, item := range env.(*PyTuple).data {
			cellvars[code.getFreevars()[i]] = item
		}
	}

	return &PyFunction{
		code:     code,
		cellvars: cellvars,
		globals:  globals,
		Dict:     dict,
	}
}

func (this *PyFunction) getType() *PyType {
	return PyTypes[PyFunctionType]
}

func (this *PyFunction) toString() string {
	return "function(" + this.callName() + ")"
}

func (this *PyFunction) callName() string {
	return this.code.getName()
}
