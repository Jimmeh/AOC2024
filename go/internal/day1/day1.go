package day1

import (
	"sort"
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
	left, right := createSortedArrays(lines)
	for i, lhs := range left {
		rhs := right[i]
		totalDifference += absoluteDifference(lhs, rhs)
	}
	return totalDifference, nil
}

func CreateLocationChecker(reader file_reading.LineReader) locationChecker {
	return locationChecker{reader}
}

func createSortedArrays(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, v := range lines {
		values := strings.Split(v, "   ")
		lhs, _ := strconv.Atoi(values[0])
		rhs, _ := strconv.Atoi(values[1])
		left = append(left, lhs)
		right = append(right, rhs)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	return left, right
}

func absoluteDifference(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
