package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Liste map[string][]Rules

type Rules struct {
	cat         string
	sign        string
	value       int
	destination string
}

type Rating map[string]int

type Ratings []Rating

func Part1(s string) int {
	file, size := utils.Read_file(s)

	workflows := make(Liste)
	var ratings Ratings
	var cat string
	var value int
	var sign string
	var destination string

	i := 0
	for len(file[i]) != 0 {
		this := file[i]
		str := strings.Split(this, "{")
		name := str[0]
		str = strings.Split(str[1][:len(str[1])-1], ",")
		for _, elt := range str {
			reste := strings.Split(elt, ":")
			if len(reste) == 1 {
				destination = reste[0]
				value = 0
				sign = ""
				cat = ""
			} else {
				cat = string(elt[0])
				sign = string(elt[1])
				destination = reste[1]
				value_str := reste[0][2:]
				value = utils.String_to_Int(value_str)
			}
			rule := Rules{string(cat), sign, value, destination}
			workflows[name] = append(workflows[name], rule)
		}
		i += 1
	}
	i += 1
	for i < size {
		this := file[i]
		this = strings.Split((strings.Split(this, "{")[1]), "}")[0]
		val := strings.Split(this, ",")
		this_rating := make(Rating)
		for _, elt := range val {
			str := strings.Split(elt, "=")
			nb := utils.String_to_Int(str[1])
			this_rating[str[0]] = nb
		}
		ratings = append(ratings, this_rating)
		i += 1
	}
	// Now we have both the ratings and the workflows
	var accepted []Rating
	for _, rating := range ratings {
		step := "in"
		for step != "A" && step != "R" {
			this := workflows[step]
			for _, rule := range this {
				if rule.cat == "" {
					step = rule.destination
					break
				}
				val := rating[rule.cat]
				if rule.sign == ">" {
					if val > rule.value {
						step = rule.destination
						break
					}
				} else {
					if val < rule.value {
						step = rule.destination
						break
					}
				}
			}
		}
		if step == "A" {
			accepted = append(accepted, rating)
		}
	}
	total := 0
	for _, rate := range accepted {
		sum := 0
		l := []string{"x", "m", "a", "s"}
		for _, letter := range l {
			sum += rate[letter]
		}
		total += sum
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
