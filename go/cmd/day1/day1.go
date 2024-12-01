package main

import (
	"fmt"
	"log"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	reader := file_reading.NewReader("./input")
	lines, err := reader.Lines()
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	for _, v := range lines {
		fmt.Printf("%s\n", v)
	}
}
