package main

import (
	"strings"
	"testing"
)

func TestXmasCount(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{
			Name:  "empty string gives 0",
			Input: "",
			Want:  0,
		},
		{
			Name:  "string with single xmas returns 1",
			Input: "XMAS",
			Want:  1,
		},
		{
			Name:  "string with multiple xmas returns the count",
			Input: "XMASXMAS123",
			Want:  2,
		},
		{
			Name:  "Also count if reverse XMAS starts from 0",
			Input: "SAMX",
			Want:  1,
		},
		{
			Name:  "also count reverse xmas in the input",
			Input: "XMASAMX",
			Want:  2,
		},
		{
			Name:  "do not count lines overflow",
			Input: "XMAS123XM\nAS123SAMX",
			Want:  2,
		},
		{
			Name:  "vertical values also needs to be counted",
			Input: "X123\nM123\nA123\nS123",
			Want:  1,
		},
		{
			Name:  "vertical reverse values also needs to be counted",
			Input: "X12S\nM12A\nA12M\nS12X",
			Want:  2,
		},
		{
			Name:  "combination of horizontal and vertical should work",
			Input: "XMAS\nM12A\nA12M\nSAMX",
			Want:  4,
		},
		{
			Name: "need to count diagonal values",
			Input: `X12S
1MA2
1MA2
X12S`,
			Want: 2,
		},
		{
			Name: "need to count diagonal values",
			Input: `S12X
1AM2
1AM2
S12X`,
			Want: 2,
		},
		{
			Name: "For given test input, give 18",
			Input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			Want: 18,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := XmasCount(testcase.Input)
			if got != testcase.Want {
				t.Errorf("got invalid output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestIsXmasCross(t *testing.T) {
	input := `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	lines := strings.Split(input, "\n")[1:]
	testcases := []struct {
		Name  string
		Input []string
		I, J  int
		Want  bool
	}{
		{"no x-mas if only one diagonal", []string{"MXY", "XAY", "XYS"}, 0, 0, false},
		{"no x-mas at boundary (out of bound)", lines, 0, len(lines) - 1, false},
		{"no x-mas at (0,0)", lines, 0, 0, false},
		{"x-mas at (1,5)", lines, 1, 5, true},
		{"x-mas at (0,1)", lines, 0, 1, true},
		{"x-mas at (6,0)", lines, 6, 0, true},
		{"x-mas at (1,6)", lines, 1, 6, true},
		{"x-mas at (2,1)", lines, 2, 1, true},
		{"x-mas at (2,3)", lines, 2, 3, true},
		{"x-mas at (6,2)", lines, 6, 2, true},
		{"x-mas at (6,4)", lines, 6, 4, true},
		{"x-mas at (6,6)", lines, 6, 6, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			rules := GetMasDirectionRulesForLines(testcase.Input)
			got := IsXmasCross(testcase.Input, testcase.I, testcase.J, rules)
			if got != testcase.Want {
				t.Errorf("got invalid output: got %v, want %v", got, testcase.Want)
			}
		})
	}
}

func TestCount_X_Mas_Cross(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	got := Count_X_mas_Cross(input)
	if got != 9 {
		t.Errorf("got invalid output: got %d, want %d", got, 9)
	}
}
