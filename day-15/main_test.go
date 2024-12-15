package main

import (
	"reflect"
	"slices"
	"testing"
)

var smallTestInput = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

var testInput = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

func TestRobotMovesInDirectionSpecified(t *testing.T) {
	grid, _ := ReadInputPart1(testInput)
	testcases := []struct {
		Name         string
		Grid         Grid
		RobotPos     Position
		Dir          Direction
		WantRobotPos Position
	}{
		{
			Name:         "move robot in the specified direction",
			Grid:         grid,
			RobotPos:     Position{4, 2},
			Dir:          UP,
			WantRobotPos: Position{4, 1},
		},
		{
			Name:         "move robot in the specified direction",
			Grid:         grid,
			RobotPos:     Position{4, 2},
			Dir:          DOWN,
			WantRobotPos: Position{4, 3},
		},
		{
			Name:         "move robot in the specified direction",
			Grid:         grid,
			RobotPos:     Position{4, 2},
			Dir:          RIGHT,
			WantRobotPos: Position{5, 2},
		},
		{
			Name:         "move robot in the specified direction",
			Grid:         grid,
			RobotPos:     Position{4, 2},
			Dir:          LEFT,
			WantRobotPos: Position{3, 2},
		},
		{
			Name:         "do not move robot if there is wall in new position",
			Grid:         grid,
			RobotPos:     Position{1, 1},
			Dir:          UP,
			WantRobotPos: Position{1, 1},
		},
		{
			Name:         "do not move robot if there is wall in new position",
			Grid:         grid,
			RobotPos:     Position{1, 1},
			Dir:          LEFT,
			WantRobotPos: Position{1, 1},
		},
		{
			Name:         "do not move robot if there is box and then wall",
			Grid:         grid,
			RobotPos:     Position{7, 1},
			Dir:          RIGHT,
			WantRobotPos: Position{7, 1},
		},
		{
			Name:         "do not move robot if there is box and then wall",
			Grid:         grid,
			RobotPos:     Position{3, 2},
			Dir:          UP,
			WantRobotPos: Position{3, 2},
		},
		{
			Name:         "do not move robot if there are multiple boxes and then wall",
			Grid:         grid,
			RobotPos:     Position{5, 6},
			Dir:          DOWN,
			WantRobotPos: Position{5, 6},
		},
		{
			Name:         "do not move robot if there are multiple boxes and then wall",
			Grid:         grid,
			RobotPos:     Position{6, 7},
			Dir:          RIGHT,
			WantRobotPos: Position{6, 7},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			grid = grid.Copy()
			grid.Set(grid.Robot, '.')

			grid.Robot = testcase.RobotPos
			grid.Set(grid.Robot, '@')

			MoveRobot(&grid, testcase.Dir)
			if grid.Robot != testcase.WantRobotPos {
				t.Errorf("Got wrong position of robot: got %v, want %v", grid.Robot, testcase.WantRobotPos)
			}

			if testcase.RobotPos != testcase.WantRobotPos {
				if grid.At(testcase.RobotPos) == '@' {
					t.Errorf("Robot not moved from old position")
				}
			}
			if grid.At(testcase.WantRobotPos) != '@' {
				t.Errorf("Robot not moved to new position: %c, %v", grid.At(testcase.WantRobotPos), testcase.WantRobotPos)
			}
		})
	}
}

func CheckIfElementsAreSame[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Given inputs have different lengths: got %v, want %v", got, want)
	}
	for _, x := range got {
		if !slices.Contains(want, x) {
			t.Errorf("Unnecessary value present: %v", x)
		}
	}
	for _, x := range want {
		if !slices.Contains(got, x) {
			t.Errorf("Necessary value not present: %v", x)
		}
	}
}

func TestRobotMovesTheBoxesInMiddle(t *testing.T) {
	grid, _ := ReadInputPart1(testInput)
	testcases := []struct {
		Name         string
		Grid         Grid
		RobotPos     Position
		Dir          Direction
		WantRobotPos Position
		OldBoxPos    []Position
		NewBoxPos    []Position
	}{
		{
			Name:         "move the single box in the direction",
			Grid:         grid,
			RobotPos:     Position{4, 4},
			Dir:          LEFT,
			WantRobotPos: Position{3, 4},
			OldBoxPos:    []Position{{3, 4}},
			NewBoxPos:    []Position{{2, 4}},
		},
		{
			Name:         "move the multiple boxes in the direction",
			Grid:         grid,
			RobotPos:     Position{4, 3},
			Dir:          LEFT,
			WantRobotPos: Position{3, 3},
			OldBoxPos:    []Position{{3, 3}, {2, 3}},
			NewBoxPos:    []Position{{2, 3}, {1, 3}},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			grid := testcase.Grid.Copy()
			grid.Set(grid.Robot, '.')

			grid.Robot = testcase.RobotPos
			grid.Set(grid.Robot, '@')

			MoveRobot(&grid, testcase.Dir)
			if grid.Robot != testcase.WantRobotPos {
				t.Errorf("Got wrong position of robot: got %v, want %v", grid.Robot, testcase.WantRobotPos)
			}

			for _, box := range testcase.OldBoxPos {
				if !slices.Contains(testcase.NewBoxPos, box) && grid.At(box) == 'O' {
					t.Errorf("Box should not be present at this position: %v. Has '%c'", box, grid.At(box))
				}
			}

			for _, box := range testcase.NewBoxPos {
				if grid.At(box) != 'O' {
					t.Errorf("Box should be present at this position: %v. Has '%c'", box, grid.At(box))
				}
			}

		})
	}
}

