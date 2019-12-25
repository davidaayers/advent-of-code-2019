package main

import (
	"fmt"
	"io/ioutil"
)

// Part1 Part 1 of puzzle
func Part1(intCodeStr string) string {
	intCode := ParseIntCode(intCodeStr)
	memory := make([]int, 10000)
	copy(memory, intCode)

	//reader := bufio.NewReader(os.Stdin)
	instructionPointer := 0
	relativeBase := 0
	terminated := false
	var output []int
	var lastWord = ""
	var currentWord = ""

	roboCommands := []string{
		"north",
		"take easter egg",
		"east",
		"take astrolabe",
		"south",
		"take space law space brochure",
		"north",
		"north",
		"north",
		"take fuel cell",
		"south",
		"south",
		"west",
		"north",
		"take manifold",
		"north",
		"north",
		"take hologram",
		"north",
		"take weather machine",
		"north",
		"take antenna",
		"west",
	}

	// now append the commands that try every combination of items
	items := []string{
		"easter egg",
		"astrolabe",
		"space law space brochure",
		"fuel cell",
		"manifold",
		"hologram",
		"weather machine",
		"antenna",
	}
	subsets := All(items)

	for _, subset := range subsets {
		// drop everything
		for _, item := range items {
			roboCommands = append(roboCommands, "drop "+item)
		}
		// pick up everything in this subset
		for _, subsubset := range subset {
			roboCommands = append(roboCommands, "take "+subsubset)
		}
		// try to move south
		roboCommands = append(roboCommands, "south")
	}

	roboCommand := 0

	i := make([]int, 0)
	for {
		output, instructionPointer, relativeBase, terminated = RunIntCode(memory, i, instructionPointer, relativeBase, true)

		if terminated {
			break
		}

		letter := string(output[0])
		fmt.Print(letter)
		if letter == "\n" {
			lastWord = currentWord
			currentWord = ""
		} else {
			currentWord += letter
		}

		if lastWord == "Command?" {
			if roboCommand > len(roboCommands)-1 {
				// we need to try all the things
				// first, issue command to drop everything
				//fmt.Print(" ")
				//command, _ := reader.ReadString('\n')
				//i = toIntSlice(command)
				break
			} else {
				nextRoboCommand := roboCommands[roboCommand]
				fmt.Printf(" [%v](%v)\n", roboCommand, nextRoboCommand)
				i = toIntSlice(nextRoboCommand + "\n")
				roboCommand++
			}
		}
	}

	return "Answer: "
}

func toIntSlice(str string) []int {
	sl := make([]int, 0)
	for _, r := range str {
		sl = append(sl, int(r))
	}
	return sl
}

// from:
// https://github.com/mxschmitt/golang-combinations/blob/master/combinations.go
func All(set []string) (subsets [][]string) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
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
