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

func (dir Dir) Turn(turn string) Dir {
	if dir == north {
		if turn == "R" {
			return east
		}
		return west
	}
	if dir == south {
		if turn == "R" {
			return west
		}
		return east
	}
	if dir == east {
		if turn == "R" {
			return south
		}
		return north
	}
	// west
	if turn == "R" {
		return north
	}
	return south
}

var north = Dir{dx: 0, dy: -1}
var south = Dir{dx: 0, dy: 1}
var east = Dir{dx: 1, dy: 0}
var west = Dir{dx: -1, dy: 0}

var dirs = []Dir{north, south, east, west}

func ParseMap(toParse string) ScaffoldMap {
	mapLines := strings.Split(strings.TrimSpace(strings.ReplaceAll(toParse, "\r\n", "\n")), "\n")
	scaffoldMap := NewScaffoldMap(len(mapLines[0]), len(mapLines))

	for y, line := range mapLines {
		for x, ch := range line {
			scaffoldMap.grid[y][x] = Point{x: x, y: y, tile: string(ch)}
			if string(ch) == "^" {
				scaffoldMap.startX = x
				scaffoldMap.startY = y
			}
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

	scaffoldMap := ParseMap(str)
	alignment := CheckAlignment(scaffoldMap)

	return "Answer: " + strconv.Itoa(alignment)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	intCode := ParseIntCode(input)
	memory := make([]int, 5000)
	copy(memory, intCode)

	output, _, _, _ := RunIntCode(memory, []int{}, 0, 0, false)

	str := ""
	for _, i := range output {
		str += string(i)
	}

	scaffoldMap := ParseMap(str)

	movements := make([]string, 0)
	facing := north
	robotX, robotY := scaffoldMap.startX, scaffoldMap.startY
	turn := "R"
	steps := 0

	for {

		//fmt.Printf("Steps: %v, Facing: %v, x:%v y: %v\n", steps, facing, robotX, robotY)

		cx := robotX + facing.dx
		cy := robotY + facing.dy

		if cx < 0 || cx > scaffoldMap.width-1 || cy < 0 || cy > scaffoldMap.height-1 ||
			scaffoldMap.grid[cy][cx].tile == "." || scaffoldMap.grid[cy][cx].tile == "" {
			if steps != 0 {
				movements = append(movements, turn)
				movements = append(movements, strconv.Itoa(steps))
			}
			steps = 0

			var ok bool
			facing, turn, ok = determineTurnDirection(scaffoldMap, facing, robotX, robotY)

			// we we couldn't turn, we hit the wall, and we're done
			if !ok {
				break
			}
		} else if scaffoldMap.grid[cy][cx].tile == "#" {
			robotX += facing.dx
			robotY += facing.dy
			steps++
		}
	}

	fmt.Printf("%v: %v\n", movements, len(movements)/2)

	memory2 := make([]int, 5000)
	copy(memory2, intCode)
	memory2[0] = 2

	intCodeInput := []int{}

	// used this regex (found on reditt):
	// ^(.{1,20})\1*(.{1,20})(?:\1|\2)*(.{1,20})(?:\1|\2|\3)*$
	// to solve manually; I'll come back and try to do this programatically.

	// R,8,L,10,R,8,R,12,R,8,L,8,L,12,R,8,L,10,R,8,L,12,L,10,L,8,R,8,L,10,R,8,R,12,R,8,L,8,L,12,L,12,L,10,L,8,L,12,L,10,L,8,R,8,L,10,R,8,R,12,R,8,L,8,L,12
	order := "A,B,A,C,A,B,C,C,A,B\n"
	a := "R,8,L,10,R,8\n"
	b := "R,12,R,8,L,8,L,12\n"
	c := "L,12,L,10,L,8\n"

	intCodeInput = append(intCodeInput, toIntSlice(order)...)
	fmt.Printf("%v\n", intCodeInput)
	intCodeInput = append(intCodeInput, toIntSlice(a)...)
	fmt.Printf("%v\n", intCodeInput)
	intCodeInput = append(intCodeInput, toIntSlice(b)...)
	fmt.Printf("%v\n", intCodeInput)
	intCodeInput = append(intCodeInput, toIntSlice(c)...)
	fmt.Printf("%v\n", intCodeInput)

	// add the interactive flag
	intCodeInput = append(intCodeInput, toIntSlice("n\n")...)

	fmt.Printf("%v\n", intCodeInput)

	output2, _, _, _ := RunIntCode(memory2, intCodeInput, 0, 0, false)

	fmt.Printf("%v\n", output2)

	str = ""
	for _, i := range output2 {
		if i > 46 {
			fmt.Printf("***")
		}
		r := string(rune(i))
		fmt.Printf("%v", r)
		str += string(i)
	}

	fmt.Println(str)

	scaffoldMap = ParseMap(str)
	scaffoldMap.RenderMap()

	return "Answer: "
}

func toIntSlice(str string) []int {
	sl := make([]int, 0)
	for _, r := range str {
		sl = append(sl, int(r))
	}
	return sl
}

func determineTurnDirection(scaffoldMap ScaffoldMap, facing Dir, x int, y int) (Dir, string, bool) {
	// now figure out if it's a right or left turn
	tryTurn := []string{"R", "L"}

	var checkTurn Dir
	for _, t := range tryTurn {
		tryDir := facing.Turn(t)

		tx := x + tryDir.dx
		ty := y + tryDir.dy

		if tx < 0 || tx > scaffoldMap.width-1 || ty < 0 || ty > scaffoldMap.height-1 ||
			scaffoldMap.grid[ty][tx].tile != "#" {
			continue
		}

		return tryDir, t, true
	}
	return checkTurn, "", false
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	//fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
