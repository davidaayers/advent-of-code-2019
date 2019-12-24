package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	board         string
	expectedSteps int
}{
	{
		`#########
#b.A.@.a#
#########`,
		8,
	},
}

func TestFindLeastSteps(t *testing.T) {
	for _, testCase := range testCases {
		neptune := buildNeptune(testCase.board)
		steps := FindLeastSteps(neptune)
		if steps != testCase.expectedSteps {
			t.Errorf("Error, expected %v got %v", testCase.expectedSteps, steps)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: WRONG"
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
