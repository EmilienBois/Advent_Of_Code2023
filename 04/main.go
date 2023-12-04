package main

import (
	"strings"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

func Part2(s string) int {
	test, size := utils.Read_file(s)

	nb_copied := initialize_1_slice(size)

	for i := 0; i < size; i++ {

		matches := 0
		this := test[i]

		new_str := strings.Split(this, ":")
		interest := strings.Split(new_str[1], "|")

		win := Get_numbers(interest[0])
		numbers := Get_numbers(interest[1])

		for _, nbr := range numbers {
			for _, w := range win {
				if nbr == w {
					matches += 1
				}
			}
		}
		for index := 1; index <= matches; index++ {
			nb_copied[i+index] += nb_copied[i]

		}
	}
	return utils.Somme(nb_copied)
}

func Part1(s string) int {

	test, size := utils.Read_file(s)
	global_points := 0
	for i := 0; i < size; i++ {
		card_point := 0
		this := test[i]
		new_str := strings.Split(this, ":")
		interest := strings.Split(new_str[1], "|")

		win := Get_numbers(interest[0])
		numbers := Get_numbers(interest[1])

		for _, nbr := range numbers {
			for _, w := range win {
				if nbr == w {
					if card_point == 0 {
						card_point = 1
					} else {
						card_point = card_point * 2
					}
				}
			}
		}

		global_points += card_point
	}

	return global_points
}

func Get_numbers(s string) []int {

	var liste []int
	for i := 0; i < len(s); i++ {
		end := 0
		if utils.Is_number(s[i]) {
			end = 1
			if i+1 < len(s) && utils.Is_number(s[i+1]) {
				end = 2
			}
			number := utils.Byte_to_Int(s[i : i+end])

			liste = append(liste, number)
		}
		i = i + end
	}
	return liste
}

func initialize_1_slice(size int) []int {
	var to_return []int
	for i := 0; i < size; i++ {
		to_return = append(to_return, 1)
	}
	return to_return
}
func main() {
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)
	result_one := Part1("input.txt")
	println("The result of the input test for the part 1 is : ", result_one)
	result_two_test := Part2("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_two_test)
	result_two := Part2("input.txt")
	println("The result of the input test for the part 1 is : ", result_two)
}
