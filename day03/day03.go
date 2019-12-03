package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"strconv"
	//"strings"
)

type Point struct {
	x,y int
}

func (p Point) ManhattanDistanceTo(other Point) int {
	return Abs(p.x - other.x) + Abs(p.y - other.y) 
}

type Wire struct {
	points []Point
} 

func (w Wire) AddPoint(p Point) {
	w.points = append(w.points,p)
}

func (w Wire) ContainsPoint(p Point) bool {
	for _,other := range w.points {
		if p == other {
			return true
		}
	}
	return false
}

func (w Wire) Intersections(other Wire) []Point {
	intersections := make([]Point,0)

	for _,p := range(other.points) {
		if w.ContainsPoint(p) {
			intersections = append(intersections,p)
		}
	}

	return intersections
}

func NewWire() Wire {
	return Wire{make([]Point,0)}
}

func MakeWire(wirePath []string) Wire {
	wire := Wire{}

	lastPoint := Point{0,0}

	wire.AddPoint(lastPoint)

	for _,i := range(wirePath) {
		dir := i[:1]
		steps,_ := strconv.Atoi(i[1:])
		xd := 0
		yd := 0

		switch (string(dir)) {
		case "U":
			yd = -1
		case "D":
			yd = 1
		case "L":
			xd = -1
		case "R":
			xd = 1
		}

		fmt.Printf("i: %v Dir: %v xd: %v yd: %v steps: %v\n", i,dir,xd,yd,steps)

		for cnt := 0; cnt == steps; cnt ++ {
			nextPoint := lastPoint
			nextPoint.x += xd
			nextPoint.y += yd
			lastPoint = nextPoint
			wire.AddPoint(nextPoint)
		}
	}

	return wire
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func FindClosestIntersectionDistance(wire1Path,wire2Path []string) int {

	wire1 := MakeWire(wire1Path)
	fmt.Printf("Wire1: %v\n", wire1)

	wire2 := MakeWire(wire2Path)
	fmt.Printf("Wire2: %v\n", wire2)

	intersections := wire1.Intersections(wire2)

	fmt.Printf("Intersections: %v\n", intersections)

	return 0
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
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
