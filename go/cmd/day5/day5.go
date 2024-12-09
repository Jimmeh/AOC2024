package main

import (
	"fmt"
	"os"

	"github.com/Jimmeh/AOC2024/go/internal/day5"
	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	filepath := os.Args[1]
	reader := file_reading.NewReader(filepath)

	checker := day5.CreateRuleChecker(reader)

	sumOfMiddle := checker.MiddleNumberSum()

	fmt.Printf("Sum of middle number in valid updates is %d", sumOfMiddle)
	fmt.Println()
}
