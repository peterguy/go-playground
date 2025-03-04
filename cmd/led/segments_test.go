package main

import (
	"testing"
)

func TestSegments(t *testing.T) {
	tests := []struct {
		name     string
		digits   []Digit
		height   int
		width    int
		expected []string
	}{
		{
			name:     "numeral 0",
			digits:   []Digit{Zero},
			height:   1,
			width:    1,
			expected: []string{" _ ", "| |", "|_|"},
		},
		{
			name:     "numeral 1",
			digits:   []Digit{One},
			height:   1,
			width:    1,
			expected: []string{"   ", "  |", "  |"},
		},
		{
			name:     "numeral 2",
			digits:   []Digit{Two},
			height:   1,
			width:    1,
			expected: []string{" _ ", " _|", "|_ "},
		},
		{
			name:     "numeral 3",
			digits:   []Digit{Three},
			height:   1,
			width:    1,
			expected: []string{" _ ", " _|", " _|"},
		},
		{
			name:     "numeral 4",
			digits:   []Digit{Four},
			height:   1,
			width:    1,
			expected: []string{"   ", "|_|", "  |"},
		},
		{
			name:     "numeral 5",
			digits:   []Digit{Five},
			height:   1,
			width:    1,
			expected: []string{" _ ", "|_ ", " _|"},
		},
		{
			name:     "numeral 6",
			digits:   []Digit{Six},
			height:   1,
			width:    1,
			expected: []string{" _ ", "|_ ", "|_|"},
		},
		{
			name:     "numeral 7",
			digits:   []Digit{Seven},
			height:   1,
			width:    1,
			expected: []string{" _ ", "  |", "  |"},
		},
		{
			name:     "numeral 8",
			digits:   []Digit{Eight},
			height:   1,
			width:    1,
			expected: []string{" _ ", "|_|", "|_|"},
		},
		{
			name:     "numeral 9",
			digits:   []Digit{Nine},
			height:   1,
			width:    1,
			expected: []string{" _ ", "|_|", " _|"},
		},
		{
			name:     "all numerals + width and height",
			digits:   []Digit{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine},
			height:   3,
			width:    3,
			expected: []string{" ___       ___  ___       ___  ___  ___  ___  ___ ", "|   |    |    |    ||   ||    |        ||   ||   |", "|   |    |    |    ||   ||    |        ||   ||   |", "|   |    | ___| ___||___||___ |___     ||___||___|", "|   |    ||        |    |    ||   |    ||   |    |", "|   |    ||        |    |    ||   |    ||   |    |", "|___|    ||___  ___|    | ___||___|    ||___| ___|"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := write_all_segments(test.digits, test.height, test.width)
			for i := 0; i < len(result); i++ {
				if result[i] != test.expected[i] {
					t.Errorf("write_all_segments(%q, %d, %d) = %q, want %q", test.digits, test.height, test.width, result, test.expected)
					break
				}
			}
		})
	}
}
