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

func (recipe ReactionRecipe) requiresOre() bool {
	// must have only 1 input, and that input must be ore
	if len(recipe.inputs) == 1 && recipe.inputs[0].result == "ORE" {
		return true
	}
	return false
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

	fmt.Printf("Recipes: %v\n", recipes)

	ingredients := make(map[string]int)

	MakeReaction("FUEL", recipes, ingredients, 1)

	ore := 0
	for k, v := range ingredients {
		if recipes[k].requiresOre() {
			ore += CalcOre(recipes[k], v)
		}
		fmt.Printf("ingredient: %v qty: %v\n", k, v)
	}

	return ore
}

func CalcOre(recipe ReactionRecipe, numDesired int) int {
	fmt.Printf("-> Recipe %v NumDesired: %v\n", recipe, numDesired)
	batches := int(math.Ceil(float64(numDesired) / float64(recipe.count)))
	oreProduced := batches * recipe.inputs[0].count
	fmt.Printf("-> %v need %v batches of ore at %v ore per batch: %v\n", recipe.result, batches, recipe.inputs[0].count, oreProduced)
	return oreProduced
}

func MakeReaction(desiredElement string, recipes map[string]ReactionRecipe, ingredients map[string]int, numDesired int) {
	recipe := recipes[desiredElement]
	//"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
	// need 53 STKFG, so
	// 2 * 53 = 106 VPVL
	// 7 * 53
	for _, input := range recipe.inputs {
		if input.result != "ORE" {
			// how many batches do I need
			batches := int(math.Ceil(float64(numDesired) / float64(recipe.count)))
			ingredients[input.result] += batches * input.count
			MakeReaction(input.result, recipes, ingredients, input.count)
		}
	}
}

func parseReaction(reaction string) ReactionRecipe {
	parts := strings.Split(reaction, "=>")
	reactionResult, resultCnt := parseElement(parts[1])
	//fmt.Printf("%v: %v\n", reactionResult, resultCnt)

	recipe := ReactionRecipe{
		Ingredient: Ingredient{reactionResult, resultCnt},
		inputs:     make([]Ingredient, 0),
	}

	for _, inputStr := range strings.Split(parts[0], ",") {
		input, inputCnt := parseElement(inputStr)
		//fmt.Printf("-> %v: %v\n", input, inputCnt)
		recipe.inputs = append(recipe.inputs, Ingredient{
			result: input,
			count:  inputCnt,
		})
	}

	//fmt.Printf("Recipe: %v\n", recipe)
	return recipe
}

func MakeReaction2(desiredElement string, recipes map[string]ReactionRecipe, cnt int) int {
	//recipe := recipes[desiredElement]
	//
	//ore := 0
	//for _, input := range recipe.inputs {
	//	if input.result == "ORE" {
	//		ore += input.count * cnt
	//	} else {
	//		ore += MakeReaction(input.result, recipes, input.count)
	//	}
	//}
	return 0
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
	return "Answer: "
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
