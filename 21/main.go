package main

import (
	"fmt"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Pos struct {
	X, Y int
}

type Grid map[Pos]bool

type Reach map[Pos]int

func Part1(nb_step int, s string) int {
	file, size := utils.Read_file(s)
	nb_col := len(file[0])
	grid := make(Grid)
	var initial Pos
	for i := 0; i < size; i++ {
		this := file[i]
		for j := 0; j < len(this); j++ {
			pos := Pos{i, j}
			if this[j] == 'S' {
				initial = pos
			} else if this[j] != '.' {
				grid[pos] = true
			}
		}
	}
	reach := make(Reach)
	reach[initial] = 2
	var reach_list []Pos
	reach_list = append(reach_list, initial)

	for step := 1; step <= nb_step; step++ {
		for _, elt := range reach_list {
			n := neighbors(elt)
			for _, neighbor := range n {
				if !InList(reach_list, neighbor) {
					if neighbor.X < 0 || neighbor.Y < 0 || neighbor.Y >= nb_col || neighbor.X >= size {

					} else {
						if grid[neighbor] {

						} else {
							reach[neighbor] = step
							reach_list = append(reach_list, neighbor)
						}
					}
				}
			}

		}
	}
	total := 0
	parity := nb_step % 2
	for i := 0; i < size; i++ {
		for j := 0; j < nb_col; j++ {
			pos := Pos{i, j}
			if reach[pos] != 0 {
				if reach[pos]%2 == parity {
					total += 1
				}
			}
		}
	}
	return total
}

func neighbors(pos Pos) []Pos {
	return []Pos{{pos.X + 1, pos.Y}, {pos.X, pos.Y + 1}, {pos.X - 1, pos.Y}, {pos.X, pos.Y - 1}}
}

func InList(list []Pos, t_elt Pos) bool {
	for _, elt := range list {
		if elt == t_elt {
			return true
		}
	}
	return false
}

func Part2(nb_step int, s string) int {
	file, size := utils.Read_file(s)
	nb_col := len(file[0])
	grid := make(Grid)
	var initial Pos
	for i := 0; i < size; i++ {
		this := file[i]
		for j := 0; j < len(this); j++ {
			pos := Pos{i, j}
			if this[j] == 'S' {
				initial = pos
			} else if this[j] != '.' {
				grid[pos] = true
			}
		}
	}
	reach := make(Reach)
	reach[initial] = 2
	var reach_list []Pos
	reach_list = append(reach_list, initial)
	no_change := false
	step := 0
	for !no_change {
		step += 1
		no_change = true
		for _, elt := range reach_list {
			n := neighbors(elt)
			for _, neighbor := range n {
				if !InList(reach_list, neighbor) {
					if neighbor.X < 0 || neighbor.Y < 0 || neighbor.Y >= nb_col || neighbor.X >= size {

					} else {
						no_change = false
						if grid[neighbor] {

						} else {
							reach[neighbor] = step
							reach_list = append(reach_list, neighbor)
						}
					}
				}
			}

		}
	}
	total := 0
	parity := nb_step % 2
	for i := 0; i < size; i++ {
		for j := 0; j < nb_col; j++ {
			pos := Pos{i, j}
			if reach[pos] != 0 {
				if reach[pos]%2 == parity {
					total += 1
				}
			}
		}
	}
	return total
}

func main() {
	start := time.Now()
	result_one_test := Part1(6, "inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)

	end := time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_one := Part1(64, "input.txt")
	println("The result of the input for the part 1 is : ", result_one)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two_test := Part2(2000, "inputtest.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two := Part2(2000, "input.txt")
	println("The result of the input for the part 2 is : ", result_two)

	end = time.Now()
	fmt.Printf("%s \n", end.Sub(start))
}
