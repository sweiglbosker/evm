package main

import (
	"fmt"

	"github.com/holiman/uint256"
)

func main() {
	s := NewStack()
	x, _ := uint256.FromHex("0xfeedface")
	s.Push(x)
	y := s.Pop()
	fmt.Println(y.String())
}
