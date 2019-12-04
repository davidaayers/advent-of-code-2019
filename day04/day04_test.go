package main

import (
	"testing"
)

// go test -timeout 30s day03 -run '^(TestPart1)$'
func TestPart1(t *testing.T) {
	start,end := 197487,673251
	expected := "Answer: "
	answer := Part1(start,end)
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

// go test -timeout 30s day03 -run '^(TestPart2)$'
func TestPart2(t *testing.T) {
	start,end := 197487,673251
	expected := "Answer: "
	answer := Part2(start,end)
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}