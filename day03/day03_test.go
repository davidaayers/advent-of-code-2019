package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	wire1Path []string
	wire2Path []string
	closestIntersectionDistance int
} {
	{
		[]string{"R8","U5","L5","D3"},
		[]string{"U7","R6","D4","L4"},
		6,
	},
	{
		[]string{"R75","D30","R83","U83","L12","D49","R71","U7","L72"},
		[]string{"U62","R66","U55","R34","D71","R55","D58","R83"},
		159,
	},
	{
		[]string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"},
		[]string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"},
		135,
	},
}

func TestFindClosestIntersectionDistance(t *testing.T) {
	for _, testCase := range testCases {
		distance := FindClosestIntersectionDistance(testCase.wire1Path,testCase.wire2Path)
		if distance != testCase.closestIntersectionDistance {
			t.Errorf("Error, expected %v got %v", testCase.closestIntersectionDistance, distance)	
		}
	}
}

// go test -timeout 30s day03 -run '^(TestRunIntCode)$'
func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: "
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