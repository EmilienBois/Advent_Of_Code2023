package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest.txt")
	expected := 35
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGet_numbers(t *testing.T) {
	test := "45 6873  4"
	expected := []int{45, 6873, 4}
	result := Get_numbers(test)
	for i := 0; i < 3; i++ {
		if result[i] != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
		}
	}
}
