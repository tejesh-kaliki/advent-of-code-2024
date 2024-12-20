package main

import (
	"testing"
)

func TestIsThePatternPossible(t *testing.T) {
	testcases := []struct {
		Name    string
		Pattern string
		Towels  []string
		Want    bool
	}{
		{
			Name:    "empty pattern is always possible",
			Pattern: "",
			Towels:  []string{},
			Want:    true,
		},
		{
			Name:    "not possible if pattern is present, but towels are empty",
			Pattern: "rgb",
			Towels:  []string{},
			Want:    false,
		},
		{
			Name:    "possible if pattern matches exactly with a towel",
			Pattern: "rgb",
			Towels:  []string{"rgb"},
			Want:    true,
		},
		{
			Name:    "not possible if starts with towel, but remaining is not possible",
			Pattern: "rgbabc",
			Towels:  []string{"rgb"},
			Want:    false,
		},
		{
			Name:    "possible from given test case",
			Pattern: "brwrr",
			Towels:  []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			Want:    true,
		},
		{
			Name:    "not possible from given test case",
			Pattern: "bbrgwb",
			Towels:  []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			Want:    false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := CountPossibilities(testcase.Pattern, testcase.Towels, map[string]int{"": 1}) > 0
			if got != testcase.Want {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.Want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	towels, patterns := ReadInput(input)
	got, _ := SolveParts(patterns, towels)
	want := 6

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func TestCountPossibilities(t *testing.T) {
	testcases := []struct {
		Name    string
		Pattern string
		Towels  []string
		Want    int
	}{
		{
			Name:    "empty pattern is always possible",
			Pattern: "",
			Towels:  []string{},
			Want:    1,
		},
		{
			Name:    "not possible if pattern is present, but towels are empty",
			Pattern: "rgb",
			Towels:  []string{},
			Want:    0,
		},
		{
			Name:    "possible if pattern matches exactly with a towel",
			Pattern: "rgb",
			Towels:  []string{"rgb"},
			Want:    1,
		},
		{
			Name:    "not possible if starts with towel, but remaining is not possible",
			Pattern: "rgbabc",
			Towels:  []string{"rgb"},
			Want:    0,
		},
		{
			Name:    "possible from given test case",
			Pattern: "brwrr",
			Towels:  []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			Want:    2,
		},
		{
			Name:    "not possible from given test case",
			Pattern: "bbrgwb",
			Towels:  []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			Want:    0,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := CountPossibilities(testcase.Pattern, testcase.Towels, map[string]int{"": 1})
			if got != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	towels, patterns := ReadInput(input)
	_, got := SolveParts(patterns, towels)
	want := 16

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}
