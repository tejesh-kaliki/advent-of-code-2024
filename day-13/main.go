package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Vector struct {
	X, Y int
}

type MachineInfo struct {
	A, B  Vector
	Prize Vector
}

func (info MachineInfo) GetMinimumScore(countExcludeFn func(int) bool) (int, bool) {
	ax, ay := info.A.X, info.A.Y
	bx, by := info.B.X, info.B.Y
	px, py := info.Prize.X, info.Prize.Y

	// ax*l + bx*m = px
	// ay*l + by*m = py
	// Parallel if ax/bx == ay/by

	if math.Abs(float64(ax)/float64(bx)-float64(ay)/float64(by)) < 0.01 {
		panic(fmt.Sprintf("Parallel lines %v", info))
	}

	//    ax * ay * l + bx * ay * m = px * ay
	//    ax * ay * l + by * ax * m = py * ax
	// => (bx * ay - by * ax) * m = px * ay - py * ax
	// => m = (px * ay - py * ax) / (bx * ay - by * ax)
	// => l = (py * bx - px * by) / (bx * ay - by * ax)

	den := bx*ay - by*ax
	aNum := py*bx - px*by
	bNum := px*ay - py*ax

	if aNum%den != 0 || bNum%den != 0 {
		return 0, false
	}

	aCount := aNum / den
	bCount := bNum / den

	if countExcludeFn(aCount) || countExcludeFn(bCount) {
		return 0, false
	}

	return aCount*3 + bCount, true

}

func (info MachineInfo) CalibirateForPart2() MachineInfo {
	return MachineInfo{
		A:     info.A,
		B:     info.B,
		Prize: Vector{info.Prize.X + 10000000000000, info.Prize.Y + 10000000000000},
	}
}

func SolvePart1(infos []MachineInfo) int {
	total := 0
	for _, info := range infos {
		if score, possible := info.GetMinimumScore(func(i int) bool { return i > 100 }); possible {
			total += score
		}
	}

	return total
}

func alwaysFalse[T any](x T) bool {
	return false
}

func SolvePart2(infos []MachineInfo) int {
	total := 0
	for _, info := range infos {
		if score, possible := info.CalibirateForPart2().GetMinimumScore(alwaysFalse); possible {
			total += score
		}
	}
	return total
}

func ReadInput(input string) []MachineInfo {
	parts := strings.Split(input, "\n\n")
	res := make([]MachineInfo, len(parts))
	for i, part := range parts {
		lines := strings.Split(part, "\n")
		info := MachineInfo{}
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &info.A.X, &info.A.Y)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &info.B.X, &info.B.Y)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &info.Prize.X, &info.Prize.Y)

		res[i] = info
	}

	return res
}

func main() {
	infos := ReadInput(input)
	fmt.Println("Part 1 Solution:", SolvePart1(infos))
	fmt.Println("Part 2 Solution:", SolvePart2(infos))
}
