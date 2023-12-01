package main

import (
	"bufio"
	"os"
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

func get_first(s string) int {
	for i, elt := range s {
		if elt == '0' {
			return 0
		}
		if elt == '1' {
			return 1
		}
		if elt == '2' {
			return 2
		}
		if elt == '3' {
			return 3
		}
		if elt == '4' {
			return 4
		}
		if elt == '5' {
			return 5
		}
		if elt == '6' {
			return 6
		}
		if elt == '7' {
			return 7
		}
		if elt == '8' {
			return 8
		}
		if elt == '9' {
			return 9
		}
		if i < len(s)-5 {
			if s[i:i+5] == "three" {
				return 3
			}
			if s[i:i+5] == "seven" {
				return 7
			}
			if s[i:i+5] == "eight" {
				return 8
			}
		}
		if i < len(s)-4 {
			if s[i:i+4] == "nine" {
				return 9
			}
			if s[i:i+4] == "four" {
				return 4
			}
			if s[i:i+4] == "five" {
				return 5
			}
		}
		if i < len(s)-3 {
			if s[i:i+3] == "one" {
				return 1
			}
			if s[i:i+3] == "two" {
				return 2
			}
			if s[i:i+3] == "six" {
				return 6
			}
		}
	}
	return -1
}

func get_last(s string) int {
	for i, _ := range s {
		if s[len(s)-i-1] == '0' {
			return 0
		}
		if s[len(s)-i-1] == '1' {
			return 1
		}
		if s[len(s)-i-1] == '2' {
			return 2
		}
		if s[len(s)-i-1] == '3' {
			return 3
		}
		if s[len(s)-i-1] == '4' {
			return 4
		}
		if s[len(s)-i-1] == '5' {
			return 5
		}
		if s[len(s)-i-1] == '6' {
			return 6
		}
		if s[len(s)-i-1] == '7' {
			return 7
		}
		if s[len(s)-i-1] == '8' {
			return 8
		}
		if s[len(s)-i-1] == '9' {
			return 9
		}

		if i < len(s)-5 {
			if s[len(s)-i-5:len(s)-i] == "three" {
				return 3
			}
			if s[len(s)-i-5:len(s)-i] == "seven" {
				return 7
			}
			if s[len(s)-i-5:len(s)-i] == "eight" {
				return 8
			}
		}
		if i < len(s)-4 {
			if s[len(s)-i-4:len(s)-i] == "nine" {
				return 9
			}
			if s[len(s)-i-4:len(s)-i] == "four" {
				return 4
			}
			if s[len(s)-i-4:len(s)-i] == "five" {
				return 5
			}
		}
		if i < len(s)-3 {
			if s[len(s)-i-3:len(s)-i] == "one" {
				return 1
			}
			if s[len(s)-i-3:len(s)-i] == "two" {
				return 2
			}
			if s[len(s)-i-3:len(s)-i] == "six" {
				return 6
			}
		}

	}
	return -1
}
func somme(liste []int) int {
	somme := 0
	for _, elt := range liste {
		somme += elt
	}
	return somme
}

func main() {
	test, size := read_file("input.txt")
	println(test)
	var digits []int
	//Then we work on the test list to find the digit
	for i := 0; i < size; i++ {
		actual := test[i]
		f := get_first(actual)
		l := get_last(actual)
		digits = append(digits, f*10+l)
	}
	println(somme(digits))
}
