package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day1/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	var calories int
	var maxCalories int
	for _, str := range strs {
		if str == "" {
			if calories > maxCalories {
				maxCalories = calories
			}

			calories = 0
		}

		food, _ := strconv.Atoi(str)

		calories += food
	}

	return maxCalories
}

func part2(strs []string) int {
	var calories int
	top3 := make([]int, 3)

	for _, str := range strs {
		food, _ := strconv.Atoi(str)

		calories += food

		if str == "" {
			if calories > top3[0] {
				top3[0] = calories
				sort.Ints(top3)
			}

			calories = 0
		}
	}

	total := 0

	for _, cals := range top3 {
		total += cals
	}

	return total
}
