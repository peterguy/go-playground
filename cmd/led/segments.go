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
var segments = [7]string{"_", "|", "|", "_", "|", "|", "_"}

// toggle on the segments that display each digit
var digit_segments = map[rune][7]bool{
	'0': {true, true, true, true, true, true, false},
	'1': {false, true, true, false, false, false, false},
	'2': {true, true, false, true, true, false, true},
	'3': {true, true, true, true, false, false, true},
	'4': {false, true, true, false, false, true, true},
	'5': {true, false, true, true, false, true, true},
	'6': {true, false, true, true, true, true, true},
	'7': {true, true, true, false, false, false, false},
	'8': {true, true, true, true, true, true, true},
	'9': {true, true, true, true, false, true, true},
}

func write_top_segments(digit rune, width int) string {
	out := " "
	for i := 0; i < width; i++ {
		if digit_segments[digit][0] {
			out += "_"
		} else {
			out += " "
		}
	}
	out += " "
	return out
}

func write_middle_segments(digit rune, width int, verticals_only bool) string {
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
	return out
}

func write_bottom_segments(digit rune, width int, verticals_only bool) string {
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
	return out
}
