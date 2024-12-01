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
	firstLine := lines[0]
	lhs, _ := strconv.Atoi(strings.Split(firstLine, "   ")[0])
	rhs, _ := strconv.Atoi(strings.Split(firstLine, "   ")[1])
	return absoluteDifference(lhs, rhs), nil
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
