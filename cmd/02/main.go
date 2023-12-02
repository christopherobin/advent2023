package main

import (
	"fmt"

	"github.com/christopherobin/advent2023/pkg/utils"
)

func main() {
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	power := 0
	for line := range utils.ReadInput() {
		fewest := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		gameId := 0
		count := 0
		color := ""
		seenColon := false
		bad := false

		for _, c := range line + ";" {
			if !seenColon && c == ':' {
				seenColon = true
				gameId = count
				count = 0
				color = ""
				continue
			}
			if c >= '0' && c <= '9' {
				count = count*10 + int(c-'0')
				continue
			}
			if c >= 'a' && c <= 'z' {
				color += string(c)
				continue
			}
			if c == ',' || c == ';' || c == 0 {
				if bag[color] < count {
					bad = true
				}
				if count > fewest[color] {
					fewest[color] = count
				}
				count = 0
				color = ""
			}
		}

		if !bad {
			sum += gameId
		}
		power += fewest["red"] * fewest["green"] * fewest["blue"]
	}

	fmt.Println(sum)
	fmt.Println(power)
}
