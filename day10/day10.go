package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Universe struct {
	universeMap [][]string
	asteroids   []Asteroid
}

func getLine(from, to Point) (points []Point) {
	x1, y1 := from.x, from.y
	x2, y2 := to.x, to.y

	isSteep := Abs(y2-y1) > Abs(x2-x1)
	if isSteep {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	reversed := false
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
		reversed = true
	}

	deltaX := x2 - x1
	deltaY := Abs(y2 - y1)
	err := deltaX / 2
	y := y1
	var ystep int

	if y1 < y2 {
		ystep = 1
	} else {
		ystep = -1
	}

	for x := x1; x < x2+1; x++ {
		if isSteep {
			points = append(points, Point{y, x})
		} else {
			points = append(points, Point{x, y})
		}
		err -= deltaY
		if err < 0 {
			y += ystep
			err += deltaX
		}
	}

	if reversed {
		//Reverse the slice
		for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
			points[i], points[j] = points[j], points[i]
		}
	}

	return
}

func (universe Universe) CanObserve(from, to Asteroid) bool {
	if from == to {
		return false
	}
	line := getLine(from.point, to.point)
	fmt.Printf("Line from %v to %v: %v ", from, to, line)
	line = line[1 : len(line)-1]
	fmt.Printf("Chomped: %v\n", line)

	for _, checkPoint := range line {
		if universe.universeMap[checkPoint.y][checkPoint.x] == "#" {
			return false
		}
	}

	return true
}

func (universe Universe) CanObserve2(from, to Asteroid) bool {
	// use Bresenham's to draw a line from -> to. If we hit
	// another asteroid on the way, we can't see to's position
	// implementation below inspired by
	// http://www.roguebasin.com/index.php?title=Bresenham%27s_Line_Algorithm#JavaScript
	x1, y1 := from.point.x, from.point.y
	x2, y2 := to.point.x, to.point.y

	isSteep := Abs(y2-y1) > Abs(x2-x1)
	if isSteep {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	deltaX := x2 - x1
	deltaY := Abs(y2 - y1)
	err := deltaX / 2
	y := y1
	var ystep int

	if y1 < y2 {
		ystep = 1
	} else {
		ystep = -1
	}

	for x := x1; x < x2+1; x++ {
		if isSteep {
			if universe.universeMap[x][y] == "#" {
				return false
			}
		} else {
			if universe.universeMap[y][x] == "#" {
				return false
			}
		}
		err -= deltaY
		if err < 0 {
			y += ystep
			err += deltaX
		}
	}

	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x, y int
}

type Asteroid struct {
	point                  Point
	numObservableAsteroids int
}

func ParseMap(input string) Universe {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	universeMap := make([][]string, len(lines))
	universe := Universe{universeMap: universeMap}
	for y, r := range lines {
		row := strings.Split(r, "")
		universeMap[y] = make([]string, len(row))
		for x, s := range row {
			universeMap[y][x] = s
			if s == "#" {
				asteroid := Asteroid{Point{x: x, y: y}, 0}
				universe.asteroids = append(universe.asteroids, asteroid)
			}
		}
	}
	//PrintMap(universe)
	return universe
}

func PrintMap(universe Universe) {
	for _, row := range universe.universeMap {
		for _, column := range row {
			fmt.Printf("%v", column)
		}
		fmt.Printf("\n")
	}
}

func FindBestLocation(universe *Universe) Asteroid {
	highestNumObservableAsteroids := 0
	var bestLocation Asteroid

	for i := 0; i < len(universe.asteroids)-1; i++ {
		for j := 0; j < len(universe.asteroids)-1; j++ {
			if universe.CanObserve(universe.asteroids[i], universe.asteroids[j]) {
				universe.asteroids[i].numObservableAsteroids++
				if universe.asteroids[i].numObservableAsteroids > highestNumObservableAsteroids {
					highestNumObservableAsteroids = universe.asteroids[i].numObservableAsteroids
					bestLocation = universe.asteroids[i]
				}
			}
		}
	}
	return bestLocation
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
