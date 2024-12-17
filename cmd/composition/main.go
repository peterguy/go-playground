package main

import "fmt"

type Parent interface {
	Hello() string
}

type Child struct {
}

func (c *Child) Hello() string {
	return "Hello from Child"
}

func (c *Child) Goodbye() string {
	return "Goodbye from Child"
}

func main() {
	c := Child{}
	var p Parent = &c
	fmt.Println(p.Hello())
	fmt.Println(c.Goodbye())

}
