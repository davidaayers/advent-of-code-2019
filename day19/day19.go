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
			memory := make([]int, 5000)
			copy(memory, intCode)
			output, _, _, _ := RunIntCode(memory, []int{x, y}, 0, 0, true)
			fmt.Printf("x=%v,y=%v,output[0]=%v\n", x, y, output[0])
			if output[0] == 1 {
				cnt++
			}
		}
	}

	return "Answer: " + strconv.Itoa(cnt)
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
