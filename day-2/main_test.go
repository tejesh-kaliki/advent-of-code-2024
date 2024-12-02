package main

import "testing"

func TestSafeReportCount(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{"Single number is safe", "1", 1},
		{"2 same numbers is not safe", "1 1", 0},
		{"2 same numbers with slight increase is safe", "1 2", 1},
		{"2 same numbers with slight increase is 3", "1 4", 1},
		{"2 same numbers with >3 increase is unsafe", "1 5", 0},
		{"2 same numbers with >3 decrease is unsafe", "5 1", 0},
		{"3 numbers are unsafe if anyone has large increase", "1 2 6", 0},
		{"3 numbers are unsafe if it increase and decrease", "1 3 2", 0},
		{"3 numbers are unsafe if it decrease and increase", "5 2 4", 0},
		{"For 2 line input, give total safe lines", "1 2\n4 2", 2},
		{"For 2 line input, give total safe lines", "1 2\n4 0", 1},
		{"Output 2 for given example", `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, 2},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			safeCount := SafeReportCount(testcase.Input, isLineSafe)
			if safeCount != testcase.Want {
				t.Errorf("Got wront output: got %d, want %d", safeCount, testcase.Want)
			}
		})
	}
}

func BenchmarkSafeReportCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeReportCount(input, isLineSafe)
	}
}
func BenchmarkSafeReportCountPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeReportCount(input, isLineSafeWithRemove)
	}
}

func TestIsLineSafeWithRemove(t *testing.T) {
	testcases := []struct {
		Name   string
		Input  string
		IsSafe bool
	}{
		{"single number is safe", "1", true},
		{"2 same numbers is not safe, because remove", "1 1", true},
		{"2 same numbers with slight increase is safe", "1 2", true},
		{"2 same numbers with increase 3 is safe", "1 4", true},
		{"2 same numbers with >3 increase is safe, because remove", "1 5", true},
		{"3 numbers are safe if only one has large increase", "1 2 6", true},
		{"3 numbers are unsafe if there are large increases", "1 5 9", false},
		{"4 numbers are safe if it increase and decrease once", "1 3 2 4", true},
		{"4 numbers are unsafe if there are 2 large increases", "1 6 8 14", false},
		{"4 numbers are safe if first num can be removed", "3 1 4 5", true},
		{"line from given example is safe", "7 6 4 2 1", true},
		{"line from given example is unsafe", "1 2 7 8 9", false},
		{"line from given example is unsafe", "9 7 6 2 1", false},
		{"line from given example is safe", "1 3 2 4 5", true},
		{"line from given example is safe", "8 6 4 4 1", true},
		{"line from given example is safe", "1 3 6 7 9", true},
	}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			isSafe := isLineSafeWithRemove(testcase.Input)
			if isSafe != testcase.IsSafe {
				t.Errorf("Got wrong output: got %v, want %v", isSafe, testcase.IsSafe)
			}
		})
	}

}
