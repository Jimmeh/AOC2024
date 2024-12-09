package day5

import (
	"strconv"
	"strings"

	"github.com/Jimmeh/AOC2024/go/internal/file_reading"
)

func CreateRuleChecker(reader file_reading.LineReader) ruleChecker {
	inputs, _ := reader.Lines()
	doingRules := true
	rules := []rule{}
	updates := [][]int{}
	for _, input := range inputs {
		if strings.TrimSpace(input) == "" {
			doingRules = false
			continue
		}
		if doingRules {
			rules = append(rules, CreateRule(input))
		} else {
			updates = append(updates, convertToIntArray(input))
		}
	}
	return ruleChecker{rules, updates}
}

func convertToIntArray(update string) []int {
	strs := strings.Split(update, ",")
	result := []int{}
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		result = append(result, num)
	}
	return result
}

type ruleChecker struct {
	rules   []rule
	updates [][]int
}

func (c ruleChecker) MiddleNumberSum() int {
	return 3
}

func CreateRule(input string) rule {
	split := strings.Split(input, "|")
	left, _ := strconv.Atoi(split[0])
	right, _ := strconv.Atoi(split[1])
	return rule{left, right}
}

type rule struct {
	first  int
	second int
}

func (r rule) Passes(update []int) bool {
	firstIndex := -1
	secondIndex := -1
	for i, page := range update {
		if page == r.first {
			firstIndex = i
		}
		if page == r.second {
			secondIndex = i
		}
		if firstIndex != -1 && secondIndex != -1 {
			break
		}
	}
	return firstIndex == -1 || secondIndex == -1 || firstIndex < secondIndex
}
