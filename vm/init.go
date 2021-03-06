package vm

func init() {

	// PyTypes init. Declaration in pytype.go
	PyTypes = make(map[PyTypeId]*PyType)
	PyTypes[PyBoolType] = NewPyType("bool", PyBoolType)
	PyTypes[PyCodeType] = NewPyType("code", PyCodeType)
	PyTypes[PyFloatType] = NewPyType("float", PyFloatType)
	PyTypes[PyFunctionType] = NewPyType("function", PyFunctionType)
	PyTypes[PyIntType] = NewPyType("int", PyIntType)
	PyTypes[PyListType] = NewPyType("list", PyListType)
	PyTypes[PyNoneType] = NewPyType("none", PyNoneType)
	PyTypes[PyStrType] = NewPyType("str", PyStrType)
	PyTypes[PyTupleType] = NewPyType("tuple", PyTupleType)
	PyTypes[PyTypeType] = NewPyType("type", PyTypeType)
	// TODO
	// other type

	// Init opIndex and cmpOp. Declaration in pybytecode.go
	opIndex = map[string]PyOpCode{
		"STOP_CODE":            STOP_CODE,
		"NOP":                  NOP,
		"POP_TOP":              POP_TOP,
		"ROT_TWO":              ROT_TWO,
		"ROT_THREE":            ROT_THREE,
		"DUP_TOP":              DUP_TOP,
		"DUP_TOP_TWO":          DUP_TOP_TWO,
		"UNARY_POSITIVE":       UNARY_POSITIVE,
		"UNARY_NEGATIVE":       UNARY_NEGATIVE,
		"UNARY_NOT":            UNARY_NOT,
		"UNARY_INVERT":         UNARY_INVERT,
		"GET_ITER":             GET_ITER,
		"BINARY_POWER":         BINARY_POWER,
		"BINARY_MULTIPLY":      BINARY_MULTIPLY,
		"BINARY_FLOOR_DIVIDE":  BINARY_FLOOR_DIVIDE,
		"BINARY_TRUE_DIVIDE":   BINARY_TRUE_DIVIDE,
		"BINARY_MODULO":        BINARY_MODULO,
		"BINARY_ADD":           BINARY_ADD,
		"BINARY_SUBTRACT":      BINARY_SUBTRACT,
		"BINARY_SUBSCR":        BINARY_SUBSCR,
		"BINARY_LSHIFT":        BINARY_LSHIFT,
		"BINARY_RSHIFT":        BINARY_RSHIFT,
		"BINARY_AND":           BINARY_AND,
		"BINARY_XOR":           BINARY_XOR,
		"BINARY_OR":            BINARY_OR,
		"INPLACE_POWER":        INPLACE_POWER,
		"INPLACE_MULTIPLY":     INPLACE_MULTIPLY,
		"INPLACE_FLOOR_DIVIDE": INPLACE_FLOOR_DIVIDE,
		"INPLACE_TRUE_DIVIDE":  INPLACE_TRUE_DIVIDE,
		"INPLACE_MODULO":       INPLACE_MODULO,
		"INPLACE_ADD":          INPLACE_ADD,
		"INPLACE_SUBTRACT":     INPLACE_SUBTRACT,
		"INPLACE_LSHIFT":       INPLACE_LSHIFT,
		"INPLACE_RSHIFT":       INPLACE_RSHIFT,
		"INPLACE_AND":          INPLACE_AND,
		"INPLACE_XOR":          INPLACE_XOR,
		"INPLACE_OR":           INPLACE_OR,
		"STORE_SUBSCR":         STORE_SUBSCR,
		"DELETE_SUBSCR":        DELETE_SUBSCR,
		"PRINT_EXPR":           PRINT_EXPR,
		"BREAK_LOOP":           BREAK_LOOP,
		"CONTINUE_LOOP":        CONTINUE_LOOP,
		"SET_ADD":              SET_ADD,
		"LIST_APPEND":          LIST_APPEND,
		"MAP_ADD":              MAP_ADD,
		"RETURN_VALUE":         RETURN_VALUE,
		"YIELD_VALUE":          YIELD_VALUE,
		"IMPORT_STAR":          IMPORT_STAR,
		"POP_BLOCK":            POP_BLOCK,
		"POP_EXCEPT":           POP_EXCEPT,
		"END_FINALLY":          END_FINALLY,
		"LOAD_BUILD_CLASS":     LOAD_BUILD_CLASS,
		"SETUP_WITH":           SETUP_WITH,
		"WITH_CLEANUP":         WITH_CLEANUP,
		"STORE_LOCALS":         STORE_LOCALS,
		"STORE_NAME":           STORE_NAME,
		"DELETE_NAME":          DELETE_NAME,
		"UNPACK_SEQUENCE":      UNPACK_SEQUENCE,
		"UNPACK_EX":            UNPACK_EX,
		"STORE_ATTR":           STORE_ATTR,
		"DELETE_ATTR":          DELETE_ATTR,
		"STORE_GLOBAL":         STORE_GLOBAL,
		"DELETE_GLOBAL":        DELETE_GLOBAL,
		"LOAD_CONST":           LOAD_CONST,
		"LOAD_NAME":            LOAD_NAME,
		"BUILD_TUPLE":          BUILD_TUPLE,
		"SELECT_TUPLE":         SELECT_TUPLE,
		"BUILD_LIST":           BUILD_LIST,
		"BUILD_SET":            BUILD_SET,
		"BUILD_MAP":            BUILD_MAP,
		"LOAD_ATTR":            LOAD_ATTR,
		"COMPARE_OP":           COMPARE_OP,
		"IMPORT_NAME":          IMPORT_NAME,
		"IMPORT_FROM":          IMPORT_FROM,
		"JUMP_FORWARD":         JUMP_FORWARD,
		"POP_JUMP_IF_TRUE":     POP_JUMP_IF_TRUE,
		"POP_JUMP_IF_FALSE":    POP_JUMP_IF_FALSE,
		"JUMP_IF_TRUE_OR_POP":  JUMP_IF_TRUE_OR_POP,
		"JUMP_IF_FALSE_OR_POP": JUMP_IF_FALSE_OR_POP,
		"JUMP_ABSOLUTE":        JUMP_ABSOLUTE,
		"FOR_ITER":             FOR_ITER,
		"LOAD_GLOBAL":          LOAD_GLOBAL,
		"SETUP_LOOP":           SETUP_LOOP,
		"SETUP_EXCEPT":         SETUP_EXCEPT,
		"SETUP_FINALLY":        SETUP_FINALLY,
		"STORE_MAP":            STORE_MAP,
		"LOAD_FAST":            LOAD_FAST,
		"STORE_FAST":           STORE_FAST,
		"DELETE_FAST":          DELETE_FAST,
		"LOAD_CLOSURE":         LOAD_CLOSURE,
		"LOAD_DEREF":           LOAD_DEREF,
		"STORE_DEREF":          STORE_DEREF,
		"DELETE_DEREF":         DELETE_DEREF,
		"RAISE_VARARGS":        RAISE_VARARGS,
		"CALL_FUNCTION":        CALL_FUNCTION,
		"MAKE_FUNCTION":        MAKE_FUNCTION,
		"MAKE_CLOSURE":         MAKE_CLOSURE,
		"BUILD_SLICE":          BUILD_SLICE,
		"EXTENDED_ARG":         EXTENDED_ARG,
		"CALL_FUNCTION_VAR":    CALL_FUNCTION_VAR,
		"CALL_FUNCTION_KW":     CALL_FUNCTION_KW,
		"CALL_FUNCTION_VAR_KW": CALL_FUNCTION_VAR_KW,
		"HAVE_ARGUMENT":        HAVE_ARGUMENT,
	}
	opName = []string{
		"STOP_CODE",
		"NOP",
		"POP_TOP",
		"ROT_TWO",
		"ROT_THREE",
		"DUP_TOP",
		"DUP_TOP_TWO",
		"UNARY_POSITIVE",
		"UNARY_NEGATIVE",
		"UNARY_NOT",
		"UNARY_INVERT",
		"GET_ITER",
		"BINARY_POWER",
		"BINARY_MULTIPLY",
		"BINARY_FLOOR_DIVIDE",
		"BINARY_TRUE_DIVIDE",
		"BINARY_MODULO",
		"BINARY_ADD",
		"BINARY_SUBTRACT",
		"BINARY_SUBSCR",
		"BINARY_LSHIFT",
		"BINARY_RSHIFT",
		"BINARY_AND",
		"BINARY_XOR",
		"BINARY_OR",
		"INPLACE_POWER",
		"INPLACE_MULTIPLY",
		"INPLACE_FLOOR_DIVIDE",
		"INPLACE_TRUE_DIVIDE",
		"INPLACE_MODULO",
		"INPLACE_ADD",
		"INPLACE_SUBTRACT",
		"INPLACE_LSHIFT",
		"INPLACE_RSHIFT",
		"INPLACE_AND",
		"INPLACE_XOR",
		"INPLACE_OR",
		"STORE_SUBSCR",
		"DELETE_SUBSCR",
		"PRINT_EXPR",
		"BREAK_LOOP",
		"CONTINUE_LOOP",
		"SET_ADD",
		"LIST_APPEND",
		"MAP_ADD",
		"RETURN_VALUE",
		"YIELD_VALUE",
		"IMPORT_STAR",
		"POP_BLOCK",
		"POP_EXCEPT",
		"END_FINALLY",
		"LOAD_BUILD_CLASS",
		"SETUP_WITH",
		"WITH_CLEANUP",
		"STORE_LOCALS",
		"STORE_NAME",
		"DELETE_NAME",
		"UNPACK_SEQUENCE",
		"UNPACK_EX",
		"STORE_ATTR",
		"DELETE_ATTR",
		"STORE_GLOBAL",
		"DELETE_GLOBAL",
		"LOAD_CONST",
		"LOAD_NAME",
		"BUILD_TUPLE",
		"SELECT_TUPLE",
		"BUILD_LIST",
		"BUILD_SET",
		"BUILD_MAP",
		"LOAD_ATTR",
		"COMPARE_OP",
		"IMPORT_NAME",
		"IMPORT_FROM",
		"JUMP_FORWARD",
		"POP_JUMP_IF_TRUE",
		"POP_JUMP_IF_FALSE",
		"JUMP_IF_TRUE_OR_POP",
		"JUMP_IF_FALSE_OR_POP",
		"JUMP_ABSOLUTE",
		"FOR_ITER",
		"LOAD_GLOBAL",
		"SETUP_LOOP",
		"SETUP_EXCEPT",
		"SETUP_FINALLY",
		"STORE_MAP",
		"LOAD_FAST",
		"STORE_FAST",
		"DELETE_FAST",
		"LOAD_CLOSURE",
		"LOAD_DEREF",
		"STORE_DEREF",
		"DELETE_DEREF",
		"RAISE_VARARGS",
		"CALL_FUNCTION",
		"MAKE_FUNCTION",
		"MAKE_CLOSURE",
		"BUILD_SLICE",
		"EXTENDED_ARG",
		"CALL_FUNCTION_VAR",
		"CALL_FUNCTION_KW",
		"CALL_FUNCTION_VAR_KW",
		"HAVE_ARGUMENT",
	}
	cmpOp = []string{
		"__lt__",
		"__le__",
		"__eq__",
		"__ne__",
		"__gt__",
		"__ge__",
	}

	// Init callStack. Declaration in pyframe.go
	callStack = NewPyStack()

	// Init Globals. Declaration in pyframe.go
	Globals = make(map[string]PyObject)
	Globals["print"] = NewPyBuiltInPrint()
}
