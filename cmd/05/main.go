package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/christopherobin/advent2023/pkg/utils"
)

type Map struct {
	target string
	ranges []RangeMap
}

func (m Map) Convert(r []Range) (string, []Range) {
	queue := slices.Clone(r)
	result := []Range{}

	for _, rng := range m.ranges {
		if len(queue) == 0 {
			break
		}

		nextQueue := []Range{}
		for _, r := range queue {
			converted, unconverted := rng.Convert(r)
			if converted.end != 0 {
				result = append(result, converted)
			}
			nextQueue = append(nextQueue, unconverted...)
		}
		queue = nextQueue
	}

	result = append(result, queue...)

	return m.target, result
}

type Range struct {
	start, end int
}

func NewRange(s, l int) Range {
	return Range{s, s + l}
}

type RangeMap struct {
	src   Range
	shift int
}

func (rm RangeMap) Convert(srcRange Range) (Range, []Range) {
	unconvertedRanges := []Range{}
	if srcRange.end < rm.src.start || srcRange.start > rm.src.end {
		return Range{}, []Range{srcRange}
	}

	if srcRange.start < rm.src.start {
		unconvertedRanges = append(unconvertedRanges, Range{srcRange.start, rm.src.start - 1})
	}

	if srcRange.end > rm.src.end {
		unconvertedRanges = append(unconvertedRanges, Range{rm.src.end, srcRange.end})
	}

	overlapStart := max(srcRange.start, rm.src.start)
	overlapEnd := min(srcRange.end, rm.src.end)
	return Range{overlapStart + rm.shift, overlapEnd + rm.shift}, unconvertedRanges
}

func FindLowest(maps map[string]*Map, values []Range) int {
	current := "seed"
	lowest := 4294967296
	ranges := slices.Clone(values)

	for {
		current, ranges = maps[current].Convert(ranges)
		if current == "location" {
			break
		}
	}

	for _, val := range ranges {
		if val.start < lowest {
			lowest = val.start
		}
	}

	return lowest
}

func main() {
	s := time.Now()
	var seeds []int
	var currentMap *Map
	maps := map[string]*Map{}

	for line := range utils.ReadInput() {
		if len(line) == 0 {
			continue
		}

		_, extra, found := strings.Cut(line, ":")
		if found {
			if len(extra) > 0 {
				seeds = utils.ParseNumbersString(extra)
				continue
			}

			mapType := strings.Split(line[:strings.Index(line, " ")], "-")
			currentMap = &Map{target: mapType[2]}
			maps[mapType[0]] = currentMap
			continue
		}

		numbers := utils.ParseNumbersString(line)
		currentMap.ranges = append(currentMap.ranges, RangeMap{
			NewRange(numbers[1], numbers[2]),
			numbers[0] - numbers[1],
		})
	}

	// sort the ranges
	for _, rmap := range maps {
		slices.SortFunc(rmap.ranges, func(a, b RangeMap) int {
			return a.src.start - b.src.start
		})
	}

	part1Ranges := []Range{}
	part2Ranges := []Range{}
	for idx, seed := range seeds {
		part1Ranges = append(part1Ranges, NewRange(seed, 1))

		if idx%2 > 0 {
			part2Ranges = append(part2Ranges, NewRange(seeds[idx-1], seed))
		}
	}

	fmt.Printf("[part1] Lowest location among all seeds: %d\n", FindLowest(maps, part1Ranges))
	fmt.Printf("[part2] Lowest location among all seeds: %d\n", FindLowest(maps, part2Ranges))
	fmt.Println("Execution time:", time.Since(s))
}
