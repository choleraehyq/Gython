# Gython

## Overview 

Gython is a simplified python virtual machine implemented in Go. It reads a source file compiled in a special bytecode format named CoCo which is easier to be parsed than .pyc format. The compiler is example/disassembler.py. 

Gython implement a subset of Python 3 bytecodes(not compatible with Python 2).the feature supported now are below:

* basic data structures: int/float/string/list/tuple/function
* control flow statements: for/while/if
* closure

I wrote this just for understanding python internel and practicing Go, so it will never be completely compatible with CPython.

## TODO:

- [ ] Read .pyc file directly
- [ ] Exception
- [ ] iterator
- [ ] class
- [ ] module

## License

BSD 2-Clause 

## Reference

* [GoPy](https://github.com/flosch/GoPy)
  Another python virtual machine implemented in Go. My implementation is more similar to CPython than it.

* [CoCo](https://github.com/kentdlee/CoCo)
  A C++ implementation of python virtual machine. The CoCo bytecode format is defined by it.
