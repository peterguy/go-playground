package main

import (
	"fmt"

	"github.com/go-enry/go-enry/v2"
)

func language_from_file_name(fileName string) {
	// Get possible languages for the given file name
	languages := enry.GetLanguages(fileName, nil)

	// Print the list of languages
	fmt.Printf("Possible languages for %s:\n", fileName)
	for i, lang := range languages {
		fmt.Printf("%d. %s\n", i+1, lang)
	}

	// Get the most probable language
	language := enry.GetLanguage(fileName, nil)
	fmt.Printf("\nMost probable language: %s\n", language)
}
