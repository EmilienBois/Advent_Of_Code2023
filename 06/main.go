package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

func Part1(s string) int {
	//We read the file entirely
	test, _ := utils.Read_file(s)
	// We separate distance and
	total := 1
	file_time := utils.Get_liste_numbers(strings.Split(test[0], ": ")[1])
	file_dist := utils.Get_liste_numbers(strings.Split(test[1], ": ")[1])

	t := 0
	dist := 0

	run_succed := 0
	for runs := 0; runs < len(file_time); runs++ {
		run_succed = 0
		t = file_time[runs]
		dist = file_dist[runs]
		s1, s2, nb_solution := Resolve_second_degree(1, -float64(t), float64(dist))
		var r1, r2 int

		r1 = int(s1) + 1

		if float64(int(s2)) == s2 {
			r2 = int(s2) - 1
		} else {
			r2 = int(s2)
		}

		if nb_solution == 2 {
			if r1 > 0 && r2 < t {
				run_succed += r2 - r1 + 1
			} else if r1 > 0 {
				run_succed += t - r1
			} else if r2 < t {
				run_succed += r2
			}
		}
		total = total * run_succed
	}
	return total
}

func Resolve_second_degree(a, b, c float64) (float64, float64, int) {
	delta := b*b - 4*a*c
	if delta < 0 {
		return 0, 0, 0
	} else if delta == 0 {
		return -b / (2 * a), -b / (2 * a), 1
	} else {
		sdelta := math.Sqrt(float64(delta))
		s1 := (-float64(b) + sdelta) / (2 * float64(a))
		s2 := (-float64(b) - sdelta) / (2 * float64(a))
		return min(s1, s2), max(s1, s2), 2
	}
}

func Part2(s string) int {
	//We read the file entirely
	test, _ := utils.Read_file(s)
	// We separate distance and

	time := Get_the_number(test[0])
	dist := Get_the_number(test[1])
	s1, s2, _ := Resolve_second_degree(1, -float64(time), float64(dist))

	return int(s2) - int(s1)
}

func Get_the_number(s string) int {
	liste_number := utils.Get_liste_numbers(strings.Split(s, ": ")[1])
	var time_int []string
	for _, elt := range liste_number {
		time_int = append(time_int, strconv.Itoa(elt))
	}
	str := strings.Join(time_int, "")

	return utils.Byte_to_Int(str)

}
func main() {
	println("--2023 day 06 solution--")
	start := time.Now()
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)

	println(time.Since(start))
	result_one := Part1("input.txt")
	println("The result of the input for the part 1 is : ", result_one)
	println(time.Since(start))
	result_two_test := Part2("inputtest.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)
	println(time.Since(start))
	result_two := Part2("input.txt")
	println("The result of the input for the part 2 is : ", result_two)
	println(time.Since(start))
	end := time.Now()
	fmt.Printf("%s", end.Sub(start))
}
