package main

import (
	 "github.com/holiman/uint256"
)

type Handler func(vm *Evm)

type Instruction struct {
	name string
	n    int // number of params
	handler Handler  
}

type Inst Instruction // shorthand

// i miss c macros :(
var Instructions = [256]Instruction{
	STOP: {
        	name: "STOP",
		n: 0,
		handler: func(vm *Evm) { vm.stopped = true },
	} ,
	ADD: {
		name: "ADD", 
		n: 2, 
		handler: func(vm *Evm) { 
			o := uint256.NewInt(0)
			vm.stack.Push(o.Add(vm.stack.Pop(), vm.stack.Pop()))
		},
	},
	MUL: {
		n: 2, 
		name: "MUL",
		handler: func(vm *Evm) {
			o := uint256.NewInt(0)
			vm.stack.Push(o.Mul(vm.stack.Pop(), vm.stack.Pop())) 
		},
	},
	SUB: {
		n: 2, 
		name: "SUB",
		handler: func(vm *Evm) {
			o := uint256.NewInt(0)
			vm.stack.Push(o.Sub(vm.stack.Pop(), vm.stack.Pop()))
		},
	},
	DIV: {
		n: 2, 
		name: "DIV",
		handler: func(vm *Evm) { 
			o := uint256.NewInt(0)
			vm.stack.Push(o.Div(vm.stack.Pop(), vm.stack.Pop()))
		},	
	},
	SDIV: { 
		n: 2, 
		name: "SDIV", 
		handler: func(vm *Evm) {
			o := uint256.NewInt(0)
			vm.stack.Push(o.SDiv(vm.stack.Pop(), vm.stack.Pop())) 
		},
	},
	MOD: {
		n: 2, 
		name: "MOD",
		handler: func(vm *Evm) { 
			o := uint256.NewInt(0)
			vm.stack.Push(o.Mod(vm.stack.Pop(), vm.stack.Pop())) 
		},
	},
	SMOD: {
		n: 2, 
		name: "SMOD", 
		handler: func(vm *Evm) { 
			o := uint256.NewInt(0)
			vm.stack.Push(o.SMod(vm.stack.Pop(), vm.stack.Pop())) 
		},
	},
	ADDMOD: { 
		n: 3, 
		name: "ADDMOD",
		handler: func(vm *Evm) { 
			o := uint256.NewInt(0)
			vm.stack.Push(o.AddMod(vm.stack.Pop(), vm.stack.Pop(), vm.stack.Pop())) 
		},
	},
	MULMOD: {
		n: 3, 
		name: "MULMOD", 
		handler: func(vm *Evm) {
			o := uint256.NewInt(0)
			vm.stack.Push(o.MulMod(vm.stack.Pop(), vm.stack.Pop(), vm.stack.Pop()))
		},
	},
	POP: { 
		n: 1, 
		name: "POP", 
		handler: func(vm *Evm) { 
			vm.stack.Pop() 
		},
	}, 
	MLOAD: { 
		n: 1, 
		name: "MLOAD",
	},
	MSTORE: {
		n: 2, 
		name: "MSTORE", 
	},
	MSTORE8: { 
		n: 2, 
		name: "MSTORE8", 
	},
	PUSH1:  { name: "PUSH1" },
	PUSH2:  { name: "PUSH2" },
	PUSH3:  { name: "PUSH3" },
	PUSH4:  { name: "PUSH4" },
	PUSH5:  { name: "PUSH5" },
	PUSH6:  { name: "PUSH6" },
	PUSH7:  { name: "PUSH7" },
	PUSH8:  { name: "PUSH8" },
	PUSH9:  { name: "PUSH9" },
	PUSH10: { name: "PUSH0" },
	PUSH11: { name: "PUSH11" },
	PUSH12: { name: "PUSH12" },
	PUSH13: { name: "PUSH13" },
	PUSH14: { name: "PUSH14" },
	PUSH15: { name: "PUSH15" },
	PUSH16: { name: "PUSH16" },
	PUSH17: { name: "PUSH17" },
	PUSH18: { name: "PUSH18" },
	PUSH19: { name: "PUSH19" },
	PUSH20: { name: "PUSH20" },
	PUSH21: { name: "PUSH21" },
	PUSH22: { name: "PUSH22" },
	PUSH23: { name: "PUSH23" },
	PUSH24: { name: "PUSH24" },
	PUSH25: { name: "PUSH25" },
	PUSH26: { name: "PUSH26" },
	PUSH27: { name: "PUSH27" },
	PUSH28: { name: "PUSH28" },
	PUSH29: { name: "PUSH29" },
	PUSH30: { name: "PUSH30" },
	PUSH31: { name: "PUSH31" },
	PUSH32: { name: "PUSH32" },
}
