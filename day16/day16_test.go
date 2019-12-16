package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	signal    string
	numPhases int
	output    string
}{
	{
		"12345678",
		4,
		"01029498",
	},
	{
		"80871224585914546619083218645595",
		100,
		"24176176",
	},
	{
		"19617804207202209144916044189917",
		100,
		"73745418",
	},
	{
		"69317163492948606335995924319873",
		100,
		"52432133",
	},
}

func TestCleanupSignal(t *testing.T) {
	for _, testCase := range testCases {
		output := CleanupSignal(testCase.signal, testCase.numPhases, 0)
		if output[0:8] != testCase.output {
			t.Errorf("Error, expected %v got %v", testCase.output, output)
		}
	}
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 52611030"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: WRONG"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
