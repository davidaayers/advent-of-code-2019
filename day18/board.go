package main

import (
	"fmt"
	"sort"
	"strings"
)

type Board struct {
	width, height int
	grid          [][]Point
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

func (board *Board) AddTile(x, y int, tile string) {
	board.grid[y][x] = Point{
		x:    x,
		y:    y,
		tile: tile,
	}
}

func NewBoard(width, height int) (board Board) {
	board = Board{
		width:  width,
		height: height,
	}

	board.grid = make([][]Point, board.height)
	for y := 0; y < board.height; y++ {
		board.grid[y] = make([]Point, board.width)
		for x := 0; x < board.width; x++ {
			board.AddTile(x, y, unexplored)
		}
	}

	return
}

func (board Board) RenderBoard() {
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			fmt.Print(board.grid[y][x].tile)
		}
		fmt.Print("\n")
	}
}

func (board Board) Neighbors(from Point) []Point {
	neighbors := make([]Point, 0)
	for _, dir := range dirs {
		pointInDir := board.grid[from.y+dir.dy][from.x+dir.dx]
		if pointInDir.tile == floor {
			neighbors = append(neighbors, pointInDir)
		}
	}
	return neighbors
}

func (board Board) Path(from, to Point) (path []Point, ok bool) {
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

		for _, pointInDir := range board.Neighbors(currentNode.point) {
			children = append(children, PathNode{
				point:  pointInDir,
				parent: &currentNode,
			})
		}

		for _, child := range children {
			if containsPathNode(closedList, child) {
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

func containsPathNode(s []PathNode, e PathNode) bool {
	for _, a := range s {
		if a.point == e.point {
			return true
		}
	}
	return false
}

func containsPoint(s []Point, e Point) bool {
	for _, a := range s {
		if a.x == e.x && a.y == e.y {
			return true
		}
	}
	return false
}

var unexplored = " "
var wall = "#"
var floor = "."

type Dir struct {
	dx, dy int
}

var north = Dir{dx: 0, dy: -1}
var south = Dir{dx: 0, dy: 1}
var east = Dir{dx: 1, dy: 0}
var west = Dir{dx: -1, dy: 0}

var dirs = []Dir{north, south, east, west}

func ParseBoard(input string, callback func(point Point), printBoard bool) Board {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	board := NewBoard(len(lines[0]), len(lines))
	for y, r := range lines {
		row := strings.Split(r, "")
		for x, s := range row {
			point := Point{x: x, y: y, tile: s}
			board.grid[y][x] = point
			callback(point)
		}
	}
	if printBoard {
		board.RenderBoard()
	}
	return board
}
