package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Program name required\n")
		os.Exit(1)
	}
	program := os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)
	switch program {
	case "pipe_commands":
		pipe_commands()
	case "lipsum":
		lipsum()
	case "struct_pointers":
		struct_pointers()
	default:
		fmt.Printf("No function for program %s\n", program)
	}
}
