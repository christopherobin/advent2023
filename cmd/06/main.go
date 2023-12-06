package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/christopherobin/advent2023/pkg/utils"
	"github.com/samber/lo"
)

// quadratic formula solver
func solve(time, distance int) int {
	// we bump the distance by a tiny bit because we want to make sure we beat the current record
	timef, distancef := float64(time), float64(distance)+0.001
	s1 := (timef - math.Sqrt(math.Pow(timef, 2)-4*distancef)) / 2
	s2 := (timef + math.Sqrt(math.Pow(timef, 2)-4*distancef)) / 2

	return int(math.Floor(s2)) - int(math.Ceil(s1)) + 1
}

func main() {
	lines := utils.ReadInput()
	timesStr := <-lines
	distancesStr := <-lines
	races := lo.Zip2(utils.ParseNumbersString(timesStr), utils.ParseNumbersString(distancesStr))

	acc := 1
	for _, race := range races {
		acc = acc * solve(race.A, race.B)
	}

	fmt.Println("part1:", acc)

	// kerning lul
	time := utils.ParseNumbersString(strings.Replace(timesStr, " ", "", -1))[0]
	distance := utils.ParseNumbersString(strings.Replace(distancesStr, " ", "", -1))[0]

	fmt.Println("part2:", solve(time, distance))
}
