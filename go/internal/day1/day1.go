package day1

import "github.com/Jimmeh/AOC2024/go/internal/file_reading"

type locationChecker struct {
	reader file_reading.LineReader
}

func (locationChecker) TotalDistance() int {
	return 0
}

func CreateLocationChecker(reader file_reading.LineReader) locationChecker {
	return locationChecker{reader}
}
