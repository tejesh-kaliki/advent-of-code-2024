package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func CountPossibilities(pattern string, towels []string, countMap map[string]int) int {
	if possible, found := countMap[pattern]; found {
		return possible
	}

	total := 0
	for _, towel := range towels {
		if remaining, found := strings.CutPrefix(pattern, towel); found {
			total += CountPossibilities(remaining, towels, countMap)
		}
	}

	countMap[pattern] = total
	return total
}

func SolveParts(patterns, towels []string) (int, int) {
	possible := 0
	total := 0
	countMap := map[string]int{"": 1}
	for _, pattern := range patterns {
		if possibilities := CountPossibilities(pattern, towels, countMap); possibilities > 0 {
			total += possibilities
			possible++
		}
	}
	return possible, total
}

func GetTowelMap(towels []string) map[byte][]string {
	towelMap := make(map[byte][]string)
	for _, towel := range towels {
		values, found := towelMap[towel[0]]
		if !found {
			values = []string{towel}
		} else {
			values = append(values, towel)
		}
		towelMap[towel[0]] = values
	}
	return towelMap
}

func ReadInput(input string) (towels, patterns []string) {
	towelsInput, remaining, _ := strings.Cut(input, "\n\n")
	towels = strings.Split(towelsInput, ", ")
	patterns = strings.Split(remaining, "\n")
	return
}

func main() {
	towels, patterns := ReadInput(input)
	part1Sol, part2Sol := SolveParts(patterns, towels)
	fmt.Println("Part 1 Solution:", part1Sol)
	fmt.Println("Part 1 Solution:", part2Sol)
}
