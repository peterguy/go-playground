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
	case "struct_types":
		struct_types()
	case "stuff":
		stuff()
	case "untar":
		untar(os.Args[1])
	case "url_encoding":
		url_encoding()
	case "file_types":
		file_types(os.Args[1])
	case "file_read":
		file_read(os.Args[1])
	case "pointers":
		pointers()
	case "unique":
		unique()
	case "file_name_from_language":
		file_name_from_language()
	case "composition":
		composition()
	case "execute":
		execute()
	case "secure_store":
		secret := os.Args[1]
		err := secure_store("https://example.com", "peterguy", []byte(secret))
		if err != nil {
			fmt.Printf("ERROR storing secret: %v\n", err)
		}
	case "secure_retrieve":
		secret, err := secure_retrieve("https://exmaple.com", "peterguy")
		if err != nil {
			fmt.Printf("ERROR retrieving secret: %v\n", err)
		} else {
			fmt.Printf("secret: %s\n", string(secret))
		}
	default:
		fmt.Printf("No function for program %s\n", program)
	}
}
