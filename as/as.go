package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/holiman/uint256"
)

func Assemble(code string) []byte {
	bytecode := make([]byte, 0)

	t := strings.Fields(code)
	push := byte(0)

	for i := range t {
		x, found := OpcodeStrings[t[i]]

		if found {
			push = (&x).IsPush()
			fmt.Printf("token: %s, opcode: 0x%x, push = %d\n", t[i], int(x), int(push))
			bytecode = append(bytecode, byte(x))
		} else {
			if push == 0 {
				fmt.Fprintf(os.Stderr, "unexpected token: %s\n", t[i])
				os.Exit(1)
			}
			var n *uint256.Int = nil
			if len(t[i]) > 1 && t[i][1] == 'x' {
				n, _ = uint256.FromHex(t[i])
			} else {
				n, _ = uint256.FromDecimal(t[i])
			} 
			
			if (t[i] == "PUSH") {
				push = byte(n.ByteLen())
			} 


			b := n.Bytes32()
			fmt.Printf("token: NUMBER, value: %s, push = %d\n", t[i], int(push), b[32 - push:])
			bytecode = append(bytecode, b[32-push:]...)
		}
	}
	return bytecode
}
