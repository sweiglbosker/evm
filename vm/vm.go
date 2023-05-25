package main

import (
	"fmt"
	"github.com/holiman/uint256"
)

type Evm struct {
	code Rom
	stack Stack
	memory Memory
	pc uint64
	stopped bool
	returndata []byte
}

func NewEvm(_code []byte) *Evm {
	return &Evm{
		pc: 0,
		stopped: true,
		stack : Stack{},
		code: _code,
		returndata: nil,
	}
}

func (vm *Evm) Run() []byte {
	vm.stopped = false
	fmt.Printf("code: ", vm.code)
	for !(vm.stopped) {
		s := vm.code.Fetch(&(vm.pc), 1)
		op := Opcode(STOP)
		if (s != nil) {
			op = Opcode(s[0])
		}

		fmt.Printf("pc: %d | opcode: 0x%x -> string: %s\n", vm.pc, op, Instructions[op].name)
		push := (&op).IsPush()
		if (push != 0) {
			b := vm.code.Fetch(&(vm.pc), uint64(push))
		        x := uint256.NewInt(0)
			x = x.SetBytes(b)
			vm.stack.Push(x)
		} else {
			vm.Execute(byte(op))
		}

		if (op == RETURN) {
			return vm.returndata
		} 
	}
	return nil
}

func (vm *Evm) Execute(op byte) { 
	Instructions[op].handler(vm)
}
