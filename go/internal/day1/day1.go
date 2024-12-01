package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func CreateLocationChecker(reader file_reading.LineReader) (locationChecker, error) {
	lines, err := reader.Lines()
	if err != nil {
		return locationChecker{}, err
	}
	left, right := createSortedArrays(lines)
	return locationChecker{left, right}, nil
}

type locationChecker struct {
	left  []int
	right []int
}

func (l locationChecker) TotalDistance() int {
	totalDifference := 0
	for i, lhs := range l.left {
		totalDifference += absoluteDifference(lhs, l.right[i])
	}
	return totalDifference
}

func (l locationChecker) TotalSimilarity() int {
	totalSimilarity := 0
	for _, lhs := range l.left {
		count := 0
		for _, rhs := range l.right {
			if lhs == rhs {
				count++
			}
		}
		totalSimilarity += lhs * count
	}
	return totalSimilarity
}

func createSortedArrays(lines []string) ([]int, []int) {
	left, right := parseArrays(lines)

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	return left, right
}

func parseArrays(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, v := range lines {
		values := strings.Split(v, "   ")
		lhs, _ := strconv.Atoi(values[0])
		rhs, _ := strconv.Atoi(values[1])
		left = append(left, lhs)
		right = append(right, rhs)
	}
	return left, right
}

func absoluteDifference(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
