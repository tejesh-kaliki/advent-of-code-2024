package main

import (
	"testing"
)

func TestIsTotalPossible(t *testing.T) {
	testcases := []struct {
		Name    string
		Total   int64
		Numbers []int64
		IsValid bool
	}{
		{
			Name:    "true if single number, which is same as total",
			Total:   10,
			Numbers: []int64{10},
			IsValid: true,
		},
		{
			Name:    "false if single number, which is not same as total",
			Total:   10,
			Numbers: []int64{5},
			IsValid: false,
		},
		{
			Name:    "false if 2 numbers, cannot be combined to get result",
			Total:   10,
			Numbers: []int64{10, 2},
			IsValid: false,
		},
		{
			Name:    "true if 2 numbers can be added to get result",
			Total:   10,
			Numbers: []int64{8, 2},
			IsValid: true,
		},
		{
			Name:    "true if product of numbers is equal to total",
			Total:   10,
			Numbers: []int64{5, 2},
			IsValid: true,
		},
		{
			Name:    "Valid in given example",
			Total:   3267,
			Numbers: []int64{81, 40, 27},
			IsValid: true,
		},
		{
			Name:    "Valid in given example",
			Total:   292,
			Numbers: []int64{11, 6, 16, 20},
			IsValid: true,
		},
		{
			Name:    "Valid in given example",
			Total:   190,
			Numbers: []int64{19, 10},
			IsValid: true,
		},
		{
			Name:    "Invalid in given example",
			Total:   83,
			Numbers: []int64{17, 5},
			IsValid: false,
		},
		{
			Name:    "Invalid in given example",
			Total:   156,
			Numbers: []int64{15, 6},
			IsValid: false,
		},
		{
			Name:    "Invalid in given example",
			Total:   7290,
			Numbers: []int64{6, 8, 6, 15},
			IsValid: false,
		},
		{
			Name:    "Invalid in given example",
			Total:   161011,
			Numbers: []int64{16, 10, 13},
			IsValid: false,
		},
		{
			Name:    "Invalid in given example",
			Total:   192,
			Numbers: []int64{17, 8, 14},
			IsValid: false,
		},
		{
			Name:    "Invalid in given example",
			Total:   21037,
			Numbers: []int64{9, 7, 18, 13},
			IsValid: false,
		},
	}

	ops := []Operation{AddOp{}, MulOp{}}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got, _ := IsTheTotalPossible(testcase.Total, testcase.Numbers, ops)
			if got != testcase.IsValid {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.IsValid)
			}
		})
	}
}

func TestFindTotalOfValidEquations(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	want := int64(3749)
	eqs := ParseEquations(input)
	got := FindTotalOfValidEquations(eqs, []Operation{AddOp{}, MulOp{}})

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func TestFindTotalOfValidEquationsPart2(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	want := int64(11387)
	eqs := ParseEquations(input)
	got := FindTotalOfValidEquations(eqs, []Operation{AddOp{}, MulOp{}, ConcatOp{}})

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func TestConcatOperationReverse(t *testing.T) {
	testcases := []struct {
		Name      string
		Total     int64
		LastNum   int64
		IsValid   bool
		PrevTotal int64
	}{
		{
			Name:      "false if total does not end with number",
			Total:     125,
			LastNum:   10,
			IsValid:   false,
			PrevTotal: 0,
		},
		{
			Name:      "valid if the total ends with last num",
			Total:     123,
			LastNum:   3,
			IsValid:   true,
			PrevTotal: 12,
		},
	}

	op := ConcatOp{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			gotValid, gotPrevTotal := op.Reverse(testcase.Total, testcase.LastNum)
			if gotValid != testcase.IsValid {
				t.Errorf("Got wrong output: got %v, want %v", gotValid, testcase.IsValid)
			}
			if gotPrevTotal != testcase.PrevTotal {
				t.Errorf("Got wrong output: got %d, want %d", gotPrevTotal, testcase.PrevTotal)
			}
		})
	}
}

func BenchmarkFindTotalOfValidEquationsWithConcat(b *testing.B) {
	eqs := ParseEquations(input)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FindTotalOfValidEquations(eqs, []Operation{AddOp{}, MulOp{}, ConcatOp{}})
	}
}
