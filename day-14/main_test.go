package main

import (
	"testing"
)

var testInput = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestPositionAfterTime(t *testing.T) {
	space := Space{11, 7}
	testcases := []struct {
		Name    string
		Robot   Robot
		Time    int
		WantPos Vector
	}{
		{
			Name:    "return initial pos if time is 0",
			Robot:   Robot{Vector{1, 2}, Vector{3, 4}},
			Time:    0,
			WantPos: Vector{1, 2},
		},
		{
			Name:    "add velocity and initial pos if time is 1",
			Robot:   Robot{Vector{1, 2}, Vector{3, 4}},
			Time:    1,
			WantPos: Vector{4, 6},
		},
		{
			Name:    "add velocity multiple times based on time",
			Robot:   Robot{Vector{1, 2}, Vector{3, 2}},
			Time:    2,
			WantPos: Vector{7, 6},
		},
		{
			Name:    "wrap position if it overflows",
			Robot:   Robot{Vector{2, 4}, Vector{2, -3}},
			Time:    5,
			WantPos: Vector{1, 3},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := testcase.Robot.PositionAfter(testcase.Time, space)
			if got != testcase.WantPos {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.WantPos)
			}
		})
	}
}

func TestFindQuadrantOfPos(t *testing.T) {
	testcases := []struct {
		Name     string
		Position Vector
		Space    Space
		Want     int
	}{
		{
			Name:     "return 2 if the position is 0,0",
			Position: Vector{0, 0},
			Space:    Space{5, 5},
			Want:     2,
		},
		{
			Name:     "return 1 for top right corner",
			Position: Vector{4, 0},
			Space:    Space{5, 5},
			Want:     1,
		},
		{
			Name:     "return -1 if exactly on the middle of x",
			Position: Vector{2, 0},
			Space:    Space{5, 5},
			Want:     -1,
		},
		{
			Name:     "return 3 for bottom left corner",
			Position: Vector{0, 4},
			Space:    Space{5, 5},
			Want:     3,
		},
		{
			Name:     "return 4 for bottom right corner",
			Position: Vector{4, 4},
			Space:    Space{5, 5},
			Want:     4,
		},
		{
			Name:     "return -1 if exactly on the middle of y",
			Position: Vector{1, 2},
			Space:    Space{5, 5},
			Want:     -1,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := testcase.Space.FindQuadrant(testcase.Position)
			if got != testcase.Want {
				t.Errorf("Got wrong quadrant: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestPart1Solution(t *testing.T) {
	robots := ReadInput(testInput)
	space := Space{11, 7}
	want := 12
	got := SolvePart1(robots, space)

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}
