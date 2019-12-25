package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Part1 Part 1 of puzzle
func Part1(intCodeStr string) string {
	intCode := ParseIntCode(intCodeStr)
	memory := make([]int, 10000)
	copy(memory, intCode)

	reader := bufio.NewReader(os.Stdin)
	instructionPointer := 0
	relativeBase := 0
	terminated := false
	var output []int
	var lastWord = ""
	var currentWord = ""
	i := make([]int, 0)
	for {
		output, instructionPointer, relativeBase, terminated = RunIntCode(memory, i, instructionPointer, relativeBase, true)

		if terminated {
			break
		}

		letter := string(output[0])
		fmt.Print(letter)
		if letter == "\n" {
			lastWord = currentWord
			currentWord = ""
		} else {
			currentWord += letter
		}

		if lastWord == "Command?" {
			fmt.Print(" ")
			command, _ := reader.ReadString('\n')
			i = toIntSlice(command)
		}

	}

	return "Answer: "
}

func toIntSlice(str string) []int {
	sl := make([]int, 0)
	for _, r := range str {
		sl = append(sl, int(r))
	}
	return sl
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
