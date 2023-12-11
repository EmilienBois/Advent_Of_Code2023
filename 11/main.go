package main

import (
	"fmt"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Pos struct {
	X int
	Y int
}

type Pair struct {
	pos1 Pos
	pos2 Pos
}

type Grid map[Pos]bool

func Part1(s string) int {
	test, nb_line := utils.Read_file(s)
	nb_col := len(test[0])
	grid := make(Grid)

	for i := 0; i < nb_line; i++ {
		this := test[i]
		for j := 0; j < nb_col; j++ {
			if this[j] == '#' {
				grid[Pos{i, j}] = true
			}
		}
	}

	// Expansion of the universe
	expanded_grid, new_line, new_col := Expand(grid, nb_line, nb_col, 1)

	var liste_stars []Pos
	for line := 0; line < new_line; line++ {
		for colomns := 0; colomns < new_col; colomns++ {
			pos := Pos{line, colomns}
			if expanded_grid[pos] == true {
				liste_stars = append(liste_stars, pos)
			}
		}
	}

	var liste_pairs []Pair
	for i := 0; i < len(liste_stars); i++ {
		for j := i + 1; j < len(liste_stars); j++ {
			liste_pairs = append(liste_pairs, Pair{liste_stars[i], liste_stars[j]})
		}
	}

	tot_dist := 0
	for pair := 0; pair < len(liste_pairs); pair++ {
		dist := Dist(liste_pairs[pair])
		tot_dist += dist
	}
	return tot_dist
}

func Dist(pair Pair) int {
	dist_x := utils.Abs(pair.pos1.X - pair.pos2.X)
	dist_y := utils.Abs(pair.pos1.Y - pair.pos2.Y)
	return dist_x + dist_y
}

func Expand(grid Grid, nb_line, nb_col, expansion int) (Grid, int, int) {
	empty_lines := Get_empty_lines(grid, nb_line, nb_col)
	empty_colons := Get_empty_colons(grid, nb_line, nb_col)
	new_grid := make(Grid)
	for line := 0; line < nb_line; line++ {
		for colomns := 0; colomns < nb_col; colomns++ {
			pos := Pos{line, colomns}
			if grid[pos] == true {
				line, col := Nb_empty_before(pos, empty_lines, empty_colons)
				newpos := Pos{pos.X + line*expansion, pos.Y + col*expansion}
				new_grid[newpos] = true
			}
		}
	}
	return new_grid, nb_line + len(empty_lines)*expansion, nb_col + len(empty_colons)*expansion
}

func Nb_empty_before(pos Pos, empty_lines, empty_col []int) (int, int) {
	line := 0
	col := 0
	for _, elt := range empty_lines {
		if elt < pos.X {
			line += 1
		} else {
			break
		}
	}
	for _, elt := range empty_col {
		if elt < pos.Y {
			col += 1
		} else {
			break
		}
	}
	return line, col
}

func Get_empty_lines(grid Grid, nb_line, nb_col int) []int {
	var empty_lines []int
	for line := 0; line < nb_line; line++ {
		empty := true
		for colomns := 0; colomns < nb_col; colomns++ {
			if grid[Pos{line, colomns}] == true {
				empty = false
				break
			}
		}
		if empty {
			empty_lines = append(empty_lines, line)
		}
	}
	return empty_lines
}

func Get_empty_colons(grid Grid, nb_line, nb_col int) []int {
	var empty_col []int
	for colomns := 0; colomns < nb_col; colomns++ {
		empty := true
		for line := 0; line < nb_col; line++ {
			if grid[Pos{line, colomns}] == true {
				empty = false
				break
			}
		}
		if empty {
			empty_col = append(empty_col, colomns)
		}
	}
	return empty_col
}
func Part2(s string, expansion int) int {
	test, nb_line := utils.Read_file(s)
	nb_col := len(test[0])
	grid := make(Grid)

	for i := 0; i < nb_line; i++ {
		this := test[i]
		for j := 0; j < nb_col; j++ {
			if this[j] == '#' {
				grid[Pos{i, j}] = true
			}
		}
	}
	var liste_stars []Pos
	for line := 0; line < nb_line; line++ {
		for colomns := 0; colomns < nb_col; colomns++ {
			pos := Pos{line, colomns}
			if grid[pos] == true {
				liste_stars = append(liste_stars, pos)
			}
		}
	}
	var expanded_liste_stars []Pos
	empty_lines := Get_empty_lines(grid, nb_line, nb_col)
	empty_colons := Get_empty_colons(grid, nb_line, nb_col)
	for _, pos := range liste_stars {
		line, col := Nb_empty_before(pos, empty_lines, empty_colons)
		newpos := Pos{pos.X + line*(expansion-1), pos.Y + col*(expansion-1)}
		expanded_liste_stars = append(expanded_liste_stars, newpos)
	}

	var liste_pairs []Pair
	for i := 0; i < len(expanded_liste_stars); i++ {
		for j := i + 1; j < len(expanded_liste_stars); j++ {
			liste_pairs = append(liste_pairs, Pair{expanded_liste_stars[i], expanded_liste_stars[j]})
		}
	}
	tot_dist := 0
	for pair := 0; pair < len(liste_pairs); pair++ {
		dist := Dist(liste_pairs[pair])
		tot_dist += (dist)
	}
	return tot_dist
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

	result_two_test := Part2("inputtest.txt", 10)
	println("The result of the input test for the part 2 is : ", result_two_test)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two_test = Part2("inputtest.txt", 100)
	println("The result of the input test for the part 2 is : ", result_two_test)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two := Part2("input.txt", 1000000)
	println("The result of the input for the part 2 is : ", result_two)

	end = time.Now()
	fmt.Printf("%s \n", end.Sub(start))
}
