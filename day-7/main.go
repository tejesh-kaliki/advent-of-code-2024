package main

import (
	_ "embed"
	"fmt"
	"log"
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
	totalString := fmt.Sprint(total)
	numString := fmt.Sprint(num)
	prevTotalString, found := strings.CutSuffix(totalString, numString)
	if !found || prevTotalString == "" {
		return false, 0
	}

	prevTotal, err := strconv.ParseInt(prevTotalString, 10, 64)
	if err != nil {
		log.Println("Error processing previous total:", err, total, num)
		return false, 0
	}
	return true, prevTotal
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

func FindTotalOfValidEquations(input string, ops []Operation) int64 {
	lines := strings.Split(input, "\n")
	result := int64(0)
	for _, line := range lines {
		eq := ParseEquationText(line)
		if valid, _ := IsTheTotalPossible(eq.Total, eq.Numbers, ops); valid {
			// fmt.Printf("%d = %s\n", eq.Total, repr)
			result += eq.Total
		}
	}
	return result
}

func main() {
	fmt.Println("Solution to 1st part:", FindTotalOfValidEquations(input, []Operation{AddOp{}, MulOp{}}))
	fmt.Println("Solution to 2nd part:", FindTotalOfValidEquations(input, []Operation{AddOp{}, MulOp{}, ConcatOp{}}))
}
