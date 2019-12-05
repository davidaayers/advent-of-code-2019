package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	startingCode []int
	input        int
	expected     []int
}{
	// Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
	{
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		8,
		[]int{1},
	},
	{
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		9,
		[]int{0},
	},
	// Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
	{
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		9,
		[]int{0},
	},
	{
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		3,
		[]int{1},
	},
	// 3,3,1108,-1,8,3,4,3,99 - Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
	{
		[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		8,
		[]int{1},
	},
	{
		[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		9,
		[]int{0},
	},
	// 3,3,1107,-1,8,3,4,3,99 - Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
	{
		[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		3,
		[]int{1},
	},
	{
		[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		9,
		[]int{0},
	},
	// Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero:
	// (using position mode)
	{
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		0,
		[]int{0},
	},
	{
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		1,
		[]int{1},
	},
	// (using immediate mode)
	{
		[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		0,
		[]int{0},
	},
	{
		[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		1,
		[]int{1},
	},
	// This example program uses an input instruction to ask for a single number. The program will then
	// output `999` if the input value is below `8`, output `1000` if the input value is equal to `8`, or
	// output `1001` if the input value is greater than 8.
	//
	{
		[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		1,
		[]int{999},
	},
	{
		[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		8,
		[]int{1000},
	},
	{
		[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		9,
		[]int{1001},
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

// go test -timeout 30s day05 -run '^(TestRunIntCode)$'
func TestRunIntCode(t *testing.T) {
	for _, testCase := range testCases {
		output := RunIntCode(testCase.startingCode, testCase.input)
		if !Equal(output, testCase.expected) {
			t.Errorf("Error, expected %v got %v", testCase.expected, output)
		}
	}
}

// go test -timeout 30s day05 -run '^(TestPart1)$'
func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 12440243"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

// go test -timeout 30s day05 -run '^(TestPart2)$'
func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 15486302"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
