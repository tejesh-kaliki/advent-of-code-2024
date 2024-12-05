package main

import (
	"reflect"
	"strings"
	"testing"
)

var exampleText = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestIsUpdateInRightOrder(t *testing.T) {
	ruleSection, _, _ := strings.Cut(exampleText, "\n\n")
	exampleRules := MapLines(ruleSection, getPageRule)

	testcases := []struct {
		Name    string
		Rules   []PageRule
		Updates string
		IsRight bool
	}{
		{
			Name:    "always correct when rules are empty",
			Rules:   []PageRule{},
			Updates: "1,2,3",
			IsRight: true,
		},
		{
			Name:    "not correct when does not follow order of rule",
			Rules:   []PageRule{{"2", "1"}},
			Updates: "1,2,3",
			IsRight: false,
		},
		{
			Name:    "correct when follow order of rule",
			Rules:   []PageRule{{"2", "1"}},
			Updates: "3,2,1",
			IsRight: true,
		},
		{
			Name:    "not correct when page number is not exact",
			Rules:   []PageRule{{"2", "1"}},
			Updates: "21,1,2",
			IsRight: false,
		},
		{
			Name:    "not correct when any of the rules are not followed",
			Rules:   []PageRule{{"2", "1"}, {"3", "2"}},
			Updates: "2,1,3",
			IsRight: false,
		},
		{
			Name:    "ignore rule if number is not present",
			Rules:   []PageRule{{"2", "1"}, {"1", "3"}},
			Updates: "2,1,4,5",
			IsRight: true,
		},
		{
			Name:    "correct row in given example",
			Rules:   exampleRules,
			Updates: "75,47,61,53,29",
			IsRight: true,
		},
		{
			Name:    "correct row in given example",
			Rules:   exampleRules,
			Updates: "97,61,53,29,13",
			IsRight: true,
		},
		{
			Name:    "correct row in given example",
			Rules:   exampleRules,
			Updates: "75,29,13",
			IsRight: true,
		},
		{
			Name:    "wrong row in given example",
			Rules:   exampleRules,
			Updates: "75,97,47,61,53",
			IsRight: false,
		},
		{
			Name:    "wrong row in given example",
			Rules:   exampleRules,
			Updates: "61,13,29",
			IsRight: false,
		},
		{
			Name:    "wrong row in given example",
			Rules:   exampleRules,
			Updates: "97,13,75,29,47",
			IsRight: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			pages := strings.Split(testcase.Updates, ",")
			got := IsUpdateInRightOrder(testcase.Rules, pages)
			if got != testcase.IsRight {
				t.Errorf("Got wrong output: got %v, want %v", got, testcase.IsRight)
			}
		})
	}
}

func TestFindSumOfMedianOfCorrectRecords(t *testing.T) {

	testcases := []struct {
		Name  string
		Input string
		Want1 int
		Want2 int
	}{
		{
			Name:  "return 0 when no updates are correct",
			Input: "1|2\n\n2,1,3",
			Want1: 0,
			Want2: 2,
		},
		{
			Name:  "return middle value when only one update is correct",
			Input: "1|2\n\n1,3,2",
			Want1: 3,
			Want2: 0,
		},
		{
			Name:  "return middle value when only one update is correct",
			Input: "1|2\n\n1,2,3",
			Want1: 2,
			Want2: 0,
		},
		{
			Name:  "return middle value when only one update is correct",
			Input: "1|2\n3|5\n\n1,2,4,3,5",
			Want1: 4,
			Want2: 0,
		},
		{
			Name:  "return sum if there are multiple matching rows",
			Input: "1|2\n3|5\n\n1,2,3,4.6\n1,5,2,3,4\n1,2,4,3,5",
			Want1: 7,
			Want2: 2,
		},
		{
			Name:  "value of given example is 143",
			Input: exampleText,
			Want1: 143,
			Want2: 123,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got1, got2 := FindSumOfMedians(testcase.Input)
			if got1 != testcase.Want1 {
				t.Errorf("Got wrong output: got %d, want %d", got1, testcase.Want1)
			}
			if got2 != testcase.Want2 {
				t.Errorf("Got wrong output: got %d, want %d", got2, testcase.Want2)
			}
		})
	}
}

func TestReorderUpdates(t *testing.T) {
	ruleSection, _, _ := strings.Cut(exampleText, "\n\n")
	exampleRules := MapLines(ruleSection, getPageRule)

	testcases := []struct {
		Name    string
		Rules   []PageRule
		Updates []string
		Want    []string
	}{
		{
			Name:    "do nothing if rules are followed",
			Rules:   []PageRule{{"1", "2"}},
			Updates: []string{"1", "3", "2"},
			Want:    []string{"1", "3", "2"},
		},
		{
			Name:    "switch if order is different",
			Rules:   []PageRule{{"1", "2"}},
			Updates: []string{"2", "3", "1"},
			Want:    []string{"1", "3", "2"},
		},
		{
			Name:    "from given example",
			Rules:   exampleRules,
			Updates: strings.Split("75,97,47,61,53", ","),
			Want:    strings.Split("97,75,47,61,53", ","),
		},
		{
			Name:    "from given example",
			Rules:   exampleRules,
			Updates: strings.Split("61,13,29", ","),
			Want:    strings.Split("61,29,13", ","),
		},
		{
			Name:    "from given example",
			Rules:   exampleRules,
			Updates: strings.Split("97,13,75,29,47", ","),
			Want:    strings.Split("97,75,47,29,13", ","),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			ReorderUpdates(testcase.Updates, testcase.Rules)
			if !reflect.DeepEqual(testcase.Updates, testcase.Want) {
				t.Errorf("Updates not correctly ordered: got %v want %v", testcase.Updates, testcase.Want)
			}
		})
	}
}
