package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	_, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Please give me integers to display, not %q\n", input)
		os.Exit(1)
	}
	height := 1
	width := 1

	if len(os.Args) > 2 {
		height, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Please give me a height integer, not %q\n", os.Args[2])
			os.Exit(1)
		}
	}

	if len(os.Args) > 3 {
		width, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Please give me a width integer, not %q\n", os.Args[3])
			os.Exit(1)
		}
	}

	// does line-based rendering, which has been superceded by segment-based rendering
	// because segment-based rendering allows for size adjustments more easily
	write_input_lines(input)

	var lines []string = make([]string, 2*(height-1)+3)
	line_index := 0

	line := ""
	for _, digit := range input {
		line += write_top_segments(digit, width)
	}
	lines[line_index] = line
	line_index++

	for h := 1; h < height; h++ {
		line = ""
		for _, digit := range input {
			line += write_middle_segments(digit, width, true)
		}
		lines[line_index] = line
		line_index++
	}

	line = ""
	for _, digit := range input {
		line += write_middle_segments(digit, width, false)
	}
	lines[line_index] = line
	line_index++

	for h := 1; h < height; h++ {
		line = ""
		for _, digit := range input {
			line += write_bottom_segments(digit, width, true)
		}
		lines[line_index] = line
		line_index++
	}

	line = ""
	for _, digit := range input {
		line += write_bottom_segments(digit, width, false)
	}
	lines[line_index] = line

	for _, line := range lines {
		fmt.Println(line)
	}
}
