package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getValuesFromLine(line string) []int {
	nums := strings.Split(line, " ")
	values := make([]int, len(nums))
	for i, num := range nums {
		value, err := strconv.ParseInt(num, 10, 32)
		if err == nil {
			values[i] = int(value)
		} else {
			fmt.Println(err)
		}
	}
	return values
}

func getChanges(values []int) []int {
	changes := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		changes[i] = values[i+1] - values[i]
	}
	return changes
}

func getSign(value int) int {
	if value > 0 {
		return 1
	}
	return -1
}

func areValuesSafe(values []int) bool {
	changes := getChanges(values)
	if len(changes) == 0 {
		return true
	}

	sign := getSign(changes[0])
	for _, change := range changes {
		change = change * sign
		if change <= 0 || change > 3 {
			return false
		}
	}
	return true
}

func isLineSafe(line string) bool {
	values := getValuesFromLine(line)
	return areValuesSafe(values)
}

func getRemovedList(values []int, index int) []int {
	newValues := make([]int, 0, len(values)-1)
	for i, value := range values {
		if i == index {
			continue
		}
		newValues = append(newValues, value)
	}
	return newValues
}

func isLineSafeWithRemove(line string) bool {
	values := getValuesFromLine(line)

	if areValuesSafe(values) {
		return true
	}

	for i := range values {
		removedList := getRemovedList(values, i)
		if areValuesSafe(removedList) {
			return true
		}
	}

	return false
}

func SafeReportCount(input string, safeFn func(string) bool) int {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		if safeFn(line) {
			count += 1
		}
	}

	return count
}

func main() {
	fmt.Println("Answer for part 1:", SafeReportCount(input, isLineSafe))
	fmt.Println("Answer for part 2:", SafeReportCount(input, isLineSafeWithRemove))
}
