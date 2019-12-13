package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

/*
0 is an empty tile. No game object appears in this tile.
1 is a wall tile. Walls are indestructible barriers.
2 is a block tile. Blocks can be broken by the ball.
3 is a horizontal paddle tile. The paddle is indestructible.
4 is a ball tile. The ball moves diagonally and bounces off objects.
*/
var tiles = map[int]string{
	0: " ",
	1: "D",
	2: "=",
	3: "_",
	4: "O",
}

type ArcadeCabinet struct {
	memory []int
	screen Screen
}

func (arcadeCabinet *ArcadeCabinet) RunGame() {
	output, _, _, _ := RunIntCode(arcadeCabinet.memory, []int{}, 0, 0, false)

	for idx := 0; idx < len(output); idx += 3 {
		x, y, tileId := output[idx], output[idx+1], output[idx+2]
		arcadeCabinet.screen.AddTile(x, y, tiles[tileId])
	}
}

type Screen struct {
	width, height int
	grid          [][]string
	tileCounts    map[string]int
}

func (screen *Screen) AddTile(x, y int, tile string) {
	screen.grid[y][x] = tile
	screen.tileCounts[tile]++
}

func (screen Screen) RenderScreen() {
	for y := 0; y < screen.height; y++ {
		for x := 0; x < screen.width; x++ {
			fmt.Print(screen.grid[y][x])
		}
		fmt.Print("\n")
	}
}

func NewArcadeCabinet(intCode []int) (arcadeCabinet ArcadeCabinet) {
	arcadeCabinet = ArcadeCabinet{
		memory: make([]int, 3000),
		screen: Screen{
			height:     20,
			width:      44,
			tileCounts: make(map[string]int, 5),
		},
	}

	copy(arcadeCabinet.memory, intCode)

	arcadeCabinet.screen.grid = make([][]string, arcadeCabinet.screen.height)
	for y := 0; y < arcadeCabinet.screen.height; y++ {
		arcadeCabinet.screen.grid[y] = make([]string, arcadeCabinet.screen.width)
		for x := 0; x < arcadeCabinet.screen.width; x++ {
			arcadeCabinet.screen.grid[y][x] = tiles[0]
		}
	}

	return
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	intCode := ParseIntCode(input)
	arcadeCabinet := NewArcadeCabinet(intCode)
	arcadeCabinet.RunGame()
	arcadeCabinet.screen.RenderScreen()
	return "Answer: " + strconv.Itoa(arcadeCabinet.screen.tileCounts[tiles[2]])
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
