package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Universe struct {
	universeMap [][]string
	asteroids   []Asteroid
}

func Angle(from Asteroid, to Asteroid) float64 {
	dx := float64(from.x - to.x)
	dy := float64(from.y - to.y)
	theta := math.Atan2(dy, dx)
	theta *= 180 / math.Pi
	return theta
}

func Angle360(from Asteroid, to Asteroid) float64 {
	theta := Angle(from, to)
	if theta < 0 {
		theta = 360 + theta
	}

	theta = math.Round(theta*1000) / 1000
	return theta
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
	Point
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
	var bestLocation Asteroid
	highestViewableAsteroids := 0

	for i := 0; i < len(universe.asteroids)-1; i++ {
		angles := make(map[float64]int)
		for j := 0; j < len(universe.asteroids)-1; j++ {
			if universe.asteroids[i].Point == universe.asteroids[j].Point {
				continue
			}
			angle := Angle360(universe.asteroids[i], universe.asteroids[j])
			angles[angle]++
		}
		// all of the entries in angles with a count of 1 are asteroids that aren't
		// occluded by any others

		viewableAsteroids := 0
		for _, v := range angles {
			if v == 1 {
				viewableAsteroids++
			}
		}
		viewableAsteroids = len(angles)
		universe.asteroids[i].numObservableAsteroids = viewableAsteroids
		if viewableAsteroids > highestViewableAsteroids {
			bestLocation = universe.asteroids[i]
		}
	}
	return bestLocation
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	universe := ParseMap(input)
	bestLocation := FindBestLocation(&universe)
	return "Answer: " + strconv.Itoa(bestLocation.numObservableAsteroids)
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
