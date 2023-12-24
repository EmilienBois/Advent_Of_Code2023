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

type Path struct {
	pos    Pos
	dir    string
	length int
}

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
	dir := Get_dir_S(initial, grid)
	path1 := Path{initial, dir[0], 0}
	path2 := Path{initial, dir[1], 0}
	path1 = Next(path1, grid)
	path2 = Next(path2, grid)
	for path1.pos != path2.pos {
		path1 = Next(path1, grid)
		if path1.pos == path2.pos {
			break
		}
		path2 = Next(path2, grid)
	}

	return path1.length
}

func Next(path Path, grid Grid) Path {

	if path.dir == "W" {
		path.pos.Y -= 1
		dir := Get_dir(grid[path.pos])
		if len(dir) == 0 {

		} else if dir[0] == "E" {
			path.dir = dir[1]
		} else if dir[1] == "E" {
			path.dir = dir[0]
		} else {
			println("Erreur dans le chemin")
		}
	} else if path.dir == "E" {
		path.pos.Y += 1
		dir := Get_dir(grid[path.pos])
		if len(dir) == 0 {

		} else if dir[0] == "W" {
			path.dir = dir[1]
		} else if dir[1] == "W" {
			path.dir = dir[0]
		} else {
			println("Erreur dans le chemin")
		}
	} else if path.dir == "N" {
		path.pos.X -= 1
		dir := Get_dir(grid[path.pos])
		if len(dir) == 0 {

		} else if dir[0] == "S" {
			path.dir = dir[1]
		} else if dir[1] == "S" {
			path.dir = dir[0]
		} else {
			println("Erreur dans le chemin")
		}
	} else if path.dir == "S" {
		path.pos.X += 1
		dir := Get_dir(grid[path.pos])
		if len(dir) == 0 {

		} else if dir[0] == "N" {
			path.dir = dir[1]
		} else if dir[1] == "N" {
			path.dir = dir[0]
		} else {
			println("Erreur dans le chemin")
		}
	}
	path.length += 1
	return path
}

func Get_dir_S(initial Pos, grid Grid) []string {
	var in_dir []string

	n := Neighbors(initial)
	south := grid[n[0]]
	if south == "J" || south == "|" || south == "L" {
		in_dir = append(in_dir, "S")
	}
	north := grid[n[2]]
	if north == "F" || north == "|" || north == "7" {
		in_dir = append(in_dir, "N")
	}
	east := grid[n[1]]
	if east == "J" || east == "-" || east == "7" {
		in_dir = append(in_dir, "E")
	}
	west := grid[n[3]]
	if west == "F" || west == "-" || west == "L" {
		in_dir = append(in_dir, "W")
	}
	return in_dir
}

func Get_dir(s string) []string {
	if s == "F" {
		return []string{"E", "S"}
	} else if s == "J" {
		return []string{"N", "W"}
	} else if s == "7" {
		return []string{"S", "W"}
	} else if s == "|" {
		return []string{"N", "S"}
	} else if s == "-" {
		return []string{"E", "W"}
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
	dir := Get_dir_S(initial, grid)
	path := Path{initial, dir[0], 0}

	var corner []Pos

	actual_dir := path.dir
	path = Next(path, grid)
	if actual_dir != path.dir {
		corner = append(corner, path.pos)
	}

	for path.pos != initial {
		if grid[path.pos] == "S" {
			break
		}
		actual_dir := path.dir
		if len(corner) == 11 {
			println()
		}
		path = Next(path, grid)
		if actual_dir != path.dir {
			corner = append(corner, path.pos)
		}
	}
	perimeter := path.length
	// Calculating the area, with Shoelace formula
	A := 0
	for i := range corner {
		var pos1, pos2 Pos
		if i == len(corner)-1 {
			pos1 = corner[i]
			pos2 = corner[0]
		} else {
			pos1 = corner[i]
			pos2 = corner[i+1]
		}
		A += (pos1.Y - pos2.Y) * (pos1.X + pos2.X)
	}
	A = A / 2

	//using Pick's theorem
	return A + 1 - perimeter/2

}
func main() {
	start := time.Now()
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)

	end := time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()
	result_one_test2 := Part1("inputtest2.txt")
	println("The result of the input test for the part 1 is : ", result_one_test2)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_one := Part1("input.txt")
	println("The result of the input for the part 1 is : ", result_one)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two_test := Part2("inputtestpart2.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two_test2 := Part2("inputtestpart22.txt")
	println("The result of the input test for the part 2 is : ", result_two_test2)

	end = time.Now()
	fmt.Printf("%s\n", end.Sub(start))
	start = time.Now()

	result_two := Part2("input.txt")
	println("The result of the input for the part 2 is : ", result_two)

	end = time.Now()
	fmt.Printf("%s \n", end.Sub(start))
}
