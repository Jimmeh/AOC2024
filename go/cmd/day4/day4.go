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
	totalXmas := searcher.FindXmas()
	totalCrosses := searcher.FindCrosses()

	fmt.Printf("Total xmas found is %d", totalXmas)
	fmt.Println()

	fmt.Printf("Total X-Mas found is %d", totalCrosses)
	fmt.Println()
}
