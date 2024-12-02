package day2

import (
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func CreateReportSafetyChecker(reader file_reading.LineReader) (reportSafetyChecker, error) {
	reports, err := parseReports(reader)
	if err != nil {
		return reportSafetyChecker{}, err
	}
	return reportSafetyChecker{reports}, nil
}

type reportSafetyChecker struct {
	reports [][]int
}

func parseReports(reader file_reading.LineReader) ([][]int, error) {
	lines, err := reader.Lines()
	if err != nil {
		return nil, err
	}
	reports := [][]int{}
	for _, line := range lines {
		report := []int{}
		for _, entry := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(entry)
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func (c reportSafetyChecker) TotalSafeReports() int {

	safeReports := 0
	for _, report := range c.reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func isSafe(report []int) bool {
	reportType := Undetermined
	for i := range report {
		if i == 0 {
			continue
		}
		reportType = determineReportType(reportType, report[i-1], report[i])
		if reportType == Unsafe {
			return false
		}
	}
	return true
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
