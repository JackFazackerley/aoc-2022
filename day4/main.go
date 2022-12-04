package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day4/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	contained := 0
	for _, str := range strs {
		elf := strings.Split(str, ",")
		section1 := strings.Split(elf[0], "-")
		section1a, _ := strconv.Atoi(section1[0])
		section1b, _ := strconv.Atoi(section1[1])

		section2 := strings.Split(elf[1], "-")
		section2a, _ := strconv.Atoi(section2[0])
		section2b, _ := strconv.Atoi(section2[1])

		if section1a <= section2a && section1b >= section2b {
			contained++
		} else if section2a <= section1a && section2b >= section1b {
			contained++
		}
	}

	return contained
}

func part2(strs []string) int {
	overlapped := 0
	for _, str := range strs {
		elf := strings.Split(str, ",")
		section1 := strings.Split(elf[0], "-")
		section1a, _ := strconv.Atoi(section1[0])
		section1b, _ := strconv.Atoi(section1[1])

		section2 := strings.Split(elf[1], "-")
		section2a, _ := strconv.Atoi(section2[0])
		section2b, _ := strconv.Atoi(section2[1])

		if section2b >= section1a && section2b <= section1b {
			overlapped++
		} else if section2a >= section1a && section2a <= section1b {
			overlapped++
		} else if section1b >= section2a && section1b <= section2b {
			overlapped++
		}

	}

	return overlapped
}
