package main

import (
	"fmt"
	"slices"

	"github.com/christopherobin/advent2023/pkg/utils"
)

func main() {
	total := 0
	ntotal := 0

	for line := range utils.ReadInput() {
		numbers := utils.ParseNumbersString(line)
		current := slices.Clone(numbers)
		firsts := []int{}
		acc := 0

		// part 1
		for {
			acc += current[len(current)-1]

			next := make([]int, len(current)-1)
			allz := true
			firsts = append(firsts, current[0])

			for i := 1; i < len(current); i++ {
				diff := current[i] - current[i-1]
				next[i-1] = diff

				if diff != 0 {
					allz = false
				}
			}

			if allz || len(next) == 0 {
				break
			}

			current = next
		}

		// part 2
		prev := 0
		for i := len(firsts) - 1; i >= 0; i-- {
			prev = firsts[i] - prev
		}

		ntotal += prev
		total += acc
	}

	fmt.Println("[part1]", total)
	fmt.Println("[part2]", ntotal)
}
