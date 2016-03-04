package vm

type PyCode struct {
	name         string
	nestedFunc   []*PyCode
	locals       []string
	cellvars     []string
	freevars     []string
	globals      []string
	constants    []PyObject
	instructions []*PyByteCode
	argcount     int
	Dict         map[string]func(this *PyCode, args []PyObject) PyObject
}

func (this *PyCode) Name() string {
	return this.getName()
}

func (this *PyCode) getType() *PyType {
	return PyTypes[PyCodeType]
}

func (this *PyCode) toString() string {
	return "code(" + this.name + ")"
}

func (this *PyCode) getName() string {
	return this.name
}

func (this *PyCode) getNestedFunc() []*PyCode {
	return this.nestedFunc
}

func (this *PyCode) getLocals() []string {
	return this.locals
}

func (this *PyCode) getFreevars() []string {
	return this.freevars
}

func (this *PyCode) getCellvars() []string {
	return this.cellvars
}

func (this *PyCode) getGlobals() []string {
	return this.globals
}

func (this *PyCode) getConstants() []PyObject {
	return this.constants
}

func (this *PyCode) getInstructions() []*PyByteCode {
	return this.instructions
}

func (this *PyCode) getArgCount() int {
	return this.argcount
}

func (this *PyCode) numLocals() int {
	return len(this.locals)
}

func NewPyCode(name string, nestedFunc []*PyCode, locals []string, cellvars []string,
	freevars []string, globals []string, constants []PyObject, instrs []*PyByteCode, argc int) *PyCode {
	dict := make(map[string]func(this *PyCode, args []PyObject) PyObject)
	dict["__type__"] = func(this *PyCode, args []PyObject) PyObject {
		return this.getType()
	}
	dict["__str__"] = func(this *PyCode, args []PyObject) PyObject {
		return NewPyStr(this.toString())
	}
	return &PyCode{
		name:         name,
		nestedFunc:   nestedFunc,
		locals:       locals,
		cellvars:     cellvars,
		freevars:     freevars,
		globals:      globals,
		constants:    constants,
		instructions: instrs,
		argcount:     argc,
		Dict:         dict,
	}
}
