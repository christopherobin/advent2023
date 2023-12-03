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

func parseNumber(line []byte, idx int) (int, int, int) {
	number := 0
	s := 0
	e := 0

	// find the start
	for s = idx; s >= 1; s-- {
		if line[s-1] < '0' || line[s-1] > '9' {
			break
		}
	}

	for e = s; e < len(line); e++ {
		if line[e] < '0' || line[e] > '9' {
			break
		}

		number = number*10 + int(line[e]-'0')
	}

	return s, e, number
}

func neighbors(lines [3]sline, idx int) [][3]int {
	parts := [][3]int{}

	for y := 0; y < 3; y++ {
		if len(lines[y].data) == 0 {
			continue
		}

		for x := idx - 1; x <= idx+1; x++ {
			if x < 0 || x > len(lines[y].data) {
				continue
			}
			if lines[y].data[x] >= '0' && lines[y].data[x] <= '9' {
				s, e, num := parseNumber(lines[y].data, x)
				parts = append(parts, [3]int{y, s, num})

				// no need to read if the end of the number overlaps the entire row
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

	view[0], view[1], view[2] = view[1], view[2], sline{data: line}
	for idx, c := range view[1].data {
		if (c >= '0' && c <= '9') || c == '.' {
			continue
		}

		parts := neighbors(*view, idx)
		for _, part := range parts {
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
