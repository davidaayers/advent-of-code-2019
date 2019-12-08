package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func sifChecksum(sif string, width int, height int) int {
	leastZeros := math.MaxInt64
	leastZeroLayer := map[int]int{}

	for i := 0; i < len(sif); i += width * height {
		layerStr := sif[i : i+width*height]
		layer := make(map[int]int)
		pixels := strings.Split(layerStr, "")

		for _, pixelStr := range pixels {
			pixel, _ := strconv.Atoi(pixelStr)
			layer[pixel]++
		}

		if layer[0] < leastZeros {
			leastZeros = layer[0]
			leastZeroLayer = layer
		}
	}
	return leastZeroLayer[1] * leastZeroLayer[2]
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	width, height := 25, 6
	return "Answer: " + strconv.Itoa(sifChecksum(input, width, height))
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
