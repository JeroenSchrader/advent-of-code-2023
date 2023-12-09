package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	part1 := SolvePart1(inputFile)
	part2 := SolvePart2(inputFile)
	fmt.Printf("Day 1 results\n")
	fmt.Printf("-- Part 1: %d\n", part1)
	fmt.Printf("-- Part 2: %d\n", part2)
}

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	calibrationValueTotal := 0
	for _, line := range lines {
		res := calculateCalibrationValue(line)
		calibrationValueTotal += res
	}
	return calibrationValueTotal
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	calibrationValueTotal := 0
	for _, line := range lines {
		res := calculateCalibrationValuePart2(line)
		calibrationValueTotal += res
	}
	return calibrationValueTotal
}

func calculateCalibrationValuePart2(line string) int {
	firstNum := ""
	lastNum := ""

	startIndex := 0
	for true {
		number, nextIndex := getNextNumber(line, startIndex)
		if nextIndex == -1 {
			if startIndex == len(line) {
				break
			}
			startIndex++
			continue
		}
		startIndex = nextIndex

		if firstNum == "" {
			firstNum = number
		} else {
			lastNum = number
		}
	}

	if lastNum == "" {
		res, _ := strconv.Atoi(firstNum + firstNum)
		return res
	}

	res, _ := strconv.Atoi(firstNum + lastNum)
	return res
}

func getNextNumber(line string, startIndex int) (string, int) {
	numbers := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	keys := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	for _, key := range keys {
		endIndex := startIndex+len(key)

		if (endIndex > len(line)) {
			continue
		}

		if strings.Contains(line[startIndex:endIndex], key) {
			if (endIndex > 1 && len(key) > 1) {
				// -1 because some letters overlap, for example 'nineight' (should result in nine and eight)
				endIndex -= 1
			}
			return numbers[key], endIndex
		}
	}

	return "0", -1
}

func calculateCalibrationValue(line string) int {
	numbers := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	firstNum := ""
	lastNum := ""
	for _, char := range line {
		number := string(char)
		if slices.Contains(numbers, char) {
			if firstNum == "" {
				firstNum = number
			} else {
				lastNum = number
			}
		}
	}

	if lastNum == "" {
		res, _ := strconv.Atoi(firstNum + firstNum)
		return res
	}

	res, _ := strconv.Atoi(firstNum + lastNum)
	return res
}
