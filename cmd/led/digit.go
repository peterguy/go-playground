package main

import (
	"fmt"
	"strconv"
)

const (
	Zero  Digit = 0
	One   Digit = 1
	Two   Digit = 2
	Three Digit = 3
	Four  Digit = 4
	Five  Digit = 5
	Six   Digit = 6
	Seven Digit = 7
	Eight Digit = 8
	Nine  Digit = 9
)

type Digit uint8

func IsIntADigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func IsRuneADigit(b rune) bool {
	return b >= '0' && b <= '9'
}

func DigitFromChar(val byte) (Digit, error) {
	if !IsIntADigit(val) {
		return 0, fmt.Errorf("%d is not a digit", val)
	}
	return Digit(val - '0'), nil
}

func DigitFromRune(val rune) (Digit, error) {
	if !IsRuneADigit(val) {
		return 0, fmt.Errorf("%c is not a digit", val)
	}
	return Digit(val - '0'), nil
}

func DigitsFromString(val string) ([]Digit, error) {
	digits := make([]Digit, len(val))
	for i, d := range val {
		digit, err := DigitFromRune(d)
		if err != nil {
			return nil, fmt.Errorf("invalid digit at position %d: %w", i, err)
		}
		digits[i] = digit
	}
	return digits, nil
}

// implement the Stringer interface
func (d Digit) String() string {
	return strconv.Itoa(int(d))
}
