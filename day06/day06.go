package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	//"sort"
	//"strconv"
	//"strings"
)

type SpaceObject struct {
	designation string
	parent      *SpaceObject
	child       *SpaceObject
}

func MakeUniverse(orbitMap []string) map[string]*SpaceObject {
	universe := make(map[string]*SpaceObject)

	for _, orbitingSpaceObjects := range orbitMap {
		objs := strings.Split(orbitingSpaceObjects, ")")
		parentDesignation := objs[0]
		childDesignation := objs[1]

		parent, ok := universe[parentDesignation]
		if !ok {
			parent = &SpaceObject{designation: parentDesignation}
			universe[parentDesignation] = parent
		}

		child, ok := universe[childDesignation]
		if !ok {
			child = &SpaceObject{designation: childDesignation}
			universe[childDesignation] = child
		}

		parent.child = child
		child.parent = parent
	}
	return universe
}

func CountOrbits(universe map[string]*SpaceObject) int {
	numOrbits := 0

	for _, spaceObject := range universe {
		numOrbits += countToCOM(spaceObject)
	}

	return numOrbits
}

func countToCOM(object *SpaceObject) int {
	cnt := 0
	lastObj := object

	for {
		if lastObj.parent == nil {
			break
		}
		cnt++
		lastObj = lastObj.parent
	}

	return cnt
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	orbitMap := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	universe := MakeUniverse(orbitMap)
	return "Answer: " + strconv.Itoa(CountOrbits(universe))
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
