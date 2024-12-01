package main

import "testing"

func TestTotalDistanceBetweenLocations(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{"return 0 if single same location in each", "1  1", 0},
		{"return difference if single location in each", "3  1", 2},
		{"return difference if single location in each (reverse)", "1  3", 2},
		{"return sum of difference if 2 locations in each", "1  3\n3  4", 3},
		{"return sum of difference if 2 locations in each, after sorting", "1  3\n4  2", 2},
		{"return 11 for given example input", `3   4
4   3
2   5
1   3
3   9
3   3`, 11},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			distance := TotalDistanceBetweenLocations(testcase.Input)
			if distance != testcase.Want {
				t.Errorf("Got wrong distance. Got %d, want %d", distance, testcase.Want)
			}
		})
	}
}

func TestSimilarityScoresBetweenLocations(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{"return x if single same location x in each", "1  1", 1},
		{"return x if single same location x in each", "2  2", 2},
		{"return 0 if single different location in each", "1  2", 0},
		{"return sum of locations if location lists are same, with 2 in each", "1  1\n2  2", 3},
		{"return 2 if one location from list 1 is repeated twice in list 2", "1  1\n2  1", 2},
		{"return 4 if 2 locations in each list, and all are same", "1  1\n1  1", 4},
		{"return 31 for given example input", `3   4
4   3
2   5
1   3
3   9
3   3`, 31},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			similarity := SimilarityScoresBetweenLocations(testcase.Input)
			if similarity != testcase.Want {
				t.Errorf("Got wrong similarity. Got %d, want %d", similarity, testcase.Want)
			}
		})
	}
}
