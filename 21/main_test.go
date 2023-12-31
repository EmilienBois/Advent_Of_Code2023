package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1(6, "inputtest.txt")
	expected := 19114
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2(t *testing.T) {
	result := Part2(500, "inputtest.txt")
	expected := 94
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
