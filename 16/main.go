package main

import (
	"fmt"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Pos struct {
	X, Y int
}

type Grid map[Pos]string

type Point struct {
	pos       Pos
	direction string
}

func Part1(s string) int {
	first := Point{Pos{0, 0}, "right"}
	return Nb_energized(s, first)
}

func Nb_energized(s string, first Point) int {
	file, size := utils.Read_file(s)

	grid := make(Grid)
	for i := 0; i < size; i++ {
		this := file[i]
		for j := 0; j < len(this); j++ {
			pos := Pos{i, j}
			if this[j] != '.' {
				grid[pos] = string(this[j]) + "nd"
			}
		}
	}
	//now we have the position of non empty
	//We initialize the list of point
	var liste_points []Point

	liste_points = append(liste_points, first)
	notend := true
	times_no_change := 0
	for notend {
		notend = false
		change := 0
		for i_point := 0; i_point < len(liste_points); i_point++ {
			this := liste_points[i_point]
			actual_pos := this.pos
			do := true
			if this.pos.Y == len(file[0]) || this.pos.Y == -1 || this.pos.X == -1 || this.pos.X == size {
				do = false
			} else {
				notend = true
			}
			if do {
				if grid[actual_pos] == "" || grid[actual_pos] == "#" {
					switch this.direction {
					case "right":
						liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y + 1}
					case "left":
						liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y - 1}
					case "down":
						liste_points[i_point].pos = Pos{this.pos.X + 1, this.pos.Y}
					case "up":
						liste_points[i_point].pos = Pos{this.pos.X - 1, this.pos.Y}
					}
					if grid[actual_pos] == "" {
						change += 1
					}
					grid[actual_pos] = "#"
				} else {
					if grid[actual_pos] == "/" || grid[actual_pos] == "/nd" {
						grid[actual_pos] = "/"
						switch this.direction {
						case "right":
							liste_points[i_point].pos = Pos{this.pos.X - 1, this.pos.Y}
							liste_points[i_point].direction = "up"
						case "left":
							liste_points[i_point].pos = Pos{this.pos.X + 1, this.pos.Y}
							liste_points[i_point].direction = "down"
						case "down":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y - 1}
							liste_points[i_point].direction = "left"
						case "up":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y + 1}
							liste_points[i_point].direction = "right"
						}

					}
					if grid[actual_pos] == "\\" || grid[actual_pos] == "\\nd" {
						grid[actual_pos] = "\\"
						switch this.direction {
						case "right":
							liste_points[i_point].pos = Pos{this.pos.X + 1, this.pos.Y}
							liste_points[i_point].direction = "down"
						case "left":
							liste_points[i_point].pos = Pos{this.pos.X - 1, this.pos.Y}
							liste_points[i_point].direction = "up"
						case "down":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y + 1}
							liste_points[i_point].direction = "right"
						case "up":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y - 1}
							liste_points[i_point].direction = "left"
						}

					}
					if grid[actual_pos] == "|" || grid[actual_pos] == "|nd" {
						grid[actual_pos] = "|"
						switch this.direction {
						case "right":
							liste_points[i_point].pos = Pos{this.pos.X - 1, this.pos.Y}
							liste_points[i_point].direction = "up"
							new_pos := Pos{this.pos.X + 1, this.pos.Y}
							new_direction := "down"
							liste_points = append(liste_points, Point{new_pos, new_direction})
						case "left":
							liste_points[i_point].pos = Pos{this.pos.X - 1, this.pos.Y}
							liste_points[i_point].direction = "up"
							new_pos := Pos{this.pos.X + 1, this.pos.Y}
							new_direction := "down"
							liste_points = append(liste_points, Point{new_pos, new_direction})
						case "down":
							liste_points[i_point].pos = Pos{this.pos.X + 1, this.pos.Y}
						case "up":
							liste_points[i_point].pos = Pos{this.pos.X - 1, this.pos.Y}
						}
					}
					if grid[actual_pos] == "-" || grid[actual_pos] == "-nd" {
						grid[actual_pos] = "-"
						switch this.direction {
						case "right":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y + 1}
						case "left":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y - 1}
						case "down":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y - 1}
							liste_points[i_point].direction = "left"
							new_pos := Pos{this.pos.X, this.pos.Y + 1}
							new_direction := "right"
							liste_points = append(liste_points, Point{new_pos, new_direction})
						case "up":
							liste_points[i_point].pos = Pos{this.pos.X, this.pos.Y - 1}
							liste_points[i_point].direction = "left"
							new_pos := Pos{this.pos.X, this.pos.Y + 1}
							new_direction := "right"
							liste_points = append(liste_points, Point{new_pos, new_direction})
						}
					}
				}
			}

		}
		if change == 0 {
			times_no_change += 1
		} else {
			times_no_change = 0
		}
		if times_no_change >= 3 {
			notend = false
		}
	}
	total := 0
	for line := 0; line < size; line++ {
		for col := 0; col < len(file[0]); col++ {
			pos := Pos{line, col}
			if grid[pos] == "#" || grid[pos] == "|" || grid[pos] == "-" || grid[pos] == "\\" || grid[pos] == "/" {
				total += 1
			}
		}
	}
	return total
}

func (grid Grid) PrintGrid(nb_line, nb_col int) {
	for line := 0; line < nb_line; line++ {
		for col := 0; col < nb_col; col++ {
			if grid[Pos{line, col}] == "" {
				print(".")
			} else {
				print(grid[Pos{line, col}])
			}
		}
		print("\n")
	}
}

func Part2(s string) int {
	max := 0
	// pour la première ligne
	for col := 0; col < 110; col++ {
		pos := Pos{0, col}
		direction := "down"
		nb := Nb_energized(s, Point{pos, direction})
		if nb > max {
			max = nb
		}
	}
	for col := 0; col < 110; col++ {
		pos := Pos{109, col}
		direction := "up"
		nb := Nb_energized(s, Point{pos, direction})
		if nb > max {
			max = nb
		}
	}
	for line := 0; line < 110; line++ {
		pos := Pos{line, 0}
		direction := "right"
		nb := Nb_energized(s, Point{pos, direction})
		if nb > max {
			max = nb
		}
	}
	for line := 0; line < 110; line++ {
		pos := Pos{line, 109}
		direction := "left"
		nb := Nb_energized(s, Point{pos, direction})
		if nb > max {
			max = nb
		}
	}
	return max
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
