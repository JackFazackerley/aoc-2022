package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

var hand = map[string]int{
	"X": 1, // Rock
	"A": 1, // Rock
	"Y": 2, // Paper
	"B": 2, // Paper
	"Z": 3, // Scissors
	"C": 3, // Scissors
}

func main() {
	absPath, _ := filepath.Abs("day2/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	var score int
	for _, str := range strs {
		split := strings.Split(str, " ")
		opponent := hand[split[0]]
		me := hand[split[1]]

		score += me

		if opponent == me {
			score += 3
			continue
		}

		switch me {
		case 1: // rock
			if opponent != 2 { // paper
				score += 6
			}
		case 2: // paper
			if opponent != 3 { // scissors
				score += 6
			}
		case 3: // scissors
			if opponent != 1 { // rock
				score += 6
			}
		}
	}
	return score
}

func part2(strs []string) int {
	var score int
	for _, str := range strs {
		split := strings.Split(str, " ")
		opponent := hand[split[0]]
		me := hand[split[1]]

		switch me {
		case 1: // rock
			score += (opponent+1)%3 + 1
		case 2: // paper
			score += 3
			score += opponent
		case 3: // scissors
			score += 6
			score += (opponent % 3) + 1
		}
	}

	return score
}
