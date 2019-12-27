package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
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
				adjacentBugs := len(board.NeighborsForTile(Point{x: x, y: y}, "#"))
				if tile == "." {
					// An empty space becomes infested with a bug if exactly one or two bugs are adjacent to it
					if adjacentBugs == 1 || adjacentBugs == 2 {
						newTile = "#"
					}
				} else if tile == "#" {
					// A bug dies (becoming an empty space) unless there is exactly one bug adjacent to it
					if adjacentBugs != 1 {
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

func CountBugsRecursive(startingBoard Board, numMinutes int) int {
	boards := map[int]Board{0: startingBoard}
	//boards[-1] = NewBoardWithStartingTile(startingBoard.width, startingBoard.height, ".")
	//boards[-1].grid[2][2].tile = "?"
	//boards[1] = NewBoardWithStartingTile(startingBoard.width, startingBoard.height, ".")
	//boards[1].grid[2][2].tile = "?"

	for minute := 0; minute < numMinutes; minute++ {
		newBoards := make(map[int]Board, 0)
		for level, board := range boards {
			nextBoard := NewBoard(board.width, board.height)
			for y := 0; y < board.height; y++ {
				for x := 0; x < board.width; x++ {
					// count the bugs on this level first
					tile := board.grid[y][x].tile
					newTile := tile
					adjacentBugs := len(board.NeighborsForTile(Point{x: x, y: y}, "#"))

					// create or access boards one level above and below, we'll need them for counting
					containingBoard, ok := boards[level-1]
					if !ok {
						containingBoard = NewBoardWithStartingTile(startingBoard.width, startingBoard.height, ".")
						containingBoard.grid[2][2].tile = "?"
						boards[level-1] = containingBoard
						newBoards[level-1] = containingBoard
					}

					containedBoard, ok := boards[level+1]
					if !ok {
						containedBoard = NewBoardWithStartingTile(startingBoard.width, startingBoard.height, ".")
						containedBoard.grid[2][2].tile = "?"
						boards[level+1] = containedBoard
						newBoards[level+1] = containedBoard
					}

					// first, special handling for outer edges, need to look at the board one level higher for adjacent
					if x == 0 && containingBoard.grid[2][1].tile == "#" {
						adjacentBugs++
					}
					if x == board.width-1 && containingBoard.grid[2][3].tile == "#" {
						adjacentBugs++
					}
					if y == 0 && containingBoard.grid[1][2].tile == "#" {
						adjacentBugs++
					}
					if y == board.width-1 && containingBoard.grid[3][2].tile == "#" {
						adjacentBugs++
					}

					// special handling for inner edges, need to look a the board one level lower for adjacent
					//containedBoard := boards[level+1]
					if x == 2 && y == 1 {
						// whole top row of contained
						for x2 := 0; x2 < containedBoard.width; x2++ {
							if containedBoard.grid[0][x2].tile == "#" {
								adjacentBugs++
							}
						}
					}
					if x == 2 && y == 3 {
						// whole bottom row of contained
						for x2 := 0; x2 < containedBoard.width; x2++ {
							if containedBoard.grid[containingBoard.height-1][x2].tile == "#" {
								adjacentBugs++
							}
						}
					}

					if x == 1 && y == 2 {
						// whole left side
						for y2 := 0; y2 < containedBoard.width; y2++ {
							if containedBoard.grid[y2][0].tile == "#" {
								adjacentBugs++
							}
						}
					}
					if x == 3 && y == 2 {
						// whole left side
						for y2 := 0; y2 < containedBoard.width; y2++ {
							if containedBoard.grid[y2][containedBoard.width-1].tile == "#" {
								adjacentBugs++
							}
						}
					}

					// handle the game of life rules
					if tile == "." {
						// An empty space becomes infested with a bug if exactly one or two bugs are adjacent to it
						if adjacentBugs == 1 || adjacentBugs == 2 {
							newTile = "#"
						}
					} else if tile == "#" {
						// A bug dies (becoming an empty space) unless there is exactly one bug adjacent to it
						if adjacentBugs != 1 {
							newTile = "."
						}
					}
					nextBoard.grid[y][x].tile = newTile
				}
			}
			newBoards[level] = nextBoard
		}
		boards = newBoards
	}

	keys := make([]int, len(boards))

	i := 0
	for k := range boards {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	bugCnt := 0
	for _, k := range keys {
		boardBugCnt := 0
		board := boards[k]
		for y := 0; y < board.height; y++ {
			for x := 0; x < board.width; x++ {
				if board.grid[y][x].tile == "#" {
					boardBugCnt++
				}
			}
		}
		if boardBugCnt > 0 {
			fmt.Printf("Level: %v\n", k)
			board.RenderBoard()
			bugCnt += boardBugCnt
		}
	}
	return bugCnt
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	board := BuildBoard(input)
	board.grid[2][2].tile = "?"
	numBugs := CountBugsRecursive(board, 200)
	return "Answer: " + strconv.Itoa(numBugs)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
