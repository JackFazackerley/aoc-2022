package main

import (
	"fmt"
	"path/filepath"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day9/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	return positions(2, strs)
}

func part2(strs []string) int {
	return positions(10, strs)
}

func positions(count int, strs []string) int {
	visited := map[string]int{}
	knots := make([][2]int, count)

	for _, str := range strs {
		var direction string
		var move int
		_, _ = fmt.Sscanf(str, "%s %d", &direction, &move)

		for i := 0; i < move; i++ {
			switch direction {
			case "R":
				knots[0][0] += 1
			case "L":
				knots[0][0] -= 1
			case "U":
				knots[0][1] -= 1
			case "D":
				knots[0][1] += 1
			}

			for j := 1; j < len(knots); j++ {
				diffX := knots[j-1][0] - knots[j][0]
				diffY := knots[j-1][1] - knots[j][1]

				if diffX > 1 || abs(diffY) > 1 && diffX >= 1 {
					knots[j][0]++
				} else if diffX < -1 || abs(diffY) > 1 && diffX <= -1 {
					knots[j][0]--
				}

				if diffY > 1 || abs(diffX) > 1 && diffY >= 1 {
					knots[j][1]++
				} else if diffY < -1 || abs(diffX) > 1 && diffY <= -1 {
					knots[j][1]--
				}

				if j == len(knots)-1 {
					visited[fmt.Sprintf("%d-%d", knots[j][0], knots[j][1])] += 1
				}
			}
		}
	}

	return len(visited)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
