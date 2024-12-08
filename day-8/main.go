package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Position struct {
	X, Y int
}

type Slope struct {
	Dx, Dy int
}

type Size struct {
	Width, Heignt int
}

func FindSlope(pos1, pos2 Position) Slope {
	return Slope{pos1.X - pos2.X, pos1.Y - pos2.Y}
}

func (pos Position) AlongSlope(slope Slope, mul int) Position {
	return Position{pos.X + slope.Dx*mul, pos.Y + slope.Dy*mul}
}

func FindAntiNodeLocations(node1, node2 Position, size Size) []Position {
	slope := FindSlope(node1, node2)
	return []Position{node1.AlongSlope(slope, 1), node2.AlongSlope(slope, -1)}
}

func FindAllPointsAlongSlope(node1, node2 Position, size Size) []Position {
	slope := FindSlope(node1, node2)
	antiNodes := make([]Position, 0)

	for i := 0; ; i++ {
		antiNode := node1.AlongSlope(slope, i)
		if !antiNode.IsInBounds(size) {
			break
		}
		antiNodes = append(antiNodes, antiNode)
	}
	for i := 0; ; i++ {
		antiNode := node2.AlongSlope(slope, -i)
		if !antiNode.IsInBounds(size) {
			break
		}
		antiNodes = append(antiNodes, antiNode)
	}
	return antiNodes
}

func FindAllPairs[T any](elements []T) [][2]T {
	pairCount := len(elements) * (len(elements) - 1) / 2
	result := make([][2]T, 0, pairCount)
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			result = append(result, [2]T{elements[i], elements[j]})
		}
	}
	return result
}

type Grid struct {
	Size  Size
	Nodes map[rune][]Position
}

func (pos Position) IsInBounds(size Size) bool {
	return pos.X >= 0 && pos.X < size.Width && pos.Y >= 0 && pos.Y < size.Heignt
}

func ReadInputGrid(input string) Grid {
	nodePositions := make(map[rune][]Position)

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				positions, ok := nodePositions[char]
				if !ok {
					positions = make([]Position, 0)
				}
				positions = append(positions, Position{j, i})
				nodePositions[char] = positions
			}
		}
	}
	return Grid{
		Size:  Size{len(lines[0]), len(lines)},
		Nodes: nodePositions,
	}
}

func FindAllAntiNodes(grid Grid, findNodesFn func(node1, node2 Position, size Size) []Position) []Position {
	antiNodes := make([]Position, 0)

	insertIfValid := func(node Position) {
		if slices.Contains(antiNodes, node) {
			return
		}
		if !node.IsInBounds(grid.Size) {
			return
		}
		antiNodes = append(antiNodes, node)
	}

	for _, nodes := range grid.Nodes {
		combinations := FindAllPairs(nodes)
		for _, pair := range combinations {
			foundAntiNodes := findNodesFn(pair[0], pair[1], grid.Size)
			for _, node := range foundAntiNodes {
				insertIfValid(node)
			}
		}
	}
	return antiNodes
}

func main() {
	grid := ReadInputGrid(input)
	fmt.Println("Part 1 Solution:", len(FindAllAntiNodes(grid, FindAntiNodeLocations)))
	fmt.Println("Part 2 Solution:", len(FindAllAntiNodes(grid, FindAllPointsAlongSlope)))
}
