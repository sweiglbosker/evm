package main

type Opcode byte

const (
	STOP = iota
	ADD 
	MUL
	SUB
	DIV
	SDIV
	MOD
	SMOD
	ADDMOD
	MULMOD
	POP = 0x50
	MLOAD = 0x51
	MSTORE = 0x52
	MSTORE8 = 0x53
	PUSH1 = 0x60
	PUSH2 = 0x61
	PUSH3 = 0x62
	PUSH4 = 0x63
	PUSH5 = 0x64
	PUSH6 = 0x65
	PUSH7 = 0x66
	PUSH8 = 0x67
	PUSH9 = 0x68
	PUSH10 = 0x69
	PUSH11 = 0x6a
	PUSH12 = 0x6b
	PUSH13 = 0x6c
	PUSH14 = 0x6d 
	PUSH15 = 0x6e
	PUSH16 = 0x6f
	PUSH17 = 0x70
	PUSH18 = 0x71
	PUSH19 = 0x72
	PUSH20 = 0x73
	PUSH21 = 0x74
	PUSH22 = 0x75
	PUSH23 = 0x76
	PUSH24 = 0x77
	PUSH25 = 0x78
	PUSH26 = 0x79
	PUSH27 = 0x7a
	PUSH28 = 0x7b
	PUSH29 = 0x7c
	PUSH30 = 0x7d
	PUSH31 = 0x7e
	PUSH32 = 0x7f
	RETURN = 0xf3
)

var OpcodeStrings = map[string]Opcode {
	"STOP":    STOP,
	"ADD":     ADD,
	"MUL":     MUL,
	"SUB":     SUB,
	"DIV":     DIV,
	"SDIV":    SDIV,
	"MOD":     MOD,
	"SMOD":    SMOD,
	"ADDMOD":  ADDMOD,
	"MULMOD":  MULMOD,
	"POP":     POP,
	"MLOAD":   MLOAD,
	"MSTORE":  MSTORE,
	"MSTORE8": MSTORE8,
	"PUSH":    PUSH1,
	"PUSH1":   PUSH1,
	"PUSH2":   PUSH2,
	"PUSH3":   PUSH3,
	"PUSH4":   PUSH4,
	"PUSH5":   PUSH5,
	"PUSH6":   PUSH6,
	"PUSH7":   PUSH7,
	"PUSH8":   PUSH8,
	"PUSH9":   PUSH9,
	"PUSH10":   PUSH10,
	"PUSH11":   PUSH11,
	"PUSH12":   PUSH12,
	"PUSH13":   PUSH13,
	"PUSH14":   PUSH14,
	"PUSH15":   PUSH15,
	"PUSH16":   PUSH16,
	"PUSH17":   PUSH17,
	"PUSH18":   PUSH18,
	"PUSH19":   PUSH19,
	"PUSH20":   PUSH20,
	"PUSH21":   PUSH21,
	"PUSH22":   PUSH22,
	"PUSH23":   PUSH23,
	"PUSH24":   PUSH24,
	"PUSH25":   PUSH25,
	"PUSH26":   PUSH26,
	"PUSH27":   PUSH27,
	"PUSH28":   PUSH28,
	"PUSH29":   PUSH29,
	"PUSH30":   PUSH30,
	"PUSH31":   PUSH31,
	"PUSH32":   PUSH32,
	"RETURN":   RETURN,
}

// returns 0 if not push, otherwise returns number of bytes
func (o *Opcode) IsPush() byte {
	d := *o - PUSH1 
	if d >= 0 && d < 31 {
		return byte(d) + 1
	}
	return 0
}
