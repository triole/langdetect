package main

import (
	"fmt"

	"github.com/pemistahl/lingua-go"
)

func main() {
	parseArgs()

	content := readFile(CLI.Filename)

	languages := []lingua.Language{
		lingua.English,
		lingua.French,
		lingua.German,
		lingua.Italian,
		lingua.Spanish,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	if language, exists := detector.DetectLanguageOf(content); exists {
		fmt.Println(language)
	}

	// Output: English
}
