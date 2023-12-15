package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest.txt")
	expected := 1320
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGet_HASH(t *testing.T) {
	result := Get_HASH("HASH")
	expected := 52
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
