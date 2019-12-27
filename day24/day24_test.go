package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	board          string
	expectedRating int
}{
	{
		`....#
#..#.
#..##
..#..
#....`,
		2129920,
	},
}

var testCases2 = []struct {
	board        string
	numMinutes   int
	expectedBugs int
}{
	{
		`....#
#..#.
#.?##
..#..
#....`,
		10,
		99,
	},
}

func TestCalculateBiodiversityRating(t *testing.T) {
	for _, testCase := range testCases {
		board := BuildBoard(testCase.board)
		rating := CalculateBiodiversityRating(board)
		if rating != testCase.expectedRating {
			t.Errorf("Error, expected %v got %v", testCase.expectedRating, rating)
		}
	}
}

func TestCountBugsRecursive(t *testing.T) {
	for _, testCase := range testCases2 {
		board := BuildBoard(testCase.board)
		bugCnt := CountBugsRecursive(board, testCase.numMinutes)
		if bugCnt != testCase.expectedBugs {
			t.Errorf("Error, expected %v got %v", testCase.expectedBugs, bugCnt)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 18842609"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 2059"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
