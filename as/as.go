package main

import (
	"fmt"
	"strings"
	"github.com/holiman/uint256"

)

func Assemble(code string) []byte {
	bytecode := make([]byte, 0)

	t := strings.Fields(code)

	for i := range t {
		x, found := OpcodeStrings[t[i]]

		if found {
			bytecode = append(bytecode, x)
		} else {
			var n *uint256.Int = nil
			if len(t[i]) > 1 && t[i][1] == 'x' {
				n, _ = uint256.FromHex(t[i])
			} else {
				n, _ = uint256.FromDecimal(t[i])
			} 
			b := n.Bytes32()
			bytecode = append(bytecode, b[:]...)
		}
	}
	return bytecode
}
