package main

import "testing"

func TestPart1(t *testing.T) {
	answer := Part1("test")
	if answer != "Answer" {
		t.Errorf("Error")
	}
}
