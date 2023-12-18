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

type Point struct {
	pos       Pos
	direction string
}

type QueueElement struct {
	p        Point
	priority int
}

func neighbors(p Point, min, max, nb_line, nb_col int) []Point {
	var neighbors []Point
	if p.direction == "left" || p.direction == "right" {
		for i := min; i < max+1; i++ {
			if p.pos.X+i < nb_line {
				neighbors = append(neighbors, Point{Pos{p.pos.X + i, p.pos.Y}, "down"})
			}
			if p.pos.X-i >= 0 {
				neighbors = append(neighbors, Point{Pos{p.pos.X - i, p.pos.Y}, "up"})
			}
		}
	} else {
		for i := min; i < max+1; i++ {
			if p.pos.Y+i < nb_col {
				neighbors = append(neighbors, Point{Pos{p.pos.X, p.pos.Y + i}, "right"})
			}
			if p.pos.Y-i >= 0 {
				neighbors = append(neighbors, Point{Pos{p.pos.X, p.pos.Y - i}, "left"})
			}
		}
	}
	return neighbors
}

func InPoint(l []Point, elt Point) bool {
	for _, point := range l {
		if point == elt {
			return true
		}
	}
	return false
}

func GraphCost(cost Grid, current, next Point) int {
	c := 0
	if next.direction == "right" {
		for i := current.pos.Y + 1; i < next.pos.Y+1; i++ {
			c += cost[Pos{current.pos.X, i}]
		}
	}
	if next.direction == "left" {
		for i := current.pos.Y - 1; i > next.pos.Y-1; i-- {
			c += cost[Pos{current.pos.X, i}]
		}
	}
	if next.direction == "down" {
		for i := current.pos.X + 1; i < next.pos.X+1; i++ {
			c += cost[Pos{i, current.pos.Y}]
		}
	}
	if next.direction == "up" {
		for i := current.pos.X - 1; i > next.pos.X-1; i-- {
			c += cost[Pos{i, current.pos.Y}]
		}
	}
	return c
}

func Part1(s string) int {
	file, size := utils.Read_file(s)

	cost := make(Grid)
	for i := 0; i < size; i++ {
		this := file[i]
		for j := 0; j < len(this); j++ {
			pos := Pos{i, j}
			cost[pos] = utils.Byte_to_Int(string(this[j]))
		}
	}

	start := []Point{{Pos{0, 0}, "down"}, {Pos{0, 0}, "right"}}
	end := []Point{{Pos{size - 1, len(file[0]) - 1}, "down"}, {Pos{size - 1, len(file[0]) - 1}, "right"}}

	var queue PriorityQueue

	queue.Enqueue(start[0], 0)
	queue.Enqueue(start[1], 0)

	cost_so_far := make(map[Point]int)
	came_from := make(map[Point]Point)

	cost_so_far[start[0]] = 0
	cost_so_far[start[1]] = 0

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if InPoint(end, current.p) {
			return cost_so_far[current.p]
		}

		for _, next := range neighbors(current.p, 1, 3, size, len(file[0])) {
			new_cost := cost_so_far[current.p] + GraphCost(cost, current.p, next)
			if cost_so_far[next] == 0 || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				priority := new_cost
				queue.Enqueue(next, priority)
				came_from[next] = current.p
			}
		}
	}
	return 0
}

func Part2(s string) int {
	file, size := utils.Read_file(s)

	cost := make(Grid)
	for i := 0; i < size; i++ {
		this := file[i]
		for j := 0; j < len(this); j++ {
			pos := Pos{i, j}
			cost[pos] = utils.Byte_to_Int(string(this[j]))
		}
	}

	start := []Point{{Pos{0, 0}, "down"}, {Pos{0, 0}, "right"}}
	end := []Point{{Pos{size - 1, len(file[0]) - 1}, "down"}, {Pos{size - 1, len(file[0]) - 1}, "right"}}

	var queue PriorityQueue

	queue.Enqueue(start[0], 0)
	queue.Enqueue(start[1], 0)

	cost_so_far := make(map[Point]int)
	came_from := make(map[Point]Point)

	cost_so_far[start[0]] = 0
	cost_so_far[start[1]] = 0

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if InPoint(end, current.p) {
			return cost_so_far[current.p]
		}

		for _, next := range neighbors(current.p, 4, 10, size, len(file[0])) {
			new_cost := cost_so_far[current.p] + GraphCost(cost, current.p, next)
			if cost_so_far[next] == 0 || new_cost < cost_so_far[next] {
				cost_so_far[next] = new_cost
				priority := new_cost
				queue.Enqueue(next, priority)
				came_from[next] = current.p
			}
		}
	}
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

// Removes and returns the element with the min-priority from the PQ
func (p *PriorityQueue) Dequeue() QueueElement {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}

	idx := 0
	// Traverse through the unsorted array to find the element with the smallest min-priority
	for i := 0; i < p.size; i++ {
		if p.data[i].priority < p.data[idx].priority {
			idx = i
		}
	}
	dequeuedEl := p.data[idx]
	// Remove the dequeued element from the PQ
	p.data = append(p.data[:idx], p.data[idx+1:]...)
	p.size--
	return dequeuedEl
}

// Peek (Returns the element with the min-priority from the PQ)
func (p *PriorityQueue) Peek() QueueElement {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}
	dequeuedEl := p.data[0]
	// Traverse through the unsorted array to find the element with the min-priority
	for _, el := range p.data {
		if el.priority < dequeuedEl.priority {
			dequeuedEl = el
		}
	}
	return dequeuedEl
}

// Display the elements of the queue in an array form
func (p *PriorityQueue) Display() {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}

	arr := make([]interface{}, p.size)
	i := 0
	for i < p.size {
		arr[i] = p.data[i].p
		i++
	}
	fmt.Println("Priority Queue elements: ", arr)
}
