package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/christopherobin/advent2023/pkg/utils"
	"github.com/samber/lo"
)

// tracks the amount of cards we have
type multipliers struct {
	internal []int64
}

func (m *multipliers) add(score, mult int64) {
	for i := int64(0); i < score; i++ {
		if i > int64(len(m.internal)-1) {
			m.internal = append(m.internal, mult)
		} else {
			m.internal[i] += mult
		}
	}
}

func (m *multipliers) pop() int64 {
	if len(m.internal) == 0 {
		return 1 // no extra cards
	}
	var n int64
	n, m.internal = m.internal[0], m.internal[1:]
	return n + 1 // add 1 for the original card
}

func main() {
	sum := int64(0)
	copies := int64(0)
	mults := &multipliers{}
	for line := range utils.ReadInput() {
		_, numbersRaw, _ := strings.Cut(line, ":")
		numbers := lo.Filter(lo.Map(strings.Split(numbersRaw, " "), func(numberStr string, _ int) int64 {
			n, _ := strconv.Atoi(numberStr)
			return int64(n)
		}), func(n int64, _ int) bool { return n > 0 })
		winningNumbers := len(lo.FindDuplicates(numbers))
		mult := mults.pop()
		copies += mult

		if winningNumbers > 0 {
			sum += int64(math.Pow(2, float64(winningNumbers-1)))
			mults.add(int64(winningNumbers), mult)
		}
	}
	fmt.Println("Points:", sum)
	fmt.Println("Copies:", copies)
}
