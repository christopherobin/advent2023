package main

import (
	"fmt"

	"github.com/christopherobin/advent2023/pkg/utils"
)

type partNumber struct {
	line     int
	number   int
	location [2]int
}

func (p *partNumber) accumulate(idx int, c rune) {
	if p.number == 0 {
		p.location[0] = idx
	}
	p.number = p.number*10 + int(c-'0')
}

func (p *partNumber) finalize(idx int) {
	p.location[1] = idx - 1
}

func (p partNumber) hasNearbySymbol(s schematic) bool {
	symbols := []int{}
	for i := p.line - 1; i <= p.line+1; i++ {
		if i >= 0 && i < len(s.lines) {
			symbols = append(symbols, s.lines[i].symbolsIdx...)
		}
	}

	for _, symbol := range symbols {
		if symbol >= p.location[0]-1 && symbol <= p.location[1]+1 {
			return true
		}
	}

	return false
}

type schematicLine struct {
	line        int
	partNumbers []partNumber
	symbolsIdx  []int
	symbols     []rune
}

func (sl *schematicLine) addPartNumber(idx int, p partNumber) {
	if p.number == 0 {
		return
	}

	p.finalize(idx)
	sl.partNumbers = append(sl.partNumbers, p)
}

func (sl *schematicLine) addSymbol(idx int, c rune) {
	sl.symbolsIdx = append(sl.symbolsIdx, idx)
	sl.symbols = append(sl.symbols, c)
}

func (sl schematicLine) partsSum(s schematic) int {
	sum := 0
	for _, pn := range sl.partNumbers {
		if pn.hasNearbySymbol(s) {
			sum += pn.number
		}
	}
	return sum
}

func (sl schematicLine) gearRatio(s schematic) int {
	ratio := 0
	for idx, symbol := range sl.symbols {
		if symbol != '*' {
			continue
		}
		symbolIdx := sl.symbolsIdx[idx]

		nearbyPartNumbers := []partNumber{}
		for i := sl.line - 1; i <= sl.line+1; i++ {
			if i >= 0 && i < len(s.lines) {
				nearbyPartNumbers = append(nearbyPartNumbers, s.lines[i].partNumbers...)
			}
		}

		adjacentPartNumbers := []partNumber{}
		for _, p := range nearbyPartNumbers {
			if symbolIdx >= p.location[0]-1 && symbolIdx <= p.location[1]+1 {
				adjacentPartNumbers = append(adjacentPartNumbers, p)
			}
		}

		if len(adjacentPartNumbers) == 2 {
			ratio += adjacentPartNumbers[0].number * adjacentPartNumbers[1].number
		}
	}
	return ratio
}

type schematic struct {
	lines []schematicLine
}

func (s *schematic) parseLine(line string) {
	currentLine := len(s.lines)
	currentPartNumber := partNumber{line: currentLine}
	currentSchematicLine := schematicLine{line: currentLine}

	for idx, c := range line + "." {
		if c >= '0' && c <= '9' {
			currentPartNumber.accumulate(idx, c)
			continue
		} else {
			currentSchematicLine.addPartNumber(idx, currentPartNumber)
			currentPartNumber = partNumber{line: currentLine}
		}

		if c != '.' {
			currentSchematicLine.addSymbol(idx, c)
		}
	}

	s.lines = append(s.lines, currentSchematicLine)
}

func (s schematic) partsSum() int {
	sum := 0
	for _, sl := range s.lines {
		sum += sl.partsSum(s)
	}
	return sum
}

func (s schematic) gearRatio() int {
	ratio := 0
	for _, sl := range s.lines {
		ratio += sl.gearRatio(s)
	}
	return ratio
}

func main() {
	schematic := schematic{}

	for line := range utils.ReadInput() {
		schematic.parseLine(line)
	}

	fmt.Println("Parts:", schematic.partsSum())
	fmt.Println("Gear ratio:", schematic.gearRatio())
}
