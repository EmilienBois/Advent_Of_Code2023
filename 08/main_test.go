package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest1.txt")
	expected := 2
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
	result = Part1("inputtest2.txt")
	expected = 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2(t *testing.T) {
	result := Part2fast("inputtestpart2.txt")
	expected := 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
