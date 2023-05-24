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
}

func NewEvm(_code []byte) *Evm {
	return &Evm{
		pc: 0,
		stopped: true,
		stack : Stack{},
		code: _code,
	}
}

func (vm *Evm) Start() {
	vm.stopped = false
	for !(vm.stopped) {
		op := vm.code.Fetch(&(vm.pc), 1)[0] 
		fmt.Printf("pc: %d | opcode: %x -> string: %s\n", vm.pc, op, Instructions[op].name)
		if op >= PUSH1 && op <= PUSH32 {
			nb := op - PUSH1 + 1
			fmt.Printf("pushing %d byte value to the stack!\n", vm.pc, nb)
			b := vm.code.Fetch(&(vm.pc), uint64(nb))
		        x := uint256.NewInt(0)
			x = x.SetBytes(b)
			vm.stack.Push(x)
		} else {
			vm.Execute(op)
		}
	}
}

func (vm *Evm) Execute(op byte) { 
	Instructions[op].handler(vm)
}