func TestGridAfterApplyingMoves(t *testing.T) {
	afterGrid := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`
	grid, moves := ReadInputPart1(testInput)

	ApplyMoves(&grid, moves)

	finalGrid := ReadGridText(afterGrid)
	if !reflect.DeepEqual(grid, finalGrid) {
		t.Errorf("Expected grid is not achieved")
	}
}

func TestTotalScoreOfTheGrid(t *testing.T) {
	afterGrid := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`

	finalGrid := ReadGridText(afterGrid)
	got := finalGrid.FindTotalScore()
	want := 10092
	if got != want {
		t.Errorf("Got wrong score: got %d, want %d", got, want)
	}
}

func TestPart1Solution(t *testing.T) {
	got := SolveForPart1(testInput)
	want := 10092
	if got != want {
		t.Errorf("Got wrong score: got %d, want %d", got, want)
	}
}

func TestRobotMovesInDirectionSpecifiedPart2(t *testing.T) {
	grid, _ := ReadInputPart2(testInput)
	testcases := []struct {
		Name         string
		Grid         Grid
		RobotPos     Position
		Dir          Direction
		WantRobotPos Position
	}{
		{
			Name:         "move robot in specified direction if it is empty",
			Grid:         grid,
			RobotPos:     Position{3, 2},
			Dir:          UP,
			WantRobotPos: Position{3, 1},
		},
		{
			Name:         "move robot in specified direction if it is empty",
			Grid:         grid,
			RobotPos:     Position{3, 2},
			Dir:          RIGHT,
			WantRobotPos: Position{4, 2},
		},
		{
			Name:         "do not move the robot if next is a wall",
			Grid:         grid,
			RobotPos:     Position{3, 1},
			Dir:          UP,
			WantRobotPos: Position{3, 1},
		},
		{
			Name:         "move the robot if next is a box and can move box in direction",
			Grid:         grid,
			RobotPos:     Position{8, 4},
			Dir:          LEFT,
			WantRobotPos: Position{7, 4},
		},
		{
			Name:         "do not move if next is box, followed by wall",
			Grid:         grid,
			RobotPos:     Position{4, 6},
			Dir:          LEFT,
			WantRobotPos: Position{4, 6},
		},
		{
			Name: "do not move if one of the boxes touches a wall",
			Grid: ReadGridText(`##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`),
			RobotPos:     Position{7, 4},
			Dir:          UP,
			WantRobotPos: Position{7, 4},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			grid = testcase.Grid.Copy()
			grid.Set(grid.Robot, '.')

			grid.Robot = testcase.RobotPos
			grid.Set(grid.Robot, '@')

			MoveRobot(&grid, testcase.Dir)
			if grid.Robot != testcase.WantRobotPos {
				t.Errorf("Got wrong position of robot: got %v, want %v", grid.Robot, testcase.WantRobotPos)
			}

			if testcase.RobotPos != testcase.WantRobotPos {
				if grid.At(testcase.RobotPos) == '@' {
					t.Errorf("Robot not moved from old position")
				}
			}
			if grid.At(testcase.WantRobotPos) != '@' {
				t.Errorf("Robot not moved to new position: %c, %v", grid.At(testcase.WantRobotPos), testcase.WantRobotPos)
			}
		})
	}
}

func TestWhetherBoxesMoveCorrectlyForPart2(t *testing.T) {
	testcases := []struct {
		Name    string
		Grid    string
		Dir     Direction
		NewGrid string
	}{
		{
			Name: "move a single attached box",
			Grid: `##############
##......##..##
##...[][]...##
##...@[]....##
##..........##
##..........##
##############`,
			Dir: UP,
			NewGrid: `##############
##...[].##..##
##...@.[]...##
##....[]....##
##..........##
##..........##
##############`,
		},
		{
			Name: "move a single attached box",
			Grid: `##############
##.......#..##
##...[][]...##
##....[]@...##
##..........##
##..........##
##############`,
			Dir: UP,
			NewGrid: `##############
##.....[]#..##
##...[].@...##
##....[]....##
##..........##
##..........##
##############`,
		},
		{
			Name: "move a single attached box horizontally",
			Grid: `##############
##.......#..##
##...[][]...##
##....[]@...##
##..........##
##..........##
##############`,
			Dir: LEFT,
			NewGrid: `##############
##.......#..##
##...[][]...##
##...[]@....##
##..........##
##..........##
##############`,
		},
		{
			Name: "move multiple attached boxes",
			Grid: `##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.....@....##
##############`,
			Dir: UP,
			NewGrid: `##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############`,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {

			grid := ReadGridText(testcase.Grid)
			newGrid := ReadGridText(testcase.NewGrid)

			MoveCellsAlongDirection(&grid, grid.Robot, testcase.Dir)

			if grid.State() != newGrid.State() {
				t.Errorf("Final grid is not as expected: \nGot: \n%s\n\nWant: \n%s", grid.State(), newGrid.State())
			}
		})
	}
}

func TestGridAfterApplyingMovesPart2(t *testing.T) {
	afterGrid := `####################
##[].......[].[][]##
##[]...........[].##
##[]........[][][]##
##[]......[]....[]##
##..##......[]....##
##..[]............##
##..@......[].[][]##
##......[][]..[]..##
####################
`
	grid, moves := ReadInputPart2(testInput)

	ApplyMoves(&grid, moves)

	if grid.State() != afterGrid {
		t.Errorf("Final grid is not as expected: \nGot: \n%s\n\nWant: \n%s", grid.State(), afterGrid)
	}
}

func TestPart2Solution(t *testing.T) {
	got := SolveForPart2(testInput)
	want := 9021
	if got != want {
		t.Errorf("Got wrong score: got %d, want %d", got, want)
	}
}
