package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jimmeh/AOC2024/go/internal/day3"
	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	filepath := os.Args[1]
	reader := file_reading.NewReader(filepath)

	calculator, err := day3.CreateInstructionCalculator(reader)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	instructionTotal := calculator.Calculate()
	withConditionsTotal := calculator.CalculateWithConditions()

	fmt.Printf("Sum of all instructions is %d", instructionTotal)
	fmt.Println()
	fmt.Printf("Sum of all instructions with conditions is %d", withConditionsTotal)
	fmt.Println()
}
