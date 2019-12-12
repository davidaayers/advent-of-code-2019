package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	coords         string
	expectedEnergy int
	timeSteps      int
}{
	{
		`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
		179,
		10,
	},
	{
		`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
		1940,
		100,
	},
}

func TestFindBestLocation(t *testing.T) {
	for _, testCase := range testCases {
		moons := ParseCoordinates(testCase.coords)
		//fmt.Printf("Moons (start): %v\n", moons)
		MoveTime(testCase.timeSteps, moons)
		energy := CalculateTotalEnergy(moons)
		if energy != testCase.expectedEnergy {
			t.Errorf("Error, expected %v got %v", testCase.expectedEnergy, energy)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 6227"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 331346071640472"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
