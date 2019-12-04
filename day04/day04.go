package main

import (
	"fmt"
	"strconv"
)

func IsValidPassword(password int) bool {
	passStr := strconv.Itoa(password)

	// It is a six-digit number.
	if len(passStr) != 6 {
		return false
	}

	foundAdjacent := false
	for i := 1; i < len(passStr); i++ {
		d1,_ := strconv.Atoi(string(passStr[i-1]))
		d2,_ := strconv.Atoi(string(passStr[i]))

		// Two adjacent digits are the same (like `22` in `122345`).
		if d1 == d2 {
			foundAdjacent = true
		}

		// Going from left to right, the digits never decrease; they only ever increase or stay the same (like `111123` or `135679`).
		if d1 > d2 {
			return false
		}
	}

	if !foundAdjacent {
		return false
	}

	return true
}

func IsMoreValidPassword(password int) bool {
	passStr := strconv.Itoa(password)

	// It is a six-digit number.
	if len(passStr) != 6 {
		return false
	}

	for i := 1; i < len(passStr); i++ {
		d1,_ := strconv.Atoi(string(passStr[i-1]))
		d2,_ := strconv.Atoi(string(passStr[i]))

		// Going from left to right, the digits never decrease; they only ever increase or stay the same (like `111123` or `135679`).
		if d1 > d2 {
			return false
		}
	}

	// break it into chunks based on repeating numbers
	chunks := make([]string,0)
	chunk := string(passStr[0])

	for i := 1; i < len(passStr); i++ {
		if passStr[i] == passStr[i-1] {
			chunk += string(passStr[i])
		} else {
			chunks = append(chunks,chunk)
			chunk = string(passStr[i])
		}
	}

	// append last chunk
	chunks = append(chunks,chunk)

	// one of the chunks must be 2
	foundDouble := false
	for _,ch := range(chunks) {
		if len(ch) == 2 {
			foundDouble = true
		}
	}

	if !foundDouble {
		return false
	}

	return true
}

// Part1 Part 1 of puzzle
func Part1(start, end int) string {
	validPasswords := 0

	for i := start; i < end + 1 ; i ++ {
		if IsValidPassword(i) {
			validPasswords++
		}
	}

	return "Answer: " + strconv.Itoa(validPasswords)
}

// Part2 Part2 of puzzle
func Part2(start, end int) string {
	validPasswords := 0

	for i := start; i < end + 1 ; i ++ {
		if IsMoreValidPassword(i) {
			validPasswords++
		}
	}

	return "Answer: " + strconv.Itoa(validPasswords)
}

func main() {
	start,end := 197487,673251

	fmt.Println("Part 1: " + Part1(start,end))
	fmt.Println("Part 2: " + Part2(start,end))
}
