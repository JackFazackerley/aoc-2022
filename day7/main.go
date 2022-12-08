package main

import (
	"fmt"
	"path/filepath"

	"github.com/JackFazackerley/aoc-2022/pkg/reader"
)

func main() {
	absPath, _ := filepath.Abs("day7/input.txt")
	strs := reader.ReadStings(absPath)

	fmt.Println(part1(strs))
	fmt.Println(part2(strs))
}

func part1(strs []string) int {
	root := &node{
		left:   nil,
		right:  make(map[string]*node),
		amount: 0,
		dir:    "/",
	}

	current := root

	for _, str := range strs[2:] {
		node := current.insert(str)
		current = node
	}

	return findTotalSize(root)
}

func part2(strs []string) int {
	root := &node{
		left:   nil,
		right:  make(map[string]*node),
		amount: 0,
		dir:    "/",
	}

	current := root

	for _, str := range strs[2:] {
		node := current.insert(str)
		current = node
	}

	required := 30000000 - (70000000 - root.amount)

	return findSpace(root, root, required).amount
}

func findTotalSize(node *node) int {
	var total int
	for _, node := range node.right {
		if node.amount >= 0 && node.amount <= 100000 {
			total += node.amount
		}
		total += findTotalSize(node)
	}

	return total
}

func findSpace(n, selected *node, required int) *node {
	for _, inner := range n.right {
		diff := n.amount - required

		if diff < selected.amount-required && diff >= 0 {
			selected = n
		}

		selected = findSpace(inner, selected, required)
	}

	return selected
}

type node struct {
	left   *node
	right  map[string]*node
	amount int
	dir    string
}

func (n *node) insert(str string) *node {
	if str[0] == '$' {
		if str[:4] == "$ cd" {
			if str[5:] == ".." {
				return n.left
			} else {
				return n.right[str[5:]]
			}
		}
	} else if str[:3] == "dir" {
		if n.right[str[:4]] == nil {
			n.right[str[4:]] = &node{
				left:  n,
				right: make(map[string]*node),
				dir:   str[4:],
			}
		}
	} else {
		var size int

		_, _ = fmt.Sscanf(str, "%d ", &size)

		n.amount += size
		if n.left != nil {
			n.left.insert(str)
		}
	}

	return n
}
