package main

import (
	"fmt"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Pos struct {
	X, Y int
}

type Grid map[Pos]int

func (grid Grid) PrintGrid(nb_line, nb_col int) {
	for line := 0; line < nb_line; line++ {
		for col := 0; col < nb_col; col++ {
			if grid[Pos{line, col}] == 0 {
				print(".")
			} else if grid[Pos{line, col}] == 1 {
				print("O")
			} else {
				print("#")
			}
		}
		print("\n")
	}
}

func Part1(s string) int {
	file, size := utils.Read_file(s)
	// Now let's have
	grille := make(Grid)

	for i := 0; i < size; i++ {
		for j := 0; j < len(file[0]); j++ {
			this := file[i][j]
			pos := Pos{i, j}
			if this == 'O' {
				grille[pos] = 1
			} else if this == '#' {
				grille[pos] = 2
			}
		}
	}
	grille.PrintGrid(len(file[0]), size)
	//Now we have the whole grid with 0 for ., 1 for # and 2 for O

	//Let's tilt the 0 to the top
	grille = Tilt_up(grille, len(file[0]), size)

	//Let's count now
	total := 0
	for line := 0; line < size; line++ {
		value := size - line
		for col := 0; col < len(file[0]); col++ {
			if grille[Pos{line, col}] == 1 {
				total += value
			}
		}
	}
	return total
}

func Tilt_up(grid Grid, nb_col, nb_line int) Grid {
	new_grid := make(Grid)
	for col := 0; col < nb_col; col++ {
		for line := 0; line < nb_line; line++ {
			pos := Pos{line, col}
			if grid[pos] == 1 {
				new_line := line
				for new_line > 0 {
					if (new_grid[Pos{new_line - 1, col}] == 0) && (grid[Pos{new_line - 1, col}] != 2) {
						new_line -= 1
					} else {
						break
					}
				}
				new_grid[Pos{new_line, col}] = 1
			}
			if grid[pos] == 2 {
				new_grid[pos] = 2
			}
		}
	}
	return new_grid
}
func Tilt_down(grid Grid, nb_col, nb_line int) Grid {
	new_grid := make(Grid)
	for col := nb_col - 1; col >= 0; col-- {
		for line := nb_line - 1; line >= 0; line-- {
			pos := Pos{line, col}
			if grid[pos] == 1 {
				new_line := line
				for new_line < nb_line-1 {
					if (new_grid[Pos{new_line + 1, col}] == 0) && (grid[Pos{new_line + 1, col}] != 2) {
						new_line += 1
					} else {
						break
					}
				}
				new_grid[Pos{new_line, col}] = 1
			}
			if grid[pos] == 2 {
				new_grid[pos] = 2
			}
		}
	}
	return new_grid
}

func Tilt_left(grid Grid, nb_col, nb_line int) Grid {
	new_grid := make(Grid)
	for line := 0; line < nb_line; line++ {
		for col := 0; col < nb_col; col++ {
			pos := Pos{line, col}
			if grid[pos] == 1 {
				new_col := col
				for new_col > 0 {
					if (new_grid[Pos{line, new_col - 1}] == 0) && (grid[Pos{line, new_col - 1}] != 2) {
						new_col -= 1
					} else {
						break
					}
				}
				new_grid[Pos{col, new_col}] = 1
			}
			if grid[pos] == 2 {
				new_grid[pos] = 2
			}
		}
	}
	return new_grid
}

func Tilt_right(grid Grid, nb_col, nb_line int) Grid {
	new_grid := make(Grid)
	for line := nb_line - 1; line >= 0; line-- {
		for col := nb_col - 1; col >= 0; col-- {
			pos := Pos{line, col}
			if grid[pos] == 1 {
				new_col := col
				for new_col < nb_col-1 {
					if (new_grid[Pos{line, new_col + 1}] == 0) && (grid[Pos{line, new_col + 1}] != 2) {
						new_col += 1
					} else {
						break
					}
				}
				new_grid[Pos{col, new_col}] = 1
			}
			if grid[pos] == 2 {
				new_grid[pos] = 2
			}
		}
	}
	return new_grid
}

func Part2(s string) int {
	file, size := utils.Read_file(s)
	// Now let's have
	grille := make(Grid)

	for i := 0; i < size; i++ {
		for j := 0; j < len(file[0]); j++ {
			this := file[i][j]
			pos := Pos{i, j}
			if this == 'O' {
				grille[pos] = 1
			} else if this == '#' {
				grille[pos] = 2
			}
		}
	}
	//Now we have the whole grid with 0 for ., 1 for # and 2 for O

	//Let's tilt the 0 to the top
	for i := 0; i < 1; i++ {
		grille = Cycle(grille, len(file[0]), size)
	}
	grille.PrintGrid(len(file[0]), size)

	//Let's count now
	total := 0
	for line := 0; line < size; line++ {
		value := size - line
		for col := 0; col < len(file[0]); col++ {
			if grille[Pos{line, col}] == 1 {
				total += value
			}
		}
	}
	return total
}

func Cycle(grille Grid, nb_col, nb_line int) Grid {
	grille = Tilt_up(grille, nb_col, nb_line)
	grille = Tilt_left(grille, nb_col, nb_line)
	grille = Tilt_down(grille, nb_col, nb_line)
	grille = Tilt_right(grille, nb_col, nb_line)
	return grille
}

func main() {
	start := time.Now()
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)

	end := time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	//result_one := Part1("input.txt")
	//println("The result of the input for the part 1 is : ", result_one)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two_test := Part2("inputtest.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	//result_two := Part2("input.txt")
	//println("The result of the input for the part 2 is : ", result_two)

	end = time.Now()
	fmt.Printf("%s \n", end.Sub(start))
}
