package main

import (
	"fmt"
	"os"
)

func main() {
	
	if (len(os.Args) != 2) {
		fmt.Fprintf(os.Stderr, "usage: %s bytecode.bin\n")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s\n", os.Args[1])
		os.Exit(1)
	}

	vm := NewEvm(b)

	returndata := vm.Run()
	
	fmt.Printf("finished execution.\n")

	if (returndata != nil) {
		fmt.Printf("callee returned: %v\n", returndata)
	}
}

