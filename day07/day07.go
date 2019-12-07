package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	//"sort"
	//"strconv"
	//"strings"
)

type Amplifier struct {
	intCode            []int
	phase              int
	instructionPointer int
}

func CalculateThrust(intCode []int, phaseSequence []int, inputSignal int) int {
	for _, p := range phaseSequence {
		input := []int{p, inputSignal}
		ic := intCode
		output, _, _ := RunIntCode(ic, input, 0, false)
		inputSignal = output[0]
	}
	return inputSignal
}

func CalculateThrustWithFeedbackLoop(intCode []int, phaseSequence []int, inputSignal int) int {
	amplifiers := make([]Amplifier, len(phaseSequence))
	for i, p := range phaseSequence {
		sliceCopy := make([]int, len(intCode))
		copy(sliceCopy, intCode)
		amplifiers[i] = Amplifier{
			intCode:            sliceCopy,
			phase:              p,
			instructionPointer: 0,
		}
	}

	// initialize and run the first loop
	for i, p := range phaseSequence {
		input := []int{p, inputSignal}
		output, lastPointer, _ := RunIntCode(amplifiers[i].intCode, input, 0, true)
		amplifiers[i].instructionPointer = lastPointer
		inputSignal = output[0]
	}

	// now, let the feedback loop run
	runningAmplifier := 0
	for {
		// last inputSignal is the output from the last amplifier, feed that into the next
		input := []int{inputSignal}
		output, lastPointer, terminated := RunIntCode(amplifiers[runningAmplifier].intCode, input, amplifiers[runningAmplifier].instructionPointer, true)
		amplifiers[runningAmplifier].instructionPointer = lastPointer
		if terminated {
			break
		}
		inputSignal = output[0]
		runningAmplifier++
		if runningAmplifier > 4 {
			runningAmplifier = 0
		}
	}

	return inputSignal
}

var opCodeLengths = map[int]int{
	1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4,
}

// RunIntCode Our intCode interpreter
func RunIntCode(code []int, input []int, instructionPointer int, shouldPauseOnOutput bool) (output []int, lastPointer int, terminated bool) {
	output = make([]int, 0)
	inputIdx := 0
	for {

		paddedIntCode := fmt.Sprintf("%05d", code[instructionPointer])
		opCode, _ := strconv.Atoi(paddedIntCode[3:5])

		if opCode == 99 {
			return output, instructionPointer, true
		}

		paramModeMap := map[int]bool{
			1: paddedIntCode[2:3] == "0",
			2: paddedIntCode[1:2] == "0",
			3: paddedIntCode[0:1] == "0",
		}

		getParam := func(pos int) int {
			param := code[instructionPointer+pos]
			if paramModeMap[pos] {
				param = code[code[instructionPointer+pos]]
			}
			return param
		}

		if opCode == 1 || opCode == 2 {
			// Opcode `1` adds together numbers read from two positions and stores the result in a third position.
			// Opcode `2` works exactly like opcode `1`, except it multiplies the two inputs instead of adding them.
			if opCode == 1 {
				code[code[instructionPointer+3]] = getParam(1) + getParam(2)
			} else if opCode == 2 {
				code[code[instructionPointer+3]] = getParam(1) * getParam(2)
			}
			instructionPointer += opCodeLengths[opCode]
		} else if opCode == 3 {
			// Opcode `3` takes a single integer as input and saves it to the address given by its only parameter.
			code[code[instructionPointer+1]] = input[inputIdx]
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
			if (opCode == 7 && getParam(1) < getParam(2)) || (opCode == 8 && getParam(1) == getParam(2)) {
				code[code[instructionPointer+3]] = 1
			} else {
				code[code[instructionPointer+3]] = 0
			}
			instructionPointer += opCodeLengths[opCode]
		} else {
			panic("Unexpected op code")
		}
	}
}

// ParseIntCode Parses a comma delimted input string into an array of intCode
func ParseIntCode(input string) []int {
	strs := strings.Split(strings.ReplaceAll(input, "\r\n", ""), ",")
	intCode := make([]int, len(strs))
	for idx, s := range strs {
		i, _ := strconv.Atoi(s)
		intCode[idx] = i
	}
	return intCode
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)

	permutations := permutations([]int{0, 1, 2, 3, 4})
	highestThrust := 0

	for _, permutation := range permutations {
		thrust := CalculateThrust(intCode, permutation, 0)
		if thrust > highestThrust {
			highestThrust = thrust
		}
	}

	return "Answer: " + strconv.Itoa(highestThrust)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	intCode := ParseIntCode(input)

	permutations := permutations([]int{5, 6, 7, 8, 9})
	highestThrust := 0

	for _, permutation := range permutations {
		thrust := CalculateThrustWithFeedbackLoop(intCode, permutation, 0)
		if thrust > highestThrust {
			highestThrust = thrust
		}
	}

	return "Answer: " + strconv.Itoa(highestThrust)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
