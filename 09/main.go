package main

import "github.com/EmilienBois/Advent_Of_Code2023/utils"

func Part1(s string) int {
	test, size := utils.Read_file(s)

	total := 0

	for i := 0; i < size; i++ {
		numbers := utils.Get_liste_numbers(test[i])
		next_value := Get_next(numbers)
		total += next_value
	}
	return total
}

func Get_next(nb []int) int {
	empty := true
	for _, elt := range nb {
		if elt != 0 {
			empty = false
			break
		}
	}
	if empty {
		return 0
	}
	var under_nb []int
	for i := 0; i < len(nb)-1; i++ {
		under_nb = append(under_nb, nb[i+1]-nb[i])
	}
	nb = append(nb, Get_next(under_nb)+nb[len(nb)-1])
	return nb[len(nb)-1]

}

func Part2(s string) int {
	test, size := utils.Read_file(s)

	total := 0

	for i := 0; i < size; i++ {
		numbers := utils.Get_liste_numbers(test[i])
		previous_value := Get_previous(numbers)
		total += previous_value
	}
	return total
}

func Get_previous(nb []int) int {
	empty := true
	for _, elt := range nb {
		if elt != 0 {
			empty = false
			break
		}
	}
	if empty {
		return 0
	}
	var under_nb []int
	for i := 1; i < len(nb); i++ {
		under_nb = append(under_nb, nb[i]-nb[i-1])
	}
	nb = append(nb, nb[0]-Get_previous(under_nb))
	return nb[len(nb)-1]

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
