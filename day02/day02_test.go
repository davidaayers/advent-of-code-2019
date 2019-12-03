package main

import (
	"io/ioutil"
	"testing"
)

// go test -timeout 30s day01 -run '^(TestRunIntCode)$'
var testCases = []struct {
	startingCode []int 
	endingCode []int
} {
	{
		[]int{1,9,10,3,2,3,11,0,99,30,40,50},
		[]int{3500,9,10,70,2,3,11,0,99,30,40,50},
	},
	{
		[]int{1,0,0,0,99}, 
		[]int{2,0,0,0,99},
	},
	{
		[]int{2,3,0,3,99}, 
		[]int{2,3,0,6,99},
	},
	{
		[]int{2,4,4,5,99,0}, 
		[]int{2,4,4,5,99,9801},
	},
	{
		[]int{1,1,1,4,99,5,6,0,99}, 
		[]int{30,1,1,4,2,5,6,0,99},
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
		RunIntCode(testCase.startingCode)
		if !Equal(testCase.startingCode,testCase.endingCode) {
			t.Errorf("Error, expected %v got %v", testCase.endingCode, testCase.startingCode)	
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}