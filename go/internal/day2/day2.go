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
		reportType = determineReportType(report[i-1], report[i])
		if reportType == Unsafe {
			return 0, nil
		}
	}
	return 1, nil
}

func determineReportType(v1, v2 int) int {
	if v1 < v2 {
		return Increasing
	}
	if v1 > v2 {
		return Decreasing
	}
	return Unsafe
}

const (
	Undetermined = iota
	Increasing   = iota
	Decreasing   = iota
	Unsafe       = iota
)
