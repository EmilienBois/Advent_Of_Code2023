package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Step struct {
	dist      int
	direction string
}

type Pos struct {
	X, Y int
}

type Grid map[Pos]bool

func Part1(s string) int {
	file, _ := utils.Read_file(s)

	nb_line := 0
	nb_col := 0

	var liste_steps []Step
	for _, line := range file {
		str := strings.Split(line, " ")
		liste_steps = append(liste_steps, Step{utils.String_to_Int(str[1]), str[0]})
	}
	grid := make(Grid)
	perimeter := 0
	actual := Pos{0, 0}
	grid[actual] = true

	for _, step := range liste_steps {
		for i := 0; i < step.dist; i++ {
			if step.direction == "R" {
				actual = Pos{actual.X, actual.Y + 1}
				if actual.Y+1 > nb_col {
					nb_col = actual.Y + 1
				}

			} else if step.direction == "L" {
				actual = Pos{actual.X, actual.Y - 1}
			} else if step.direction == "D" {
				actual = Pos{actual.X + 1, actual.Y}
				if actual.X+1 > nb_line {
					nb_line = actual.X + 1
				}
			} else if step.direction == "U" {
				actual = Pos{actual.X - 1, actual.Y}
			}
			grid[actual] = true
			perimeter += 1
		}
	}
	//Now we have the total perimeter and it is drawn
	// Let's find the surface
	interior := Surface(grid, nb_line, nb_col)
	return interior + perimeter
}

func Surface(grid Grid, nb_line, nb_col int) int {
	interior := 0
	for line := 0; line < nb_line; line++ {
		in := false
		first := ""
		for col := 0; col < nb_col; col++ {
			pos := Pos{line, col}
			if in {
				if grid[pos] == true {
					if grid[Pos{pos.X, pos.Y + 1}] == true {
						if first == "" {
							if grid[Pos{pos.X - 1, pos.Y}] {
								first = "up"
							} else if grid[Pos{pos.X + 1, pos.Y}] {
								first = "down"
							}
						}
					} else {
						if grid[Pos{pos.X - 1, pos.Y}] && grid[Pos{pos.X + 1, pos.Y}] {
							in = false
						} else if grid[Pos{pos.X - 1, pos.Y}] == true {
							if first == "down" {
								in = false
								first = ""
							} else if first == "up" {
								first = ""
							}

						} else if grid[Pos{pos.X + 1, pos.Y}] == true {
							if first == "down" {
								first = ""
							} else if first == "up" {
								in = false
								first = ""
							}
						}
					}
				} else {
					interior += 1
				}
			} else {
				if grid[pos] == true {
					if grid[Pos{pos.X, pos.Y + 1}] == true {
						if first == "" {
							if grid[Pos{pos.X - 1, pos.Y}] {
								first = "up"
							} else if grid[Pos{pos.X + 1, pos.Y}] {
								first = "down"
							}
						}
					} else {
						if grid[Pos{pos.X - 1, pos.Y}] && grid[Pos{pos.X + 1, pos.Y}] {
							in = true
						} else if grid[Pos{pos.X - 1, pos.Y}] == true {
							if first == "up" {
								first = ""
							} else if first == "down" {
								in = true
								first = ""
							}
						} else if grid[Pos{pos.X + 1, pos.Y}] == true {
							if first == "up" {
								in = true
								first = ""
							} else if first == "down" {
								first = ""
							}
						}
					}
				}
			}
		}
	}
	return interior
}

func Part2(s string) int {
	return 0
}

func (grid Grid) PrintGrid(nb_line, nb_col int) {
	for line := 0; line < nb_line; line++ {
		for col := 0; col < nb_col; col++ {
			if grid[Pos{line, col}] == false {
				print(".")
			} else {
				print("#")
			}
		}
		print("\n")
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
