package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
		numOrbits += countToCOM(universe, spaceObject)
	}

	return numOrbits
}

func countToCOM(universe map[string]*SpaceObject, object *SpaceObject) int {
	course := MapCourse(universe, object.designation, "COM")
	return len(course) - 1
}

func CountTransfers(universe map[string]*SpaceObject) int {
	myCourse := MapCourse(universe, "YOU", "COM")
	santaCourse := MapCourse(universe, "SAN", "COM")

	//PrintCourse("myCourse", myCourse)
	//PrintCourse("santaCourse", santaCourse)

	for i, spaceObj := range myCourse {
		pos, ok := PositionInCourse(santaCourse, spaceObj.designation)
		if ok {
			// found common point on path; need to subtract two because path includes
			// the start and end planets, which don't count when we're counting
			// transfers
			return pos + i - 2
		}
	}

	return 0
}

func PositionInCourse(course []*SpaceObject, searchDesignation string) (int, bool) {
	for pos, spaceObject := range course {
		if spaceObject.designation == searchDesignation {
			return pos, true
		}
	}

	return 0, false
}

func PrintCourse(prefix string, course []*SpaceObject) {
	fmt.Print(prefix)
	for _, obj := range (course) {
		fmt.Printf(" --> %v", obj.designation)
	}
	fmt.Printf("\n")
}

func MapCourse(universe map[string]*SpaceObject, start string, end string) []*SpaceObject {
	course := make([]*SpaceObject, 0)
	lastObj := universe[start]

	if start == end {
		course = append(course, lastObj)
		return course
	}

	for {
		course = append(course, lastObj)
		if lastObj.parent.designation == end {
			break
		}
		lastObj = lastObj.parent
	}

	// add the last object
	course = append(course, lastObj.parent)

	return course
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	orbitMap := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	universe := MakeUniverse(orbitMap)
	return "Answer: " + strconv.Itoa(CountOrbits(universe))
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	orbitMap := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	universe := MakeUniverse(orbitMap)
	return "Answer: " + strconv.Itoa(CountTransfers(universe))
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: " + Part2(string(bytes)))
}
