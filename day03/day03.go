package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var dirs = map[string][]int{
	"U": {0,-1},
	"D": {0,1},
	"L": {-1,0},
	"R": {1,0},
}

var zero = Point{0,0}

type Point struct {
	x,y int
}

func (p Point) ManhattanDistanceTo(other Point) int {
	return Abs(p.x - other.x) + Abs(p.y - other.y) 
}

type Wire struct {
	points []Point
} 

func (w *Wire) AddPoint(p Point) {
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
		if w.ContainsPoint(p) && p != zero {
			intersections = append(intersections,p)
		}
	}

	return intersections
}

func NewWire() Wire {
	return Wire{make([]Point,0)}
}

func MakeWire(wirePath []string) Wire {
	wire := NewWire()

	lastPoint := zero

	wire.AddPoint(lastPoint)

	for _,i := range(wirePath) {
		dir := i[:1]
		steps,_ := strconv.Atoi(i[1:])
		for cnt := 0; cnt != steps; cnt ++ {
			nextPoint := lastPoint
			nextPoint.x += dirs[dir][0]
			nextPoint.y += dirs[dir][1]
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
	wire2 := MakeWire(wire2Path)

	intersections := wire1.Intersections(wire2)
	sort.Slice(intersections, func(i, j int) bool { return zero.ManhattanDistanceTo(intersections[i]) < zero.ManhattanDistanceTo(intersections[j]) })

	return zero.ManhattanDistanceTo(intersections[0])
}

func StepsToPoint(wire Wire, point Point) int {
	for i,p := range(wire.points) {
		if p == point {
			return i
		}
	}

	return 0
}

func FindClosestIntersectionDistanceBySteps(wire1Path,wire2Path []string) int {
	wire1 := MakeWire(wire1Path)
	wire2 := MakeWire(wire2Path)

	intersections := wire1.Intersections(wire2)

	stepCounts := make([]int,0)

	// now go through each intersection and find the one with the least number of steps
	for _,intersection := range(intersections) {
		stepsWire1 := StepsToPoint(wire1,intersection)
		stepsWire2 := StepsToPoint(wire2,intersection)
		stepCounts = append(stepCounts,stepsWire1+stepsWire2)
	}

	sort.Ints(stepCounts)

	return stepCounts[0]
}


// Part1 Part 1 of puzzle
func Part1(input string) string {
	wires := strings.Split(strings.ReplaceAll(input,"\r\n","\n"),"\n")
	answer := FindClosestIntersectionDistance(strings.Split(wires[0],","),strings.Split(wires[1],","))
	return "Answer: " + strconv.Itoa(answer)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	wires := strings.Split(strings.ReplaceAll(input,"\r\n","\n"),"\n")
	answer := FindClosestIntersectionDistanceBySteps(strings.Split(wires[0],","),strings.Split(wires[1],","))
	return "Answer: " + strconv.Itoa(answer)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
