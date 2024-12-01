package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func main() {
	filepath := os.Args[1]
	reader := file_reading.NewReader(filepath)
	lines, err := reader.Lines()
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	for _, v := range lines {
		fmt.Printf("%s\n", v)
	}
}
