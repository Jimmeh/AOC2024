package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jimmeh/AOC2024/go/internal/day4"
	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	filepath := os.Args[1]
	reader := file_reading.NewReader(filepath)

	searcher, err := day4.CreateWordSearcher(reader)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	total := searcher.Find("XMAS")
	fmt.Printf("Total words found is %d", total)
	fmt.Println()
}
