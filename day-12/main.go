package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Grid struct {
	Width, Height int
	Values        [][]rune
}

func (grid Grid) At(pos Position) rune {
	return grid.Values[pos.Y][pos.X]
}

func (grid Grid) IsInBounds(pos Position) bool {
	return pos.X >= 0 && pos.X < grid.Width && pos.Y >= 0 && pos.Y < grid.Height
}

type Direction struct {
	Dx, Dy int
}

func (dir Direction) FindPerpDirs() [2]Direction {
	if dir == UP || dir == DOWN {
		return [2]Direction{RIGHT, LEFT}
	}

	return [2]Direction{UP, DOWN}
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

func FindAdjacentValidCells(pos Position, isValid func(pos Position) bool) []Position {
	positions := make([]Position, 0, 4)
	for _, dir := range ALL_DIRS {
		nextPos := pos.MoveAlong(dir)
		if isValid(nextPos) {
			positions = append(positions, nextPos)
		}
	}
	return positions
}

func FindAdjacentCells(pos Position) []PosDir {
	positions := make([]PosDir, 0, 4)
	for _, dir := range ALL_DIRS {
		nextPos := pos.MoveAlong(dir)
		positions = append(positions, PosDir{nextPos, dir})
	}
	return positions
}

type Region []Position

func (region Region) Area() int {
	return len(region)
}

func alwaysTrue[T any](x T) bool {
	return true
}

func (region Region) Perimeter(grid Grid) int {
	total := 0
	for _, pos := range region {
		for _, adjPos := range FindAdjacentValidCells(pos, alwaysTrue) {
			if !slices.Contains(region, adjPos) {
				total += 1
			}
		}
	}
	return total
}

type PosDir struct {
	Pos Position
	Dir Direction
}

func (region Region) NumSides(grid Grid) int {
	total := 0
	posDirs := make([]PosDir, 0)
	for _, pos := range region {
		for _, adjPos := range FindAdjacentCells(pos) {
			if slices.Contains(region, adjPos.Pos) {
				continue
			}

			posDirs = append(posDirs, adjPos)

			perpDirs := adjPos.Dir.FindPerpDirs()
			if slices.Contains(posDirs, PosDir{adjPos.Pos.MoveAlong(perpDirs[0]), adjPos.Dir}) {
				continue
			}
			if slices.Contains(posDirs, PosDir{adjPos.Pos.MoveAlong(perpDirs[1]), adjPos.Dir}) {
				continue
			}

			total += 1
		}
	}
	return total
}

func (region Region) Has(pos Position) bool {
	return slices.Contains(region, pos)
}

func (grid Grid) FindContainingRegion(pos Position) Region {
	region := Region{pos}

	isValidNextPos := func(nextPos Position) bool {
		return grid.IsInBounds(nextPos) && grid.At(pos) == grid.At(nextPos) && !region.Has(nextPos)
	}

	for queue := []Position{pos}; len(queue) > 0; queue = queue[1:] {
		nextPositions := FindAdjacentValidCells(queue[0], isValidNextPos)
		queue = append(queue, nextPositions...)
		region = append(region, nextPositions...)
	}
	return region
}

func (grid Grid) FindTotalScore(scoreFn func(grid Grid, region Region) int) int {
	visited := make([]Position, 0)

	total := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if slices.Contains(visited, Position{x, y}) {
				continue
			}

			region := grid.FindContainingRegion(Position{x, y})
			total += scoreFn(grid, region)
			visited = append(visited, region...)
		}
	}
	return total
}

func (grid Grid) SolveForPart1() int {
	return grid.FindTotalScore(func(grid Grid, region Region) int {
		return region.Area() * region.Perimeter(grid)
	})
}

func (grid Grid) SolveForPart2() int {
	return grid.FindTotalScore(func(grid Grid, region Region) int {
		return region.Area() * region.NumSides(grid)
	})
}

func ReadInput(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := Grid{
		Width:  len(lines[0]),
		Height: len(lines),
		Values: make([][]rune, len(lines)),
	}

	for i, line := range lines {
		grid.Values[i] = []rune(line)
	}
	return grid
}

func main() {
	grid := ReadInput(input)
	fmt.Println("Part 1 Solution:", grid.SolveForPart1())
	fmt.Println("Part 2 Solution:", grid.SolveForPart2())
}
