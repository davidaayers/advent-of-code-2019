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
	score  int
}

func (arcadeCabinet *ArcadeCabinet) RunGame(freePlay, renderScreen bool) {
	if !freePlay {
		arcadeCabinet.RunGameRenderScreenOnly(renderScreen)
		return
	}

	// free games!
	arcadeCabinet.memory[0] = 2

	lastInstructionPointer, lastRelativeBase := 0, 0
	terminated := false
	var output []int
	var instr []int
	var input []int
	ballX, paddleX := 0, 0
	var gameStarted = false

	for {
		output, lastInstructionPointer, lastRelativeBase, terminated =
			RunIntCode(arcadeCabinet.memory, input, lastInstructionPointer, lastRelativeBase, true)
		if terminated {
			break
		}

		// in freeplay mode, output will be one character at a time, so we need to build up our instructions
		instr = append(instr, output[0])
		if len(instr) == 3 {
			x, y, tileId := instr[0], instr[1], instr[2]
			if x == -1 {
				arcadeCabinet.score = tileId
			} else {
				arcadeCabinet.screen.AddTile(x, y, tiles[tileId])

				// remember where the ball and paddle are
				if tileId == 4 {
					ballX = x
				}

				if tileId == 3 {
					paddleX = x
				}
			}

			instr = nil
		}

		if renderScreen {
			arcadeCabinet.RenderScreen()
		}

		if !gameStarted && lastInstructionPointer == 73 {
			gameStarted = true
		}

		if gameStarted {
			input = nil
			// adjust the paddle position to move it toward the ball
			/*
				If the joystick is in the neutral position, provide 0.
				If the joystick is tilted to the left, provide -1.
				If the joystick is tilted to the right, provide 1.
			*/

			joystickCommand := 0
			if ballX > paddleX {
				joystickCommand = 1
			} else if ballX < paddleX {
				joystickCommand = -1
			}
			input = append(input, joystickCommand)
		}

	}
}

func (arcadeCabinet *ArcadeCabinet) RunGameRenderScreenOnly(renderScreen bool) {
	output, _, _, _ := RunIntCode(arcadeCabinet.memory, []int{}, 0, 0, false)

	for idx := 0; idx < len(output); idx += 3 {
		x, y, tileId := output[idx], output[idx+1], output[idx+2]
		arcadeCabinet.screen.AddTile(x, y, tiles[tileId])
	}
	if renderScreen {
		arcadeCabinet.RenderScreen()
	}
}

func (arcadeCabinet ArcadeCabinet) RenderScreen() {
	fmt.Printf("-------------- Score: %05d ----------------\n", arcadeCabinet.score)
	arcadeCabinet.screen.RenderScreen()
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
	arcadeCabinet.RunGame(false, false)
	return "Answer: " + strconv.Itoa(arcadeCabinet.screen.tileCounts[tiles[2]])
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	intCode := ParseIntCode(input)
	arcadeCabinet := NewArcadeCabinet(intCode)
	arcadeCabinet.RunGame(true, false)
	return "Answer: " + strconv.Itoa(arcadeCabinet.score)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
