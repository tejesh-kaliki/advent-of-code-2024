package main

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
)

//go:embed input.txt
var input string

type Vector struct {
	X, Y int
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

func (vec Vector) Scale(scale int) Vector {
	return Vector{vec.X * scale, vec.Y * scale}
}

type Robot struct {
	InitialPos Vector
	Velocity   Vector
}

func (r Robot) PositionAfter(time int, space Space) Vector {
	pos := r.InitialPos.Add(r.Velocity.Scale(time))
	return space.WrapPosition(pos)
}

type Space struct {
	Width, Height int
}

func (s Space) WrapPosition(vec Vector) Vector {
	x := vec.X % s.Width
	if x < 0 {
		x = x + s.Width
	}

	y := vec.Y % s.Height
	if y < 0 {
		y = y + s.Height
	}
	return Vector{x, y}
}

// Returns the quadrant of a position.
// The quadrants are 1, 2, 3, 4 in order of top-right, top-left, bottom-left and bottom-right
// Or return -1 if it is exactly on axis.
func (s Space) FindQuadrant(pos Vector) int {
	if s.Width%2 == 0 || s.Height%2 == 0 {
		panic("even number of rows/columns not handled")
	}

	midX := (s.Width - 1) / 2
	midY := (s.Height - 1) / 2

	switch {
	case pos.X > midX && pos.Y < midY:
		return 1
	case pos.X < midX && pos.Y < midY:
		return 2
	case pos.X < midX && pos.Y > midY:
		return 3
	case pos.X > midX && pos.Y > midY:
		return 4
	default:
		return -1
	}
}

func SolvePart1(robots []Robot, space Space) int {
	quadrantCounts := make([]int, 4)
	for _, robot := range robots {
		pos := robot.PositionAfter(100, space)
		quadrant := space.FindQuadrant(pos)
		if quadrant == -1 {
			continue
		}
		quadrantCounts[quadrant-1]++
	}

	product := 1
	for _, count := range quadrantCounts {
		product *= count
	}
	return product
}

func GenerateSpaceImage(robots []Robot, space Space, time int) {
	canvas := image.NewRGBA(image.Rect(0, 0, space.Width, space.Height))
	for i := 0; i < space.Width; i++ {
		for j := 0; j < space.Height; j++ {
			canvas.Set(i, j, color.Black)
		}
	}
	for _, robot := range robots {
		pos := robot.PositionAfter(time, space)
		canvas.SetRGBA(pos.X, pos.Y, color.RGBA{0x00, 0xFF, 0x00, 0xFF})
	}
	err := draw2dimg.SaveToPngFile(fmt.Sprintf("day-14/generated/frame-%05d.png", time), canvas)
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadInput(input string) []Robot {
	lines := strings.Split(input, "\n")
	res := make([]Robot, len(lines))
	for i, line := range lines {
		fmt.Sscanf(
			line,
			"p=%d,%d v=%d,%d",
			&res[i].InitialPos.X,
			&res[i].InitialPos.Y,
			&res[i].Velocity.X,
			&res[i].Velocity.Y,
		)
	}
	return res
}

func main() {
	robots := ReadInput(input)
	fmt.Println("Part 1 solution:", SolvePart1(robots, Space{101, 103}))
	space := Space{101, 103}

	for i := 0; i < 10000; i++ {
		GenerateSpaceImage(robots, space, i)
	}
}
