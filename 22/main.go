package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Pos struct {
	X, Y, Z int
}

type Brick struct {
	pos1, pos2  Pos
	standing_on []Brick
}

type Grid map[Pos]bool

func MakeGridFromBricks(liste map[int][]Brick, nb_floor int) Grid {
	new_grid := make(Grid)
	for i := 1; i <= nb_floor; i++ {
		for _, bricks := range liste[i] {
			if bricks.pos1.Z != -1 {
				for x := bricks.pos1.X; x <= bricks.pos2.X; x++ {
					for y := bricks.pos1.Y; y <= bricks.pos2.Y; y++ {
						for z := bricks.pos1.Z; z <= bricks.pos2.Z; z++ {
							pos := Pos{x, y, z}
							new_grid[pos] = true
						}
					}
				}
			}
		}
	}
	return new_grid
}

func UpgradeGrid(grid Grid, ancient_brick, changed_brick Brick) Grid {

	for x := ancient_brick.pos1.X; x <= ancient_brick.pos2.X; x++ {
		for y := ancient_brick.pos1.Y; y <= ancient_brick.pos2.Y; y++ {
			for z := ancient_brick.pos1.Z; z <= ancient_brick.pos2.Z; z++ {
				pos := Pos{x, y, z}
				grid[pos] = false
			}
		}
	}
	for x := changed_brick.pos1.X; x <= changed_brick.pos2.X; x++ {
		for y := changed_brick.pos1.Y; y <= changed_brick.pos2.Y; y++ {
			for z := changed_brick.pos1.Z; z <= changed_brick.pos2.Z; z++ {
				pos := Pos{x, y, z}
				grid[pos] = true
			}
		}
	}
	return grid
}

func (brick1 Brick) IsOver(brick2 Brick) bool {
	if brick1.pos1.Z != -1 {
		maxx1 := max(brick1.pos1.X, brick1.pos2.X)
		maxy1 := max(brick1.pos1.Y, brick1.pos2.Y)
		maxx2 := max(brick2.pos1.X, brick2.pos2.X)
		maxy2 := max(brick2.pos1.Y, brick2.pos2.Y)
		maxZ2 := max(brick2.pos1.Z, brick2.pos2.Z)

		minx1 := min(brick1.pos1.X, brick1.pos2.X)
		miny1 := min(brick1.pos1.Y, brick1.pos2.Y)
		minZ1 := min(brick1.pos1.Z, brick1.pos2.Z)
		minx2 := min(brick2.pos1.X, brick2.pos2.X)
		miny2 := min(brick2.pos1.Y, brick2.pos2.Y)

		if !(minZ1 == maxZ2+1) {
			return false
		}
		if maxx1 < minx2 {
			return false
		}
		if minx1 > maxx2 {
			return false
		}
		if maxy1 < miny2 {
			return false
		}
		if miny1 > maxy2 {
			return false
		}
		return true
	}
	return false
}

func NumberHeightEmpty(grid Grid, brick Brick) int {
	z := brick.pos1.Z - 1
	empty := true
	for empty {
		for x := brick.pos1.X; x <= brick.pos2.X; x++ {
			for y := brick.pos1.Y; y <= brick.pos2.Y; y++ {
				if grid[Pos{x, y, z}] {
					empty = false
				}
			}
		}
		if empty {
			z -= 1
		}
	}
	return brick.pos1.Z - z - 1
}

func InListe(brick Brick, list []Brick) bool {
	for _, elt := range list {
		if brick.pos1 == elt.pos1 && brick.pos2 == elt.pos2 {
			return true
		}
	}
	return false
}

func Part1(s string) int {
	file, nb_line := utils.Read_file(s)

	liste_bricks_ord := make(map[int][]Brick)
	nb_floor := 0

	for i := 0; i < nb_line; i++ {
		str := strings.Split(file[i], "~")
		coord1 := strings.Split(str[0], ",")
		coord2 := strings.Split(str[1], ",")
		X1 := utils.String_to_Int(coord1[0])
		Y1 := utils.String_to_Int(coord1[1])
		Z1 := utils.String_to_Int(coord1[2])

		X2 := utils.String_to_Int(coord2[0])
		Y2 := utils.String_to_Int(coord2[1])
		Z2 := utils.String_to_Int(coord2[2])
		pos1 := Pos{min(X1, X2), min(Y1, Y2), min(Z1, Z2)}
		pos2 := Pos{max(X1, X2), max(Y1, Y2), max(Z1, Z2)}

		brick := Brick{pos1, pos2, []Brick{}}

		minZ := min(Z1, Z2)
		liste_bricks_ord[minZ] = append(liste_bricks_ord[minZ], brick)
		if minZ > nb_floor {
			nb_floor = minZ
		}

	}

	print("reached")
	//Now we make them fall
	grid := MakeGridFromBricks(liste_bricks_ord, nb_floor)

	for floor := 2; floor <= nb_floor; floor++ {
		for b := range liste_bricks_ord[floor] {
			brick := liste_bricks_ord[floor][b]
			z := NumberHeightEmpty(grid, brick)
			brick.pos1.Z -= z
			brick.pos2.Z -= z
			new_floor := floor - z
			liste_bricks_ord[new_floor] = append(liste_bricks_ord[new_floor], brick)
			liste_bricks_ord[floor][b].pos1.Z = -1
			grid = UpgradeGrid(grid, liste_bricks_ord[floor][b], brick)
		}
	}
	// Now we check how many is under
	total := 0
	var checked_bricks []Brick
	for floor := 2; floor <= nb_floor; floor++ {
		for _, brick1 := range liste_bricks_ord[floor] {
			nb := 0
			for _, brick2 := range liste_bricks_ord[floor-1] {
				if brick1.IsOver(brick2) {
					nb += 1
					brick1.standing_on = append(brick1.standing_on, brick2)
				}
			}
			if nb == 0 {
				println("il y a eu un problème lors de la retombée de la brick")
			}
			if nb >= 2 {
				for _, brick2 := range brick1.standing_on {
					if !InListe(brick2, checked_bricks) {
						total += 1
						checked_bricks = append(checked_bricks, brick2)
					}

				}

			}
		}
	}

	return total
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
