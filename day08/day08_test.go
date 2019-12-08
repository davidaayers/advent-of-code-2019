package main

import (
	"io/ioutil"
	"testing"
)

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: 1064"
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
