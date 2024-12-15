package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Position struct {
	X, Y int
}

type Direction struct {
	Dx, Dy int
}

var (
	RIGHT = Direction{1, 0}
	LEFT  = Direction{-1, 0}
	UP    = Direction{0, -1}
	DOWN  = Direction{0, 1}
)

type Grid struct {
	Width  int
	Height int
	Cells  [][]rune
	Robot  Position
}

func (grid Grid) Copy() Grid {
	cells := make([][]rune, len(grid.Cells))
	for i, row := range grid.Cells {
		newRow := make([]rune, len(row))
		copy(newRow, row)
		cells[i] = newRow
	}
	return Grid{
		Width:  grid.Width,
		Height: grid.Height,
		Robot:  grid.Robot,
		Cells:  cells,
	}
}

func ReadGridText(gridText string) Grid {
	lines := strings.Split(gridText, "\n")
	grid := Grid{
		Width:  len(lines[0]),
		Height: len(lines),
		Cells:  make([][]rune, len(lines)),
		Robot:  Position{},
	}

	for i, line := range lines {
		row := make([]rune, len(line))
		for j, char := range line {
			row[j] = char
			if char == '@' {
				grid.Robot = Position{j, i}
			}
		}
		grid.Cells[i] = row
	}
	return grid
}

func ReadInputPart1(input string) (Grid, string) {
	gridText, moves, _ := strings.Cut(input, "\n\n")
	return ReadGridText(gridText), moves
}

func ReadGridTextPart2(gridText string) Grid {
	lines := strings.Split(gridText, "\n")
	grid := Grid{
		Width:  len(lines[0]) * 2,
		Height: len(lines),
		Cells:  make([][]rune, len(lines)),
		Robot:  Position{},
	}

	for i, line := range lines {
		row := make([]rune, 0, grid.Width)
		for j, char := range line {
			switch char {
			case 'O':
				row = append(row, []rune("[]")...)
			case '#':
				row = append(row, []rune("##")...)
			case '@':
				row = append(row, []rune("@.")...)
				grid.Robot = Position{2 * j, i}
			case '.':
				row = append(row, []rune("..")...)
			}
		}
		grid.Cells[i] = row
	}
	return grid
}

func ReadInputPart2(input string) (Grid, string) {
	gridText, moves, _ := strings.Cut(input, "\n\n")
	return ReadGridTextPart2(gridText), moves
}

func (grid Grid) Set(pos Position, char rune) {
	grid.Cells[pos.Y][pos.X] = char
}

func (grid Grid) At(pos Position) rune {
	return grid.Cells[pos.Y][pos.X]
}

func (grid Grid) State() string {
	state := ""
	for _, row := range grid.Cells {
		state += string(row) + "\n"
	}
	return state
}

func (pos Position) MoveAlong(dir Direction) Position {
	return Position{pos.X + dir.Dx, pos.Y + dir.Dy}
}

func ApplyMoves(grid *Grid, moves string) {
	for _, move := range moves {
		switch move {
		case '>':
			MoveRobot(grid, RIGHT)
		case '<':
			MoveRobot(grid, LEFT)
		case '^':
			MoveRobot(grid, UP)
		case 'v':
			MoveRobot(grid, DOWN)
		}
	}
}

func CanTheCellBeMoved(grid Grid, pos Position, dir Direction) bool {
	nextPos := pos.MoveAlong(dir)

	switch grid.At(nextPos) {
	case '.': // If next cell is free, can move
		return true
	case '#': // If next cell is wall, cannot move
		return false
	case 'O': // If next cell is box, recursively check if it can be moved
		return CanTheCellBeMoved(grid, nextPos, dir)

	// Part 2 Cases
	case '[': // If next cell is left half of box, recursively check if both halves can be moved
		if dir == LEFT { // To handle infinite recursion
			return CanTheCellBeMoved(grid, nextPos, dir)
		} else {
			return CanTheCellBeMoved(grid, nextPos, dir) && CanTheCellBeMoved(grid, nextPos.MoveAlong(RIGHT), dir)
		}
	case ']': // If next cell is right half of box, recursively check if both halves can be moved
		if dir == RIGHT { // To handle infinite recursion
			return CanTheCellBeMoved(grid, nextPos, dir)
		} else {
			return CanTheCellBeMoved(grid, nextPos, dir) && CanTheCellBeMoved(grid, nextPos.MoveAlong(LEFT), dir)
		}
	default:
		return true
	}
}

// Recursively move all the cells to move the robot in specified direction
func MoveCellsAlongDirection(grid *Grid, pos Position, dir Direction) {
	nextPos := pos.MoveAlong(dir)
	nextObj := grid.At(nextPos)

	// recursively move the connected cells
	switch nextObj {
	case 'O': // If box, recursively move the box forward
		MoveCellsAlongDirection(grid, nextPos, dir)

	// Part 2 Cases
	case '[': // If left half of box, recursively move both the halves. (Make sure we do not cause infinite loop)
		if dir != LEFT {
			MoveCellsAlongDirection(grid, nextPos.MoveAlong(RIGHT), dir)
		}
		MoveCellsAlongDirection(grid, nextPos, dir)
	case ']': // If right half of box, recursively move both the halves. (Make sure we do not cause infinite loop)
		if dir != RIGHT {
			MoveCellsAlongDirection(grid, nextPos.MoveAlong(LEFT), dir)
		}
		MoveCellsAlongDirection(grid, nextPos, dir)
	}

	grid.Set(nextPos, grid.At(pos))
	grid.Set(pos, '.')
}

// Move the robot in specified direction if possible. Otherwise do not do anything
func MoveRobot(grid *Grid, dir Direction) {
	if CanTheCellBeMoved(*grid, grid.Robot, dir) {
		MoveCellsAlongDirection(grid, grid.Robot, dir)

		nextPos := grid.Robot.MoveAlong(dir)
		grid.Robot = nextPos
	}
}

func (pos Position) GPS() int {
	return 100*pos.Y + pos.X
}

func (grid Grid) FindTotalScore() int {
	total := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			pos := Position{x, y}
			if grid.At(pos) == 'O' || grid.At(pos) == '[' {
				total += pos.GPS()
			}
		}
	}
	return total
}

func SolveForPart1(input string) int {
	grid, moves := ReadInputPart1(input)

	ApplyMoves(&grid, moves)
	return grid.FindTotalScore()
}

func SolveForPart2(input string) int {
	grid, moves := ReadInputPart2(input)

	ApplyMoves(&grid, moves)
	return grid.FindTotalScore()
}

func main() {
	fmt.Println("Part 1 solution:", SolveForPart1(input))
	fmt.Println("Part 2 solution:", SolveForPart2(input))
}
