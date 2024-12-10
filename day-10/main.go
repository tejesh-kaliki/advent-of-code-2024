package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Grid struct {
	Width, Height int
	Values        [][]int
}

func (grid Grid) ValueAt(pos Position) int {
	return grid.Values[pos.Y][pos.X]
}

type Direction struct {
	Dx, Dy int
}

var (
	UP    = Direction{0, -1}
	DOWN  = Direction{0, 1}
	RIGHT = Direction{1, 0}
	LEFT  = Direction{-1, 0}

	ALL_DIRS = []Direction{UP, RIGHT, DOWN, LEFT}
)

type Position struct {
	X, Y int
}

func (pos Position) MoveAlong(dir Direction) Position {
	return Position{pos.X + dir.Dx, pos.Y + dir.Dy}
}

func (grid Grid) IdentifyStartingPositions() []Position {
	positions := make([]Position, 0)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.ValueAt(Position{x, y}) == 0 {
				positions = append(positions, Position{x, y})
			}
		}
	}
	return positions
}

func (grid Grid) IsInBounds(pos Position) bool {
	return (pos.X >= 0 && pos.X < grid.Width) && (pos.Y >= 0 && pos.Y < grid.Height)
}

func (grid Grid) FindNextPossibleLocations(pos Position) []Position {
	positions := make([]Position, 0, 4)
	for _, dir := range ALL_DIRS {
		nextPos := pos.MoveAlong(dir)
		if !grid.IsInBounds(nextPos) {
			continue
		}

		if grid.ValueAt(nextPos) != grid.ValueAt(pos)+1 {
			continue
		}

		positions = append(positions, nextPos)
	}
	return positions
}

func (grid Grid) FindReachableTops(start Position) int {
	visited := make([]Position, 0)
	count := 0
	for queue := []Position{start}; len(queue) != 0; queue = queue[1:] {
		pos := queue[0]
		if !slices.Contains(visited, pos) {
			visited = append(visited, pos)
			if grid.Values[pos.Y][pos.X] != 9 {
				nextPositions := grid.FindNextPossibleLocations(pos)
				queue = append(queue, nextPositions...)
			} else {
				count++
			}
		}
	}
	return count
}

func (grid Grid) FindPossibleTrails(start Position) int {
	count := 0
	for queue := []Position{start}; len(queue) != 0; queue = queue[1:] {
		if grid.ValueAt(queue[0]) != 9 {
			nextPositions := grid.FindNextPossibleLocations(queue[0])
			queue = append(queue, nextPositions...)
		} else {
			count++
		}
	}
	return count
}

func (grid Grid) FindTotalScore(scoreFinder func(Position) int) int {
	scoreCh := make(chan int)

	startingPos := grid.IdentifyStartingPositions()
	for _, pos := range startingPos {
		go func() {
			scoreCh <- scoreFinder(pos)
		}()
	}

	total := 0
	for i := 0; i < len(startingPos); i++ {
		total += <-scoreCh
	}
	return total
}

func ReadInput(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := Grid{
		Height: len(lines),
		Width:  len(lines[0]),
		Values: make([][]int, len(lines)),
	}
	for i, line := range lines {
		row := make([]int, len(line))
		for j := range line {
			v, err := strconv.ParseInt(line[j:j+1], 10, 8)
			if err != nil {
				row[j] = -1
			} else {
				row[j] = int(v)
			}
		}
		grid.Values[i] = row
	}

	return grid
}

func main() {
	grid := ReadInput(input)
	fmt.Println("Part 1 Solution:", grid.FindTotalScore(grid.FindReachableTops))
	fmt.Println("Part 2 Solution:", grid.FindTotalScore(grid.FindPossibleTrails))
}
