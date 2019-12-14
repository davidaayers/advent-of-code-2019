package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type ReactionRecipe struct {
	Ingredient
	inputs []Ingredient
}

type Ingredient struct {
	result string
	count  int
}

func DetermineRequiredOre(reactions []string) int {
	recipes := make(map[string]ReactionRecipe, len(reactions))
	for _, reaction := range reactions {
		recipe := parseReaction(reaction)
		recipes[recipe.result] = recipe
	}

	return Produce("FUEL", 1, recipes, make(map[string]int))
}

func Produce(desiredElement string, numDesired int, recipes map[string]ReactionRecipe, excess map[string]int) int {
	// we make ore for free!
	if desiredElement == "ORE" {
		return numDesired
	}

	// if we have enough excess already, consume it
	if excess[desiredElement] >= numDesired {
		excess[desiredElement] -= numDesired
		return 0
	}

	// if we don't have enough in excess, use what we have
	if excess[desiredElement] > 0 {
		numDesired -= excess[desiredElement]
		excess[desiredElement] = 0
	}

	// how many batches must we make?
	recipe := recipes[desiredElement]
	batches := int(math.Ceil(float64(numDesired) / float64(recipe.count)))

	// consume the necessary ingredients to produce this element
	ore := 0
	for _, input := range recipe.inputs {
		ore += Produce(input.result, input.count*batches, recipes, excess)
	}

	// produce, and store any excess for later use
	numProduced := batches * recipe.count
	excess[desiredElement] += numProduced - numDesired

	return ore
}

func parseReaction(reaction string) ReactionRecipe {
	parts := strings.Split(reaction, "=>")
	reactionResult, resultCnt := parseElement(parts[1])

	recipe := ReactionRecipe{
		Ingredient: Ingredient{reactionResult, resultCnt},
		inputs:     make([]Ingredient, 0),
	}

	for _, inputStr := range strings.Split(parts[0], ",") {
		input, inputCnt := parseElement(inputStr)
		recipe.inputs = append(recipe.inputs, Ingredient{
			result: input,
			count:  inputCnt,
		})
	}

	return recipe
}

func parseElement(element string) (elementName string, count int) {
	element = strings.TrimSpace(element)
	parts := strings.Split(element, " ")
	elementName = parts[1]
	count, _ = strconv.Atoi(parts[0])
	return
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	ore := DetermineRequiredOre(lines)
	return "Answer: " + strconv.Itoa(ore)
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
