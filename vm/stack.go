package main

import (
	"log"

	 "github.com/holiman/uint256"
)

const STACK_CAP = (1 << 10)

type Stack []uint256.Int

func (s *Stack) Push(x *uint256.Int) {
	*s = append(*s, *x)
	if len(*s) + 1 > STACK_CAP {
		log.Fatal("stack overflow")
	}
}

func (s *Stack) Pop() *uint256.Int {
	if len(*s) <= 0 {
		log.Fatal("stack underflow")
	}

	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &r
}

func (s *Stack) Peek() *uint256.Int {
	return &(*s)[len(*s)-1]	
}
