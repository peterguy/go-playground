package main

// in a 7-segment display, the segments are labeled a-g, spiraling clockwise
// starting with the top segment. We'll use 0-6 instead of a-g
//
//	 -- 0 --
//	|       |
//	5       1
//	|       |
//	 -- 6 --
//	|       |
//	4       2
//	|       |
//	 -- 3 --

// toggle on the segments that display each digit
var digit_segments = map[Digit][7]bool{
	Zero:  {true, true, true, true, true, true, false},
	One:   {false, true, true, false, false, false, false},
	Two:   {true, true, false, true, true, false, true},
	Three: {true, true, true, true, false, false, true},
	Four:  {false, true, true, false, false, true, true},
	Five:  {true, false, true, true, false, true, true},
	Six:   {true, false, true, true, true, true, true},
	Seven: {true, true, true, false, false, false, false},
	Eight: {true, true, true, true, true, true, true},
	Nine:  {true, true, true, true, false, true, true},
}

func write_top_segments(digits []Digit, width int) string {
	var line string
	for _, digit := range digits {
		out := " "
		for i := 0; i < width; i++ {
			if digit_segments[digit][0] {
				out += "_"
			} else {
				out += " "
			}
		}
		out += " "
		line += out
	}
	return line
}

func write_middle_segments(digits []Digit, width int, verticals_only bool) string {
	var line string
	for _, digit := range digits {
		var out string
		if digit_segments[digit][5] {
			out += "|"
		} else {
			out += " "
		}
		for i := 0; i < width; i++ {
			if !verticals_only && digit_segments[digit][6] {
				out += "_"
			} else {
				out += " "
			}
		}
		if digit_segments[digit][1] {
			out += "|"
		} else {
			out += " "
		}
		line += out
	}
	return line
}

func write_bottom_segments(digits []Digit, width int, verticals_only bool) string {
	var line string
	for _, digit := range digits {
		var out string
		if digit_segments[digit][4] {
			out += "|"
		} else {
			out += " "
		}
		for i := 0; i < width; i++ {
			if !verticals_only && digit_segments[digit][3] {
				out += "_"
			} else {
				out += " "
			}
		}
		if digit_segments[digit][2] {
			out += "|"
		} else {
			out += " "
		}
		line += out
	}
	return line
}

func write_middle_segments_with_height(digits []Digit, height, width int) []string {
	var lines []string = make([]string, height)
	line_index := 0
	for h := 1; h < height; h++ {
		lines[line_index] = write_middle_segments(digits, width, true)
		line_index++
	}
	lines[line_index] = write_middle_segments(digits, width, false)
	return lines
}

func write_bottom_segments_with_height(digits []Digit, height, width int) []string {
	var lines []string = make([]string, height)
	line_index := 0
	for h := 1; h < height; h++ {
		lines[line_index] = write_bottom_segments(digits, width, true)
		line_index++
	}
	lines[line_index] = write_bottom_segments(digits, width, false)
	return lines
}

func write_all_segments(digits []Digit, height, width int) []string {

	var lines []string = make([]string, 2*(height-1)+3)

	lines[0] = write_top_segments(digits, width)

	copy(lines[1:], write_middle_segments_with_height(digits, height, width))

	copy(lines[1+height:], write_bottom_segments_with_height(digits, height, width))

	return lines
}
