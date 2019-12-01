package main

import (
	"io/ioutil"
	"testing"
)

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
		result := CalcFuelIncludingWeightOfFuel(testCase.moduleWeight)
		if result != testCase.expectedFuel {
			t.Errorf("Error, expected %d fuel, got %d", testCase.expectedFuel, result)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer 3219099"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer 4825810"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}