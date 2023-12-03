package main

import (
	"fmt"
	"slices"

	"github.com/christopherobin/advent2023/pkg/utils"
)

type sline struct {
	data []byte
	seen [][3]int
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func parseNumber(line []byte, idx int) (int, int, int) {
	number := 0
	s := 0
	e := 0

	// find the start
	for s = idx; s >= 1; s-- {
		if !isDigit(line[s-1]) {
			break
		}
	}

	for e = s; e < len(line); e++ {
		if !isDigit(line[e]) {
			break
		}

		number = number*10 + int(line[e]-'0')
	}

	return s, e, number
}

func neighbors(lines [3]sline, idx int) [][3]int {
	parts := [][3]int{}

	// iterate on a 3x3 grid to find numbers
	for y := 0; y < 3; y++ {
		if len(lines[y].data) == 0 {
			continue
		}

		for x := idx - 1; x <= idx+1; x++ {
			if x < 0 || x > len(lines[y].data) {
				continue
			}
			if isDigit(lines[y].data[x]) {
				s, e, num := parseNumber(lines[y].data, x)
				parts = append(parts, [3]int{y, s, num})

				// no need to read if the end of the number overlaps the scanned area
				if e >= idx+1 {
					break
				}
			}
		}
	}

	return parts
}

func parse(view *[3]sline, line []byte) (int, int) {
	sum := 0
	ratio := 0

	// this is a sliding 3 line window with the middle one being the one scanned
	view[0], view[1], view[2] = view[1], view[2], sline{data: line}
	for idx, c := range view[1].data {
		if isDigit(c) || c == '.' {
			continue
		}

		// found a symbol, check neighbors
		parts := neighbors(*view, idx)
		for _, part := range parts {
			// if not seen before, sum and add to seen list
			if !slices.Contains(view[part[0]].seen, part) {
				sum += part[2]
				view[part[0]].seen = append(view[part[0]].seen, part)
			}
		}

		if c == '*' && len(parts) == 2 {
			ratio += parts[0][2] * parts[1][2]
		}
	}

	return sum, ratio
}

func main() {
	view := [3]sline{}
	sum := 0
	ratio := 0

	for line := range utils.ReadInputByteAddEmptyLine() {
		lsum, lratio := parse(&view, line)
		sum += lsum
		ratio += lratio
	}

	fmt.Println("Parts:", sum)
	fmt.Println("Gear ratio:", ratio)
}
