package main

import (
	"reflect"
	"slices"
	"testing"
)

var part1TestInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestReadGrid(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  Grid
	}{
		{
			Name: "read example grid",
			Input: `0123
1234
8765
9876`,
			Want: Grid{
				Height: 4,
				Width:  4,
				Values: [][]int{
					{0, 1, 2, 3},
					{1, 2, 3, 4},
					{8, 7, 6, 5},
					{9, 8, 7, 6},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := ReadInput(testcase.Input)
			if !reflect.DeepEqual(got, testcase.Want) {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.Want)
			}
		})
	}
}

func TestIdentifStartingPoints(t *testing.T) {
	testcases := []struct {
		Name  string
		Input Grid
		Want  []Position
	}{
		{
			Name:  "read example grid",
			Input: ReadInput("0123\n1234\n8765\n9876"),
			Want:  []Position{{0, 0}},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := testcase.Input.IdentifyStartingPositions()
			if !reflect.DeepEqual(got, testcase.Want) {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.Want)
			}
		})
	}
}

func CheckIfElementsAreSame[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Given inputs have different lengths: got %v, want %v", got, want)
	}
	for _, x := range got {
		if !slices.Contains(want, x) {
			t.Errorf("Unnecessary value present: %v", x)
		}
	}
	for _, x := range want {
		if !slices.Contains(got, x) {
			t.Errorf("Necessary value not present: %v", x)
		}
	}
}

func TestFindNextPossibleLocations(t *testing.T) {
	testcases := []struct {
		Name  string
		Input Grid
		Pos   Position
		Want  []Position
	}{
		{
			Name:  "Can move either up or right at (0,0)",
			Input: ReadInput("0123\n1234\n8765\n9876"),
			Pos:   Position{0, 0},
			Want:  []Position{{0, 1}, {1, 0}},
		},
		{
			Name:  "Can move either up or right at (1,0)",
			Input: ReadInput("0123\n1234\n8765\n9876"),
			Pos:   Position{1, 0},
			Want:  []Position{{2, 0}, {1, 1}},
		},
		{
			Name:  "Can move only UP at (4,2)",
			Input: ReadInput(part1TestInput),
			Pos:   Position{4, 2},
			Want:  []Position{{4, 1}},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := testcase.Input.FindNextPossibleLocations(testcase.Pos)
			CheckIfElementsAreSame(t, got, testcase.Want)
		})
	}
}

func TestFindReachableTops(t *testing.T) {
	testcases := []struct {
		Name  string
		Input Grid
		Start Position
		Want  int
	}{
		{
			Name:  "Can reach one 9 from (0,0)",
			Input: ReadInput("0123\n1234\n8765\n9876"),
			Start: Position{0, 0},
			Want:  1,
		},
		{
			Name:  "Can move to 5 9s from (2,0)",
			Input: ReadInput(part1TestInput),
			Start: Position{2, 0},
			Want:  5,
		},
		{
			Name:  "Can move to 6 9s from (2,0)",
			Input: ReadInput(part1TestInput),
			Start: Position{4, 0},
			Want:  6,
		},
		{
			Name: "Example input",
			Input: ReadInput(`...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`),
			Start: Position{4, 0},
			Want:  2,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := testcase.Input.FindReachableTops(testcase.Start)
			if got != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestPart1Solution(t *testing.T) {
	grid := ReadInput(part1TestInput)
	want := 36
	got := grid.FindTotalScore(grid.FindReachableTops)
	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func TestPart2Solution(t *testing.T) {
	grid := ReadInput(part1TestInput)
	want := 81
	got := grid.FindTotalScore(grid.FindPossibleTrails)
	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}
