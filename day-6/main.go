package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Direction struct {
	X, Y int
}

var (
	UP    = Direction{0, -1}
	DOWN  = Direction{0, 1}
	RIGHT = Direction{1, 0}
	LEFT  = Direction{-1, 0}
)

func (dir Direction) Rotate90() Direction {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	panic("This is not supposed to happen")
}

type Position struct {
	X, Y int
}

func NullPosition() Position {
	return Position{-1, -1}
}

func (pos Position) IsNull() bool {
	return pos.X == -1 && pos.Y == -1
}

type Guard struct {
	Pos  Position
	Dir  Direction
	Path []Position
}

type Size struct {
	Width, Height int
}

type Obstacles struct {
	Locations []Position
}

func (size Size) IsInBounds(pos Position) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= size.Width || pos.Y >= size.Height {
		return false
	}
	return true
}

func (guard *Guard) IsInTheWay(obstacle Position) bool {
	if guard.Dir.X == 0 {
		return guard.Pos.X == obstacle.X && (guard.Pos.Y-obstacle.Y)*guard.Dir.Y < 0
	}
	return guard.Pos.Y == obstacle.Y && (guard.Pos.X-obstacle.X)*guard.Dir.X < 0
}

// Excluding current guard pos
func (guard *Guard) FindDistanceToEdge(size Size) int {
	switch guard.Dir {
	case UP:
		return guard.Pos.Y
	case DOWN:
		return size.Height - guard.Pos.Y - 1
	case RIGHT:
		return size.Width - guard.Pos.X - 1
	case LEFT:
		return guard.Pos.X
	}
	panic("Not Possible to happen")
}

func (guard *Guard) FindNextObstacle(obs Obstacles) *Position {
	var nearestObstacle *Position
	minDistance := 100000000
	for _, obstacle := range obs.Locations {
		if guard.IsInTheWay(obstacle) {
			distance := AbsDiff(guard.Pos.X, obstacle.X) + AbsDiff(guard.Pos.Y, obstacle.Y)
			if distance < minDistance {
				nearestObstacle = &obstacle
				minDistance = distance
			}
		}
	}
	return nearestObstacle
}

func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// Get the distance to next position, and whether it is an obstacle
func (guard *Guard) FindDistanceToNextPos(obs Obstacles, size Size) (int, bool) {
	nextObstacle := guard.FindNextObstacle(obs)
	if nextObstacle != nil {
		return AbsDiff(nextObstacle.X, guard.Pos.X) + AbsDiff(nextObstacle.Y, guard.Pos.Y) - 1, true
	}

	return guard.FindDistanceToEdge(size), false
}

func (guard *Guard) NextPosAfter1Time() Position {
	return Position{guard.Pos.X + guard.Dir.X, guard.Pos.Y + guard.Dir.Y}
}

func (guard *Guard) MoveToNextPos(obs Obstacles, size Size) bool {
	nextObstacle := guard.FindNextObstacle(obs)
	nextPos := guard.NextPosAfter1Time()
	for size.IsInBounds(nextPos) && (nextObstacle == nil || nextPos != *nextObstacle) {
		guard.Path = append(guard.Path, nextPos)
		guard.Pos = nextPos
		nextPos = guard.NextPosAfter1Time()
	}
	guard.Dir = guard.Dir.Rotate90()
	return size.IsInBounds(nextPos)
}

func (guard *Guard) CountPositions() int {
	positions := make([]Position, 0)
	total := 0
	total2 := 0
	for _, pos := range guard.Path {
		if !slices.Contains(positions, pos) {
			total += 1
			positions = append(positions, pos)
		} else {
			total2 += 1
		}
	}
	fmt.Println(total2)
	return total
}

func GetInputGrid(input string) (Obstacles, Guard, Size) {
	rows := strings.Split(input, "\n")
	obs := Obstacles{Locations: make([]Position, 0)}
	guard := Guard{NullPosition(), UP, []Position{}}
	for i, row := range rows {
		for j, rune := range row {
			switch rune {
			case '#':
				obs.Locations = append(obs.Locations, Position{j, i})
			case '^':
				guard.Pos = Position{j, i}
				guard.Path = append(guard.Path, guard.Pos)
			}
		}
	}
	return obs, guard, Size{Height: len(rows), Width: len(rows[0])}
}

func FindGaurdPathLength(obs Obstacles, guard Guard, size Size) int {
	for guard.MoveToNextPos(obs, size) {
	}
	return guard.CountPositions()
}

func main() {
	obs, guard, size := GetInputGrid(input)
	fmt.Println(FindGaurdPathLength(obs, guard, size))
}
