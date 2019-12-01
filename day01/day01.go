package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// CalcFuel Calculate the fuel equired for oe module
func CalcFuel(weight int) int {
	return weight / 3 - 2
}

func CalcFuelIncludingWeightOfFuel(weight,lastWeight int) int {
	fuel := CalcFuel(weight)

	if fuel <= 0 {
		return lastWeight
	}

	return CalcFuelIncludingWeightOfFuel(fuel,lastWeight+fuel)
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
 	var totalFuel = 0

	for _, s := range strings.Split(input, "\r\n") {
		w, _ := strconv.Atoi(s) 
		totalFuel += CalcFuel(w)
	}

	return "Answer " + strconv.Itoa(totalFuel)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
 	var totalFuel = 0

	for _, s := range strings.Split(input, "\r\n") {
		w, _ := strconv.Atoi(s) 
		totalFuel += CalcFuelIncludingWeightOfFuel(w,0)
	}

	return "Answer " + strconv.Itoa(totalFuel)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
