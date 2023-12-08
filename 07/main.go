package main

import (
	"strconv"
	"strings"

	"github.com/EmilienBois/Advent_Of_Code2023/utils"
)

type Hand struct {
	hand  string
	bid   int
	power float64 // 1 pour rien, 2 pour double, 2.5 pour double pair, 3 pour 3 pareil, 3.5 pour 3+2, 4 pour 4 pareil, 5 pour 5 pareil
	rank  int
}

func Part1(s string) int {
	//We read the file entirely
	test, size := utils.Read_file(s)
	var cartes [7][]Hand

	for i := 0; i < size; i++ {
		sep := strings.Split(test[i], " ")
		bid, err := strconv.Atoi(sep[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		hand := sep[0]
		power := Getpower(hand)
		if power == 2.5 {
			power = 0.
		} else if power == 3.5 {
			power = 6
		}
		newCard := Hand{sep[0], bid, power, 0}
		cartes[int(power)] = append(cartes[int(power)], newCard)
	}
	tot_winnings := 0
	rank := 1

	power := []int{1, 2, 0, 3, 6, 4, 5}
	for _, i_power := range power { //Pour tous les pouvoirs
		cartes[i_power] = Compare(cartes[i_power])
		for _, elt := range cartes[i_power] {
			tot_winnings += elt.bid * (rank + elt.rank)
		}
		rank += len(cartes[i_power])
	}

	return tot_winnings
}

func Compare(this []Hand) []Hand {
	size := len(this)
	var new []Hand
	for i := 0; i < size; i++ {
		rank := 0
		h1 := this[i]
		for j := 0; j < size; j++ {
			if j != i {
				h2 := this[j]
				if h1.sup(h2) {
					rank += 1
				}
			}
		}
		h1.rank = rank
		new = append(new, h1)
	}
	return new
}

func (h1 Hand) sup(h2 Hand) bool {
	i := 0
	for {
		ind1 := Get_equivalent(h1.hand[i])
		ind2 := Get_equivalent(h2.hand[i])
		if ind1 > ind2 {
			return true
		} else if ind1 < ind2 {
			return false
		}
		if i < 4 {
			i += 1
		} else {
			return false
		}
	}
}

func Get_equivalent(elt byte) int {
	switch elt {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	}
	return 0
}

func Getpower(hand string) float64 {
	cards := make(map[rune]float64)
	for _, elt := range hand {
		cards[elt] += 1
	}
	max := 0.
	test := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	for _, elt := range test {
		if cards[elt] > max {
			max = cards[elt]
		} else if cards[elt] == 2 {
			if max == 2 {
				max = 2.5
			} else if max == 3 {
				max = 3.5
			}
		}
	}
	return max
}

func Part2(s string) int {
	return 0
}

func main() {
	result_one_test := Part1("inputtest.txt")
	println("The result of the input test for the part 1 is : ", result_one_test)
	result_one := Part1("input.txt")
	println("The result of the input for the part 1 is : ", result_one)
	result_two_test := Part2("inputtest.txt")
	println("The result of the input test for the part 2 is : ", result_two_test)
	result_two := Part2("input.txt")
	println("The result of the input for the part 2 is : ", result_two)
}
