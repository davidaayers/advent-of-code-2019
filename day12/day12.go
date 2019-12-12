package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Moon struct {
	moonId     int
	x, y, z    int
	vx, vy, vz int
}

func (moon Moon) String() string {
	return fmt.Sprintf("{[%v]: (%v,%v,%v) (%v,%v,%v)}",
		moon.moonId, moon.x, moon.y, moon.z, moon.vx, moon.vy, moon.vz)
}

func ParseCoordinates(input string) []Moon {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	moons := make([]Moon, len(lines))
	regex := regexp.MustCompile(`<x=(.*), y=(.*), z=(.*)>`)

	for i, line := range lines {
		parts := regex.FindStringSubmatch(line)
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[2])
		z, _ := strconv.Atoi(parts[3])
		moons[i] = Moon{moonId: i, x: x, y: y, z: z}
	}
	return moons
}

func MoveTime(steps int, moons []Moon) {
	for t := 0; t < steps; t++ {
		MoveTimeOneStep(moons)
	}
}

func MoveTimeOneStep(moons []Moon) {
	// apply gravity first
	for m1 := 0; m1 < len(moons); m1++ {
		for m2 := m1 + 1; m2 < len(moons); m2++ {
			if moons[m1].x > moons[m2].x {
				moons[m1].vx--
				moons[m2].vx++
			} else if moons[m1].x < moons[m2].x {
				moons[m1].vx++
				moons[m2].vx--
			}

			if moons[m1].y > moons[m2].y {
				moons[m1].vy--
				moons[m2].vy++
			} else if moons[m1].y < moons[m2].y {
				moons[m1].vy++
				moons[m2].vy--
			}

			if moons[m1].z > moons[m2].z {
				moons[m1].vz--
				moons[m2].vz++
			} else if moons[m1].z < moons[m2].z {
				moons[m1].vz++
				moons[m2].vz--
			}
		}
	}

	// now move the moons
	for m := 0; m < len(moons); m++ {
		moons[m].x += moons[m].vx
		moons[m].y += moons[m].vy
		moons[m].z += moons[m].vz
	}
}

func CalculateTotalEnergy(moons []Moon) int {
	totalEnergy := 0

	for _, m := range moons {
		energy := (Abs(m.x) + Abs(m.y) + Abs(m.z)) * (Abs(m.vx) + Abs(m.vy) + Abs(m.vz))
		totalEnergy += energy
	}

	return totalEnergy
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	moons := ParseCoordinates(input)
	MoveTime(1000, moons)
	energy := CalculateTotalEnergy(moons)
	return "Answer: " + strconv.Itoa(energy)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	moons := ParseCoordinates(input)

	xs := make(map[string]bool)
	ys := make(map[string]bool)
	zs := make(map[string]bool)

	rx, ry, rz := 0, 0, 0

	i := 0
	for {
		MoveTime(1, moons)

		if rx == 0 {
			x := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].x, moons[1].x, moons[2].x, moons[3].x,
				moons[0].vx, moons[1].vx, moons[2].vx, moons[3].vx)
			if xs[x] {
				rx = i
			}
			xs[x] = true
		}

		if ry == 0 {
			y := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].y, moons[1].y, moons[2].y, moons[3].y,
				moons[0].vy, moons[1].vy, moons[2].vy, moons[3].vy)
			if ys[y] {
				ry = i
			}
			ys[y] = true
		}

		if rz == 0 {
			z := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].z, moons[1].z, moons[2].z, moons[3].z,
				moons[0].vz, moons[1].vz, moons[2].vz, moons[3].vz)
			if zs[z] {
				rz = i
			}
			zs[z] = true
		}

		if rx != 0 && ry != 0 && rz != 0 {
			break
		}
		i++
	}

	return "Answer: " + strconv.Itoa(LCM(rx, ry, rz))
}

// LCM code from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
