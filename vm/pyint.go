package vm

import (
	"strconv"
)

type PyInt struct {
	val  int
	Dict map[string]func(this *PyInt, args []PyObject) PyObject
}

func NewPyInt(val int) *PyInt {
	dict := make(map[string]func(this *PyInt, args []PyObject) PyObject)
	// TODO
	// calculate with PyFloat
	dict["__add__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyInt(this.val + args[0].(*PyInt).val)
	}
	dict["__sub__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyInt(this.val - args[0].(*PyInt).val)
	}
	dict["__mul__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyInt(this.val * args[0].(*PyInt).val)
	}
	dict["__floordiv__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyInt(this.val / args[0].(*PyInt).val)
	}
	dict["__eq__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyBool(this.val == args[0].(*PyInt).val)
	}
	dict["__lt__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyBool(this.val < args[0].(*PyInt).val)
	}
	dict["__gt__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyBool(this.val > args[0].(*PyInt).val)
	}
	dict["__le__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyBool(this.val <= args[0].(*PyInt).val)
	}
	dict["__ge__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyBool(this.val >= args[0].(*PyInt).val)
	}
	dict["__type__"] = func(this *PyInt, args []PyObject) PyObject {
		return this.getType()
	}
	dict["__str__"] = func(this *PyInt, args []PyObject) PyObject {
		return NewPyStr(this.toString())
	}
	return &PyInt{
		val:  val,
		Dict: dict,
	}
}

func (this *PyInt) getType() *PyType {
	return PyTypes[PyIntType]
}

func (this *PyInt) toString() string {
	return strconv.FormatInt(int64(this.val), 10)
}

func (this *PyInt) getVal() int {
	return this.val
}
