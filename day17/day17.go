package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ScaffoldMap struct {
	width, height  int
	startX, startY int
	grid           [][]Point
}

func (scaffoldMap ScaffoldMap) RenderMap() {
	for y := 0; y < scaffoldMap.height; y++ {
		for x := 0; x < scaffoldMap.width; x++ {
			fmt.Print(scaffoldMap.grid[y][x].tile)
		}
		fmt.Print("\n")
	}
}

func (scaffoldMap ScaffoldMap) IsIntersection(p Point) bool {
	// must have # in all 4 dirs
	for _, dir := range dirs {
		cx := p.x + dir.dx
		cy := p.y + dir.dy
		if cx > scaffoldMap.width-1 || cx < 0 || cy > scaffoldMap.height-1 || cy < 0 {
			return false
		}

		if scaffoldMap.grid[cy][cx].tile != "#" {
			return false
		}
	}
	return true
}

func NewScaffoldMap(width, height int) ScaffoldMap {
	scaffoldMap := ScaffoldMap{
		width:  width,
		height: height,
	}

	scaffoldMap.grid = make([][]Point, height)
	for y := 0; y < height; y++ {
		scaffoldMap.grid[y] = make([]Point, width)
	}

	return scaffoldMap
}

type Point struct {
	x, y int
	tile string
}

type Dir struct {
	dx, dy int
}

var north = Dir{dx: 0, dy: -1}
var south = Dir{dx: 0, dy: 1}
var east = Dir{dx: 1, dy: 0}
var west = Dir{dx: -1, dy: 0}

var dirs = []Dir{north, south, east, west}

func ParseMap(toParse string) ScaffoldMap {
	mapLines := strings.Split(strings.ReplaceAll(toParse, "\r\n", "\n"), "\n")
	scaffoldMap := NewScaffoldMap(len(mapLines[0]), len(mapLines))

	for y, line := range mapLines {
		for x, ch := range line {
			scaffoldMap.grid[y][x] = Point{x: x, y: y, tile: string(ch)}
		}
	}
	scaffoldMap.RenderMap()
	return scaffoldMap
}

func CheckAlignment(scaffoldMap ScaffoldMap) int {
	alignment := 0
	for y := 0; y < scaffoldMap.height; y++ {
		for x := 0; x < scaffoldMap.width; x++ {
			checkPoint := scaffoldMap.grid[y][x]
			if checkPoint.tile == "#" {
				isIntersection := scaffoldMap.IsIntersection(checkPoint)
				if isIntersection {
					alignment += x * y
				}
			}
		}
	}

	return alignment
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)
	memory := make([]int, 5000)
	copy(memory, intCode)

	output, _, _, _ := RunIntCode(memory, []int{}, 0, 0, false)

	str := ""
	for _, i := range output {
		str += string(i)
	}

	alignment := CheckAlignment(ParseMap(str))

	return "Answer: " + strconv.Itoa(alignment)
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
