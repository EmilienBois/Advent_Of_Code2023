package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	//We read the file line by line
	test, size := read_file("input.txt")
	compteur_global := 0
	//For every line we're gonna get the index of every symbols
	var tab_index [141][]int
	var tab_index_stars [141][]int

	for i := 0; i < size; i++ { // Here we initialise the tab index containing the number
		s := test[i]
		for j := 0; j < len(s); j++ {
			number := 0

			new_j := j
			notend := is_number(s[new_j])
			if notend {
				for notend { // on récupère la taille d'un nombre
					number += 1
					new_j += 1
					if new_j == 140 {
						notend = false
					} else {
						notend = is_number(s[new_j])
					}
				}
				tab_index[i] = append(tab_index[i], j)      //We have the index of the number
				tab_index[i] = append(tab_index[i], number) //We have the size of the number
				num_number, err := strconv.Atoi(s[j : j+number])
				if err != nil {
					// ... handle error
					panic(err)
				}
				tab_index[i] = append(tab_index[i], num_number)
				j = new_j
			}
		}

	}
	for i := 0; i < size; i++ {
		tab_index_stars[i] = get_index_stars(test[i])
	}
	for i := 0; i < size; i++ {
		for _, elt := range tab_index_stars[i] {
			t_number, number := touched_number(i, elt, tab_index)
			if t_number == 2 {
				compteur_global += number[0] * number[1]
			}
		}
	}
	println(compteur_global)
}

func touched_number(i int, elt int, tab_index [141][]int) (int, []int) {
	max := 1
	min := -1
	if i == 0 {
		min = 0
	}
	if i == 140 {
		max = 0
	}
	var number_touched []int
	number_stars := 0
	for l := min; l < max+1; l++ {
		this := tab_index[l+i]
		for c := 0; c < len(this); c += 3 {
			index := this[c]
			size := this[c+1]
			number := this[c+2]
			for j := 0; j < size; j++ {
				real_index := j + index
				if real_index == elt-1 || real_index == elt || real_index == elt+1 {
					number_stars += 1
					number_touched = append(number_touched, number)
					j = size
				}
			}
		}
	}
	return number_stars, number_touched
}

func get_index_stars(s string) []int {
	var new_str []int
	for i, elt := range s {
		if elt == '*' {
			new_str = append(new_str, i)
		}
	}
	return new_str
}

func main2() {
	//We read the file line by line
	test, size := read_file("input.txt")
	compteur := 0
	//For every line we're gonna get the index of every symbols
	var tab_index [141][]int
	for i := 0; i < size; i++ {
		tab_index[i] = get_index(test[i])
	}
	for i := 0; i < size; i++ {
		s := test[i]
		for j := 0; j < len(s); j++ {
			number := 0
			new_j := j
			notend := is_number(s[new_j])
			for notend { // on récupère la taille d'un nombre
				number += 1
				new_j += 1
				if new_j == 140 {
					notend = false
				} else {
					notend = is_number(s[new_j])
				}
			}
			//Ensuite on check si le numéro touche un symbole
			touch := false
			for index := 0; index < number; index++ {
				if is_touching(index+j, i, tab_index) {
					touch = true
				}
			}
			if touch == true {
				num_number, err := strconv.Atoi(s[j : j+number])
				if err != nil {
					// ... handle error
					panic(err)
				}
				compteur = compteur + num_number
			}
			j = new_j

		}

	}
	println(compteur)
}

func is_touching(j int, i int, tab_index [141][]int) bool {
	max := 1
	min := -1
	if i == 0 {
		min = 0
	}
	if i == 140 {
		max = 0
	}
	for l := min; l < max+1; l++ {
		this := tab_index[l+i]
		for _, elt := range this {
			if j == elt-1 || j == elt || j == elt+1 {
				return true
			}
		}
	}
	return false
}

func is_number(elt byte) bool {
	if elt == '0' || elt == '1' || elt == '2' || elt == '3' || elt == '4' || elt == '5' || elt == '6' || elt == '7' || elt == '8' || elt == '9' {
		return true
	}
	return false
}

func get_index(s string) []int {
	var new_str []int
	for i, elt := range s {
		if elt == '.' {
		} else if is_number(s[i]) {
		} else {
			new_str = append(new_str, i)
		}

	}
	return new_str
}

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
