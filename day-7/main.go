package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Operation interface {
	Reverse(total, num int64) (bool, int64)
	Repr(repr string, lastNum int64) string
}

type AddOp struct{}

func (op AddOp) Reverse(total, num int64) (bool, int64) {
	return true, total - num
}

func (op AddOp) Repr(repr string, lastNum int64) string {
	return fmt.Sprintf("%s + %d", repr, lastNum)
}

type MulOp struct{}

func (op MulOp) Reverse(total, num int64) (bool, int64) {
	if total%num != 0 {
		return false, 0
	}
	return true, total / num
}

func (op MulOp) Repr(repr string, lastNum int64) string {
	return fmt.Sprintf("%s * %d", repr, lastNum)
}

type ConcatOp struct{}

func (op ConcatOp) Reverse(total, num int64) (bool, int64) {
	// Find number of digits in num
	numDigits := int(math.Floor(math.Log10(float64(num))) + 1)
	power := int64(math.Pow10(numDigits))

	// If the total does not end with num, can't undo concatenation
	if power == 0 || total%power != num {
		return false, 0
	}

	return true, total / power
}

func (op ConcatOp) Repr(repr string, lastNum int64) string {
	return fmt.Sprintf("%s || %d", repr, lastNum)
}

func IsTheTotalPossible(total int64, nums []int64, ops []Operation) (bool, string) {
	// Base case
	if len(nums) == 1 {
		if total == nums[0] {
			return true, fmt.Sprintf("%d", nums[0])
		}
		return false, ""
	}

	for _, op := range ops {
		canReverse, prevTotal := op.Reverse(total, nums[len(nums)-1])
		if !canReverse {
			continue
		}

		possible, repr := IsTheTotalPossible(prevTotal, nums[:len(nums)-1], ops)
		if !possible {
			continue
		}

		return possible, op.Repr(repr, nums[len(nums)-1])
	}

	return false, ""
}

type Equation struct {
	Total   int64
	Numbers []int64
}

func ParseEquationText(line string) Equation {
	totalText, numsText, _ := strings.Cut(line, ": ")
	total, err := strconv.ParseInt(totalText, 10, 64)
	if err != nil {
		log.Fatalln("Not able to parse total:", err)
	}

	numsTextSplit := strings.Split(numsText, " ")
	nums := make([]int64, len(numsTextSplit))
	for i, numText := range numsTextSplit {
		num, err := strconv.ParseInt(numText, 10, 64)
		if err != nil {
			log.Fatalln("Not able to parse total:", err)
		}
		nums[i] = num
	}

	return Equation{total, nums}

}

func ParseEquations(input string) []Equation {

	lines := strings.Split(input, "\n")
	eqs := make([]Equation, len(lines))
	for i, line := range lines {
		eqs[i] = ParseEquationText(line)
	}
	return eqs
}

func FindTotalOfValidEquations(eqs []Equation, ops []Operation) int64 {
	result := int64(0)
	for _, eq := range eqs {
		if valid, _ := IsTheTotalPossible(eq.Total, eq.Numbers, ops); valid {
			// fmt.Printf("%d = %s\n", eq.Total, repr)
			result += eq.Total
		}
	}
	return result
}

func main() {
	eqs := ParseEquations(input)
	fmt.Println("Solution to 1st part:", FindTotalOfValidEquations(eqs, []Operation{AddOp{}, MulOp{}}))
	fmt.Println("Solution to 2nd part:", FindTotalOfValidEquations(eqs, []Operation{AddOp{}, MulOp{}, ConcatOp{}}))
}
