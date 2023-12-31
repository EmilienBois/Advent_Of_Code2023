package utils

import (
	"bufio"
	"os"
	"strconv"
)

// Lecture de fichier
func Read_file(path string) ([]string, int) {
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

func Is_number(elt byte) bool {
	if elt == '0' || elt == '1' || elt == '2' || elt == '3' || elt == '4' || elt == '5' || elt == '6' || elt == '7' || elt == '8' || elt == '9' {
		return true
	}
	return false
}

func String_to_Int(elt string) int {
	number, err := strconv.Atoi(elt)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return number
}

func Somme(liste []int) int {
	somme := 0
	for _, elt := range liste {
		somme += elt
	}
	return somme
}

func Get_liste_numbers(s string) []int {
	var numbers []int
	for j := 0; j < len(s); j++ {
		new_j := j
		negative := false
		if s[j] == '-' {
			new_j += 1
			negative = true
		}
		notend := Is_number(s[new_j])
		for notend { // on récupère la taille d'un nombre
			new_j += 1
			if new_j == len(s) {
				notend = false
			} else {
				notend = Is_number(s[new_j])
			}
		}
		if new_j == j {

		} else {
			if !negative {
				number := String_to_Int(s[j:new_j])
				numbers = append(numbers, number)
				j = new_j
			} else {
				number := String_to_Int(s[j+1 : new_j])
				numbers = append(numbers, -number)
				j = new_j
			}
		}
	}
	return numbers
}

func Min_tab(tab []int) int {
	min := tab[0]
	for i := 1; i < len(tab); i++ {
		if tab[i] < min {
			min = tab[i]
		}
	}
	return min
}
func Max_tab(tab []int) int {
	max := tab[0]
	for i := 1; i < len(tab); i++ {
		if tab[i] > max {
			max = tab[i]
		}
	}
	return max
}

func Abs(entier int) int {
	if entier < 0 {
		return -entier
	} else {
		return entier
	}
}

func InString(l []string, elt string) bool {
	for _, elt := range l {
		if elt == elt {
			return true
		}
	}
	return false
}
