package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func struct_pointers() {
	// create a new person
	p := Person{"Alice", 25}

	// get a pointer to the person
	ptr := &p

	fmt.Println(p)
	fmt.Println(*ptr)

	// update the person via the pointer
	//p.Age = 26
	ptr.Age = 26

	fmt.Println(p)
	fmt.Println(*ptr)
}
