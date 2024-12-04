package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type XmasCountRule func(i, j int) bool

func isInBound(width, height, i, j int) bool {
	return i >= 0 && j >= 0 && i < height && j < width
}

// Args:
//
//	width, dx corresponds to x, i.e., within line
//	height, dy corresponds to y, i.e., vertical
func CheckStartsWithUsingDirection(lines []string, dx, dy int, word string) XmasCountRule {
	if len(lines) == 0 {
		return func(i, j int) bool { return false }
	}

	width := len(lines[0])
	height := len(lines)
	s := len(word) - 1
	return func(i, j int) bool {
		if !isInBound(width, height, i, j) || !isInBound(width, height, i+dy*s, j+dx*s) {
			return false
		}

		for c := range word {
			if word[c] != lines[i+c*dy][j+c*dx] {
				return false
			}
		}

		return true
	}
}

func XmasCount(input string) int {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return 0
	}

	getDirectionRule := func(dx, dy int) XmasCountRule {
		return CheckStartsWithUsingDirection(lines, dx, dy, "XMAS")
	}

	dirs := []struct{ dx, dy int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	rules := make([]XmasCountRule, len(dirs))
	for i, dir := range dirs {
		rules[i] = getDirectionRule(dir.dx, dir.dy)
	}

	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			for _, rule := range rules {
				if rule(i, j) {
					total += 1
				}
			}
		}
	}
	return total
}

type MasCornerRules struct {
	TopRight    XmasCountRule
	TopLeft     XmasCountRule
	BottomRight XmasCountRule
	BottomLeft  XmasCountRule
}

func GetMasDirectionRulesForLines(lines []string) MasCornerRules {
	getDirectionRule := func(dx, dy int) XmasCountRule {
		return CheckStartsWithUsingDirection(lines, dx, dy, "MAS")
	}

	return MasCornerRules{
		TopRight:    getDirectionRule(-1, 1),
		TopLeft:     getDirectionRule(1, 1),
		BottomRight: getDirectionRule(-1, -1),
		BottomLeft:  getDirectionRule(1, -1),
	}
}

func IsXmasCross(lines []string, i, j int, rules MasCornerRules) bool {
	if len(lines) == 0 {
		return false
	}
	width := len(lines[1])
	height := len(lines)

	if !isInBound(width, height, i, j) || !isInBound(width, height, i+2, j+2) {
		return false
	}

	if (rules.TopLeft(i, j) || rules.BottomRight(i+2, j+2)) && (rules.TopRight(i, j+2) || rules.BottomLeft(i+2, j)) {
		return true
	}
	return false
}

func Count_X_mas_Cross(input string) int {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return 0
	}

	rules := GetMasDirectionRulesForLines(lines)
	total := 0
	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < len(lines[i])-2; j++ {
			if IsXmasCross(lines, i, j, rules) {
				total += 1
			}
		}
	}
	return total
}

func main() {
	fmt.Println("Solution to part 1:", XmasCount(input))
	fmt.Println("Solution to part 2:", Count_X_mas_Cross(input))
}
