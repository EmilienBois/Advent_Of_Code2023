package main

import (
	"strings"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

func Part2(s string) int {
	test, size := utils.Read_file(s)
	//Retrieving initial seeds
	init_seeds := Get_numbers(strings.Split(test[0], ": ")[1])
	//retrieving links

	links := make(map[string][]int)

	actual_str := ""
	var values []int
	for i := 2; i < size; i++ {
		this := test[i]
		if len(this) == 0 {
			links[actual_str] = values
			values = []int{}
		} else if !utils.Is_number(this[0]) {
			actual_str = strings.Split(this, " ")[0]
		} else {
			numbers := Get_numbers(this)
			values = append(values, numbers...)
			if i == size-1 {
				links[actual_str] = values
			}
		}
	}
	//So now we have the entire linking
	min := Get_location(links, init_seeds[0])
	for i := 0; i < len(init_seeds); i = i + 2 {
		for j := 0; j < init_seeds[i+1]; j++ {
			location := Get_location(links, init_seeds[i]+j)

			if location < min {
				min = location
			}
		}
	}
	return min
}

func Part1(s string) int {

	test, size := utils.Read_file(s)
	//Retrieving initial seeds
	init_seeds := Get_numbers(strings.Split(test[0], ": ")[1])
	//retrieving links

	links := make(map[string][]int)

	actual_str := ""
	var values []int
	for i := 2; i < size; i++ {
		this := test[i]
		if len(this) == 0 {
			links[actual_str] = values
			values = []int{}
		} else if !utils.Is_number(this[0]) {
			actual_str = strings.Split(this, " ")[0]
		} else {
			numbers := Get_numbers(this)
			values = append(values, numbers...)
			if i == size-1 {
				links[actual_str] = values
			}
		}
	}
	//So now we have the entire linking
	min := Get_location(links, init_seeds[0])
	for _, elt := range init_seeds {
		location := Get_location(links, elt)
		if location < min {
			min = location
		}
	}
	return min
}

func Get_location(links map[string][]int, init int) int {
	liste_steps := [7]string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	actual := init
	next := init
	for _, e := range liste_steps {
		this := links[e]
		for i := 0; i < len(this); i = i + 3 {
			min := this[1+i]
			max := this[2+i] + this[1+i]
			if min <= actual && actual <= max {
				next = this[0+i] + actual - this[1+i]
				i = len(this)
			}
		}
		actual = next
	}
	return next
}

func Get_numbers(s string) []int {
	var numbers []int
	for j := 0; j < len(s); j++ {
		new_j := j
		notend := utils.Is_number(s[new_j])
		for notend { // on récupère la taille d'un nombre
			new_j += 1
			if new_j == len(s) {
				notend = false
			} else {
				notend = utils.Is_number(s[new_j])
			}
		}
		if new_j == j {

		} else {
			number := utils.Byte_to_Int(s[j:new_j])
			numbers = append(numbers, number)
			j = new_j
		}
	}
	return numbers
}

func main() {
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)
	result_one := Part1("input.txt")
	println("The result of the input for the part 1 is : ", result_one)
	result_two_test := Part2("inputtest.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)
	result_two := Part2("input.txt")
	println("The result of the input for the part 2 is : ", result_two)
}
