package main

import "testing"

var testCases = []struct {
	moduleWeight int
	expectedFuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestCalcFuel(t *testing.T) {
	for _, testCase := range testCases {
		result := CalcFuel(testCase.moduleWeight)
		if result != testCase.expectedFuel {
			t.Errorf("Error, expected %d fuel, got %d", testCase.expectedFuel, result)
		}
	}
}

func TestPart1(t *testing.T) {
	answer := Part1("test")
	if answer != "Answer" {
		t.Errorf("Error")
	}
}
