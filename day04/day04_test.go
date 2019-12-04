package main

import (
	"testing"
)

var testCases = []struct {
	password int
	valid bool
} {
	{ 111111, true  },
	{ 223450, false },
	{ 123789, false },
}

var testCases2 = []struct {
	password int
	valid bool
} {
	{ 112233, true  },
	{ 123444, false },
	{ 111122, true },
}

// go test -timeout 30s day04 -run '^(TestIsValidPassword)$'
func TestIsValidPassword(t *testing.T) {
	for _, testCase := range testCases {
		isValid := IsValidPassword(testCase.password)
		if isValid != testCase.valid {
			t.Errorf("For %v, expected %v got %v", testCase.password, testCase.valid, isValid)	
		}
	}
}

// go test -timeout 30s day04 -run '^(TestIsMoreValidPassword)$'
func TestIsMoreValidPassword(t *testing.T) {
	for _, testCase := range testCases2 {
		isValid := IsMoreValidPassword(testCase.password)
		if isValid != testCase.valid {
			t.Errorf("For %v, expected %v got %v", testCase.password, testCase.valid, isValid)	
		}
	}
}

// go test -timeout 30s day04 -run '^(TestPart1)$'
func TestPart1(t *testing.T) {
	start,end := 197487,673251
	expected := "Answer: 1640"
	answer := Part1(start,end)
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

// go test -timeout 30s day04 -run '^(TestPart2)$'
func TestPart2(t *testing.T) {
	start,end := 197487,673251
	expected := "Answer: 1126"
	answer := Part2(start,end)
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}