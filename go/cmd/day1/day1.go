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

	checker := day1.CreateLocationChecker(reader)
	result, err := checker.TotalDistance()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Total distance when ordered is %d", result)
	fmt.Println()
}
