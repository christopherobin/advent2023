package main

import (
	"fmt"

	"github.com/christopherobin/advent2023/pkg/utils"
)

func navigate(tree map[string][2]string, directions string) int {
	current := tree["AAA"]
	steps := 0
	for {
		for _, direction := range directions {
			target := current[0]

			if direction == 'R' {
				target = current[1]
			}

			steps++
			if target == "ZZZ" {
				return steps
			}
			current = tree[target]
		}
	}
}

func ghostNavigate(tree map[string][2]string, directions string) int {
	dests := map[string]string{}
	zs := map[string][]int{}
	from := []string{}
	for start := range tree {
		current := tree[start]
		if start[2] == 'A' {
			from = append(from, start)
		}
		var last string
		for step, direction := range directions {
			target := current[0]

			if direction == 'R' {
				target = current[1]
			}

			if target[2] == 'Z' {
				zs[start] = append(zs[start], step)
			}
			last = target
			current = tree[target]
		}

		dests[start] = last
	}

	steps := 0
	oneIt := len(directions)
	for {
		win := true
		for idx, key := range from {
			if _, ok := zs[key]; !ok {
				win = false
			}
			from[idx] = dests[key]
		}
		if win {
			break
		}

		steps += oneIt
	}

	fmt.Println(steps)
	return 0
}

func main() {
	reader := utils.ReadInput()
	directions := <-reader
	<-reader

	tree := map[string][2]string{}
	for line := range reader {
		var node, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &node, &left, &right)
		tree[node] = [2]string{left, right}
	}

	//fmt.Println(navigate(tree, directions))
	fmt.Println(ghostNavigate(tree, directions))
}
