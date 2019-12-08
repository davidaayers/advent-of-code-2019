package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	black = iota
	white
	transparent
)

type Sif struct {
	width, height int
	layers        [][]string
}

func decodeSif(undecodedSif string, width int, height int) Sif {
	sif := Sif{width: width, height: height}
	sif.layers = make([][]string, 0)

	for i := 0; i < len(undecodedSif); i += width * height {
		layerStr := undecodedSif[i : i+width*height]
		pixels := strings.Split(layerStr, "")
		sif.layers = append(sif.layers, pixels)
	}
	return sif
}

func renderSif(sif Sif) {
	for y := 0; y < sif.height; y++ {
		for x := 0; x < sif.width; x++ {
			firstNonTransparentPixel := findFirstNonTransparentPixelAt(sif, x, y)
			if firstNonTransparentPixel == "0" {
				firstNonTransparentPixel = " "
			}
			fmt.Printf("%v", firstNonTransparentPixel)
		}
		fmt.Println("")
	}
}

func findFirstNonTransparentPixelAt(sif Sif, x int, y int) string {
	for _, layer := range sif.layers {
		pixel := layer[y*(sif.width)+x]
		pixelValue, _ := strconv.Atoi(pixel)
		if pixelValue != transparent {
			return pixel
		}
	}

	panic("Non transparent pixel not found")
}

func sifChecksum(sif string, width int, height int) int {
	leastZeros := math.MaxInt64
	leastZeroLayer := map[int]int{}

	for i := 0; i < len(sif); i += width * height {
		layerStr := sif[i : i+width*height]
		layer := make(map[int]int)
		pixels := strings.Split(layerStr, "")

		for _, pixelStr := range pixels {
			pixel, _ := strconv.Atoi(pixelStr)
			layer[pixel]++
		}

		if layer[0] < leastZeros {
			leastZeros = layer[0]
			leastZeroLayer = layer
		}
	}
	return leastZeroLayer[1] * leastZeroLayer[2]
}

// Part1 Part 1 of puzzle
func Part1(input string) string {
	width, height := 25, 6
	return "Answer: " + strconv.Itoa(sifChecksum(input, width, height))
}

// Part2 Part2 of puzzle
func Part2(input string) {
	width, height := 25, 6
	sif := decodeSif(input, width, height)
	renderSif(sif)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1: " + Part1(string(bytes)))
	fmt.Println("Part 2: ")
	Part2(string(bytes))
}
