package main

import (
	"os"
	"fmt"
)

func usage(arg0 string) {
	fmt.Fprintf(os.Stderr, "usage: %s [-o file] file.easm\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if (len(os.Args) <= 2) {
		usage(os.Args[0])
	}
	
	
	var aout int = 0
	var ain int = 0

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		switch os.Args[i] {
		case "-o": 
			i = i + 1
			aout = i
		default: 
			ain = i
		}
	}

	if (ain == 0) {
		usage(os.Args[0])
	}
	
	var outf string = "out.bin"
	if (aout != 0) {
		outf = os.Args[aout]
	}

	b, errin := os.ReadFile(os.Args[ain])
	out, errout := os.Create(outf)

	if errin != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %s\n", os.Args[ain])
		os.Exit(1)
	}

	if errout != nil {
		fmt.Fprintf(os.Stderr, "failed to create file for output: %s\n", outf)
		os.Exit(1)
	}

	s := string(b)
	c := Assemble(s)

	out.Write(c)
}
