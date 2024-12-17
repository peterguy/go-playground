package main

import "fmt"

type Off struct {
	something string
	other     bool
}

type Offer struct {
	Offf func() any
}

var mapOff = map[string]Offer{
	"off": {Offf: func() any { return &Off{} }},
}

func getOff(off string) any {
	return mapOff[off].Offf()
}

func main() {
	x := getOff("off")
	fmt.Printf("%p\n", x)
	y := getOff("off")
	fmt.Printf("%p\n", y)
	switch z := x.(type) {
	case *Off:
		z.something = "hello"
	}

	switch z := y.(type) {
	case *Off:
		z.something = "world"
	}

	switch z := x.(type) {
	case *Off:
		fmt.Printf("%s\n", z.something)
	}

	fmt.Printf("%s\n", x)
}
