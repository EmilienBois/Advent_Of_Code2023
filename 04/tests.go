package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest")
	expected := 4361
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func Testget_numbers(t *testing.T) {
	test := "45 63  4"
	expected := []int{45, 63, 4}
	result := Get_numbers(test)
	for i := 0; i < 3; i++ {
		if result[i] != expected[i] {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
		}
	}

}
