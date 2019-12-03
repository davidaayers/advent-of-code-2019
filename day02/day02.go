package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// RunIntCode Our intCode interpreter
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

// ParseIntCode Parses a comma delimted input string into an array of intCode
func ParseIntCode(input string) []int {
	strs := strings.Split(strings.ReplaceAll(input,"\r\n",""),",")
	intCode := make([]int, len(strs))
	for idx, s := range strs {
		i, _ := strconv.Atoi(s) 
		intCode[idx] = i
	}
	return intCode
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)

	// per the puzzle description
	intCode[1] = 12
	intCode[2] = 2

	RunIntCode(intCode)

	return "Answer: " + strconv.Itoa(intCode[0])
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	intCode := ParseIntCode(input)

	for noun:=0; noun < 100; noun++ {
		for verb:=0; verb < 100; verb++ {
			codeCopy := make([]int, len(intCode))
			copy(codeCopy,intCode)

			codeCopy[1] = noun
			codeCopy[2] = verb

			RunIntCode(codeCopy)

			if codeCopy[0] == 19690720 {
				return "Answer: " + strconv.Itoa(100 * noun + verb)
			}
		}
	}

	return "Answer Not found!"
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
