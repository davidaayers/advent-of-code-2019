package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type Neptune struct {
	board          Board
	startX, startY int
	keys           map[string]Point
	doors          map[string]Point
}

func buildNeptune(input string) Neptune {
	neptune := Neptune{
		keys:  make(map[string]Point, 0),
		doors: make(map[string]Point, 0),
	}

	var keyFinder = func(point Point) {
		keysMatch := regexp.MustCompile(`[a-z]+`)
		doorsMatch := regexp.MustCompile(`[A-Z]+`)

		if keysMatch.MatchString(point.tile) {
			fmt.Printf("Found key: %v\n", point)
			neptune.keys[point.tile] = point
		}

		if doorsMatch.MatchString(point.tile) {
			fmt.Printf("Found door: %v\n", point)
			neptune.doors[point.tile] = point
		}

		if point.tile == "@" {
			neptune.startX = point.x
			neptune.startY = point.y
		}
	}
	neptune.board = ParseBoard(input, keyFinder, true)
	return neptune
}

func FindLeastSteps(neptune Neptune) (steps int) {
	steps = 0
	return
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	neptune := buildNeptune(input)

	fmt.Printf("%v\n", neptune)

	return "Answer: "
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	return "Answer: "
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	//fmt.Println("Part 2: " + Part2(string(bytes)))
}
