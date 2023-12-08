package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("inputtest.txt")
	expected := 6440
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestGetpower(t *testing.T) {
	result := Getpower("A2345")
	if result != 1. {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 1.)
	}
	result = Getpower("AA234")
	if result != 2. {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 2.)
	}
	result = Getpower("AAA34")
	if result != 3. {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 3.)
	}
	result = Getpower("AA224")
	if result != 2.5 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 2.5)
	}

	result = Getpower("AA22A")
	if result != 3.5 {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 3.5)
	}
	result = Getpower("AAAA4")
	if result != 4. {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 4.)
	}
	result = Getpower("55555")
	if result != 5. {
		t.Errorf("Result is incorrect, got: %f, want: %f.", result, 5.)
	}
}

func TestCompare(t *testing.T) {
	hand1 := Hand{"AA234", 1, 2, 0}
	hand2 := Hand{"22345", 2, 2, 0}
	hand3 := Hand{"AA456", 3, 2, 0}
	Compare([]Hand{hand1, hand2, hand3})
	if hand1.rank != 2 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", hand1.rank, 2)
	}
}
