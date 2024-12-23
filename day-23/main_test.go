package main

import (
	"fmt"
	"slices"
	"testing"
)

var testInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

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

func TestFindInterconnectedComputers(t *testing.T) {
	type testinfo struct {
		Name  string
		Input string
		Want  [][3]string
	}
	testcases := []testinfo{
		{
			Name:  "Return empty list if no connected computers",
			Input: "ab-cd\nef-gh\nij-kl\nmn-op\nqr-st\nuv-wx",
			Want:  [][3]string{},
		},
		{
			Name:  "Return the computers if single connected",
			Input: "ab-cd\ncd-ef\nef-ab",
			Want:  [][3]string{{"ab", "cd", "ef"}},
		},
		{
			Name:  "Return 2 interconnected computer groups",
			Input: "ab-cd\ncd-ef\nef-ab\ngh-ef\nab-gh",
			Want:  [][3]string{{"ab", "cd", "ef"}, {"ab", "ef", "gh"}},
		},
		{
			Name:  "test input",
			Input: testInput,
			Want: [][3]string{
				{"aq", "cg", "yn"},
				{"aq", "vc", "wq"},
				{"co", "de", "ka"},
				{"co", "de", "ta"},
				{"co", "ka", "ta"},
				{"de", "ka", "ta"},
				{"kh", "qp", "ub"},
				{"qp", "td", "wh"},
				{"tb", "vc", "wq"},
				{"tc", "td", "wh"},
				{"td", "wh", "yn"},
				{"ub", "vc", "wq"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			graph := ReadInput(testcase.Input)
			got := FindInterconnectedComputersOfSize3(graph)
			if len(got) != len(testcase.Want) {
				t.Errorf("Got wrong length of interconnected comps: got %v, want %v", got, testcase.Want)
			}

			CheckIfElementsAreSame(t, got, testcase.Want)
		})
	}
}

func TestPart1Solution(t *testing.T) {
	graph := ReadInput(testInput)
	got := SolvePart1(graph)
	want := 7
	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func TestPart2Solution(t *testing.T) {
	graph := ReadInput(testInput)

	got := SolvePart2(graph)
	fmt.Println(got)
}
