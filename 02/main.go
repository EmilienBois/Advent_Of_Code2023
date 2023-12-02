package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func read_file(path string) ([]string, int) {
	f, _ := os.Open(path)

	scanner := bufio.NewScanner(f)

	var newMeasureString string
	var file []string
	i := 0
	for scanner.Scan() {

		newMeasureString = scanner.Text()
		// adding to the list
		file = append(file, newMeasureString)
		i += 1
	}
	return file, i
}

func better(min [3]int, s string) [3]int {
	str := strings.Split(s, ", ")
	for _, elt := range str {
		substring := strings.Split(elt, " ")

		number, err := strconv.Atoi(substring[0])
		if err != nil {
			// ... handle error
			panic(err)
		}

		switch substring[1] {
		case "red":
			if number > min[0] {
				min[0] = number
			}
		case "green":
			if number > min[1] {
				min[1] = number
			}
		case "blue":
			if number > min[2] {
				min[2] = number
			}
		}
	}
	return min
}

func main() {
	file, size := read_file("input.txt")

	compteur := 0

	for i := 0; i < size; i++ {

		intermediate := strings.Split(file[i], ": ")
		line := strings.Split(intermediate[1], "; ")

		min := [3]int{0, 0, 0}

		for _, elt := range line {
			min = better(min, elt)
		}
		println("Les minimum valent ", min[0], " ", min[1], " ", min[2])
		mult := min[0] * min[1] * min[2]
		println("La multiplication vaut ", mult)
		compteur += mult

	}
	println("Sum of all the games ID that are possible : ", compteur)

}

// For part 1 :
func main2() {
	file, size := read_file("input.txt")
	maxred := 12
	maxblue := 14
	maxgreen := 13
	max := [3]int{maxred, maxgreen, maxblue}
	compteur := 0

	for i := 0; i < size; i++ {

		intermediate := strings.Split(file[i], ": ")
		line := strings.Split(intermediate[1], "; ")
		valide := true
		for _, elt := range line {
			valide = validate(max, elt)
			if valide == false {
				break
			}
		}
		if valide == true {
			println("game ", i, "valide")
			compteur += i + 1
		}
	}
	println("Sum of all the games ID that are possible : ", compteur)

}

func validate(max [3]int, s string) bool {
	str := strings.Split(s, ", ")
	for _, elt := range str {
		substring := strings.Split(elt, " ")

		number, err := strconv.Atoi(substring[0])
		if err != nil {
			// ... handle error
			panic(err)
		}

		switch substring[1] {
		case "red":
			if number > max[0] {
				return false
			}
		case "green":
			if number > max[1] {
				return false
			}
		case "blue":
			if number > max[2] {
				return false
			}
		}
	}
	return true
}
