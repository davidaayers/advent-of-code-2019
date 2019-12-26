package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type DonutMaze struct {
	board          Board
	startX, startY int
	endX, endY     int
	portals        map[string]Portal
}

type Portal struct {
	label  string
	point1 Point
	point2 Point
}

func (portal Portal) ContainsPoint(point Point) bool {
	if portal.point1 == point {
		return true
	}
	if portal.point2 == point {
		return true
	}
	return false
}

func BuildDonutMaze(input string) DonutMaze {
	maze := DonutMaze{
		portals: make(map[string]Portal, 0),
	}

	letterMatch := regexp.MustCompile(`[A-Z]+`)
	maze.board = ParseBoard(input, func(point Point) {}, true)

	for y := 0; y < maze.board.height; y++ {
		for x := 0; x < maze.board.width; x++ {
			firstPoint := maze.board.grid[y][x]
			if letterMatch.MatchString(firstPoint.tile) {
				// this point only matters if there is another letter to the immediate right or bottom
				secondPoint, ok := maze.board.PointInDir(firstPoint, east)
				if !ok || !letterMatch.MatchString(secondPoint.tile) {
					secondPoint, ok = maze.board.PointInDir(firstPoint, south)
					if !ok || !letterMatch.MatchString(secondPoint.tile) {
						continue
					}
				}
				label := firstPoint.tile + secondPoint.tile
				portalXgress := FindXgress(maze, firstPoint, secondPoint)
				if existingPortal, ok := maze.portals[label]; ok {
					delete(maze.portals, label)
					maze.portals[label] = Portal{label: label, point1: existingPortal.point1, point2: portalXgress}
				} else {
					maze.portals[label] = Portal{label: label, point1: portalXgress}
				}
			}
		}
	}

	// grab the start and end
	start := maze.portals["AA"]
	maze.startX, maze.startY = start.point1.x, start.point1.y
	delete(maze.portals, "AA")
	end := maze.portals["ZZ"]
	maze.endX, maze.endY = end.point1.x, end.point1.y
	delete(maze.portals, "ZZ")

	fmt.Printf("Portals: %v\n", maze.portals)

	return maze
}

func FindXgress(maze DonutMaze, point1 Point, point2 Point) Point {
	neighbors := maze.board.WalkableNeighbors(point1)
	if len(neighbors) == 1 {
		return neighbors[0]
	}
	neighbors = maze.board.WalkableNeighbors(point2)
	return neighbors[0]
}

func FindLeastSteps(maze DonutMaze) (steps int) {
	var pathExpander = func(board Board, parent Point) []Point {
		neighbors := board.WalkableNeighbors(parent)

		// see if point is a portal spot, if so, add the other end of the portal as a walkable neighbor
		for _, portal := range maze.portals {
			if portal.point1 == parent {
				neighbors = append(neighbors, portal.point2)
			}
			if portal.point2 == parent {
				neighbors = append(neighbors, portal.point1)
			}
		}
		return neighbors
	}

	start := maze.board.grid[maze.startY][maze.startX]
	end := maze.board.grid[maze.endY][maze.endX]
	path, ok := maze.board.Path(start, end, pathExpander)

	if !ok {
		panic("Path not found")
	}

	fmt.Printf("Path: %v\n", path)

	return len(path)
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	maze := BuildDonutMaze(input)
	steps := FindLeastSteps(maze)
	return "Answer: " + strconv.Itoa(steps)
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
