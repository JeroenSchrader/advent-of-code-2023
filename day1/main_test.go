package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var inputTest string

//go:embed input_test_2.txt
var inputTest2 string

type testInput struct {
	input      string
	startIndex int
	expected   string
}

type testInput2 struct {
	input      string
	startIndex int
	expected   int
}

func TestDay1Part1(t *testing.T) {
	result := SolvePart1(inputTest)
	expected := 142
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay1Part2(t *testing.T) {
	result := SolvePart2(inputTest2)
	expected := 281
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}


func TestGetNextNumber(t *testing.T) {
	var tests = []testInput{
		{input: "one", startIndex: 0, expected: "1"},
		{input: "one", startIndex: 1, expected: "0"},
		{input: "onefive", startIndex: 0, expected: "1"},
		{input: "onefive", startIndex: 3, expected: "5"},
		{input: "onefive", startIndex: 4, expected: "0"},
		{input: "89", startIndex: 0, expected: "8"},
		{input: "89", startIndex: 1, expected: "9"},
		{input: "89eight", startIndex: 2, expected: "8"},
		{input: "eightwo", startIndex: 0, expected: "8"},
		{input: "eightwo", startIndex: 4, expected: "2"},
		{input: "nineightwo", startIndex: 0, expected: "9"},
		{input: "nineightwo", startIndex: 3, expected: "8"},
		{input: "nineightwo", startIndex: 7, expected: "2"},
	}

	for _, test := range tests {
		result, _ := getNextNumber(test.input, test.startIndex)
		if result != test.expected {
			t.Errorf("Expected %v, got %v for input %v", test.expected, result, test.input)
		}
	}
}

func TestCalculateCalibrationValuePart2(t *testing.T) {
	var tests = []testInput2{
		{input: "two1nine", expected: 29 },
		{input: "eightwothree", expected: 83 },
		{input: "abcone2threexyz", expected: 13 },
		{input: "xtwone3four", expected: 24 },
		{input: "4nineeightseven2", expected: 42 },
		{input: "zoneight234", expected: 14 },
		{input: "7pqrstsixteen", expected: 76 },
	}

	for _, test := range tests {
		result := calculateCalibrationValuePart2(test.input)
		if result != test.expected {
			t.Errorf("Expected %v, got %v for input %v", test.expected, result, test.input)
		}
	}
}