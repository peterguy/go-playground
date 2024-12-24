package main

import (
	"fmt"
)

// this is how the digits are laid out, broken into three "lines"
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

func write_top_line(digit rune) {
	fmt.Print(digit_led[digit][0])
}

func write_middle_line(digit rune) {
	fmt.Print(digit_led[digit][1])
}

func write_bottom_line(digit rune) {
	fmt.Print(digit_led[digit][len(digit_led[digit])-1])
}

func write_input_lines(input string) {
	for _, digit := range input {
		write_top_line(digit)
	}
	fmt.Println()
	for _, digit := range input {
		write_middle_line(digit)
	}
	fmt.Println()
	for _, digit := range input {
		write_bottom_line(digit)
	}
	fmt.Println()
}
