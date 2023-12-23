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

type QueueElement struct {
	p        Point
	priority int
}

func neighbors(point Point, grid Grid, nb_line, nb_col int) []Point {
	var n []Point
	pos := point.pos
	if grid[pos] == ">" {
		return []Point{{Pos{pos.X, pos.Y + 1}, "right"}}
	}
	if grid[pos] == "<" {
		return []Point{{Pos{pos.X, pos.Y - 1}, "left"}}
	}
	if grid[pos] == "v" {
		return []Point{{Pos{pos.X + 1, pos.Y}, "down"}}
	}
	if pos.X+1 < nb_line && grid[Pos{pos.X + 1, pos.Y}] != "#" {

		n = append(n, Point{Pos{pos.X + 1, pos.Y}, "down"})
	}
	if pos.Y+1 < nb_col && grid[Pos{pos.X, pos.Y + 1}] != "#" {

		n = append(n, Point{Pos{pos.X, pos.Y + 1}, "right"})
	}
	if pos.X-1 >= 0 && grid[Pos{pos.X - 1, pos.Y}] != "#" {
		n = append(n, Point{Pos{pos.X - 1, pos.Y}, "up"})
	}
	if pos.Y-1 >= 0 && grid[Pos{pos.X, pos.Y - 1}] != "#" {

		n = append(n, Point{Pos{pos.X, pos.Y - 1}, "left"})
	}
	return n
}

func InPoint(l []Point, elt Point) bool {
	for _, point := range l {
		if point == elt {
			return true
		}
	}
	return false
}

func heuristic(point1, point2 Point) int {
	return utils.Abs(point1.pos.X-point2.pos.X) + utils.Abs(point1.pos.Y-point2.pos.Y)
}

func Part1(s string) int {
	file, size := utils.Read_file(s)

	grid := make(Grid)
	for i := 0; i < size; i++ {
		this := file[i]
		for j := 0; j < len(this); j++ {
			pos := Pos{i, j}
			grid[pos] = string(this[j])
		}
	}

	start := []Point{{Pos{0, 1}, "down"}}
	end := []Point{{Pos{size - 1, len(file[0]) - 2}, "down"}}

	var queue PriorityQueue

	queue.Enqueue(start[0], 0)

	cost_so_far := make(map[Point]int)
	came_from := make(map[Point]Point)

	cost_so_far[start[0]] = 0

	time_reached := make(map[Pos]int)

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if InPoint(end, current.p) {
			return cost_so_far[current.p]
		}

		for _, next := range neighbors(current.p, grid, size, len(file[0])) {
			new_cost := cost_so_far[current.p] + 1
			if cost_so_far[next] == 0 || new_cost > cost_so_far[next] {
				if cost_so_far[next] != 0 {
					time_reached[next.pos] += 1
				}
				if time_reached[next.pos] < 1 {
					cost_so_far[next] = new_cost
					priority := new_cost + heuristic(Point{Pos{size - 1, len(file[0]) - 2}, "down"}, next)
					queue.Enqueue(next, priority)
					came_from[next] = current.p
				}

			}
		}
	}
	return 0
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

type PriorityQueue struct {
	data []QueueElement
	size int
}

// Checks if the priority queue is empty
func (p *PriorityQueue) IsEmpty() bool {
	return p.size == 0
}

// Adds a new element to the priority queue (unsorted array case)
func (p *PriorityQueue) Enqueue(el Point, priority int) {
	newElement := QueueElement{el, priority}
	p.data = append(p.data, newElement)
	p.size++
}

// Removes and returns the element with the max-priority from the PQ
func (p *PriorityQueue) Dequeue() QueueElement {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}

	idx := 0
	// Traverse through the unsorted array to find the element with the smallest min-priority
	for i := 0; i < p.size; i++ {
		if p.data[i].priority > p.data[idx].priority {
			idx = i
		}
	}
	dequeuedEl := p.data[idx]
	// Remove the dequeued element from the PQ
	p.data = append(p.data[:idx], p.data[idx+1:]...)
	p.size--
	return dequeuedEl
}
