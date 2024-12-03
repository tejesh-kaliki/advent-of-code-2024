package main

import "testing"

func TestTotalMulValue(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{"value of empty string is 0", "", 0},
		{"value of mul(1,2) is 2", "mul(1,2)", 2},
		{"value of mul is 0", "mul", 0},
		{"value of mul(1 is 0", "mul(1", 0},
		{"value of mul(1,3) is 3", "mul(1,3)", 3},
		{"value of mul(2,3) is 6", "mul(2,3)", 6},
		{"value of mul(a,b) is 0", "mul(a,b)", 0},
		{"value of mul(1,) is 0", "mul(1,)", 0},
		{"value of mul(1024,98) is 0 (because larger than 3 digits)", "mul(1024,98)", 0},
		{"value of mul(-1,2) is 0 (because negative)", "mul(-1,2)", 0},
		{"value of mul(-1,2!) is 0", "mul(1,2!)", 0},
		{"value of mul(1, 2) is 0 (space)", "mul(1, 2)", 0},
		{"value of mul (1,2) is 0 (space)", "mul (1,2)", 0},
		{"When there are multiple give sum", "mul(1,2)mul(3,4)", 14},
		{"When there are multiple give sum", "mul(1,2)  mul(3,4)", 14},
		{"Value for given text is 161", "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := TotalMulValue(testcase.Input)
			if got != testcase.Want {
				t.Errorf("Got wrong value: got %d want %d", got, testcase.Want)
			}
		})
	}
}

func TestTotalMulValueWithEnable(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{"value of empty string is 0", "", 0},
		{"value of mul(1,2) is 2", "mul(1,2)", 2},
		{"value of mul is 0", "mul", 0},
		{"value of mul(1 is 0", "mul(1", 0},
		{"value of mul(1,3) is 3", "mul(1,3)", 3},
		{"value of mul(2,3) is 6", "mul(2,3)", 6},
		{"value of mul(a,b) is 0", "mul(a,b)", 0},
		{"value of mul(1,) is 0", "mul(1,)", 0},
		{"value of mul(1024,98) is 0 (because larger than 3 digits)", "mul(1024,98)", 0},
		{"value of mul(-1,2) is 0 (because negative)", "mul(-1,2)", 0},
		{"value of mul(-1,2!) is 0", "mul(1,2!)", 0},
		{"value of mul(1, 2) is 0 (space)", "mul(1, 2)", 0},
		{"value of mul (1,2) is 0 (space)", "mul (1,2)", 0},
		{"When there are multiple give sum", "mul(1,2)mul(3,4)", 14},
		{"When there are multiple give sum", "mul(1,2)  mul(3,4)", 14},
		{"input following don't() is not evaluated", "xdon't()mul(1,2)", 0},
		{"input following don't() is not evaluated", "mul(2,3)don't()mul(1,2)", 6},
		{"input following do() is evaluated", "mul(2,3)don't()mul(1,2)do()mul(3,5)", 21},
		{"Value of given example is 48", "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := TotalMulValueWithEnabling(testcase.Input)
			if got != testcase.Want {
				t.Errorf("Got wrong value: got %d want %d", got, testcase.Want)
			}
		})
	}
}
