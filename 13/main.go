package main

import (
	"fmt"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Pos struct {
	X, Y int
}

type Pattern struct {
	grid            Grid
	nb_line, nb_col int
}

type Grid map[Pos]bool

func Part1(s string) int {
	file, size := utils.Read_file(s)

	var liste_pattern []Pattern
	grid := make(Grid)
	Y := 1
	for i := 0; i < size; i++ {
		if len(file[i]) == 0 {
			X := len(file[i-1])
			liste_pattern = append(liste_pattern, Pattern{grid, Y - 1, X})
			grid = make(Grid)
			Y = 1
		} else {
			this := file[i]
			for X, elt := range this {
				pos := Pos{X + 1, Y}
				if elt == '#' {
					grid[pos] = true
				} else if elt != '.' {
					println("Ce n'est pas un symbole connu qui est lu")
				}
			}
			Y += 1
		}
		if i == size-1 {
			X := len(file[i-1])
			liste_pattern = append(liste_pattern, Pattern{grid, Y - 1, X})
			grid = make(Grid)
			Y = 1
		}
	}
	total := 0
	index_done := make(map[int]bool)
	//We check for vertical mirror
	var list_vertical []int
	for i, p := range liste_pattern {
		found := false
		for !found {

			ok := true
			for mirror := 1; mirror < p.nb_col; mirror++ { // For every colomn

				//We check if a line is okay
				ok = check_vertical(p, mirror)
				if ok {
					found = true
					total += mirror
					list_vertical = append(list_vertical, mirror)
					mirror = p.nb_col + 1
					index_done[i] = true
				}
			}
			found = true
		}
	}
	var list_horizontal []int
	for i, p := range liste_pattern {
		found := false
		if index_done[i] == true {
			found = true
		}
		for !found {

			ok := true
			for mirror := 1; mirror <= p.nb_line; mirror++ { // For every line

				//We check if a line is okay
				ok = check_horizontal(p, mirror)
				if ok {
					found = true
					total += 100 * mirror
					list_horizontal = append(list_horizontal, mirror)
					mirror = p.nb_line + 1
				}
			}
		}
	}

	return total
}
func check_horizontal(p Pattern, mirror int) bool {
	for col := 1; col <= p.nb_line; col++ {
		for dist := 0; dist <= mirror; dist++ {
			if (mirror-dist) >= 1 && mirror+dist+1 <= p.nb_line {
				pos1 := Pos{col, mirror - dist}
				pos2 := Pos{col, mirror + dist + 1}
				if p.grid[pos1] != p.grid[pos2] {
					return false
				}
			}
		}
	}
	return true
}
func check_vertical(p Pattern, mirror int) bool {
	for line := 1; line <= p.nb_line; line++ {
		for dist := 0; dist <= mirror; dist++ {
			if (mirror-dist) >= 1 && mirror+dist+1 <= p.nb_col {
				pos1 := Pos{mirror - dist, line}
				pos2 := Pos{mirror + dist + 1, line}
				if p.grid[pos1] != p.grid[pos2] {
					return false
				}
			}
		}
	}
	return true
}

func Part2(s string) int {
	file, size := utils.Read_file(s)

	var liste_pattern []Pattern
	grid := make(Grid)
	Y := 1
	for i := 0; i < size; i++ {
		if len(file[i]) == 0 {
			X := len(file[i-1])
			liste_pattern = append(liste_pattern, Pattern{grid, Y - 1, X})
			grid = make(Grid)
			Y = 1
		} else {
			this := file[i]
			for X, elt := range this {
				pos := Pos{X + 1, Y}
				if elt == '#' {
					grid[pos] = true
				} else if elt != '.' {
					println("Ce n'est pas un symbole connu qui est lu")
				}
			}
			Y += 1
		}
		if i == size-1 {
			X := len(file[i-1])
			liste_pattern = append(liste_pattern, Pattern{grid, Y - 1, X})
			grid = make(Grid)
			Y = 1
		}
	}

	total := 0
	index_done := make(map[int]bool)
	//We check for vertical mirror
	var list_vertical []int
	for i, p := range liste_pattern {
		found := false
		for !found {

			ok := true
			for mirror := 1; mirror < p.nb_col; mirror++ { // For every colomn

				//We check if a line is okay
				ok = check_vertical2(p, mirror)
				if ok {
					found = true
					total += mirror
					list_vertical = append(list_vertical, mirror)
					mirror = p.nb_col + 1
					index_done[i] = true
				}
			}
			found = true
		}
	}
	println("liste verticale reached")
	var list_horizontal []int
	for i, p := range liste_pattern {
		found := false
		if index_done[i] == true {
			found = true
		}
		for !found {

			ok := true
			for mirror := 1; mirror <= p.nb_line; mirror++ { // For every line

				//We check if a line is okay
				ok = check_horizontal2(p, mirror)
				if ok {
					found = true
					total += 100 * mirror
					list_horizontal = append(list_horizontal, mirror)
					mirror = p.nb_line + 1
				}
			}
		}
		if i == 48 {
			println("pattern ", i, " is done ")
		}

	}
	println("liste horizontale reached")
	return total
}
func check_horizontal2(p Pattern, mirror int) bool {
	nb_errors := 0
	for col := 1; col <= p.nb_line; col++ {
		for dist := 0; dist <= mirror; dist++ {
			if (mirror-dist) >= 1 && mirror+dist+1 <= p.nb_line {
				pos1 := Pos{col, mirror - dist}
				pos2 := Pos{col, mirror + dist + 1}
				if p.grid[pos1] != p.grid[pos2] {
					nb_errors += 1
				}
			}
		}
	}
	if nb_errors == 1 {
		return true
	} else {
		return false
	}
}
func check_vertical2(p Pattern, mirror int) bool {
	nb_errors := 0
	for line := 1; line <= p.nb_line; line++ {
		for dist := 0; dist <= mirror; dist++ {
			if (mirror-dist) >= 1 && mirror+dist+1 <= p.nb_col {
				pos1 := Pos{mirror - dist, line}
				pos2 := Pos{mirror + dist + 1, line}
				if p.grid[pos1] != p.grid[pos2] {
					nb_errors += 1
				}
			}
		}
	}
	if nb_errors == 1 {
		return true
	} else {
		return false
	}
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
