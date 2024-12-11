package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApplyBlinkRule(t *testing.T) {
	testcases := []struct {
		Name  string
		Input int
		Want  []int64
	}{
		{
			Name:  "if input is 0, return 1",
			Input: 0,
			Want:  []int64{1},
		},
		{
			Name:  "if there are odd number of digits, multiply by 2024",
			Input: 1,
			Want:  []int64{2024},
		},
		{
			Name:  "if there are odd number of digits, multiply by 2024",
			Input: 1036288,
			Want:  []int64{1036288 * 2024},
		},
		{
			Name:  "if there are even number of digits, split into 2 parts",
			Input: 2024,
			Want:  []int64{20, 24},
		},
		{
			Name:  "if there are even number of digits, split into 2 parts",
			Input: 28676032,
			Want:  []int64{2867, 6032},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := ApplyBlinkRule(int64(testcase.Input))
			if !reflect.DeepEqual(got, testcase.Want) {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.Want)
			}
		})
	}
}

func TestGetCountAfterBlinks(t *testing.T) {
	testcases := []struct {
		Name   string
		Value  int
		Blinks int
		Want   int
	}{
		{
			Name:   "Return 1 if no blinks",
			Blinks: 0,
			Want:   1,
		},
		{
			Name:   "return 2 if 1 blink and even digits",
			Value:  2024,
			Blinks: 1,
			Want:   2,
		},
		{
			Name:   "return 4 if 2 blinks and multiple of 4 number of digits",
			Value:  2024,
			Blinks: 2,
			Want:   4,
		},
		{
			Name:   "return 5 for 125 after 5 blinks",
			Value:  125,
			Blinks: 5,
			Want:   5,
		},
		{
			Name:   "return 55312 for 125017 after 26 blinks",
			Value:  125017,
			Blinks: 26,
			Want:   55312,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := GetCountAfterBlinks(int64(testcase.Value), testcase.Blinks)
			if got != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestTotalCount(t *testing.T) {
	input := []int64{125, 17}
	want := 55312
	got := GetTotalElementsAfterBlinks(input, 25)
	fmt.Println(got, want)
}

func BenchmarkGetTotalElementsAfterBlinks(b *testing.B) {
	values := ReadInput(input)
	for i := 0; i < b.N; i++ {
		GetTotalElementsAfterBlinks(values, 25)
	}
}
