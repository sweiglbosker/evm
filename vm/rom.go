package main 

import (
	 "os"
	 "fmt"
)

type Rom []byte

func (r *Rom) Fetch(pc *uint64, n uint64) []byte {
	if (*pc + n > uint64(len(*r)) - 1) {
		fmt.Fprintf(os.Stderr, "error: trying to read outside of rom\n")
		os.Exit(1)
	}
	x := make([]byte, n)
	copy(x, (*r)[*pc:*pc + n])
	*pc += n
	return x
}
