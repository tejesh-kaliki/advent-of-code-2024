package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Direction struct {
	Dx, Dy int
}

var ALL_DIRS = []Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Position struct {
	X, Y int
}

func (pos Position) MoveAlong(dir Direction) Position {
	return Position{pos.X + dir.Dx, pos.Y + dir.Dy}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindShortestPath(start, end Position) int {
	xDiff := end.X - start.X
	yDiff := end.Y - start.Y
	return Abs(xDiff) + Abs(yDiff)
}

type Grid struct {
	Width     int
	Height    int
	Obstacles []Position
}

func (grid Grid) IsInBounds(pos Position) bool {
	return 0 <= pos.X && pos.X < grid.Width && 0 <= pos.Y && pos.Y < grid.Height
}

func FindShortestPathWithObstacles(start, end Position, grid Grid) int {
	scores := make([][]int, grid.Width)
	for i := range scores {
		scores[i] = make([]int, grid.Height)
		for j := range scores[i] {
			if i == start.X && j == start.Y {
				scores[i][j] = 0
			} else {
				scores[i][j] = math.MaxInt
			}
		}
	}

	queue := []Position{start}

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		for _, dir := range ALL_DIRS {
			point := next.MoveAlong(dir)
			if !grid.IsInBounds(point) || slices.Contains(grid.Obstacles, point) {
				continue
			}

			curScore := scores[point.X][point.Y]
			newScore := scores[next.X][next.Y] + 1
			if curScore > newScore {
				queue = append(queue, point)
				scores[point.X][point.Y] = newScore
			}
		}
	}

	return scores[end.X][end.Y]
}

func ReadInput(input string, width, height int) Grid {
	lines := strings.Split(input, "\n")
	grid := Grid{width, height, make([]Position, len(lines))}
	for i := range grid.Obstacles {
		pos := Position{}
		fmt.Sscanf(lines[i], "%d,%d", &pos.X, &pos.Y)
		grid.Obstacles[i] = pos
	}

	return grid
}

func main() {
	grid := ReadInput(input, 71, 71)
	start := Position{0, 0}
	end := Position{grid.Width - 1, grid.Height - 1}

	newGrid := Grid{grid.Width, grid.Height, grid.Obstacles[:1024]}
	fmt.Println("Part 1 Solution:", FindShortestPathWithObstacles(start, end, newGrid))

	for i := range grid.Obstacles {
		if i <= 1024 {
			continue
		}
		newGrid := Grid{grid.Width, grid.Height, grid.Obstacles[:i]}
		shortest := FindShortestPathWithObstacles(start, end, newGrid)
		if shortest == math.MaxInt {
			fmt.Println("Part 2 Solution:", grid.Obstacles[len(newGrid.Obstacles)-1])
			break
		}

		fmt.Println(i, shortest)
	}
}
