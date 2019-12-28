package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Shuffle(numCards int, instructions []string) []int {
	// build the deck
	deck := make([]int, numCards)
	for c := 0; c < numCards; c++ {
		deck[c] = c
	}

	for _, i := range instructions {
		if i == "deal into new stack" {
			deck = DealIntoNewStack(deck)
		} else if strings.HasPrefix(i, "cut") {
			deck = Cut(deck, NumAtEnd(i))
		} else if strings.HasPrefix(i, "deal with increment") {
			deck = DealWithIncrement(deck, NumAtEnd(i))
		} else {
			panic(fmt.Sprintf("Unknown shuffle instruction: %v\n", i))
		}
	}

	return deck
}

func NumAtEnd(i string) int {
	parts := strings.Split(i, " ")
	num, _ := strconv.Atoi(parts[len(parts)-1])
	return num
}

func DealIntoNewStack(deck []int) []int {
	// reverse routine from: https://github.com/golang/go/wiki/SliceTricks
	for i := len(deck)/2 - 1; i >= 0; i-- {
		opp := len(deck) - 1 - i
		deck[i], deck[opp] = deck[opp], deck[i]
	}
	return deck
}

func Cut(deck []int, n int) []int {
	if n > 0 {
		return append(deck[n:], deck[:n]...)
	}
	return append(deck[len(deck)+n:], deck[0:len(deck)+n]...)
}

func DealWithIncrement(deck []int, increment int) []int {
	newDeck := make([]int, len(deck))

	cnt := 0
	for _, card := range deck {
		newDeck[cnt] = card
		cnt += increment
		if cnt > len(deck) {
			cnt = cnt - len(deck)
		}
	}

	return newDeck
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	instructions := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	deck := Shuffle(10007, instructions)

	answer := 0
	for i, card := range deck {
		if card == 2019 {
			answer = i
			break
		}
	}

	return "Answer: " + strconv.Itoa(answer)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	return "Answer: "
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
