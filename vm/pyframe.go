package vm

// Init in init.go
var callStack *PyStack
var Globals map[string]PyObject

type PyFrame struct {
	code       *PyCode
	PC         int
	locals     map[string]PyObject
	globals    map[string]PyObject
	cellvars   map[string]PyObject
	consts     []PyObject
	opStack    *PyStack
	blockStack *IntStack
}

func NewPyFrame(code *PyCode, args []PyObject, globals map[string]PyObject, consts []PyObject, cellvars map[string]PyObject) *PyFrame {
	ret := new(PyFrame)
	ret.code, ret.globals, ret.consts, ret.cellvars = code, globals, consts, cellvars

	varNames := code.getLocals()

	ret.blockStack, ret.opStack = NewIntStack(), NewPyStack()

	// Bind the parameters' names to the arguments
	argsP := len(args) - 1
	for i := 0; i < len(args); i++ {
		ret.locals[varNames[i]] = args[argsP]
		argsP--
	}

	// Bind the cellvars to their names
	for i := 0; i < len(code.getCellvars()); i++ {
		name := code.getCellvars()[i]
		if _, ok := ret.locals[name]; ok {
			ret.cellvars[name] = ret.locals[name]
		}
	}

	return ret
}

func (this *PyFrame) getPC() int {
	return this.PC
}

func (this *PyFrame) getCode() *PyCode {
	return this.code
}

func (this *PyFrame) getCellName(idx int) string {
	var ret string
	if idx < len(this.code.getCellvars()) {
		ret = this.code.getCellvars()[idx]
	} else {
		ret = this.code.getFreevars()[idx-len(this.code.getCellvars())]
	}
	return ret
}

func (this *PyFrame) execute() PyObject {
	this.PC = 0
	pushFrame(this)
	for {
		nextInstr := this.code.getInstructions()[this.PC]
		this.PC++
		opcode := nextInstr.opcode
		operand := nextInstr.operand
		// TODO
		// python exception handler
		switch opcode {
		case LOAD_FAST:
			u := this.locals[this.code.getLocals()[operand]]
			this.opStack.push(u)
		case LOAD_CONST:
			u := this.consts[operand]
			this.opStack.push(u)
		case LOAD_GLOBAL:
			u := this.globals[this.code.getGlobals()[operand]]
			this.opStack.push(u)
		case STORE_FAST:
			u := this.opStack.pop()
			this.locals[this.code.getLocals()[operand]] = u
		case POP_TOP:
			this.opStack.pop()
		case COMPARE_OP:
			v := this.opStack.pop()
			u := this.opStack.pop()
			args := make([]PyObject, 0)
			args = append(args, v)
			r := CallAttr(u, cmpOp[operand], args)
			this.opStack.push(r)
		case SETUP_LOOP:
			this.blockStack.push(operand)
		case BREAK_LOOP:
			this.PC = this.blockStack.pop()
		case POP_BLOCK:
			this.blockStack.pop()
		case JUMP_FORWARD:
			// This instruction is the same as JUMP_ABSOLUTE in this VM
			this.PC = operand
		case JUMP_ABSOLUTE:
			this.PC = operand
		case POP_JUMP_IF_TRUE:
			u := this.opStack.pop().(*PyBool)
			if u.getVal() {
				this.PC = operand
			}
		case POP_JUMP_IF_FALSE:
			u := this.opStack.pop().(*PyBool)
			if !u.getVal() {
				this.PC = operand
			}
		case RETURN_VALUE:
			u := this.opStack.pop()
			popFrame()
			return u
		case INPLACE_ADD:
			fallthrough
			// In fact INPLACE calculate will not pop out the second topest item
		case BINARY_ADD:
			v := this.opStack.pop()
			u := this.opStack.pop()
			args := make([]PyObject, 0)
			args = append(args, v)
			r := CallAttr(u, "__add__", args)
			this.opStack.push(r)
		case BINARY_SUBTRACT:
			v := this.opStack.pop()
			u := this.opStack.pop()
			args := make([]PyObject, 0)
			args = append(args, v)
			r := CallAttr(u, "__sub__", args)
			this.opStack.push(r)
		case BINARY_MULTIPLY:
			v := this.opStack.pop()
			u := this.opStack.pop()
			args := make([]PyObject, 0)
			args = append(args, v)
			r := CallAttr(u, "__mul__", args)
			this.opStack.push(r)
		case BINARY_FLOOR_DIVIDE:
			v := this.opStack.pop()
			u := this.opStack.pop()
			args := make([]PyObject, 0)
			args = append(args, v)
			r := CallAttr(u, "__floordiv__", args)
			this.opStack.push(r)
		case ROT_TWO:
			v := this.opStack.pop()
			u := this.opStack.pop()
			this.opStack.push(v)
			this.opStack.push(u)
		case CALL_FUNCTION:
			args := make([]PyObject, 0)
			// the parameters
			for i := 0; i < operand; i++ {
				args = append(args, this.opStack.pop())
			}
			// the function
			f := this.opStack.pop()
			r := CallAttr(f, "__call__", args)
			this.opStack.push(r)
		case LOAD_ATTR:
			u := this.opStack.pop()
			v := NewPyAttr(u, this.code.getGlobals()[operand])
			this.opStack.push(v)
		case BINARY_SUBSCR:
			u := this.opStack.pop()
			v := this.opStack.pop()
			args := make([]PyObject, 0)
			args = append(args, u)
			r := CallAttr(v, "__getitem__", args)
			this.opStack.push(r)
		case STORE_SUBSCR:
			u := this.opStack.pop()
			v := this.opStack.pop()
			w := this.opStack.pop()
			args := []PyObject{u, w}
			// The returned value must be None
			_ = CallAttr(v, "__setitem__", args)
		case LOAD_CLOSURE:
			name := this.getCellName(operand)
			this.opStack.push(this.cellvars[name])
		case BUILD_TUPLE:
			args := make([]PyObject, operand)
			for i := 0; i < operand; i++ {
				args[operand-i-1] = this.opStack.pop()
			}
			this.opStack.push(NewPyTuple(args))
		case SELECT_TUPLE:
			tuple := this.opStack.pop().(*PyTuple)
			for i := tuple.size() - 1; i >= 0; i-- {
				this.opStack.push(tuple.getItem(i))
			}
		case BUILD_LIST:
			args := make([]PyObject, operand)
			for i := 0; i < operand; i++ {
				args[operand-i-1] = this.opStack.pop()
			}
			this.opStack.push(NewPyList(args))
		case MAKE_CLOSURE:
			u := this.opStack.pop()
			v := this.opStack.pop()
			f := NewPyFunction(u.(*PyCode), this.globals, v)
			this.opStack.push(f)
		case MAKE_FUNCTION:
			u := this.opStack.pop()
			f := NewPyFunction(u.(*PyCode), this.globals, nil)
			this.opStack.push(f)
		case DUP_TOP:
			this.opStack.push(this.opStack.top())
		case DELETE_FAST:
			delete(this.locals, this.code.getLocals()[operand])
		default:
			panic("UnImplemented instruction.")
		}
	}
}

// For implement PyObject
func (this *PyFrame) getType() *PyType {
	return nil
}

// For implement PyObject
func (this *PyFrame) toString() string {
	return ""
}

func pushFrame(frame *PyFrame) {
	callStack.push(frame)
	if callStack.size() > 1024 {
		panic("callStack overflow.")
	}
}

func popFrame() {
	callStack.pop()
}
