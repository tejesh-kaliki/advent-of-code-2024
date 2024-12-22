package main

import (
	"fmt"
	"testing"
)

func TestGenerateNextPseudoRandomNumber(t *testing.T) {
	nums := ReadInput(`15887950
16495136
527345
704524
1553684
12683156
11100544
12249484
7753432
5908254`)

	type testinfo struct {
		Name  string
		Input int
		Want  int
	}
	testcases := []testinfo{
		{
			Name:  "next number of 123 is 15887950",
			Input: 123,
			Want:  15887950,
		},
	}

	for i := 0; i < len(nums)-1; i++ {
		testcases = append(testcases, testinfo{
			Name:  fmt.Sprintf("next number of %d is %d", nums[i], nums[i+1]),
			Input: nums[i],
			Want:  nums[i+1],
		})
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := GenerateNextPseudoRandomNumber(testcase.Input)
			if got != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestFindNthSecretNumber(t *testing.T) {
	type testinfo struct {
		Name  string
		Input int
		N     int
		Want  int
	}
	testcases := []testinfo{
		{
			Name:  "2000th number of 1 is 8685429",
			Input: 1,
			N:     2000,
			Want:  8685429,
		},
		{
			Name:  "2000th number of 10 is 4700978",
			Input: 10,
			N:     2000,
			Want:  4700978,
		},
		{
			Name:  "2000th number of 100 is 15273692",
			Input: 100,
			N:     2000,
			Want:  15273692,
		},
		{
			Name:  "2000th number of 2024 is 8667524",
			Input: 2024,
			N:     2000,
			Want:  8667524,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := FindNthSecretNumber(testcase.Input, testcase.N)
			if got != testcase.Want {
				t.Errorf("Got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestPart1Solution(t *testing.T) {
	input := "1\n10\n100\n2024"
	nums := ReadInput(input)
	got := SolvePart1(nums)
	want := 37327623

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func TestThePriceWithChanges(t *testing.T) {
	type testinfo struct {
		Name    string
		Prices  []int
		Changes [4]int
		Want    int
	}
	testcases := []testinfo{
		{"available change in test example", []int{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}, [4]int{-1, -1, 0, 2}, 6},
		{"not available change in test example", []int{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}, [4]int{-1, -1, 0, 3}, -1},
	}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := FindThePriceWithChanges(testcase.Prices, testcase.Changes)
			if got != testcase.Want {
				t.Errorf("got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestTotalOfPricesWithChanges(t *testing.T) {
	nums := []int{1, 2, 3, 2024}
	pricesList := make([][]int, 4)

	for i, num := range nums {
		pricesList[i], _ = FindFirstNPricesAndNthSecret(num, 2000)
	}

	type testinfo struct {
		Name    string
		Changes [4]int
		Want    int
	}
	testcases := []testinfo{
		{"from the test input", [4]int{-2, 1, -1, 3}, 23},
		{"from the test input", [4]int{-2, 1, -1, 3}, 23},
	}
	totalPriceFn := GetTotalOfPricesWithChangeFn(pricesList)
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := totalPriceFn(testcase.Changes)
			if got != testcase.Want {
				t.Errorf("got wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestPart2Solution(t *testing.T) {
	input := "1\n2\n3\n2024"
	nums := ReadInput(input)
	got := SolvePart2(nums)
	want := 23

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}
