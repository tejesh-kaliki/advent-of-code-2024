package main

import (
	"slices"
	"testing"
)

var testInput = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

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

func TestFindRegion(t *testing.T) {
	grid := ReadInput(testInput)
	testcases := []struct {
		Name  string
		Input Grid
		Pos   Position
		Want  Region
	}{
		{
			Name:  "region with only single point if no adjacent places",
			Input: grid,
			Pos:   Position{7, 4},
			Want:  Region{{7, 4}},
		},
		{
			Name:  "select 2 horizontally adjacent cells with same plant",
			Input: Grid{Height: 2, Width: 3, Values: [][]rune{{'A', 'A', 'B'}, {'B', 'B', 'C'}}},
			Pos:   Position{0, 0},
			Want:  Region{{0, 0}, {1, 0}},
		},
		{
			Name:  "select 3 horizontally adjacent cells with same plant",
			Input: Grid{Height: 2, Width: 3, Values: [][]rune{{'A', 'A', 'A'}, {'B', 'B', 'C'}}},
			Pos:   Position{0, 0},
			Want:  Region{{0, 0}, {1, 0}, {2, 0}},
		},
		{
			Name:  "also need to move along towards left",
			Input: Grid{Height: 2, Width: 3, Values: [][]rune{{'A', 'A', 'A'}, {'B', 'B', 'C'}}},
			Pos:   Position{1, 0},
			Want:  Region{{0, 0}, {1, 0}, {2, 0}},
		},
		{
			Name:  "also need to move along towards DOWN",
			Input: Grid{Height: 3, Width: 2, Values: [][]rune{{'A', 'B'}, {'A', 'B'}, {'A', 'C'}}},
			Pos:   Position{0, 0},
			Want:  Region{{0, 0}, {0, 1}, {0, 2}},
		},
		{
			Name:  "also need to move along towards UP",
			Input: Grid{Height: 3, Width: 2, Values: [][]rune{{'A', 'B'}, {'A', 'B'}, {'A', 'C'}}},
			Pos:   Position{0, 1},
			Want:  Region{{0, 0}, {0, 1}, {0, 2}},
		},
		{
			Name:  "a region in given example",
			Input: grid,
			Pos:   Position{8, 0},
			Want: Region{
				{8, 0}, {9, 0},
				{9, 1},
				{9, 2}, {8, 2}, {7, 2},
				{9, 3}, {8, 3}, {7, 3},
				{8, 4},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := testcase.Input.FindContainingRegion(testcase.Pos)
			CheckIfElementsAreSame(t, got, testcase.Want)
		})
	}
}

func TestRegionSides(t *testing.T) {
	grid := ReadInput(testInput)
	testcases := []struct {
		Name      string
		Input     Region
		WantSides int
	}{
		{
			Name:      "region with only single point has 4 sides",
			Input:     Region{{7, 4}},
			WantSides: 4,
		},
		{
			Name:      "region with 2 adjacent points also has 4 sides",
			Input:     Region{{7, 4}, {7, 5}},
			WantSides: 4,
		},
		{
			Name:      "region in L shape has 6 sides",
			Input:     Region{{0, 0}, {1, 0}, {0, 1}},
			WantSides: 6,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			gotSides := testcase.Input.NumSides(grid)
			if gotSides != testcase.WantSides {
				t.Errorf("Got wrong sides: got %d, want %d", gotSides, testcase.WantSides)
			}
		})
	}

}

func TestRegionAreaAndPerimeter(t *testing.T) {
	grid := ReadInput(testInput)
	testcases := []struct {
		Name          string
		Input         Region
		WantArea      int
		WantPerimeter int
	}{
		{
			Name:          "region with only single point",
			Input:         grid.FindContainingRegion(Position{7, 4}),
			WantArea:      1,
			WantPerimeter: 4,
		},
		{
			Name:          "example region from given",
			Input:         grid.FindContainingRegion(Position{0, 0}),
			WantArea:      12,
			WantPerimeter: 18,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			gotArea := testcase.Input.Area()
			if gotArea != testcase.WantArea {
				t.Errorf("Got wrong area: got %d, want %d", gotArea, testcase.WantArea)
			}
			gotPerimeter := testcase.Input.Perimeter(grid)
			if gotPerimeter != testcase.WantPerimeter {
				t.Errorf("Got wrong perimeter: got %d, want %d", gotPerimeter, testcase.WantPerimeter)
			}
		})
	}

}

func TestPart1Solution(t *testing.T) {
	grid := ReadInput(testInput)
	want := 1930
	got := grid.SolveForPart1()
	if got != want {
		t.Errorf("Got wrong solution: got %d, want %d", got, want)
	}
}

func TestPart2Solution(t *testing.T) {
	grid := ReadInput(testInput)
	want := 1206
	got := grid.SolveForPart2()
	if got != want {
		t.Errorf("Got wrong solution: got %d, want %d", got, want)
	}
}
