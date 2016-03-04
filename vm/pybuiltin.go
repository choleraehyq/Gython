package vm

import (
	"fmt"
)

// TODO
// other builtin function

type PyBuiltInPrint struct {
	Dict map[string]func(this *PyBuiltInPrint, args []PyObject) PyObject
}

func (this *PyBuiltInPrint) getType() *PyType {
	return PyTypes[PyBuiltInType]
}

func (this *PyBuiltInPrint) toString() string {
	return "<builtin function print>"
}

func NewPyBuiltInPrint() *PyBuiltInPrint {
	ret := new(PyBuiltInPrint)
	ret.Dict = make(map[string]func(this *PyBuiltInPrint, args []PyObject) PyObject)
	ret.Dict["__call__"] = func(this *PyBuiltInPrint, args []PyObject) PyObject {
		str := ""
		for i, arg := range args {
			str += CallAttr(arg, "__str__", make([]PyObject, 0)).toString()
			if i < len(args)-1 {
				str += " "
			}
		}
		fmt.Println(str)
		return NewPyNone()
	}
	return ret
}
