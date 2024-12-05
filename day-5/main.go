package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type PageRule struct {
	Before string
	After  string
}

func (rule PageRule) IsRuleFollowed(pages []string) (bool, int, int) {
	beforeIndex := slices.Index(pages, rule.Before)
	afterIndex := slices.Index(pages, rule.After)

	isFollowed := beforeIndex == -1 || afterIndex == -1 || beforeIndex < afterIndex
	return isFollowed, beforeIndex, afterIndex
}

func getPageRule(ruleText string) PageRule {
	before, after, _ := strings.Cut(ruleText, "|")
	return PageRule{Before: before, After: after}
}

func IsUpdateInRightOrder(rules []PageRule, pages []string) bool {
	for _, rule := range rules {
		if isFollowed, _, _ := rule.IsRuleFollowed(pages); !isFollowed {
			return false
		}
	}
	return true
}

func MapLines[T any](text string, mapper func(string) T) []T {
	lines := strings.Split(text, "\n")
	values := make([]T, len(lines))
	for i, line := range lines {
		values[i] = mapper(line)
	}
	return values
}

func FindSumOfMedians(input string) (int, int) {
	ruleSection, updateSection, _ := strings.Cut(input, "\n\n")
	rules := MapLines(ruleSection, getPageRule)
	updates := MapLines(updateSection, func(s string) []string {
		return strings.Split(s, ",")
	})

	totalOfCorrect := 0
	totalOfReordered := 0
	for _, pages := range updates {
		isCorrect := IsUpdateInRightOrder(rules, pages)
		if !isCorrect {
			ReorderUpdates(pages, rules)
		}
		center := len(pages) / 2
		value, _ := strconv.ParseInt(pages[center], 10, 32)

		if isCorrect {
			totalOfCorrect += int(value)
		} else {
			totalOfReordered += int(value)
		}
	}

	return totalOfCorrect, totalOfReordered
}

func ReorderUpdates(pages []string, rules []PageRule) {
	for !IsUpdateInRightOrder(rules, pages) {
		for _, rule := range rules {
			if isFollowed, i, j := rule.IsRuleFollowed(pages); !isFollowed {
				t := pages[i]
				pages[i] = pages[j]
				pages[j] = t
			}
		}
	}

}

func main() {
	part1Sol, part2Sol := FindSumOfMedians(input)
	fmt.Println("Solution to part 1:", part1Sol)
	fmt.Println("Solution to part 2:", part2Sol)
}
