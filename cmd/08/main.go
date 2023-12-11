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
	all := [][2]string{}
	for src := range tree {
		if src[2] == 'A' {
			fmt.Println(src)
			all = append(all, tree[src])
		}
	}

	fmt.Println(directions)
	steps := 0
	for {
		for _, direction := range directions {
			allz := true
			zs := 0
			got := ""

			for idx, node := range all {
				target := node[0]

				if direction == 'R' {
					target = node[1]
				}

				got += string(target[2])

				if target[2] != 'Z' {
					allz = false
				} else {
					zs++
				}
				all[idx] = tree[target]
			}

			if zs > 2 {
				fmt.Println(steps, got)
			}

			steps++
			if allz {
				return steps
			}
		}
	}
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

	fmt.Println(navigate(tree, directions))
	fmt.Println(ghostNavigate(tree, directions))
}
