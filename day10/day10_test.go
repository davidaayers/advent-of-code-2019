package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var testCases = []struct {
	universe                               string
	numAsteroidsObservableFromBestLocation int
}{
	{
		`.#..#
.....
#####
....#
...##`,
		8,
	},
	{
		`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
		33,
	},
}

func TestFindBestLocation(t *testing.T) {
	for _, testCase := range testCases {
		fmt.Println(testCase.universe)
		universe := ParseMap(testCase.universe)
		bestLocation := FindBestLocation(&universe)

		if bestLocation.numObservableAsteroids != testCase.numAsteroidsObservableFromBestLocation {
			t.Errorf("Error, expected %v got %v", testCase.numAsteroidsObservableFromBestLocation, bestLocation.numObservableAsteroids)
		}
	}
}

// go test -timeout 30s day03 -run '^(TestPart1)$'
func TestPart1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: "
	answer := Part1(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

// go test -timeout 30s day03 -run '^(TestPart2)$'
func TestPart2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("input.txt")
	expected := "Answer: "
	answer := Part2(string(bytes))
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
