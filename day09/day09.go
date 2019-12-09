package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var opCodeLengths = map[int]int{
	1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2,
}

const positionMode, immediateMode, relativeMode = "0", "1", "2"

// RunIntCode Our intCode interpreter
func RunIntCode(initialCode []int, input []int, instructionPointer int, shouldPauseOnOutput bool) (output []int, lastPointer int, terminated bool) {

	// copy the code into a large memory buffer
	code := make([]int, 3000)
	copy(code, initialCode)

	output = make([]int, 0)
	inputIdx := 0
	relativeBase := 0
	for {

		paddedIntCode := fmt.Sprintf("%05d", code[instructionPointer])
		opCode, _ := strconv.Atoi(paddedIntCode[3:5])

		if opCode == 99 {
			return output, instructionPointer, true
		}

		paramMode := map[int]string{
			1: paddedIntCode[2:3],
			2: paddedIntCode[1:2],
			3: paddedIntCode[0:1],
		}

		getParam := func(pos int) int {
			switch paramMode[pos] {
			case positionMode:
				return code[code[instructionPointer+pos]]
			case immediateMode:
				return code[instructionPointer+pos]
			case relativeMode:
				return code[relativeBase+code[instructionPointer+pos]]
			}
			panic("Invalid param mode")
		}

		if opCode == 1 || opCode == 2 {
			// Opcode `1` adds together numbers read from two positions and stores the result in a third position.
			// Opcode `2` works exactly like opcode `1`, except it multiplies the two inputs instead of adding them.
			var answer int
			if opCode == 1 {
				answer = getParam(1) + getParam(2)
			} else if opCode == 2 {
				answer = getParam(1) * getParam(2)
			}

			if paramMode[3] == positionMode {
				code[code[instructionPointer+3]] = answer
			} else {
				code[relativeBase+code[instructionPointer+3]] = answer
			}

			instructionPointer += opCodeLengths[opCode]
		} else if opCode == 3 {
			// Opcode `3` takes a single integer as input and saves it to the address given by its only parameter.
			if paramMode[1] == positionMode {
				code[code[instructionPointer+1]] = input[inputIdx]
			} else {
				code[relativeBase+code[instructionPointer+1]] = input[inputIdx]
			}

			inputIdx++
			instructionPointer += opCodeLengths[opCode]
		} else if opCode == 4 {
			// Opcode `4` outputs the value of its only parameter.
			output = append(output, getParam(1))
			instructionPointer += opCodeLengths[opCode]

			if shouldPauseOnOutput {
				return output, instructionPointer, false
			}
		} else if opCode == 5 || opCode == 6 {
			// Opcode `5` is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the
			// value from the second parameter. Otherwise, it does nothing.
			// Opcode `6` is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the
			// value from the second parameter. Otherwise, it does nothing.
			if (opCode == 5 && getParam(1) != 0) || (opCode == 6 && getParam(1) == 0) {
				instructionPointer = getParam(2)
			} else {
				instructionPointer += opCodeLengths[opCode]
			}
		} else if opCode == 7 || opCode == 8 {
			// Opcode `7` is less than: if the first parameter is less than the second parameter, it stores 1 in
			// the position given by the third parameter. Otherwise, it stores 0.
			// Opcode `8` is equals: if the first parameter is equal to the second parameter, it stores 1 in the
			// position given by the third parameter. Otherwise, it stores 0.
			answer := 0
			if (opCode == 7 && getParam(1) < getParam(2)) || (opCode == 8 && getParam(1) == getParam(2)) {
				answer = 1
			}

			if paramMode[3] == positionMode {
				code[code[instructionPointer+3]] = answer
			} else {
				code[relativeBase+code[instructionPointer+3]] = answer
			}

			instructionPointer += opCodeLengths[opCode]
		} else if opCode == 9 {
			// Opcode 9 adjusts the relative base by the value of its only parameter. The relative base increases
			// (or decreases, if the value is negative) by the value of the parameter. For example, if the relative
			// base is 2000, then after the instruction 109,19, the relative base would be 2019. If the next
			// instruction were 204,-34, then the value at address 1985 would be output.
			relativeBase += getParam(1)
			instructionPointer += opCodeLengths[opCode]
		} else {
			panic("Unexpected op code: " + paddedIntCode)
		}
	}
}

// ParseIntCode Parses a comma delimited input string into an array of intCode
func ParseIntCode(input string) []int {
	strs := strings.Split(strings.ReplaceAll(input, "\r\n", ""), ",")
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
	output, _, _ := RunIntCode(intCode, []int{1}, 0, false)
	if len(output) != 1 {
		fmt.Printf("Output: %v\n", output)
	}
	return "Answer: " + strconv.Itoa(output[len(output)-1])
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
