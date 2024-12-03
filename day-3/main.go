package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func getMulValueIfStartsWith(text string) int {
	if !strings.HasPrefix(text, "mul(") {
		return 0
	}

	var a, b int
	n, err := fmt.Sscanf(text, "mul(%d,%d)", &a, &b)
	if err != nil || n != 2 {
		return 0
	}
	if a >= 1000 || b >= 1000 || a < 0 || b < 0 {
		return 0
	}

	// For spaces in text. SScanf ignores the spaces
	if !strings.HasPrefix(text, fmt.Sprintf("mul(%d,%d)", a, b)) {
		return 0
	}
	return a * b

}

func TotalMulValue(input string) int {
	total := 0
	for i := range input {
		total += getMulValueIfStartsWith(input[i:])
	}
	return total
}

func TotalMulValueWithEnabling(input string) int {
	total := 0
	enabled := true
	for i := range input {
		text := input[i:]
		switch {
		case strings.HasPrefix(text, "don't()"):
			enabled = false
		case strings.HasPrefix(text, "do()"):
			enabled = true
		case enabled:
			total += getMulValueIfStartsWith(text)
		}
	}
	return total
}

func main() {
	fmt.Println("1st part result:", TotalMulValue(input))
	fmt.Println("2nd part result:", TotalMulValueWithEnabling(input))
}
