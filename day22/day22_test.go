package main

import (
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	shuffleInstructions []string
	expectedOrder       []int
}{
	{
		[]string{
			"deal with increment 7",
			"deal into new stack",
			"deal into new stack",
		},
		[]int{
			0, 3, 6, 9, 2, 5, 8, 1, 4, 7,
		},
	},
	{
		[]string{
			"cut 6",
			"deal with increment 7",
			"deal into new stack",
		},
		[]int{
			3, 0, 7, 4, 1, 8, 5, 2, 9, 6,
		},
	},
	{
		[]string{
			"deal with increment 7",
			"deal with increment 9",
			"cut -2",
		},
		[]int{
			6, 3, 0, 7, 4, 1, 8, 5, 2, 9,
		},
	},
	{
		[]string{
			"deal into new stack",
			"cut -2",
			"deal with increment 7",
			"cut 8",
			"cut -4",
			"deal with increment 7",
			"cut 3",
			"deal with increment 9",
			"deal with increment 3",
			"cut -1",
		},
		[]int{
			9, 2, 5, 8, 1, 4, 7, 0, 3, 6,
		},
	},
}

func TestShuffle(t *testing.T) {
	for _, testCase := range testCases {
		cards := Shuffle(len(testCase.expectedOrder), testCase.shuffleInstructions)
		if !Equal(cards, testCase.expectedOrder) {
			t.Errorf("Error, expected %v got %v", testCase.expectedOrder, cards)
		}
	}
}

func TestCut(t *testing.T) {
	deck := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cutDeck := Cut(deck, 3)
	expected := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	if !Equal(cutDeck, expected) {
		t.Errorf("Error, expected %v got %v", expected, cutDeck)
	}

	cutDeck = Cut(deck, -4)
	expected = []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	if !Equal(cutDeck, expected) {
		t.Errorf("Error, expected %v got %v", expected, cutDeck)
	}
}

func TestDealWithIncrement(t *testing.T) {
	deck := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	dealtDeck := DealWithIncrement(deck, 3)
	expected := []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}
	if !Equal(dealtDeck, expected) {
		t.Errorf("Error, expected %v got %v", expected, dealtDeck)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: WRONG"
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
