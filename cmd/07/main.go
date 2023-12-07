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

// turn a hand into a list of tuples (hand value, bet value)
func computeHand(cards, values, pointsStr string, shouldImprove bool) [2]int {
	hand := map[rune]int{}
	points, _ := strconv.Atoi(pointsStr)

	var cardValues int
	for _, card := range cards {
		hand[card]++
		cardValues = cardValues*100 + strings.Index(values, string(card))
	}

	kind := 0
	for _, count := range hand {
		kind += count * count
	}
	if shouldImprove {
		kind = improve(hand, kind)
	}

	value := kind*1e10 + cardValues
	return [2]int{value, points}
}

// sort the list of hands returned by computeHand and sums it
func computeScore(hands [][2]int) int {
	slices.SortFunc(hands, func(a, b [2]int) int { return a[0] - b[0] })

	acc := 0
	for idx, hand := range hands {
		acc += (idx + 1) * hand[1]
	}

	return acc
}

func main() {
	handsPart1 := [][2]int{}
	handsPart2 := [][2]int{}
	for line := range utils.ReadInput() {
		cards, pointsStr, _ := strings.Cut(line, " ")
		handsPart1 = append(handsPart1, computeHand(cards, faceValues1, pointsStr, false))
		handsPart2 = append(handsPart2, computeHand(cards, faceValues2, pointsStr, true))
	}

	fmt.Println(computeScore(handsPart1))
	fmt.Println(computeScore(handsPart2))
}
