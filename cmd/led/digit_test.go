package main

import (
	"testing"
)

func TestIsIntADigit(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{"valid digit 0", '0', true},
		{"valid digit 9", '9', true},
		{"invalid digit letter", 'a', false},
		{"invalid digit symbol", '*', false},
		{"invalid digit space", ' ', false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsIntADigit(test.input); got != test.expected {
				t.Errorf("IsIntADigit() = %v, want %v", got, test.expected)
			}
		})
	}
}

func TestIsRuneADigit(t *testing.T) {
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{"valid digit 0", '0', true},
		{"valid digit 9", '9', true},
		{"invalid digit letter", 'a', false},
		{"invalid digit symbol", '*', false},
		{"invalid digit space", ' ', false},
		{"invalid unicode char", '世', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRuneADigit(tt.input); got != tt.expected {
				t.Errorf("IsRuneADigit() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDigitFromChar(t *testing.T) {
	tests := []struct {
		name        string
		input       byte
		expected    Digit
		expectError bool
	}{
		{"valid digit 0", '0', Zero, false},
		{"valid digit 9", '9', Nine, false},
		{"invalid digit letter", 'a', Zero, true},
		{"invalid digit symbol", '*', Zero, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DigitFromChar(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("DigitFromChar() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && got != tt.expected {
				t.Errorf("DigitFromChar() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDigitsFromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []Digit
		expectError bool
	}{
		{"valid digits", "123", []Digit{One, Two, Three}, false},
		{"empty string", "", []Digit{}, false},
		{"invalid digit", "12a3", nil, true},
		{"special chars", "1#2", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DigitsFromString(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("DigitsFromString() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError {
				if len(got) != len(tt.expected) {
					t.Errorf("DigitsFromString() length = %v, want %v", len(got), len(tt.expected))
					return
				}
				for i := range got {
					if got[i] != tt.expected[i] {
						t.Errorf("DigitsFromString() at index %d = %v, want %v", i, got[i], tt.expected[i])
					}
				}
			}
		})
	}
}

func TestDigit_String(t *testing.T) {
	tests := []struct {
		name     string
		digit    Digit
		expected string
	}{
		{"Zero", Zero, "0"},
		{"Five", Five, "5"},
		{"Nine", Nine, "9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.digit.String(); got != tt.expected {
				t.Errorf("Digit.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
