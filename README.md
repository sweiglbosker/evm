evm
===

components
----------

`vm/`: code for the evm interpreter
`as/`: code for an easm assembler (you can also use `evm compile` if you have geth installed)

building
--------

you can build individual go modules like usual

```sh
cd vm & go build
cd ../as & go build
```

you can also use the makefile to build all of the modules, and run the test of your choice

```sh
TEST=test/add.easm make clean test      # don't need to specify add.easm, it is the default test
```

as
----

you can assemble files by specifying the assembly file, and optionally the name of your output binary using the `-o` flag

```sh
./as test/add.src -o test/add.bin
```

if you want to compare the output with something like the `evm` tool from geth, you can use xxd with the `-p` flag

```sh
evm compile test/add.src
xxd -p test/add.bin
```

vm
---

running bytecode:

```sh
./vm/evm test/add.bin
```

programming
-----------

right now, only a small subset of evm opcodes are supported, and there is no support for calldata, storage, branching, gas fees and much more

the only instruction that takes a parameter is the `PUSH` instruction. It will take the specified sub 256b number in hex or decimal format, and push it onto the stack.

```
PUSH 4 # pushes the value 4 to the stack
```

this `PUSH` instruction is then used to pass values into other instructions

```
PUSH 4
PUSH 9
ADD
```

the result of this `ADD` instruction (13) will then be pushed to the stack 

if you need more info on some operations, check the [opcode reference](https://ethereum.org/en/developers/docs/evm/opcodes)

if you want your program to return some data, you first need store it in memory. Then you push the offset and length (in bytes) of this memory to the stack, and use the return instruction

```
PUSH 9
PUSH 4
ADD
PUSH 3  // the offset for MSTORE8
MSTORE8 // the value we want to store is already on the stack after the ADD instruction
PUSH 1  // the size in bytes of our return data
PUSH 3  // offset again
RETURN 
```

how does this connect to the ethereum network
---------------------------------------------

- nodes on the ethereum network will take in similiar bytecode, in addition to persistent contract storage, account state, world state, call data, and much more
- modifications to the state of the ethereum blockchain are verified, usually by comparing the root of hash trees computed by nodes 

more notes
----------

- solidity does its own thing, and has its own abi and runtime
- easm is slow and suffers from serious design issues, so ethereum devs want to move to wasm.
