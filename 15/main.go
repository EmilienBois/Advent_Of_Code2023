package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

func Part1(s string) int {
	file, _ := utils.Read_file(s)

	print('H')

	liste_steps := strings.Split(file[0], ",")

	total := 0

	for _, elt := range liste_steps {
		total += Get_HASH(elt)
	}
	return total
}

func Get_HASH(s string) int {

	current_value := 0
	for _, elt := range s {
		current_value += int(elt)
		current_value *= 17
		current_value = current_value % 256

	}
	return current_value
}

type Lens struct {
	label string
	focal int
}

func Part2(s string) int {
	file, _ := utils.Read_file(s)

	liste_steps := strings.Split(file[0], ",")

	var liste_box [256][]Lens

	for _, elt := range liste_steps {
		if elt[len(elt)-1] == '-' {
			box := Get_HASH(elt[:len(elt)-1])
			label := elt[:len(elt)-1]
			for i, _ := range liste_box[box] {
				if liste_box[box][i].label == label {
					liste_box[box][i].focal = 0
					liste_box[box][i].label = ""
					break
				}
			}
		} else {
			label := elt[:len(elt)-2]
			focal := utils.String_to_Int(string(elt[len(elt)-1]))
			box := Get_HASH(label)
			exist := false
			for i, _ := range liste_box[box] {
				if liste_box[box][i].label == label {
					liste_box[box][i].focal = focal

					exist = true
					break
				}
			}
			if !exist {
				liste_box[box] = append(liste_box[box], Lens{label, focal})
			}

		}
	}
	total := 0

	for i, box := range liste_box {
		real_j := 0
		for _, lens := range box {
			if lens.label == "" {
				real_j -= 1
			} else {
				total += (i + 1) * (real_j + 1) * lens.focal
			}
			real_j += 1
		}
	}

	return total
}

func main() {
	start := time.Now()
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)

	end := time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_one := Part1("input.txt")
	println("The result of the input for the part 1 is : ", result_one)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two_test := Part2("inputtest.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two := Part2("input.txt")
	println("The result of the input for the part 2 is : ", result_two)

	end = time.Now()
	fmt.Printf("%s \n", end.Sub(start))
}
