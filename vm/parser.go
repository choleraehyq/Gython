package vm

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) ([]*PyCode, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	codeReader := bufio.NewReader(f)
	codes := make([]*PyCode, 0)

	for l, err := codeReader.ReadString('\n'); true; l, err = codeReader.ReadString('\n') {
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		codes = append(codes, parseFunction(codeReader, l))
	}

	return codes, nil
}

func parseFunction(codeReader *bufio.Reader, l string) *PyCode {
	tokens := splitTokens(l)

	// tokens[0] must be "Function"
	funcName := tokens[1]
	argcount, _ := strconv.Atoi(tokens[2])

	nestedFunc := make([]*PyCode, 0)
	locals, cellvars := make([]string, 0), make([]string, 0)
	freevars, globals := make([]string, 0), make([]string, 0)
	constants := make([]PyObject, 0)
	instrs := make([]*PyByteCode, 0)

	nextLine, _ := codeReader.ReadString('\n')
	for nextLine != "END\n" {
		if strings.HasPrefix(nextLine, "Function") {
			nestedFunc = append(nestedFunc, parseFunction(codeReader, nextLine))
			nextLine, _ = codeReader.ReadString('\n')
		} else if strings.HasPrefix(nextLine, "Constants") {
			constants = parseConstants(nextLine, nestedFunc)
			nextLine, _ = codeReader.ReadString('\n')
		} else if strings.HasPrefix(nextLine, "Locals") {
			locals = splitTokens(nextLine)[1:]
			nextLine, _ = codeReader.ReadString('\n')
		} else if strings.HasPrefix(nextLine, "CellVars") {
			cellvars = splitTokens(nextLine)[1:]
			nextLine, _ = codeReader.ReadString('\n')
		} else if strings.HasPrefix(nextLine, "FreeVars") {
			freevars = splitTokens(nextLine)[1:]
			nextLine, _ = codeReader.ReadString('\n')
		} else if strings.HasPrefix(nextLine, "Globals") {
			globals = splitTokens(nextLine)[1:]
			nextLine, _ = codeReader.ReadString('\n')
		} else if strings.HasPrefix(nextLine, "BEGIN") {
			instrs = parseInstrs(codeReader, &nextLine)
		}
	}

	return NewPyCode(funcName, nestedFunc, locals, cellvars, freevars, globals, constants, instrs, argcount)
}

func parseConstants(l string, nested []*PyCode) []PyObject {
	tokens := splitTokens(l)

	ret := make([]PyObject, 0)

	// tokens[0] must be "Constants"
	i := 1
	strIdx := 0
	for i < len(tokens) {
		if tokens[i] == "None" {
			ret = append(ret, NewPyNone())
			i++
			continue
		}
		if tokens[i] == "True" {
			ret = append(ret, NewPyBool(true))
			i++
			continue
		}
		if tokens[i] == "False" {
			ret = append(ret, NewPyBool(false))
			i++
			continue
		}
		if tmpInt, err := strconv.Atoi(tokens[i]); err == nil {
			ret = append(ret, NewPyInt(tmpInt))
			i++
			continue
		}
		if tmpFloat, err := strconv.ParseFloat(tokens[i], 64); err == nil {
			ret = append(ret, NewPyFloat(tmpFloat))
			i++
			continue
		}
		if strings.HasPrefix("code(", tokens[i]) {
			for _, code := range nested {
				if code.getName() == tokens[i][5:len(tokens[i])-1] {
					ret = append(ret, code)
				}
			}
			i++
			continue
		}
		// string constant
		if strings.HasPrefix(tokens[i], "\"") {
			strIdx++
			ret = append(ret, NewPyStr(getStringInLine(l, strIdx)))
			for !strings.HasSuffix(tokens[i], "\"") {
				i++
			}
			i++
			continue
		}
		// TODO
		// constants of tuple type
	}
	return ret
}

// Two pass
func parseInstrs(codeReader *bufio.Reader, l *string) []*PyByteCode {
	s, _ := codeReader.ReadString('\n')
	lines := make([]string, 0)
	for s != "END\n" {
		lines = append(lines, s)
		s, _ = codeReader.ReadString('\n')
	}
	*l = s

	// First pass will build the label sheet
	labels := make(map[string]int)
	for i, line := range lines {
		tokens := splitTokens(line)
		if strings.HasPrefix(tokens[0], "label") {
			labels[tokens[0]] = i
		}
	}

	// Second pass will create PyByteCode slices
	ret := make([]*PyByteCode, 0)
	for _, line := range lines {
		tokens := splitTokens(line)

		var opName string
		if strings.HasPrefix(tokens[0], "label") {
			opName = tokens[1]
		} else {
			opName = tokens[0]
		}

		var operand int
		if opName == tokens[len(tokens)-1] {
			// No operand
			operand = -1
		} else if strings.HasPrefix(tokens[len(tokens)-1], "label") {
			// operand is a label
			operand = labels[tokens[len(tokens)-1]]
		} else {
			var err error
			operand, err = strconv.Atoi(tokens[len(tokens)-1])
			if err != nil {
				panic(err)
			}
		}

		ret = append(ret, NewPyByteCode(getInstrIdx(opName), operand))
	}

	return ret
}

func splitTokens(l string) []string {
	return strings.FieldsFunc(l, func(r rune) bool {
		switch r {
		case ' ', ':', ',', '/', '\t', '\n':
			return true
		}
		return false
	})
}

func getStringInLine(l string, i int) string {
	tmp := l
	var start, end int
	for j := 1; j <= i; j++ {
		start = strings.Index(tmp, "\"")
		tmp = strings.Replace(tmp, "\"", " ", 1)
		end = strings.Index(tmp, "\"")
		tmp = strings.Replace(tmp, "\"", " ", 1)
	}
	return l[start+1 : end]
}

// TODO
// parseTuple
