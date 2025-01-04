package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numerals_flag := flag.String("numerals", "0123456789", "numerals to display")
	height_flag := flag.Int("height", 1, "the height of the seven-segment display")
	width_flag := flag.Int("width", 1, "the width of the seven-segment display")

	flag.Parse()

	numerals := *numerals_flag
	height := *height_flag
	width := *width_flag

	for _, n := range numerals {
		_, err := strconv.Atoi(string(n))
		if err != nil {
			fmt.Printf("Please give me numerals to display, not %q\n", numerals)
			os.Exit(1)
		}
	}

	// does line-based rendering, which has been superceded by segment-based rendering
	// because segment-based rendering allows for size adjustments more easily
	// line-based rendering does not handle height or width
	// write_input_lines(numerals)

	var lines []string = make([]string, 2*(height-1)+3)

	// write the first line of output across all numerals
	lines[0] = write_top_segments(numerals, width)

	copy(lines[1:], write_middle_segments_with_height(numerals, height, width))

	copy(lines[1+height:], write_bottom_segments_with_height(numerals, height, width))

	for _, line := range lines {
		fmt.Println(line)
	}
}
