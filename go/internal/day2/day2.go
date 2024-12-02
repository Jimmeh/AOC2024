package day2

import (
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func CreateReportSafetyChecker(reader file_reading.LineReader) reportSafetyChecker {
	return reportSafetyChecker{reader}
}

type reportSafetyChecker struct {
	reader file_reading.LineReader
}

func (c reportSafetyChecker) TotalSafeReports() (int, error) {
	lines, err := c.reader.Lines()
	if err != nil {
		return -1, err
	}
	report := []int{}
	for _, entry := range strings.Split(lines[0], " ") {
		num, _ := strconv.Atoi(entry)
		report = append(report, num)
	}

	reportType := Undetermined
	for i := range report {
		if i == 0 {
			continue
		}
		reportType = determineReportType(reportType, report[i-1], report[i])
		if reportType == Unsafe {
			return 0, nil
		}
	}
	return 1, nil
}

func determineReportType(reportType, v1, v2 int) int {
	result := Undetermined
	if v1 < v2 {
		result = Increasing
	}
	if v1 > v2 {
		result = Decreasing
	}
	safe := isSafeChange(v1, v2) &&
		hasNotChangedDirection(reportType, result) &&
		result != Undetermined

	if !safe {
		return Unsafe
	}
	return result
}

func isSafeChange(v1, v2 int) bool {
	change := 0
	if v1 > v2 {
		change = v1 - v2
	} else {
		change = v2 - v1
	}
	return change <= 3
}

func hasNotChangedDirection(old, new int) bool {
	return old == Undetermined || new == old
}

const (
	Undetermined = iota
	Increasing   = iota
	Decreasing   = iota
	Unsafe       = iota
)
