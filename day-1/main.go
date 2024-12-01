package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseInputLine(line string) (int, int) {
	locs := strings.Split(line, " ")
	loc1, err := strconv.ParseInt(locs[0], 10, 64)
	if err != nil {
		log.Fatalln("Not able to process input:", err)
	}
	loc2, err := strconv.ParseInt(locs[len(locs)-1], 10, 64)
	if err != nil {
		log.Fatalln("Not able to process input:", err)
	}
	return int(loc1), int(loc2)
}

func getLocationLists(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	firstList := make([]int, len(lines))
	secondList := make([]int, len(lines))
	for i, line := range lines {
		firstList[i], secondList[i] = parseInputLine(line)
	}
	return firstList, secondList
}

func findDistance(loc1, loc2 int) int {
	if loc1 > loc2 {
		return loc1 - loc2
	}
	return loc2 - loc1
}

func TotalDistanceBetweenLocations(input string) int {
	firstList, secondList := getLocationLists(input)
	slices.Sort(firstList)
	slices.Sort(secondList)

	totalDistance := 0
	for i := range firstList {
		totalDistance += findDistance(firstList[i], secondList[i])
	}
	return totalDistance
}

func countElements[T comparable](items []T, element T) int {
	count := 0
	for _, x := range items {
		if x == element {
			count++
		}
	}
	return count
}

func SimilarityScoresBetweenLocations(input string) int {
	firstList, secondList := getLocationLists(input)
	total := 0
	for _, element := range firstList {
		count := countElements(secondList, element)
		total += count * element
	}
	return total
}

func main() {
	fmt.Println("Part 1 solution:", TotalDistanceBetweenLocations(input))
	fmt.Println("Part 2 solution:", SimilarityScoresBetweenLocations(input))
}
