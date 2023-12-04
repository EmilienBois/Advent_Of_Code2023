package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest.txt")
	expected := 13
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
func TestPart2(t *testing.T) {
	result := Part2("inputtest.txt")
	expected := 30
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGet_numbers(t *testing.T) {
	test := "45 63  4"
	expected := []int{45, 63, 4}
	result := Get_numbers(test)
	for i := 0; i < 3; i++ {
		if result[i] != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
		}
	}

}
