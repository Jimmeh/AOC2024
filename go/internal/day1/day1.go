package day1

import (
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

type locationChecker struct {
	reader file_reading.LineReader
}

func (l locationChecker) TotalDistance() (int, error) {
	lines, err := l.reader.Lines()
	if err != nil {
		return -1, err
	}
	totalDifference := 0
	for _, line := range lines {
		lhs, _ := strconv.Atoi(strings.Split(line, "   ")[0])
		rhs, _ := strconv.Atoi(strings.Split(line, "   ")[1])
		totalDifference += absoluteDifference(lhs, rhs)
	}
	return totalDifference, nil
}

func CreateLocationChecker(reader file_reading.LineReader) locationChecker {
	return locationChecker{reader}
}

func absoluteDifference(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
