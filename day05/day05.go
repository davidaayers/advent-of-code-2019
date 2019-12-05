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
		opCode, _ := strconv.Atoi(paddedIntCode[3:5])

		if opCode == 99 {
			return output
		} 
		
		paramModeMap := map[int] bool {
			1:paddedIntCode[2:3] == "0",
			2:paddedIntCode[1:2] == "0",
			3:paddedIntCode[0:1] == "0",
		}

		getParam := func(pos int)int {
			param := code[idx+pos]
			if paramModeMap[pos] {
				param = code[code[idx+pos]]
			}
			return param
		}

		if opCode == 1 || opCode == 2 {
			// Opcode `1` adds together numbers read from two positions and stores the result in a third position.
			// Opcode `2` works exactly like opcode `1`, except it multiplies the two inputs instead of adding them.
			if opCode == 1 {
				code[code[idx+3]] = getParam(1) + getParam(2)
			} else if opCode == 2 {
				code[code[idx+3]] = getParam(1) * getParam(2)
			}
		} else if opCode == 3 {
			// Opcode `3` takes a single integer as input and saves it to the address given by its only parameter.
			code[code[idx+1]] = input
		} else if opCode == 4 {
			// Opcode `4` outputs the value of its only parameter.
			output = append(output, getParam(1))
		} else if opCode == 5 || opCode == 6 {
			// Opcode `5` is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the 
			// value from the second parameter. Otherwise, it does nothing.
			// Opcode `6` is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the 
			// value from the second parameter. Otherwise, it does nothing.
			if (opCode == 5 && getParam(1) != 0)||(opCode == 6 && getParam(1) == 0) {
				idx = getParam(2)
			} else {
				idx += opCodeLengths[opCode]
			}
		} else if opCode == 7 || opCode == 8 {
			// Opcode `7` is less than: if the first parameter is less than the second parameter, it stores 1 in 
			// the position given by the third parameter. Otherwise, it stores 0.
			// Opcode `8` is equals: if the first parameter is equal to the second parameter, it stores 1 in the 
			// position given by the third parameter. Otherwise, it stores 0.
			if (opCode == 7 && getParam(1) < getParam(2)) || (opCode == 8 && getParam(1) == getParam(2)) {
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
