package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-enry/go-enry/v2"
	"github.com/go-enry/go-enry/v2/data"
	"github.com/grafana/regexp"
)

// UnionRegExps separates values with a | operator to create a string
// representing a union of regexp patterns.
func UnionRegExps(values []string) string {
	if len(values) == 0 {
		// As a regular expression, "()" and "" are equivalent so this
		// condition wouldn't ordinarily be needed to distinguish these
		// values. But, our internal search engine assumes that ""
		// implies "no regexp" (no values), while "()" implies "match
		// empty regexp" (all values) for file patterns.
		return ""
	}
	if len(values) == 1 {
		// Cosmetic format for regexp value, wherever this happens to be
		// pretty printed.
		return values[0]
	}
	return "(?:" + strings.Join(values, ")|(?:") + ")"
}

// filenamesFromLanguage is a map of language name to full filenames
// that are associated with it. This is different from extensions, because
// some languages (like Dockerfile) do not conventionally have an associated
// extension.
var filenamesFromLanguage = func() map[string][]string {
	res := make(map[string][]string, len(data.LanguagesByFilename))
	for filename, languages := range data.LanguagesByFilename {
		for _, language := range languages {
			res[language] = append(res[language], filename)
		}
	}
	for _, v := range res {
		sort.Strings(v)
	}
	return res
}()

// LangToFileRegexp converts a lang: parameter to its corresponding file
// patterns for file filters. The lang value must be valid, cf. validate.go
func LangToFileRegexp(lang string) string {
	lang, _ = enry.GetLanguageByAlias(lang) // Invariant: lang is valid.
	extensions := enry.GetLanguageExtensions(lang)
	patterns := make([]string, len(extensions))
	for i, e := range extensions {
		// Add `\.ext$` pattern to match files with the given extension.
		patterns[i] = regexp.QuoteMeta(e) + "$"
	}
	for _, filename := range filenamesFromLanguage[lang] {
		patterns = append(patterns, "(^|/)"+regexp.QuoteMeta(filename)+"$")
	}
	return UnionRegExps(patterns)
}

func PrintFileNamesAndExtensionsForLanguage(language string) {
	language, _ = enry.GetLanguageByAlias(language) // Normalize language name

	fmt.Printf("File names and extensions for %s:\n", language)

	// Print file names
	if fileNames, ok := filenamesFromLanguage[language]; ok && len(fileNames) > 0 {
		fmt.Println("File names:")
		for _, fileName := range fileNames {
			fmt.Printf("- %s\n", fileName)
		}
	} else {
		fmt.Println("No specific file names found for this language.")
	}

	// Print extensions
	extensions := enry.GetLanguageExtensions(language)
	if len(extensions) > 0 {
		fmt.Println("Extensions:")
		for _, ext := range extensions {
			fmt.Printf("- %s\n", ext)
		}
	} else {
		fmt.Println("No extensions found for this language.")
	}
}

func main() {
	language := os.Args[1]
	PrintFileNamesAndExtensionsForLanguage(language)
	fileRegexp := LangToFileRegexp(language)
	fmt.Println(fileRegexp)
}
