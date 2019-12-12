package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	//"sort"
	//"strconv"
	//"strings"
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
		MoveTimeOneStep(moons, t)
	}
}

func MoveTimeOneStep(moons []Moon, step int) {
	// apply gravity first
	for m1 := 0; m1 < len(moons); m1++ {
		for m2 := m1 + 1; m2 < len(moons); m2++ {
			//fmt.Printf("Comparing %v to %v\n", m1, m2)

			/*
				 For example, if m1 has an x position of 3, and m2 has a x position of 5,
				 then m1's x velocity changes by +1 (because 5 > 3) and m2's x velocity
				 changes by -1 (because 3 < 5). However, if the positions on a given axis are the same,
				 the velocity on that axis does not change for that pair of moons.

				m1 start x = -1
				m2 start x = 2

				-1 > 2

			*/
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

	//fmt.Printf("Moons (%v): %v\n", step, moons)
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
	return "Answer: "
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
