package main

import (
	"reflect"
	"testing"
)

func TestComputeDiskChecksum(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{
		{
			Name:  "return 0 if single entry, always",
			Input: "1",
			Want:  0,
		},
		{
			Name:  "return 0 if single entry, always",
			Input: "5",
			Want:  0,
		},
		{
			Name:  "for second file, index of it should be multiplied by 1",
			Input: "101",
			Want:  1,
		},
		{
			Name:  "for second file, indexes of it should be multiplied by 1 and summed",
			Input: "105",
			Want:  15,
		},
		{
			Name:  "for second file indices should start after first file",
			Input: "005",
			Want:  10,
		},
		{
			Name:  "for second file indices should start after first file",
			Input: "205",
			Want:  20,
		},
		{
			Name:  "for third file, multiply indices by 2",
			Input: "10101",
			Want:  5,
		},
		{
			Name:  "test input from aoc",
			Input: "2333133121414131402",
			Want:  1928,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := ComputeDiskChecksumPart1(GetDiskFromInput(testcase.Input))
			if got != testcase.Want {
				t.Errorf("Wrong output: got %d, want %d", got, testcase.Want)
			}
		})
	}
}

func TestRearrangeDiskUsingFragmentation(t *testing.T) {
	testcases := []struct {
		Name     string
		Disk     []int
		WantDisk []int
	}{
		{
			Name:     "Return the same disk if only single file",
			Disk:     []int{1},
			WantDisk: []int{1},
		},
		{
			Name:     "If gap in middle, fill it with file id from end",
			Disk:     []int{1, -1, 2},
			WantDisk: []int{1, 2},
		},
		{
			Name:     "If gap in middle, fill it with file id from end",
			Disk:     []int{1, -1, 2, 2, -1, 3, 3},
			WantDisk: []int{1, 3, 2, 2, 3},
		},
		{
			Name:     "If gap in middle, ignore -1 when filling from back",
			Disk:     []int{1, -1, -1, 2, 2, -1, 3},
			WantDisk: []int{1, 3, 2, 2},
		},
		{
			Name:     "example from aoc",
			Disk:     []int{0, -1, -1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
			WantDisk: []int{0, 2, 2, 1, 1, 2, 2, 2},
		},
		{
			Name:     "example from aoc",
			Disk:     GetDiskFromInput("2333133121414131402"),
			WantDisk: []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := RearrangeDiskUsingFragmentation(testcase.Disk)
			if !reflect.DeepEqual(got, testcase.WantDisk) {
				t.Errorf("Wrong output: got %v, want %v", got, testcase.WantDisk)
			}
		})
	}
}

func TestRearrangeDiskAsWholeFiles(t *testing.T) {
	input := "2333133121414131402"
	files, gaps := ReadFilesAndGapsFromInput(input)
	testcases := []struct {
		Name      string
		Files     []FileData
		Gaps      []Gap
		WantFiles []FileData
	}{
		{
			Name:      "Return the same disk if only single file",
			Files:     []FileData{{0, 0, 1}},
			Gaps:      []Gap{},
			WantFiles: []FileData{{0, 0, 1}},
		},
		{
			Name:      "If file can be rearranged, just move the file",
			Files:     []FileData{{0, 0, 1}, {1, 3, 2}},
			Gaps:      []Gap{{1, 2}},
			WantFiles: []FileData{{1, 1, 2}, {0, 0, 1}},
		},
		{
			Name:  "example input from aoc",
			Files: files,
			Gaps:  gaps,
			WantFiles: []FileData{
				{9, 2, 2},
				{8, 36, 4},
				{7, 8, 3},
				{6, 27, 4},
				{5, 22, 4},
				{4, 12, 2},
				{3, 15, 3},
				{2, 4, 1},
				{1, 5, 3},
				{0, 0, 2},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := RearrangeDiskByCopyWholeFiles(testcase.Files, testcase.Gaps)
			if !reflect.DeepEqual(got, testcase.WantFiles) {
				t.Errorf("Wrong output: got %v, want %v", got, testcase.WantFiles)
			}
		})
	}
}

func TestComputeDiskChecksumPart2(t *testing.T) {
	input := "2333133121414131402"
	want := 2858
	files, gaps := ReadFilesAndGapsFromInput(input)
	got := ComputeDiskChecksumPart2(files, gaps)

	if got != want {
		t.Errorf("Got wrong output: got %d, want %d", got, want)
	}
}

func BenchmarkPart1Solution(b *testing.B) {
	disk := GetDiskFromInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeDiskChecksumPart1(disk)
	}
}
func BenchmarkPart2Solution(b *testing.B) {
	files, gaps := ReadFilesAndGapsFromInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeDiskChecksumPart2(files, gaps)
	}
}
