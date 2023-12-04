package utils

import (
	"bufio"
	"os"
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

func Somme(liste []int) int {
	somme := 0
	for _, elt := range liste {
		somme += elt
	}
	return somme
}
