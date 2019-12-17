package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	scaffoldMap                 string
	expectedAlignmentParameters int
}{
	{
		`..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..`,
		76,
	},
}

func TestCheckAlignment(t *testing.T) {
	for _, testCase := range testCases {
		scaffoldMap := ParseMap(testCase.scaffoldMap)
		output := CheckAlignment(scaffoldMap)
		if output != testCase.expectedAlignmentParameters {
			t.Errorf("Error, expected %v got %v", testCase.expectedAlignmentParameters, output)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 2804"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: WRONG"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
