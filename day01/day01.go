package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"strconv"
	"strings"
)

// Part1 Part 1 of uzzle
func Part1(input string) string {
 	var totalFuel = 0

	for _, s := range strings.Split(input, "\r\n") {
		w, _ := strconv.Atoi(s) 
		totalFuel += CalcFuel(w)
	}

	return "Answer " + strconv.Itoa(totalFuel)
}

// CalcFuel Calculate the fuel equired for oe module
func CalcFuel(weight int) int {
	return weight / 3 - 2
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	return "Answer"
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	//fmt.Println("Part 2: " + Part2(string(bytes)))
}
