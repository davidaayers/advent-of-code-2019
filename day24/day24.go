package main

import (
	"fmt"
	"io/ioutil"
)

func BuildBoard(board string) Board {
	b := ParseBoard(board, func(point Point) {}, true)
	return b
}

func CalculateBiodiversityRating(board Board) int {
	fmt.Printf("Board: %v\n", board)

	return 0
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	return "Answer: "
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
