package main

import (
	"reflect"
	"slices"
	"testing"
)

func CheckIfNodesAreAllSame(t *testing.T, got, want []Position) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Got wrong length of nodes: got %d, want %d", len(got), len(want))
	}

	for _, node := range want {
		if !slices.Contains(got, node) {
			t.Errorf("Necessary node not present: %v", node)
		}
	}

	for _, node := range got {
		if !slices.Contains(want, node) {
			t.Errorf("Got an unnecessary node: %v", node)
		}
	}
}

func TestAntiNodeLocation(t *testing.T) {
	testcases := []struct {
		Name          string
		FirstNode     Position
		SecondNode    Position
		WantAntiNodes [2]Position
	}{
		{
			Name:          "if nodes are same position, antinodes are also same",
			FirstNode:     Position{1, 1},
			SecondNode:    Position{1, 1},
			WantAntiNodes: [2]Position{{1, 1}, {1, 1}},
		},
		{
			Name:          "if nodes are one off vertically, antinodes are also off vertically",
			FirstNode:     Position{1, 3},
			SecondNode:    Position{1, 4},
			WantAntiNodes: [2]Position{{1, 2}, {1, 5}},
		},
		{
			Name:          "if nodes are one off horizontally, antinodes are also off horizontally",
			FirstNode:     Position{3, 1},
			SecondNode:    Position{4, 1},
			WantAntiNodes: [2]Position{{5, 1}, {2, 1}},
		},
		{
			Name:          "node from example",
			FirstNode:     Position{8, 1},
			SecondNode:    Position{5, 2},
			WantAntiNodes: [2]Position{{11, 0}, {2, 3}},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := FindAntiNodeLocations(testcase.FirstNode, testcase.SecondNode, Size{100, 100})
			CheckIfNodesAreAllSame(t, got[:], testcase.WantAntiNodes[:])
		})
	}
}

func TestReadInputGrid(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
	want := Grid{
		Size: Size{12, 12},
		Nodes: map[rune][]Position{
			'0': {{8, 1}, {5, 2}, {7, 3}, {4, 4}},
			'A': {{6, 5}, {8, 8}, {9, 9}},
		},
	}

	got := ReadInputGrid(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Got wrong output: got %v, want %v", got, want)
	}
}

func TestFindAllAntiNodes(t *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`
	wantLen := 14

	grid := ReadInputGrid(input)
	antiNodes := FindAllAntiNodes(grid, FindAntiNodeLocations)

	if len(antiNodes) != wantLen {
		t.Errorf("Got wrong output length: got %d, want %d", len(antiNodes), wantLen)
	}

	antiNodes = FindAllAntiNodes(grid, FindAllPointsAlongSlope)
	wantLen = 34
	want := []Position{
		{0, 0}, {1, 0}, {6, 0}, {11, 0},
		{1, 1}, {3, 1}, {8, 1},
		{2, 2}, {4, 2}, {5, 2}, {10, 2},
		{2, 3}, {3, 3}, {7, 3},
		{4, 4}, {9, 4},
		{1, 5}, {5, 5}, {6, 5}, {11, 5},
		{3, 6}, {6, 6},
		{0, 7}, {5, 7}, {7, 7},
		{2, 8}, {8, 8},
		{4, 9}, {9, 9},
		{1, 10}, {10, 10},
		{3, 11}, {10, 11}, {11, 11},
	}

	CheckIfNodesAreAllSame(t, antiNodes, want)
}

func TestFindAllPointsAlongSlope(t *testing.T) {
	pos1 := Position{4, 5}
	pos2 := Position{3, 3}
	size := Size{10, 10}
	want := []Position{{2, 1}, {5, 7}, {6, 9}, {4, 5}, {3, 3}}

	got := FindAllPointsAlongSlope(pos1, pos2, size)

	CheckIfNodesAreAllSame(t, got, want)
}
