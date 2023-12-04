package utils

import (
	"bufio"
	"os"
)

// Lecture de fichier
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
