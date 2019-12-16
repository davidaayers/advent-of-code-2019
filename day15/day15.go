package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type RepairDroid struct {
	memory                                   []int
	lastInstructionPointer, lastRelativeBase int
	x, y                                     int
}

func (droid *RepairDroid) Move(direction int) (result int) {
	output, lastInstructionPointer, lastRelativeBase, _ :=
		RunIntCode(droid.memory, []int{direction}, droid.lastInstructionPointer, droid.lastRelativeBase, true)
	droid.lastInstructionPointer = lastInstructionPointer
	droid.lastRelativeBase = lastRelativeBase
	return output[0]
}

func NewRepairDroid(intCode []int, x int, y int) (droid RepairDroid) {
	droid = RepairDroid{
		memory:                 make([]int, 3000),
		lastInstructionPointer: 0,
		lastRelativeBase:       0,
		x:                      x,
		y:                      y,
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

func (p Point) ManhattanDistanceTo(other Point) int {
	return Abs(p.x-other.x) + Abs(p.y-other.y)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type PathNode struct {
	point   Point
	parent  *PathNode
	f, g, h int
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
		width:  44,
		height: 44,
		startX: 22,
		startY: 22,
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

func (shipMap ShipMap) RenderMap(repairDroid RepairDroid) {
	for y := 0; y < shipMap.height; y++ {
		for x := 0; x < shipMap.width; x++ {
			if repairDroid.x == x && repairDroid.y == y {
				fmt.Print(droid)
			} else if shipMap.osX != 0 && shipMap.osY != 0 && shipMap.osX == x && shipMap.osY == y {
				fmt.Print(oxygenGenerator)
			} else {
				fmt.Print(shipMap.grid[y][x].tile)
			}
		}
		fmt.Print("\n")
	}
}

func (shipMap ShipMap) Path(from, to Point) (path []Point, ok bool) {
	path = make([]Point, 0)

	startNode := PathNode{
		point:  from,
		parent: nil,
		f:      0,
		g:      0,
		h:      0,
	}
	openList := []PathNode{startNode}
	closedList := make([]PathNode, 0)

	for len(openList) != 0 {
		sort.Slice(openList, func(i, j int) bool {
			return openList[i].f < openList[j].f
		})

		currentNode := openList[0]
		openList = openList[1:]
		closedList = append(closedList, currentNode)

		if currentNode.point.x == to.x && currentNode.point.y == to.y {
			// return our path
			for currentNode.point != from {
				path = append(path, currentNode.point)
				currentNode = *currentNode.parent
			}
			return path, true
		}

		children := make([]PathNode, 0)
		for _, dir := range dirs {
			pointInDir := shipMap.grid[currentNode.point.y+dir.dy][currentNode.point.x+dir.dx]
			if pointInDir.tile == floor {
				children = append(children, PathNode{
					point:  pointInDir,
					parent: &currentNode,
				})
			}
		}

		for _, child := range children {
			if contains(closedList, child) {
				continue
			}
			child.g = currentNode.g + 1
			child.h = child.point.ManhattanDistanceTo(to)
			child.f = child.g + child.h

			found, ok := find(openList, child)
			if ok && child.g > found.g {
				continue
			}

			openList = append(openList, child)
		}

	}

	return path, false
}

func find(nodes []PathNode, n PathNode) (found PathNode, ok bool) {
	for _, a := range nodes {
		if a.point == n.point {
			return n, true
		}
	}
	return PathNode{}, false
}

func contains(s []PathNode, e PathNode) bool {
	for _, a := range s {
		if a.point == e.point {
			return true
		}
	}
	return false
}

var unexplored = " "
var wall = "#"
var floor = "."
var droid = "D"
var oxygenGenerator = "O"

type Dir struct {
	instruction int
	dx, dy      int
}

var north = Dir{dx: 0, dy: -1, instruction: 1}
var south = Dir{dx: 0, dy: 1, instruction: 2}
var east = Dir{dx: 1, dy: 0, instruction: 3}
var west = Dir{dx: -1, dy: 0, instruction: 4}

var dirs = []Dir{north, south, east, west}

// Part1 Part 1 of puzzle
func Part1(input string) (string, ShipMap) {
	intCode := ParseIntCode(input)

	shipMap := NewShipMap()
	droid := NewRepairDroid(intCode, shipMap.startX, shipMap.startY)

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	steps := 0
	for {
		steps++

		dir := r1.Intn(4)
		result := droid.Move(dirs[dir].instruction)

		switch result {
		case 0:
			// `0`: The repair droid hit a wall. Its position has not changed.
			wallX, wallY := droid.x+dirs[dir].dx, droid.y+dirs[dir].dy
			shipMap.AddTile(wallX, wallY, wall)
		case 1:
			// `1`: The repair droid has moved one step in the requested direction.
			droid.x += dirs[dir].dx
			droid.y += dirs[dir].dy
			shipMap.AddTile(droid.x, droid.y, floor)
		case 2:
			// `2`: The repair droid has moved one step in the requested direction; its new position is the
			// location of the oxygen system.
			droid.x += dirs[dir].dx
			droid.y += dirs[dir].dy
			shipMap.AddTile(droid.x, droid.y, floor)
			if shipMap.osX == 0 && shipMap.osY == 0 {
				shipMap.osX = droid.x
				shipMap.osY = droid.y
				//fmt.Printf("Found oxygen system at: %v,%v\n", shipMap.osX, shipMap.osY)
			}
		}

		if steps%100000 == 0 {
			//fmt.Printf("Step %v\n", steps)
		}

		if steps > 2000000 {
			break
		}

	}

	shipMap.RenderMap(droid)

	path, ok := shipMap.Path(Point{x: shipMap.startX, y: shipMap.startY}, Point{x: shipMap.osX, y: shipMap.osY})

	if !ok {
		panic("No path found from start to oxygen system")
	}

	return "Answer: " + strconv.Itoa(len(path)), shipMap
}

// Part2 Part2 of puzzle
func Part2(input ShipMap) string {
	return "Answer: "
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	part1, shipMap := Part1(string(bytes))
	fmt.Println("Part 1: " + part1)
	fmt.Println("Part 2: " + Part2(shipMap))
}
