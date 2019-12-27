package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

func BuildBoard(board string) Board {
	return ParseBoard(board, func(point Point) {}, true)
}

func CalculateBiodiversityRating(board Board) int {
	fmt.Printf("Board: %v\n", board)

	boards := []Board{board}

	minutes := 0
	for {
		minutes++

		// create a new empty board to populate
		nextBoard := NewBoard(board.width, board.height)

		for y := 0; y < board.height; y++ {
			for x := 0; x < board.width; x++ {
				tile := board.grid[y][x].tile
				newTile := tile
				if tile == "." {
					// An empty space becomes infested with a bug if exactly one or two bugs are adjacent to it
					adjacentBugs := board.NeighborsForTile(Point{x: x, y: y}, "#")
					if len(adjacentBugs) == 1 || len(adjacentBugs) == 2 {
						newTile = "#"
					}
				} else if tile == "#" {
					// A bug dies (becoming an empty space) unless there is exactly one bug adjacent to it
					adjacentBugs := board.NeighborsForTile(Point{x: x, y: y}, "#")
					if len(adjacentBugs) != 1 {
						newTile = "."
					}
				}
				nextBoard.grid[y][x].tile = newTile
			}
		}

		fmt.Printf("Minute %v:\n", minutes)
		nextBoard.RenderBoard()

		// add next board to our list of boards
		board = nextBoard

		// finally, see if this board matches any previous board
		if IsRepeatBoard(boards, board) {
			fmt.Printf("Found a repeat sequence at minute %v\n", minutes)
			break
		}
		boards = append(boards, nextBoard)
	}

	var cnt float64 = 0
	var squareRating float64 = 1
	var overallRating float64 = 0
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			cnt++
			if board.grid[y][x].tile == "#" {
				overallRating += squareRating
			}
			squareRating = math.Pow(2, cnt)
		}
	}

	return int(overallRating)
}

func IsRepeatBoard(boards []Board, board Board) bool {
	for _, b := range boards {
		if b.Matches(board) {
			return true
		}
	}
	return false
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	board := BuildBoard(input)
	rating := CalculateBiodiversityRating(board)

	return "Answer: " + strconv.Itoa(rating)
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
