package main

import (
	"fmt"
	"os"

	"github.com/go-enry/go-enry/v2"
)

func file_types(malDir string) {
	filePath := malDir + "impls/ocaml/.xml.template"
	langs := enry.GetLanguagesByFilename(filePath, nil, nil)
	fmt.Printf("=== languages by filename ===\n")
	for _, lang := range langs {
		fmt.Printf("%s\n", lang)
	}
	fmt.Printf("========\n")
	langs = enry.GetLanguagesByExtension(filePath, nil, nil)
	fmt.Printf("=== languages by extension ===\n")
	for _, lang := range langs {
		fmt.Printf("%s\n", lang)
	}
	fmt.Printf("========\n")
	content, err := os.ReadFile(filePath)
	if err == nil {
		langs = enry.GetLanguagesByContent(filePath, content, nil)
		fmt.Printf("=== languages by content ===\n")
		for _, lang := range langs {
			fmt.Printf("%s\n", lang)
		}
		fmt.Printf("========\n")
	} else {
		fmt.Printf("ERROR: %s\n", err)
	}
}
