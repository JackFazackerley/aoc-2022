package main

import (
	"fmt"
	"path/filepath"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day8/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	var visible int
	height := len(strs) - 1
	width := len(strs[0]) - 1
	visible += (len(strs) * 2) + ((len(strs[0]) - 2) * 2)

	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if v, _ := compareX(strs, y, x, width); v {
				visible += 1
			} else if v, _ := compareY(strs, y, x, height); v {
				visible += 1
			}
		}
	}

	return visible
}

func part2(strs []string) int {
	var score int

	height := len(strs) - 1
	width := len(strs[0]) - 1

	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			_, left := compareX(strs, y, x, width)
			_, right := compareY(strs, y, x, height)

			if left*right > score {
				score = left * right
			}
		}
	}

	return score
}

func compareX(strs []string, y, x, max int) (bool, int) {
	var isVisibleLeft bool
	var isVisibleRight bool
	var visibleLeft int
	var visibleRight int

	tree := int(strs[y][x] - '0')

	for i := x - 1; i >= 0; i-- {
		nextTree := int(strs[y][i] - '0')

		if nextTree < tree {
			isVisibleLeft = true
			visibleLeft++
		} else {
			isVisibleLeft = false
			visibleLeft++
			break
		}
	}

	for i := x + 1; i <= max; i++ {
		nextTree := int(strs[y][i] - '0')

		if nextTree < tree {
			isVisibleRight = true
			visibleRight++
		} else {
			isVisibleRight = false
			visibleRight++
			break
		}
	}

	return isVisibleLeft || isVisibleRight, visibleLeft * visibleRight
}

func compareY(strs []string, y, x, max int) (bool, int) {
	var isVisibleLeft bool
	var isVisibleRight bool
	var visibleLeft int
	var visibleRight int

	tree := int(strs[y][x] - '0')

	for i := y - 1; i >= 0; i-- {
		nextTree := int(strs[i][x] - '0')

		if nextTree < tree {
			isVisibleLeft = true
			visibleLeft++
		} else {
			isVisibleLeft = false
			visibleLeft++
			break
		}
	}

	for i := y + 1; i <= max; i++ {
		nextTree := int(strs[i][x] - '0')

		if nextTree < tree {
			isVisibleRight = true
			visibleRight++
		} else {
			isVisibleRight = false
			visibleRight++
			break
		}
	}

	return isVisibleLeft || isVisibleRight, visibleLeft * visibleRight
}
