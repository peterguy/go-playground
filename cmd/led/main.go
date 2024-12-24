package main

import (
	"fmt"
	"os"
	"strconv"
)

var digit_led = map[rune][3]string{
	'0': {" _ ", "| |", "|_|"},
	'1': {"   ", "  |", "  |"},
	'2': {" _ ", " _|", "|_ "},
	'3': {" _ ", " _|", " _|"},
	'4': {"   ", "|_|", "  |"},
	'5': {" _ ", "|_ ", " _|"},
	'6': {" _ ", "|_ ", "|_|"},
	'7': {" _ ", "  |", "  |"},
	'8': {" _ ", "|_|", "|_|"},
	'9': {" _ ", "|_|", " _|"},
}

func write_top(digit rune) {
	fmt.Print(digit_led[digit][0])
}

func write_middle(digit rune) {
	fmt.Print(digit_led[digit][1])
}

func write_bottom(digit rune) {
	fmt.Print(digit_led[digit][len(digit_led[digit])-1])
}

func write_digit(digit rune) {
	for _, row := range digit_led[digit] {
		fmt.Println(row)
	}
}

func main() {
	input := os.Args[1]
	_, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Please give me integers to display, not %s\n", input)
		os.Exit(1)
	}
	for _, digit := range input {
		write_top(digit)
	}
	fmt.Println()
	for _, digit := range input {
		write_middle(digit)
	}
	fmt.Println()
	for _, digit := range input {
		write_bottom(digit)
	}
	fmt.Println()

}
