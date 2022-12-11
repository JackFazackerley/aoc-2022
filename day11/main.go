package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day11/input.txt")
	strs := reader.SpltBy(absPath, "\n\n")

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	monkeys := parseInput(strs)

	return compute(monkeys, 20)
}

func part2(strs []string) int {
	monkeys := parseInput(strs)

	return compute(monkeys, 10000)
}

func compute(monkeys []*monkey, iterations int) int {
	multiplier := 1
	for _, m := range monkeys {
		multiplier *= m.test
	}

	for i := 0; i < iterations; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := item
				switch monkey.operation {
				case "*":
					if monkey.operationNumber != 0 {
						worryLevel *= monkey.operationNumber
					} else {
						worryLevel *= worryLevel
					}
				case "+":
					worryLevel += monkey.operationNumber
				}
				monkey.inspected++

				if iterations == 20 {
					worryLevel /= 3
				} else {
					worryLevel %= multiplier
				}

				isDivisible := worryLevel%monkey.test == 0
				throwTo := monkey.receivers[isDivisible]

				monkeys[throwTo].items = append(monkeys[throwTo].items, worryLevel)

				monkey.items = monkey.items[1:]
			}
		}
	}

	sort.Slice(
		monkeys, func(i, j int) bool {
			return monkeys[i].inspected >= monkeys[j].inspected
		},
	)
	return monkeys[0].inspected * monkeys[1].inspected
}

type monkey struct {
	items           []int
	operation       string
	operationNumber int
	test            int
	receivers       map[bool]int
	inspected       int
	monkey          int
}

func parseInput(strs []string) []*monkey {
	monkeys := make([]*monkey, len(strs))
	for i, str := range strs {
		note := strings.Split(str, "\n")
		monkeys[i] = &monkey{
			items:           []int{},
			operation:       "",
			operationNumber: 0,
			test:            0,
			receivers:       map[bool]int{},
			monkey:          i,
		}

		for _, item := range strings.Split(note[1][17:], ",") {
			n, _ := strconv.Atoi(item[1:])
			if n != 0 {
				monkeys[i].items = append(monkeys[i].items, n)
			}
		}

		_, _ = fmt.Sscanf(note[2], " Operation: new = old %s %d", &monkeys[i].operation, &monkeys[i].operationNumber)
		_, _ = fmt.Sscanf(note[3], " Test: divisible by %d", &monkeys[i].test)
		monkeys[i].receivers[true], _ = strconv.Atoi(note[4][29:])
		monkeys[i].receivers[false], _ = strconv.Atoi(note[5][30:])
	}

	return monkeys
}
