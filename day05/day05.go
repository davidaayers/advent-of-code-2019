package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var opCodeLengths = map[int]int{
	1:4,
	2:4,
	3:2,
	4:2,
	5:3,
	6:3,
	7:4,
	8:4,
}

// RunIntCode Our intCode interpreter
func RunIntCode(code []int,input int) []int {
	output := make([]int,0)	
	idx := 0
	for {

		paddedIntCode := fmt.Sprintf("%05d",code[idx])
		param1PositionMode := paddedIntCode[2:3] == "0"
		param2PositionMode := paddedIntCode[1:2] == "0"

		opCode, _ := strconv.Atoi(paddedIntCode[3:5])

		if opCode == 99 {
			return output
		} 
		
		if opCode == 1 || opCode == 2 {
			// Opcode `1` adds together numbers read from two positions and stores the result in a third position.
			// Opcode `2` works exactly like opcode `1`, except it multiplies the two inputs instead of adding them.
			param1 := code[idx+1]
			if param1PositionMode {
				param1 = code[code[idx+1]]
			}

			param2 := code[idx+2]
			if param2PositionMode {
				param2 = code[code[idx+2]]
			}

			if opCode == 1 {
				code[code[idx+3]] = param1 + param2	
			} else if opCode == 2 {
				code[code[idx+3]] = param1 * param2	
			}
		} else if opCode == 3 {
			// Opcode `3` takes a single integer as input and saves it to the address given by its only parameter.
			code[code[idx+1]] = input
		} else if opCode == 4 {
			// Opcode `4` outputs the value of its only parameter.
			param1 := code[idx+1]
			if param1PositionMode {
				param1 = code[code[idx+1]]
			}
			output = append(output,param1)
		} else if opCode == 5 || opCode == 6 {
			// Opcode `5` is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the 
			// value from the second parameter. Otherwise, it does nothing.
			// Opcode `6` is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the 
			// value from the second parameter. Otherwise, it does nothing.
			param1 := code[idx+1]
			if param1PositionMode {
				param1 = code[code[idx+1]]
			}

			param2 := code[idx+2]
			if param2PositionMode {
				param2 = code[code[idx+2]]
			}

			if (opCode == 5 && param1 != 0)||(opCode == 6 && param1 == 0) {
				idx = param2
			} else {
				idx += opCodeLengths[opCode]
			}
		} else if opCode == 7 || opCode == 8 {
			// Opcode `7` is less than: if the first parameter is less than the second parameter, it stores 1 in 
			// the position given by the third parameter. Otherwise, it stores 0.
			// Opcode `8` is equals: if the first parameter is equal to the second parameter, it stores 1 in the 
			// position given by the third parameter. Otherwise, it stores 0.
			param1 := code[idx+1]
			if param1PositionMode {
				param1 = code[code[idx+1]]
			}

			param2 := code[idx+2]
			if param2PositionMode {
				param2 = code[code[idx+2]]
			}

			if (opCode == 7 && param1 < param2) || (opCode == 8 && param1 == param2) {
				code[code[idx+3]] = 1
			} else {
				code[code[idx+3]] = 0
			}
		}

		if opCode != 5 && opCode != 6 {
			idx += opCodeLengths[opCode]
		}
	}

	panic("Abnormal termination of int code")
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

	output := RunIntCode(intCode,1)

	return "Answer: " + strconv.Itoa(output[len(output)-1])
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	intCode := ParseIntCode(input)

	output := RunIntCode(intCode,5)

	return "Answer: " + strconv.Itoa(output[len(output)-1])
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
