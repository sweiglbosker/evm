package main

import (
	"github.com/holiman/uint256"
	"log"
)

type Memory []byte

// store 256b word at offset in memory
func (m *Memory) sw(ost uint64, val *uint256.Int) {
	if ost + 32 > uint64(len(*m)) {
		*m = append(*m, make([]byte, ost + 32 - uint64(len(*m)))...)
	}

	bytes := val.Bytes32()
	copy((*m)[ost:], bytes[:])
}

// store byte at offset in memory
func (m *Memory) sb(ost uint64, val byte) {
	if ost + 1 > uint64(len(*m)) {
		*m = append(*m, make([]byte, ost + 1 - uint64(len(*m)))...)
	}
	
	(*m)[ost] = val
}

// loads 256b word fron offset in memory
func (m *Memory) ld(ost uint64) []byte {
	if ost + 32 > uint64(len(*m)) {
		log.Fatal("trying to load out of memory")
	}

	r := make([]byte, 32)
	copy(r, (*m)[ost:ost+32])
	return r
}

