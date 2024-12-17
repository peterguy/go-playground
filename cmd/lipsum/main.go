package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/loremipsum.v1"
)

func main() {
	count := 50
	unit := "paragraphs"
	if len(os.Args) >= 3 {
		c, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		count = c
		unit = strings.ToLower(os.Args[2])
	}

	generator := loremipsum.New()
	lipsum := ""
	switch unit {
	case "word", "words":
		lipsum = generator.Words(count)
	case "sentence", "sentences":
		lipsum = generator.Sentences(count)
	case "paragraph", "paragraphs":
		lipsum = generator.Paragraphs(count)
	default:
		fmt.Printf("unrecognized unit: %s\n", unit)
		os.Exit(1)
	}
	fmt.Printf("%s\n", lipsum)
}
