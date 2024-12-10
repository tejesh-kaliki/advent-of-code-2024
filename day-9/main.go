package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

// The disk contains either -1 or a value at specific id.
// -1 means empty, otherwise, value is the file id.
// Rearrange the disk so that all the files are at beginning.
// By placing the files from end at gaps from start
func RearrangeDiskUsingFragmentation(disk []int) []int {
	newDisk := make([]int, 0, len(disk))

	lastIndex := len(disk) - 1
	moveToNextLastIndex := func(curIndex int) bool {
		for ; lastIndex > curIndex; lastIndex-- {
			if disk[lastIndex] != -1 {
				return true
			}
		}
		return false
	}

	for i := 0; i <= lastIndex; i++ {
		if disk[i] != -1 {
			newDisk = append(newDisk, disk[i])
			continue
		}

		if moveToNextLastIndex(i) {
			newDisk = append(newDisk, disk[lastIndex])
			lastIndex--
		}
	}

	return newDisk
}

func GetDiskFromInput(input string) []int {
	disk := make([]int, 0)
	for i := range input {
		size, _ := strconv.ParseInt(input[i:i+1], 10, 8)
		data := make([]int, size)
		valueToFill := -1
		if i%2 == 0 {
			valueToFill = i / 2
		}

		for j := range data {
			data[j] = valueToFill
		}
		disk = append(disk, data...)
	}
	return disk
}

func ComputeDiskChecksumPart1(disk []int) int {
	if len(input) <= 2 {
		return 0
	}

	newDisk := RearrangeDiskUsingFragmentation(disk)

	total := 0
	for i, fileId := range newDisk {
		total += i * fileId
	}
	return total
}

// Sum the the [n] indices starting at [start]
func findSumOfIndices(start, n int) int {
	end := start + n
	totalTillEnd := end * (end - 1) / 2
	totalTillStart := start * (start - 1) / 2
	return totalTillEnd - totalTillStart
}

type FileData struct {
	ID    int
	Start int
	Size  int
}

type Gap struct {
	Start int
	Size  int
}

func (fd FileData) CheckSum() int {
	return findSumOfIndices(fd.Start, fd.Size) * fd.ID
}

// The disk contains either -1 or a value at specific id.
// -1 means empty, otherwise, value is the file id.
// Rearrange the disk using following logic:
//
//	Starting from end, check if entire file can be moved to some gap
//	For specific file, just check from leftmost gap, and move to first gap where it fits
func RearrangeDiskByCopyWholeFiles(files []FileData, gaps []Gap) []FileData {
	newFiles := make([]FileData, 0, len(files))

	findGapThatFits := func(size int) (*Gap, int) {
		for i, gap := range gaps {
			if gap.Size >= size {
				return &gap, i
			}
		}
		return nil, -1
	}

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		gap, gapIndex := findGapThatFits(file.Size)
		if gap != nil && gap.Start <= file.Start {
			file.Start = gap.Start
			gap.Start += file.Size
			gap.Size -= file.Size
			gaps[gapIndex] = *gap
		}
		newFiles = append(newFiles, file)
	}

	return newFiles
}

func ReadFilesAndGapsFromInput(input string) ([]FileData, []Gap) {
	files := make([]FileData, 0)
	gaps := make([]Gap, 0)
	startIndex := 0
	for i := range input {
		size, _ := strconv.ParseInt(input[i:i+1], 10, 8)
		if i%2 == 0 {
			files = append(files, FileData{ID: i / 2, Size: int(size), Start: startIndex})
		} else {
			gaps = append(gaps, Gap{Size: int(size), Start: startIndex})
		}
		startIndex += int(size)
	}
	return files, gaps
}

func ComputeDiskChecksumPart2(files []FileData, gaps []Gap) int {
	if len(input) <= 2 {
		return 0
	}

	newFiles := RearrangeDiskByCopyWholeFiles(files, gaps)

	total := 0
	for _, file := range newFiles {
		total += file.CheckSum()
	}
	return total
}

func main() {
	disk := GetDiskFromInput(input)
	fmt.Println("Solution to part 1:", ComputeDiskChecksumPart1(disk))

	files, gaps := ReadFilesAndGapsFromInput(input)
	fmt.Println("Solution to part 2:", ComputeDiskChecksumPart2(files, gaps))
}
