package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)

	cnt := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			output := RunCode(intCode, x, y)
			if output == 1 {
				//fmt.Print("#")
				cnt++
			} else {
				//fmt.Print(".")
			}
		}
		//fmt.Print("\n")
	}

	return "Answer: " + strconv.Itoa(cnt)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	intCode := ParseIntCode(input)

	cnt := 0
	for y := 500; y < 1500; y++ {
		for x := 500; x < 1500; x++ {

			// check that x,y is in beam
			output := RunCode(intCode, x, y)
			if output != 1 {
				continue
			}

			// check that right top corner is in beam
			output = RunCode(intCode, x+99, y)
			if output != 1 {
				continue
			}

			// check that left bottom corner is in beam
			output = RunCode(intCode, x, y+99)
			if output != 1 {
				continue
			}

			return "Answer: " + strconv.Itoa(x*10000+y)
		}
	}

	panic("No Answer found")
}

func RunCode(intCode []int, x int, y int) int {
	memory := make([]int, 5000)
	copy(memory, intCode)
	output, _, _, _ := RunIntCode(memory, []int{x, y}, 0, 0, true)
	return output[0]
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
