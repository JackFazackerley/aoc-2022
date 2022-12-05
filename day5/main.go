package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day5/input.txt")
	strs := reader.SpltBy(absPath, "\n\n")

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) string {
	crates := findCrates(strs[0])

	lines := strings.Split(strs[1], "\n")

	for _, line := range lines {
		var amount, from, to int

		_, _ = fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)

		for i := 1; i <= amount; i++ {
			lastCrate := crates[from-1][len(crates[from-1])-1]

			crates[from-1] = crates[from-1][:len(crates[from-1])-1]

			crates[to-1] = append(crates[to-1], lastCrate)
		}
	}

	var result string

	for _, crate := range crates {
		result += crate[len(crate)-1]
	}

	return result
}

func part2(strs []string) string {
	crates := findCrates(strs[0])

	lines := strings.Split(strs[1], "\n")

	for _, line := range lines {
		var amount, from, to int

		_, _ = fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)

		lastCrates := crates[from-1][len(crates[from-1])-amount:]

		crates[from-1] = crates[from-1][:len(crates[from-1])-amount]

		crates[to-1] = append(crates[to-1], lastCrates...)
	}

	var result string

	for _, crate := range crates {
		result += crate[len(crate)-1]
	}

	return result
}

func findCrates(crates string) [][]string {
	lines := strings.Split(crates, "\n")

	numberOfStacks := strings.ReplaceAll(lines[len(lines)-1], " ", "")
	totalStacks := len(numberOfStacks)
	stacks := make([][]string, totalStacks)

	for i := len(lines) - 2; i >= 0; i-- {
		for j := 1; j < len(lines[i]); j += 4 {
			if lines[i][j] != ' ' {
				stacks[j/4] = append(stacks[j/4], string(lines[i][j]))
			}
		}
	}

	return stacks
}
