package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	startingCode []int
	input        []int
	expected     []int
}{
	// takes no input and produces a copy of itself as output.
	{
		[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		[]int{},
		[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
	},
	// should output a 16-digit number.
	{
		[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
		[]int{},
		[]int{1219070632396864},
	},
	// should output the large number in the middle.
	{
		[]int{104, 1125899906842624, 99},
		[]int{},
		[]int{1125899906842624},
	},
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRunIntCode(t *testing.T) {
	for _, testCase := range testCases {
		output, _, _ := RunIntCode(testCase.startingCode, testCase.input, 0, false)
		if !Equal(output, testCase.expected) {
			t.Errorf("Error, expected %v got %v", testCase.expected, output)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 3989758265"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: "
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
