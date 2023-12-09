package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	part1 := SolvePart1(inputFile)
	part2 := SolvePart2(inputFile)
	fmt.Printf("Day 2 results\n")
	fmt.Printf("-- Part 1: %d\n", part1)
	fmt.Printf("-- Part 2: %d\n", part2)
}

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	totalGameIds := 0
	for _, line := range lines {
		result := getPossibleGameId(line)
		if (result != -1) {
			totalGameIds += result
		}
	}
	return totalGameIds
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		total += calculatePowerOfCubes(line)
	}
	return total
}

func calculatePowerOfCubes(line string) int {
	minColorMap := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	splits := strings.Split(line, ":")
	rounds := strings.Split(splits[1], ";")

	for _, round := range rounds {
		colors := strings.Split(strings.TrimSpace(round), ",")
		for _, color := range colors {
			cSplit := strings.Split(strings.TrimSpace(color), " ")
			cAmount, _ := strconv.Atoi(cSplit[0])
			cColor := cSplit[1]

			if (cAmount > minColorMap[cColor]) {
				minColorMap[cColor] = cAmount
			}
		}
	}

	return minColorMap["red"] * minColorMap["blue"] * minColorMap["green"]
}

func getPossibleGameId(line string) int {
	colorLimitMap := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	splits := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(strings.Split(splits[0], " ")[1])

	rounds := strings.Split(splits[1], ";")
	for _, round := range rounds {
		colors := strings.Split(strings.TrimSpace(round), ",")
		for _, color := range colors {
			cSplit := strings.Split(strings.TrimSpace(color), " ")
			cAmount, _ := strconv.Atoi(cSplit[0])
			cColor := cSplit[1]

			if (cAmount > colorLimitMap[cColor]) {
				return -1
			}
		}
	}

	return gameId;
}
