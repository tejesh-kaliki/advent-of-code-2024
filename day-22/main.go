package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func GenerateNextPseudoRandomNumber(num int) int {
	secret := num

	secret = ((secret << 6) ^ secret) & ((1 << 24) - 1)

	secret = (secret >> 5) ^ secret

	secret = ((secret << 11) ^ secret) & ((1 << 24) - 1)

	return secret
}

func FindNthSecretNumber(num, n int) int {
	secret := num
	for i := 0; i < n; i++ {
		secret = GenerateNextPseudoRandomNumber(secret)
	}
	return secret
}

func FindFirstNPricesAndNthSecret(num, n int) ([]int, int) {
	prices := make([]int, 0, n+1)
	secret := num
	prices = append(prices, secret%10)

	for i := 0; i < n; i++ {
		secret = GenerateNextPseudoRandomNumber(secret)
		prices = append(prices, secret%10)
	}

	return prices, secret
}

func SolvePart1(nums []int) int {
	total := 0

	for _, num := range nums {
		total += FindNthSecretNumber(num, 2000)
	}

	return total
}

func FindThePriceWithChanges(prices []int, changes [4]int) int {
	if len(prices) < 5 {
		return -1
	}

	for i := 0; i < len(prices)-4; i++ {
		found := true
		for j := 0; j < 4; j++ {
			if prices[i+j+1]-prices[i+j] != changes[j] {
				found = false
				break
			}
		}
		if found {
			return prices[i+4]
		}
	}

	return -1
}

func GetTotalOfPricesWithChangeFn(pricesList [][]int) func(changes [4]int) int {
	cache := make(map[[4]int]int)
	return func(changes [4]int) int {
		if total, found := cache[changes]; found {
			return total
		}
		total := 0

		for _, prices := range pricesList {
			priceFound := FindThePriceWithChanges(prices, changes)
			if priceFound != -1 {
				total += priceFound
			}
		}

		cache[changes] = total
		return total
	}
}

func SolvePart2(nums []int) int {
	pricesList := make([][]int, len(nums))
	for i, num := range nums {
		pricesList[i], _ = FindFirstNPricesAndNthSecret(num, 2000)
	}

	maxPrice := 0
	priceWithChangesFn := GetTotalOfPricesWithChangeFn(pricesList)
	fmt.Println(len(nums))
	changes := [4]int{}
	for i, prices := range pricesList {
		for i := 0; i < len(prices)-4; i++ {
			for j := range changes {
				changes[j] = prices[i+j+1] - prices[i+j]
			}
			total := priceWithChangesFn(changes)

			if total > maxPrice {
				maxPrice = total
			}
		}

		fmt.Printf("checked list %d, max price now is %d\n", i, maxPrice)
	}

	return maxPrice
}

func ReadInput(input string) []int {
	numTexts := strings.Split(input, "\n")
	nums := make([]int, len(numTexts))
	for i, numText := range numTexts {
		fmt.Sscanf(numText, "%d", &nums[i])
	}
	return nums
}

func main() {
	nums := ReadInput(input)
	fmt.Println("Part 1 solution:", SolvePart1(nums))
	fmt.Println("Part 2 solution:", SolvePart2(nums))
}
