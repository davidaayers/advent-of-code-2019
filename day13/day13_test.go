package main

import (
	"io/ioutil"
	"testing"
)

// go test -timeout 30s day03 -run '^(TestPart1)$'
func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 318"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

// go test -timeout 30s day03 -run '^(TestPart2)$'
func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 16309"
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
