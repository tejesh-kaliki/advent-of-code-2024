package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type BlinkInfo struct {
	Value  int64
	Blinks int
}

var BlinkCache = map[BlinkInfo]int{}

func findNumDigits(num int64) int {
	return len(fmt.Sprint(num))
}

func ApplyBlinkRule(value int64) []int64 {
	if value == 0 {
		return []int64{1}
	}

	numDigits := findNumDigits(value)
	if numDigits%2 == 1 {
		return []int64{value * 2024}
	}

	exp := int64(math.Pow10(numDigits / 2))

	return []int64{value / int64(exp), value % int64(exp)}
}

func GetCountAfterBlinks(value int64, blinks int) int {
	if blinks == 0 {
		return 1
	}
	score := GetTotalElementsAfterBlinks(ApplyBlinkRule(value), blinks-1)
	return score
}

func GetTotalElementsAfterBlinks(values []int64, blinks int) int {
	total := 0
	for _, value := range values {
		score, ok := BlinkCache[BlinkInfo{value, blinks}]
		if !ok {
			score = GetCountAfterBlinks(value, blinks)
			BlinkCache[BlinkInfo{value, blinks}] = score
		}

		total += score
	}
	return total
}

func ReadInput(input string) []int64 {
	numTexts := strings.Split(input, " ")
	values := make([]int64, len(numTexts))
	for i, numText := range numTexts {
		values[i], _ = strconv.ParseInt(numText, 10, 64)
	}
	return values
}

func main() {
	values := ReadInput(input)
	fmt.Println("Part 1 solution:", GetTotalElementsAfterBlinks(values, 25))
	fmt.Println("Part 2 solution:", GetTotalElementsAfterBlinks(values, 75))
}
