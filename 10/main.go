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

func Part1(s string) int {
	file, nb_line := utils.Read_file(s)
	nb_col := len(file[0])
	var initial Pos
	grid := make(Grid)
	for i := 0; i < nb_line; i++ {
		for j := 0; j < nb_col; j++ {
			pos := Pos{i, j}
			if file[i][j] == 'S' {
				initial = pos
			}
			grid[pos] = string(file[i][j])
		}
	}
	in_neighbors := Neighbors(initial)
	for _, n := range in_neighbors {
		dir := Get_dir(grid[n])
		dir = append(dir, "q")
	}

	return 0
}

func Get_dir(s string) []string {
	if s == "F" {
		return []string{"E", "S"}
	} else if s == "J" {
		return []string{"N", "W"}
	} else if s == "7" {
		return []string{"W", "S"}
	} else if s == "|" {
		return []string{"N", "S"}
	} else if s == "-" {
		return []string{"W", "E"}
	} else if s == "L" {
		return []string{"N", "E"}
	} else {
		println("demande de direction d'un point")
	}
	return []string{}
}

func Neighbors(pos Pos) []Pos {
	return []Pos{{pos.X + 1, pos.Y}, {pos.X, pos.Y + 1}, {pos.X - 1, pos.Y}, {pos.X, pos.Y - 1}}
}

func Part2(s string) int {
	return 0
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
