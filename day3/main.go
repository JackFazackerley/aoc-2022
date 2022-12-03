package main

import (
	"fmt"
	"path/filepath"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day3/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	priorities := 0

	for _, str := range strs {
		half := len(str) / 2
		first := str[:half]
		second := str[half:]

		seen := map[int]interface{}{}

		for _, char := range first {
			priority := getPriority(char)

			seen[priority] = nil
		}

		for _, char := range second {
			priority := getPriority(char)

			if _, ok := seen[priority]; ok {
				priorities += priority
				break
			}
		}

	}

	return priorities
}

func part2(strs []string) int {
	priorities := 0

	for i := 0; i+2 < len(strs); i += 3 {
		first := strs[i]
		second := strs[i+1]
		third := strs[i+2]

		seen := map[int]int{}

		for _, char := range first {
			seen[getPriority(char)] = 1
		}

		for _, char := range second {
			priority := getPriority(char)
			if _, ok := seen[priority]; ok {
				seen[priority] = 2
			}
		}

		for _, char := range third {
			priority := getPriority(char)

			if v := seen[priority]; v == 2 {
				priorities += priority
				break
			}
		}
	}

	return priorities
}

func getPriority(char rune) int {
	priority := 0

	if 'A' <= char && char <= 'Z' {
		priority = int(char - 'A' + 27)
	}

	if 'a' <= char && char <= 'z' {
		priority = int(char - 'a' + 1)
	}

	return priority
}
