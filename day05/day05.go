package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strconv"
	"strings"
)

var opCodeLengths = map[int]int{
	1:4,
	2:4,
	3:2,
	4:2,
}

// RunIntCode Our intCode interpreter
func RunIntCode(code []int,input int) []int {
	output := make([]int,0)	
	idx := 0
	for {

		paddedIntCode := fmt.Sprintf("%05d",code[idx])
		param1PositionMode := paddedIntCode[2:3] == "0"
		param2PositionMode := paddedIntCode[1:2] == "0"
		//param3PositionMode := paddedIntCode[0:1] == "0"

		opCode, _ := strconv.Atoi(paddedIntCode[3:5])

		fmt.Printf("%v > %v (%v): paramMode[1]: %v paramMode[2]: %v ", idx,paddedIntCode,opCode, param1PositionMode, param2PositionMode)

		if opCode == 99 {
			return output
		} 
		
		if opCode == 1 || opCode == 2 {
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
			} else {
				code[code[idx+3]] = param1 * param2	
			}

			fmt.Printf(" Param1: %v Param2: %v Result: %v Stored in: %v\n",param1,param2,code[code[idx+3]],code[idx+3])

		} else if opCode == 3 {
			code[code[idx+1]] = input
			fmt.Printf(" Input: %v Stored in: %v\n", input, code[idx+1])
		} else if opCode == 4 {
			param1 := code[code[idx+1]]
			if !param1PositionMode {
				param1 = code[idx+1]
			}
			fmt.Printf("Value is %v\n", param1)
			output = append(output,param1)
		}

		idx += opCodeLengths[opCode]
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

	fmt.Println(output)

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
