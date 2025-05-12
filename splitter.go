package main

import (
	"fmt"
	"github.com/tmc/langchaingo/textsplitter"
	"os"
)

func main() {

	filePath := "scrimba-info.txt"
	text, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Unable to read file %v\n", err)
	}

	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(1000),
		)

	chunks, err := splitter.SplitText(string(text))
	if err != nil {
		fmt.Printf("Unable to split text %v\n", err)
	}

	for i, chunk := range chunks {
		fmt.Printf("Chunk %d:\n%s\n\n", i+1, chunk)
	}
}
