package main

import (
	"fmt"
	"path/filepath"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day6/input.txt")
	str := reader.Read(absPath)

	fmt.Println(part1(str))
	fmt.Println(part2(str))
}

func part1(str string) int {
	return find(str, 4)
}

func part2(str string) int {
	return find(str, 14)
}

func find(str string, distinct int) int {
	for i := 0; i < len(str); i++ {
		marker := map[rune]interface{}{}

		chars := str[i : i+distinct]

		for _, char := range chars {
			if _, ok := marker[char]; ok {
				break
			} else {
				marker[char] = nil
			}
		}

		if len(marker) == distinct {
			return i + distinct
		}
	}

	return 0
}
