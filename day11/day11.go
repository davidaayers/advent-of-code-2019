package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	//"sort"
	//"strconv"
	//"strings"
)

type Dir struct {
	dx, dy int
}

var north = Dir{dx: 0, dy: -1}
var south = Dir{dx: 0, dy: 1}
var east = Dir{dx: 1, dy: 0}
var west = Dir{dx: -1, dy: 0}

var dirs = []Dir{north, east, south, west}

type Robot struct {
	x, y   int
	memory []int
	dir    Dir
}

func (robot *Robot) turn(turnDirection int) {
	// (0 = left, 1 = right)
	var newIdx int
	currIdx := indexOf(dirs, robot.dir)
	if turnDirection == 0 {
		newIdx = currIdx - 1
		if newIdx < 0 {
			newIdx = len(dirs) - 1
		}
	} else {
		newIdx = currIdx + 1
		if newIdx > len(dirs)-1 {
			newIdx = 0
		}
	}
	robot.dir = dirs[newIdx]
}

func indexOf(dirList []Dir, dir Dir) int {
	for i := 0; i < len(dirList); i++ {
		if dirList[i] == dir {
			return i
		}
	}
	return -1
}

type Hull struct {
	width, height int
	hullMap       [][]string
}

func (hull Hull) printHull() {
	for y := 0; y < hull.height; y++ {
		for x := 0; x < hull.width; x++ {
			fmt.Print(hull.hullMap[y][x])
		}
		fmt.Print("\n")
	}
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)

	robot := Robot{
		x:      100,
		y:      100,
		memory: make([]int, 3000), // allocate enough memory
		dir:    north,
	}

	// copy the code into the robot's memory
	copy(robot.memory, intCode)

	hull := Hull{width: 200, height: 200}
	hull.hullMap = make([][]string, hull.height)
	for y := 0; y < hull.height; y++ {
		hull.hullMap[y] = make([]string, hull.width)
		for x := 0; x < hull.width; x++ {
			hull.hullMap[y][x] = "_"
		}
	}

	copy(robot.memory, intCode)

	instructionPointer := 0
	numPanelsPainted := 0
	for {
		// look at the panel under the robot, and feed that color as the input
		panelColor := 0
		if hull.hullMap[robot.y][robot.x] == "#" {
			panelColor = 1
		}

		// output should be two things:
		// 0 - the color to paint the panel we're over (0 = black, 1 = white)
		// 1 - the direction to turn (0 = left, 1 = right)
		// will pause on each output, so we need to run it twice, once to get the color, then again
		// to get the turn direction
		output, lastPointer, terminated := RunIntCode(robot.memory, []int{panelColor}, instructionPointer, true)
		instructionPointer = lastPointer

		if terminated {
			break
		}

		color := output[0]

		// now, run again to get the turn direction
		output, lastPointer, terminated = RunIntCode(robot.memory, []int{panelColor}, instructionPointer, true)
		instructionPointer = lastPointer

		turnDirection := output[0]

		// if we've never painted this panel before, count it
		if hull.hullMap[robot.y][robot.x] == "_" {
			numPanelsPainted++
		}

		// paint the square
		if color == 0 {
			hull.hullMap[robot.y][robot.x] = "."
		} else {
			hull.hullMap[robot.y][robot.x] = "#"
		}

		// now turn the robot and move 1 in the new direction
		robot.turn(turnDirection)
		robot.x = robot.x + robot.dir.dx
		robot.y = robot.y + robot.dir.dy
	}

	hull.printHull()

	return "Answer: " + strconv.Itoa(numPanelsPainted)
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
