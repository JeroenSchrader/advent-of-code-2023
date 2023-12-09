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
	expected   string
}

type testInput2 struct {
	input      string
	expected   int
}

func TestDay2Part1(t *testing.T) {
	result := SolvePart1(inputTest)
	expected := 8
	if (result != expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDay2Part2(t *testing.T) {
	result := SolvePart2(inputTest)
	expected := 2286
	if (result != expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}