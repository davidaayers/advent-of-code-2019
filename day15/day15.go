package main

import (
	"fmt"
	"io/ioutil"
)

type RepairDroid struct {
	memory                                   []int
	lastInstructionPointer, lastRelativeBase int
}

func (droid *RepairDroid) Move(direction int) (result int) {
	output, lastInstructionPointer, lastRelativeBase, _ :=
		RunIntCode(droid.memory, []int{direction}, droid.lastInstructionPointer, droid.lastRelativeBase, true)
	droid.lastInstructionPointer = lastInstructionPointer
	droid.lastRelativeBase = lastRelativeBase
	return output[0]
}

func NewRepairDroid(intCode []int) (droid RepairDroid) {
	droid = RepairDroid{
		memory:                 make([]int, 3000),
		lastInstructionPointer: 0,
		lastRelativeBase:       0,
	}
	copy(droid.memory, intCode)
	return
}

type ShipMap struct {
	width, height  int
	startX, startY int
	osX, osY       int
	grid           [][]Point
}

type Point struct {
	x, y int
	tile string
}

func (shipMap *ShipMap) AddTile(x, y int, tile string) {
	shipMap.grid[y][x] = Point{
		x:    x,
		y:    y,
		tile: tile,
	}
}

func NewShipMap() (shipMap ShipMap) {
	shipMap = ShipMap{
		width:  100,
		height: 100,
		startX: 50,
		startY: 50,
	}

	shipMap.grid = make([][]Point, shipMap.height)
	for y := 0; y < shipMap.height; y++ {
		shipMap.grid[y] = make([]Point, shipMap.width)
		for x := 0; x < shipMap.width; x++ {
			shipMap.AddTile(x, y, unexplored)
		}
	}

	// droid starts on floor tile
	shipMap.AddTile(shipMap.startX, shipMap.startY, floor)
	return
}

func (shipMap ShipMap) RenderMap() {
	for y := 0; y < shipMap.height; y++ {
		for x := 0; x < shipMap.width; x++ {
			fmt.Print(shipMap.grid[y][x].tile)
		}
		fmt.Print("\n")
	}
}

var unexplored = " "
var wall = "#"
var floor = "."
var droid = "D"

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)

	droid := NewRepairDroid(intCode)
	shipMap := NewShipMap()

	fmt.Printf("Droid: %v, shipMap: %v\n", droid, shipMap)

	return "Answer: "
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
