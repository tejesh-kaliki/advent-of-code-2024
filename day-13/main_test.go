package main

import (
	"slices"
	"testing"
)

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

var testInput = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestIsPrizePossible(t *testing.T) {
	infos := ReadInput(testInput)
	testcases := []struct {
		Name  string
		Input MachineInfo
		Valid bool
		Score int
	}{
		{
			Name:  "not possible if prize is less than A and B",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{1, 2}},
			Valid: false,
		},
		{
			Name:  "possible if A is exactly same as prize",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{5, 10}},
			Valid: true,
			Score: 3,
		},
		{
			Name:  "possible if prize is exactly a multiple of A",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{25, 50}},
			Valid: true,
			Score: 15,
		},
		{
			Name:  "possible if prize is a multiple of B",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{36, 24}},
			Valid: true,
			Score: 6,
		},
		{
			Name:  "possible if prize is a sum of A and B",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{11, 14}},
			Valid: true,
			Score: 4,
		},
		{
			Name:  "possible if prize is a combination of A and B",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{22, 28}},
			Valid: true,
			Score: 8,
		},
		{
			Name:  "valid example from test input",
			Input: infos[0],
			Valid: true,
			Score: 280,
		},
		{
			Name:  "invalid example from test input",
			Input: infos[1],
			Valid: false,
		},
		{
			Name:  "valid example from test input",
			Input: infos[2],
			Valid: true,
			Score: 200,
		},
		{
			Name:  "invalid example from test input",
			Input: infos[3],
			Valid: false,
		},
		{
			Name:  "not valid if buttons need to be pressed more than 100 times",
			Input: MachineInfo{Vector{5, 10}, Vector{6, 4}, Vector{600, 1200}},
			Valid: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			score, valid := testcase.Input.GetMinimumScore(func(i int) bool { return i > 100 })
			if valid != testcase.Valid {
				t.Errorf("Got wrong output: got %v, want %v", valid, testcase.Valid)
			}
			if score != testcase.Score {
				t.Errorf("Got wrong output: got %d, want %d", score, testcase.Score)
			}
		})
	}
}

func TestPart1Solution(t *testing.T) {
	infos := ReadInput(testInput)
	got := SolvePart1(infos)
	want := 480

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	infos := ReadInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SolvePart1(infos)
	}
}

func BenchmarkPart2(b *testing.B) {
	infos := ReadInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SolvePart2(infos)
	}
}
