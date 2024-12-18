package main

import (
	"math"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	testcases := []struct {
		Name  string
		Start Position
		End   Position
		Want  int
	}{
		{
			Name:  "shortest is 0 if start and end are same",
			Start: Position{1, 2},
			End:   Position{1, 2},
			Want:  0,
		},
		{
			Name:  "shortest is change in x if start and end are in same line",
			Start: Position{1, 2},
			End:   Position{5, 2},
			Want:  4,
		},
		{
			Name:  "shortest is abs of change in x if start and end are in same line",
			Start: Position{5, 2},
			End:   Position{1, 2},
			Want:  4,
		},
		{
			Name:  "shortest is change in y if start and end are in same line vertically",
			Start: Position{3, 2},
			End:   Position{3, 5},
			Want:  3,
		},
		{
			Name:  "shortest is sum of xDiff and yDiff",
			Start: Position{1, 2},
			End:   Position{4, 6},
			Want:  7,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			shortest := FindShortestPath(testcase.Start, testcase.End)
			if shortest != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", shortest, testcase.Want)
			}
		})
	}
}

func TestFindShortestPathWithObstacles(t *testing.T) {
	testcases := []struct {
		Name      string
		Start     Position
		End       Position
		Obstacles []Position
		Want      int
	}{
		{
			Name:      "use straight line or L shape if no obstacles in middle",
			Start:     Position{0, 0},
			End:       Position{5, 0},
			Obstacles: []Position{{2, 3}, {4, 5}},
			Want:      5,
		},
		{
			Name:      "infinity if cannot reach the point",
			Start:     Position{0, 0},
			End:       Position{5, 0},
			Obstacles: []Position{{1, 0}, {0, 1}},
			Want:      math.MaxInt,
		},
		{
			Name:      "move around the obstacle if only single",
			Start:     Position{0, 0},
			End:       Position{5, 0},
			Obstacles: []Position{{1, 0}},
			Want:      7,
		},
		{
			Name:  "example input",
			Start: Position{0, 0},
			End:   Position{6, 6},
			Obstacles: []Position{
				{5, 4},
				{4, 2},
				{4, 5},
				{3, 0},
				{2, 1},
				{6, 3},
				{2, 4},
				{1, 5},
				{0, 6},
				{3, 3},
				{2, 6},
				{5, 1},
			},
			Want: 22,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			grid := Grid{7, 7, testcase.Obstacles}
			shortest := FindShortestPathWithObstacles(testcase.Start, testcase.End, grid)
			if shortest != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", shortest, testcase.Want)
			}
		})
	}
}
