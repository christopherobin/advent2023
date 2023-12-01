package main

import (
	"fmt"

	"github.com/christopherobin/advent2023/pkg/utils"
)

var digitTable map[string]int = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseLine(line string) []int {
	digits := []int{}

	for i := 0; i < len(line); i++ {
		for possibleDigitStr, possibleDigit := range digitTable {
			if i+len(possibleDigitStr) <= len(line) &&
				line[i:i+len(possibleDigitStr)] == possibleDigitStr {
				digits = append(digits, possibleDigit)
				break
			}
		}
	}

	return digits
}

func main() {
	sum := 0
	for line := range utils.ReadInput() {
		matches := parseLine(line)

		if len(matches) == 0 {
			continue
		}

		first := matches[0]
		last := matches[len(matches)-1]
		number := first*10 + last
		sum += number

		//fmt.Println(line, ":", number, "=", sum)
	}

	fmt.Println(sum)
}
