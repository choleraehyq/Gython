package vm

import (
	"reflect"
)

type PyObject interface {
	getType() *PyType
	toString() string
}

func CallAttr(obj PyObject, attrName string, args []PyObject) PyObject {
	argValues := make([]reflect.Value, 0)
	argValues = append(argValues, reflect.ValueOf(obj))
	argValues = append(argValues, reflect.ValueOf(args))
	attrValue := reflect.ValueOf(obj).Elem().FieldByName("Dict").MapIndex(reflect.ValueOf(attrName))
	return (attrValue.Call(argValues))[0].Interface().(PyObject)
}
