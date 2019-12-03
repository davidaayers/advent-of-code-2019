package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	//"strings"
)


func RunIntCode(code []int) {
	idx := 0
	for {
		opCode := code[idx]

		if opCode == 99 {
			break
		} else if opCode == 1 {
			code[code[idx+3]] = code[code[idx+1]] + code[code[idx+2]]
		} else if opCode == 2 {
			code[code[idx+3]] = code[code[idx+1]] * code[code[idx+2]]
		}

		idx += 4
	}
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	return "Answer"
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	return "Answer"
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
