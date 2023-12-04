package main

import "testing"

//go:embed inputtest.txt

var inputTest string

func TestPart1(t *testing.T) {
	result := Part1("inputtest")
	expected := 4361
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
