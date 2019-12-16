package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func CleanupSignal(signal string, numPhases int) string {
	basePattern := []int{0, 1, 0, -1}
	output := ""
	for phase := 0; phase < numPhases; phase++ {
		//fmt.Printf("Phase %v:\n", phase+1)
		for e := 0; e < len(signal); e++ {
			// build up the pattern for this element
			pattern := make([]int, 0)
			for _, a := range basePattern {
				for i := 0; i < e+1; i++ {
					pattern = append(pattern, a)
				}
			}
			//fmt.Printf("%v: pattern: %v\n", e,pattern)

			sum := 0
			patternIdx := 1
			for i := 0; i < len(signal); i++ {
				val, _ := strconv.Atoi(string(signal[i]))
				multiplier := pattern[patternIdx]
				patternIdx++
				if patternIdx > len(pattern)-1 {
					patternIdx = 0
				}

				sum += val * multiplier

				//fmt.Printf("%v*%v ", val, multiplier)
			}
			sumStr := strconv.Itoa(sum)
			sumStr = string(sumStr[len(sumStr)-1])
			output += sumStr
			//fmt.Printf(" = %v\n", sumStr)
			//signal = output
		}
		//fmt.Printf("Output: %v\n", output)
		signal = output
		output = ""
	}

	return signal
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	output := CleanupSignal(input, 100)
	return "Answer: " + output[0:8]
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
