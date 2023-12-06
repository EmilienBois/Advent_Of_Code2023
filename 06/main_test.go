package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest.txt")
	expected := 288
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGet_the_number(t *testing.T) {
	result := Get_the_number("gf: 7      78  655    0")
	expected := 7786550
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestResolve_second_degree(t *testing.T) {
	s1, s2, solutions := Resolve_second_degree(1, 2, 3)
	if s1 != 0 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", s1, 0.0)
	}
	if s2 != 0 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", s2, 0.0)
	}
	if solutions != 0 {
		t.Errorf("Result is incorrect, got: %d, want: %f.", solutions, 0.0)
	}

	s1, s2, solutions = Resolve_second_degree(1, 2, 1)
	if s1 != -1 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", s1, -1.0)
	}
	if s2 != -1 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", s2, -1.0)
	}
	if solutions != 1 {
		t.Errorf("Result is incorrect, got: %d, want: %f.", solutions, 1.0)
	}

	s1, s2, solutions = Resolve_second_degree(1, 3, 2)
	if s1 != -2.0 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", s1, -2.0)
	}
	if s2 != -1.0 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", s2, -1.0)
	}
	if solutions != 2 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", solutions, 2)
	}
}
