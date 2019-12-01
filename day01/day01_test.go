package main

import "testing"

// go test -timeout 30s day01 -run '^(TestCalcFuelIncludingWeightOfFuel)$'

var testCases = []struct {
	moduleWeight int
	expectedFuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

var testCases2 = []struct {
	moduleWeight int
	expectedFuel int
}{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func TestCalcFuel(t *testing.T) {
	for _, testCase := range testCases {
		result := CalcFuel(testCase.moduleWeight)
		if result != testCase.expectedFuel {
			t.Errorf("Error, expected %d fuel, got %d", testCase.expectedFuel, result)
		}
	}
}

func TestCalcFuelIncludingWeightOfFuel(t *testing.T) {
	for _, testCase := range testCases2 {
		result := CalcFuelIncludingWeightOfFuel(testCase.moduleWeight,0)
		if result != testCase.expectedFuel {
			t.Errorf("Error, expected %d fuel, got %d", testCase.expectedFuel, result)
		}
	}
}

