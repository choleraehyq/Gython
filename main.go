// gython project main.go
package main

import (
	"flag"
	"fmt"
	"github.com/choleraehyq/gython/vm"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	file, err := os.Create("gython.perf")
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(file); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	filename := flag.String("filename", "", "File compiled by disassembler (.casm)")
	flag.Parse()

	if *filename == "" {
		log.Println("Please provide a filename.")
		flag.Usage()
		return
	}

	codes, err := vm.Parse(*filename)
	if err != nil {
		log.Println("Error: cannot read correctly from " + *filename)
		return
	}

	foundMain := false
	for _, code := range codes {
		name := code.Name()
		if name == "main" {
			foundMain = true
		}

		vm.Globals[name] = vm.NewPyFunction(code, vm.Globals, nil)
	}

	if !foundMain {
		fmt.Println("Error: No main() function found.")
	} else {
		args := make([]vm.PyObject, 0)
		vm.CallAttr(vm.Globals["main"], "__call__", args)
	}
}
