package main

import (
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{
			Name:  "return 1 for single cell",
			Input: "^",
			Want:  1,
		},
		{
			Name:  "return length of col if single col without obstacles and start at bottom",
			Input: ".\n.\n.\n.\n^",
			Want:  5,
		},
		{
			Name:  "return distance to obstacle if single col and start at bottom",
			Input: ".\n#\n.\n.\n^",
			Want:  3,
		},
		{
			Name:  "add the spaces to move to right if more than 1 col",
			Input: "...\n...\n#..\n...\n^..",
			Want:  4,
		},
		{
			Name:  "Ignore obstacle if not in same row",
			Input: "...\n...\n.#.\n...\n^..",
			Want:  5,
		},
		{
			Name:  "Need to consider any obstacle in same row",
			Input: "...\n.#.\n#..\n...\n^..",
			Want:  4,
		},
		{
			Name:  "find nearest obstacle if 2 obstacles are in same col",
			Input: "#..\n...\n#..\n...\n^..",
			Want:  4,
		},
		{
			Name:  "Take distance from guard start pos",
			Input: "...\n...\n^..\n...\n...",
			Want:  3,
		},
		{
			Name:  "Ignore obstacle if behind guard",
			Input: "...\n...\n^..\n...\n#..",
			Want:  3,
		},
		{
			Name: "Keep rotating 90 degrees if multiple obstacles",
			Input: `...
#..
..#
...
^..`,
			Want: 6,
		},
		{
			Name: "exclude already covered cells when counting",
			Input: `#....
....#
.....
...#.
^....`,
			Want: 10,
		},
		{
			Name: "value of given example input is 41",
			Input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			Want: 41,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			obs, guard, size := GetInputGrid(testcase.Input)
			got := FindGaurdPathLength(obs, guard, size)
			if got != testcase.Want {
				t.Errorf("Wrong output returned: got %d, want %d", got, testcase.Want)
			}
		})
	}
}
