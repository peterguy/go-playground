package main

import "fmt"

func return_by_value() []byte {
	buf := make([]byte, 10)
	fmt.Printf("address of buf in return_by_value: %p\n", buf)
	return buf
}

func return_by_reference() *[]byte {
	buf := make([]byte, 10)
	fmt.Printf("address of buf in return_by_reference: %p\n", buf)
	return &buf
}

func pointers() {
	buf := return_by_value()
	fmt.Printf("address of buf returned by value: %p\n", buf)
	buf2 := return_by_reference()
	fmt.Printf("address of buf returned by reference: %p\n", *buf2)
}
