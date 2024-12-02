package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jimmeh/AOC2024/go/internal/day2"
	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	filepath := os.Args[1]
	reader := file_reading.NewReader(filepath)

	checker, err := day2.CreateReportSafetyChecker(reader)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	totalSafeReports := checker.TotalSafeReports()
	totalNearlySafeReports := checker.TotalNearlySafeReports()

	fmt.Printf("Total safe reports is %d", totalSafeReports)
	fmt.Println()
	fmt.Printf("Total nearly safe reports is %d", totalNearlySafeReports)
	fmt.Println()
}
