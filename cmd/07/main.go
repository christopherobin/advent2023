package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/christopherobin/advent2023/pkg/utils"
)

const faceValues1 = "23456789TJQKA"
const faceValues2 = "J23456789TQKA"

const (
	FIVE_OF_A_KIND  = 5 * 5
	FOUR_OF_A_KIND  = 4*4 + 1
	FULL_HOUSE      = 3*3 + 2*2
	THREE_OF_A_KIND = 3*3 + 2
	TWO_PAIR        = 2*2 + 2*2 + 1
	ONE_PAIR        = 2*2 + 3
	HIGH_CARD       = 5
)

func valueToString(value int) string {
	switch value {
	case FIVE_OF_A_KIND:
		return "FIVE_OF_A_KIND"
	case FOUR_OF_A_KIND:
		return "FOUR_OF_A_KIND"
	case FULL_HOUSE:
		return "FULL_HOUSE"
	case THREE_OF_A_KIND:
		return "THREE_OF_A_KIND"
	case TWO_PAIR:
		return "TWO_PAIR"
	case ONE_PAIR:
		return "ONE_PAIR"
	case HIGH_CARD:
		return "HIGH_CARD"
	}

	panic("WUT")
}

func improve(hand map[rune]int, value int) int {
	if _, ok := hand['J']; !ok {
		return value
	}

	switch value {
	case FIVE_OF_A_KIND:
		return FIVE_OF_A_KIND
	case FOUR_OF_A_KIND:
		return FIVE_OF_A_KIND
	case FULL_HOUSE:
		return FIVE_OF_A_KIND
	case THREE_OF_A_KIND:
		return FOUR_OF_A_KIND
	case TWO_PAIR:
		if hand['J'] == 2 {
			return FOUR_OF_A_KIND
		}
		return FULL_HOUSE
	case ONE_PAIR:
		return THREE_OF_A_KIND
	case HIGH_CARD:
		return ONE_PAIR
	}

	return value
}

func main() {
	hands1 := [][2]int{}
	hands2 := [][2]int{}
	for line := range utils.ReadInput() {
		cards, pointsStr, _ := strings.Cut(line, " ")
		hand := map[rune]int{}

		var cardValues1 int
		var cardValues2 int
		for _, card := range cards {
			hand[card]++
			cardValues1 = cardValues1*100 + strings.Index(faceValues1, string(card))
			cardValues2 = cardValues2*100 + strings.Index(faceValues2, string(card))
		}

		kind1 := 0
		for _, count := range hand {
			kind1 += count * count
		}
		kind2 := improve(hand, kind1)

		value1 := kind1*1e10 + cardValues1
		value2 := kind2*1e10 + cardValues2

		points, _ := strconv.Atoi(pointsStr)

		hands1 = append(hands1, [2]int{value1, points})
		hands2 = append(hands2, [2]int{value2, points})
	}

	slices.SortFunc(hands1, func(a, b [2]int) int { return a[0] - b[0] })
	slices.SortFunc(hands2, func(a, b [2]int) int { return a[0] - b[0] })

	acc1 := 0
	for idx, hand := range hands1 {
		acc1 += (idx + 1) * hand[1]
	}

	acc2 := 0
	for idx, hand := range hands2 {
		acc2 += (idx + 1) * hand[1]
	}

	fmt.Println(acc1)
	fmt.Println(acc2)
}
