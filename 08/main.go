package main

import (
	"strings"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type tuple struct {
	Left  string
	Right string
}

func Part1(s string) int {
	test, size := utils.Read_file(s)

	sequence := test[0]

	rules := make(map[string]tuple)

	for i := 2; i < size; i++ {
		this := strings.Split(test[i], " = ")
		str := strings.Split(strings.Split(this[1], "(")[1], ")")[0]
		left := strings.Split(str, ", ")[0]
		right := strings.Split(str, ", ")[1]

		rules[this[0]] = tuple{left, right}
	}

	end := "ZZZ"
	actual := "AAA"
	i := 0
	steps := 0
	for actual != end {
		step := sequence[i]
		if step == 'L' {
			actual = rules[actual].Left
		} else { //  if step == 'R'
			actual = rules[actual].Right
		}
		if i == len(sequence)-1 {
			i = 0
		} else {
			i += 1
		}
		steps += 1
	}
	return steps
}

func Part2(s string) int {
	test, size := utils.Read_file(s)

	sequence := test[0]

	rules := make(map[string]tuple)
	var liste_actual []string
	for i := 2; i < size; i++ {
		this := strings.Split(test[i], " = ")
		str := strings.Split(strings.Split(this[1], "(")[1], ")")[0]
		left := strings.Split(str, ", ")[0]
		right := strings.Split(str, ", ")[1]
		if this[0][2] == 'A' {
			liste_actual = append(liste_actual, this[0])
		}
		rules[this[0]] = tuple{left, right}
	}
	stop := false
	i := 0
	steps := 0
	for !stop {
		stop = true
		step := sequence[i]
		for i, elt := range liste_actual {

			if step == 'L' {
				liste_actual[i] = rules[elt].Left
			} else { //  if step == 'R'
				liste_actual[i] = rules[elt].Right
			}
			if liste_actual[i][2] != 'Z' {
				stop = false
			}
		}
		if i == len(sequence)-1 {
			i = 0
		} else {
			i += 1
		}
		steps += 1
	}
	return steps
}

func Part2fast(s string) int {
	test, size := utils.Read_file(s)

	sequence := test[0]

	rules := make(map[string]tuple)
	var liste_actual []string
	for i := 2; i < size; i++ {
		this := strings.Split(test[i], " = ")
		str := strings.Split(strings.Split(this[1], "(")[1], ")")[0]
		left := strings.Split(str, ", ")[0]
		right := strings.Split(str, ", ")[1]
		if this[0][2] == 'A' {
			liste_actual = append(liste_actual, this[0])
		}
		rules[this[0]] = tuple{left, right}
	}
	i := 0
	steps := 0
	var liste_steps []int
	for _, elt := range liste_actual {
		for elt[2] != 'Z' {
			step := sequence[i]
			if step == 'L' {
				elt = rules[elt].Left
			} else { //  if step == 'R'
				elt = rules[elt].Right
			}
			if i == len(sequence)-1 {
				i = 0
			} else {
				i += 1
			}
			steps += 1
		}
		liste_steps = append(liste_steps, steps)
	}
	to_return := 0
	end := false
	for !end {
		end = true
		for _, elt := range liste_steps {
			if to_return%elt != 0 {
				end = false
				break
			}
		}
		to_return += 1
	}

	return to_return
}
func main() {
	result_one_test1 := Part1("inputtest1.txt")
	println("The result of the input test for the part 1 is : ", result_one_test1)
	result_one_test2 := Part1("inputtest2.txt")
	println("The result of the input test for the part 1 is : ", result_one_test2)
	result_one := Part1("input.txt")
	println("The result of the input for the part 1 is : ", result_one)
	result_two_test := Part2fast("inputtestpart2.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)
	result_two := Part2fast("input.txt")
	println("The result of the input for the part 2 is : ", result_two)
}
