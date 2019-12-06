package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	orbitMap       []string
	expectedOrbits int
}{
	{
		[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
		42,
	},
}

var testCases2 = []struct {
	orbitMap          []string
	expectedTransfers int
}{
	{
		[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
		4,
	},
}

func TestCountOrbits(t *testing.T) {
	for _, testCase := range testCases {
		universe := MakeUniverse(testCase.orbitMap)
		orbits := CountOrbits(universe)
		if orbits != testCase.expectedOrbits {
			t.Errorf("Error, expected %v got %v", testCase.expectedOrbits, orbits)
		}
	}
}

func TestCountTransfers(t *testing.T) {
	for _, testCase := range testCases2 {
		universe := MakeUniverse(testCase.orbitMap)
		orbits := CountTransfers(universe)
		if orbits != testCase.expectedTransfers {
			t.Errorf("Error, expected %v got %v", testCase.expectedTransfers, orbits)
		}
	}
}

// go test -timeout 30s day03 -run '^(TestPart1)$'
func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 249308"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

// go test -timeout 30s day03 -run '^(TestPart2)$'
func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 349"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
