package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day10/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	part2(strs)
}

func part1(strs []string) int {
	signalStrengths := 0
	X := 1

	addx := 0
	processing := false

	cycles := 1
	cycle := 0

	for i := 0; i < len(strs); {
		cycle++
		if !processing {
			instruction := strs[i][:4]
			processing = true
			if instruction == "addx" {
				v, _ := strconv.Atoi(strs[i][5:])
				addx = v
				cycles = cycle + 1
			} else {
				addx = 0
				cycles += 1
			}
		}

		if cycle == 20 || (cycle+20)%40 == 0 {
			tmp := cycle * X
			signalStrengths += tmp
		}

		if cycles == cycle {
			X += addx
			i++
			processing = false
		}
	}

	return signalStrengths
}

func part2(strs []string) {
	X := 1

	addx := 0
	processing := false
	position := 0

	cycles := 1
	cycle := 0

	for i := 0; i < len(strs); {
		cycle++
		if !processing {
			instruction := strs[i][:4]
			processing = true
			if instruction == "addx" {
				v, _ := strconv.Atoi(strs[i][5:])
				addx = v
				cycles = cycle + 1
			} else {
				addx = 0
				cycles += 1
			}
		}

		if (position%40) >= X-1 && (position%40) <= X+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		position++

		if cycles == cycle {
			X += addx
			i++
			processing = false
		}

		if (cycle)%40 == 0 {
			fmt.Println()
		}
	}
}
