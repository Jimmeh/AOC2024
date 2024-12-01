package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jimmeh/AOC2024/go/internal/day1"
	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	filepath := os.Args[1]
	reader := file_reading.NewReader(filepath)

	checker, err := day1.CreateLocationChecker(reader)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	distance := checker.TotalDistance()
	similarity := checker.TotalSimilarity()
	fmt.Printf("Total distance when ordered is %d", distance)
	fmt.Println()
	fmt.Printf("Total similarity score is %d", similarity)
	fmt.Println()
}
