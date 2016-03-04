package vm

import (
	"strconv"
)

type PyFloat struct {
	val  float64
	Dict map[string]func(this *PyFloat, args []PyObject) PyObject
}

func NewPyFloat(val float64) *PyFloat {
	dict := make(map[string]func(this *PyFloat, args []PyObject) PyObject)
	dict["__add__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyFloat(this.val + args[0].(*PyFloat).val)
	}
	dict["__sub__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyFloat(this.val - args[0].(*PyFloat).val)
	}
	dict["__mul__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyFloat(this.val * args[0].(*PyFloat).val)
	}
	dict["__floordiv__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyFloat(this.val / args[0].(*PyFloat).val)
	}
	dict["__eq__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyBool(this.val == args[0].(*PyFloat).val)
	}
	dict["__lt__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyBool(this.val < args[0].(*PyFloat).val)
	}
	dict["__gt__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyBool(this.val > args[0].(*PyFloat).val)
	}
	dict["__le__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyBool(this.val <= args[0].(*PyFloat).val)
	}
	dict["__ge__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyBool(this.val >= args[0].(*PyFloat).val)
	}
	dict["__float__"] = func(this *PyFloat, args []PyObject) PyObject {
		return this
	}
	dict["__int__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyInt(int(this.val))
	}
	dict["__str__"] = func(this *PyFloat, args []PyObject) PyObject {
		return NewPyStr(this.toString())
	}
	dict["__type__"] = func(this *PyFloat, args []PyObject) PyObject {
		return this.getType()
	}
	return &PyFloat{
		val:  val,
		Dict: dict,
	}
}

func (this *PyFloat) getType() *PyType {
	return PyTypes[PyFloatType]
}

func (this *PyFloat) toString() string {
	return strconv.FormatFloat(this.val, 'f', -1, 64)
}

func (this *PyFloat) getVal() float64 {
	return this.val
}
