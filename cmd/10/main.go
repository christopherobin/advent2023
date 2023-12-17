package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/christopherobin/advent2023/pkg/utils"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// valid connections, the 4 strings are left, right, up, down
var connections = map[string][4]string{
	"|": {"", "", "|7FS", "|JLS"},
	"-": {"-FLS", "-7JS", "", ""},
	"L": {"", "-7JS", "|7FS", ""},
	"J": {"-FLS", "", "|7FS", ""},
	"F": {"", "-J7S", "", "|JLS"},
	"7": {"-FLS", "", "", "|JLS"},
	"S": {"-FL", "-J7", "|7F", "|JL"},
}

type Node struct {
	Type string
	X    int
	Y    int
}

func (n Node) Connects(to Node) bool {
	if n.X != to.X && n.Y != to.Y {
		return false
	}

	if (abs(n.X-to.X) + abs(n.Y-to.Y)) > 1 {
		return false
	}

	if to.Type == "." || n.Type == "." {
		return false
	}

	if to.X < n.X {
		return strings.Contains(connections[n.Type][0], to.Type)
	}

	if to.X > n.X {
		return strings.Contains(connections[n.Type][1], to.Type)
	}

	if to.Y < n.Y {
		return strings.Contains(connections[n.Type][2], to.Type)
	}

	return strings.Contains(connections[n.Type][3], to.Type)
}

type PipeMap struct {
	Data        string
	Width       int
	Height      int
	Start       Node
	Network     map[int]map[int]Node
	Loop        []Node
	Orientation string
}

func (pm PipeMap) Get(x, y int) (Node, bool) {
	if x < 0 || x >= pm.Width || y < 0 || y*pm.Width >= len(pm.Data) {
		return Node{".", x, y}, false
	}

	return Node{string(pm.Data[y*pm.Width+x]), x, y}, true
}

func (pm *PipeMap) Next(n Node, from Node) Node {
	neighbors := pm.Neighbors(n, true, false)
	if len(neighbors) != 2 {
		fmt.Println(n, from, neighbors)
		panic("should have 2 neighbors")
	}
	next := neighbors[0]
	if neighbors[0] == from {
		next = neighbors[1]
	}
	pm.AddToNetwork(next)

	return next
}

func (from Node) Direction(to Node) int {
	switch from.Type {
	case "|":
		if to.Type == "F" || to.Type == "J" {
			return 1
		}
		if to.Type == "7" || to.Type == "L" {
			return -1
		}
	case "-":
		if to.Type == "7" || to.Type == "L" {
			return 1
		}
		if to.Type == "F" || to.Type == "J" {
			return -1
		}
	case "L":
		if to.Type == "7" {
			if to.Y > from.Y {
				return -1
			}
			return 1
		}
		if to.Type == "F" {
			return 1
		}
		if to.Type == "J" {
			return -1
		}
	case "J":
		if to.Type == "F" {
			if to.Y > from.Y {
				return 1
			}
			return -1
		}
		if to.Type == "7" {
			return -1
		}
		if to.Type == "L" {
			return 1
		}
	case "F":
		if to.Type == "J" {
			if to.Y < from.Y {
				return 1
			}
			return -1
		}
		if to.Type == "7" {
			return 1
		}
		if to.Type == "L" {
			return -1
		}
	case "7":
		if to.Type == "L" {
			if to.Y < from.Y {
				return -1
			}
			return 1
		}
		if to.Type == "F" {
			return -1
		}
		if to.Type == "L" {
			return 1
		}
	case "S":
		if to.Type == "L" {
			if to.Y < from.Y {
				return -1
			}
			return 1
		}
		if to.Type == "J" {
			if to.Y < from.Y {
				return 1
			}
			return -1
		}
		if to.Type == "F" {
			if to.Y > from.Y {
				return 1
			}
			return -1
		}
		if to.Type == "7" {
			if to.Y > from.Y {
				return -1
			}
			return 1
		}
	}

	return 0
}

func (pm PipeMap) GetAndCheck(from Node, x, y int, connects, networkOnly bool) (Node, bool) {
	node, ok := pm.Get(x, y)
	if !ok {
		return Node{".", x, y}, false
	}
	if networkOnly {
		if networkNode, ok := pm.Network[x][y]; ok {
			return networkNode, true
		}
		return Node{".", x, y}, true
	}
	if !connects || from.Connects(node) {
		return node, true
	}
	return Node{".", x, y}, false
}

func (pm PipeMap) Neighbors(from Node, connects bool, networkOnly bool) []Node {
	res := []Node{}

	if pipe, ok := pm.GetAndCheck(from, from.X-1, from.Y, connects, networkOnly); ok {
		res = append(res, pipe)
	}
	if pipe, ok := pm.GetAndCheck(from, from.X+1, from.Y, connects, networkOnly); ok {
		res = append(res, pipe)
	}
	if pipe, ok := pm.GetAndCheck(from, from.X, from.Y-1, connects, networkOnly); ok {
		res = append(res, pipe)
	}
	if pipe, ok := pm.GetAndCheck(from, from.X, from.Y+1, connects, networkOnly); ok {
		res = append(res, pipe)
	}

	return res
}

func (pm *PipeMap) AddToNetwork(n Node) bool {
	if _, ok := pm.Network[n.X]; !ok {
		pm.Network[n.X] = map[int]Node{}
	}
	if _, ok := pm.Network[n.X][n.Y]; ok {
		return false
	}
	pm.Network[n.X][n.Y] = n
	return true
}

func (pm *PipeMap) FindFurthest() int {
	var connections []Node = pm.Neighbors(pm.Start, true, false)

	if len(connections) != 2 {
		panic("S should only have 2 connections")
	}

	loopa, loopb := []Node{pm.Start}, []Node{}
	pa, pb := pm.Start, pm.Start
	a, b := connections[0], connections[1]
	pm.AddToNetwork(a)
	pm.AddToNetwork(b)
	i := 1
	for {
		i++
		loopa = append(loopa, a)
		loopb = append([]Node{b}, loopb...)
		na := pm.Next(a, pa)
		nb := pm.Next(b, pb)
		pa, pb, a, b = a, b, na, nb
		if a == b {
			loopa = append(loopa, a)
			break
		}
	}

	pm.Loop = append(loopa, loopb...)
	return i
}

func (pm *PipeMap) FindContained() int {
	// find the direction of the loop, >0 = clockwise, <0 = ccw
	direction := 0
	for i := 1; i < len(pm.Loop); i++ {
		direction += pm.Loop[i-1].Direction(pm.Loop[i])
	}
	if direction == 0 {
		panic("loop must have a direction")
	}

	// then go again and try to find a . that is on the inside relative to the direction
	for i := 1; i < len(pm.Loop); i++ {
		prev := pm.Loop[i-1]
		current := pm.Loop[i]
		if !strings.Contains("-|7LJ", current.Type) {
			continue
		}

		maybeDot := Node{}
		if current.Type == "-" {
			revert := 1
			if prev.X > current.X {
				revert = -1
			}
			if direction > 0 { // cw
				if neighbor, ok := pm.GetAndCheck(current, current.X, current.Y+1*revert, false, true); ok {
					maybeDot = neighbor
				}
			} else {
				if neighbor, ok := pm.GetAndCheck(current, current.X, current.Y-1*revert, false, true); ok {
					maybeDot = neighbor
				}
			}
		}
		if current.Type == "|" {
			revert := 1
			if prev.Y < current.Y {
				revert = -1
			}
			if direction > 0 { // cw
				if neighbor, ok := pm.GetAndCheck(current, current.X+1*revert, current.Y, false, true); ok {
					maybeDot = neighbor
				}
			} else {
				if neighbor, ok := pm.GetAndCheck(current, current.X-1*revert, current.Y, false, true); ok {
					maybeDot = neighbor
				}
			}
		}
		// edge cases for 7LFJ
		if current.Type == "7" {
			if direction > 0 && prev.X < current.X {
				continue
			}

			if direction < 0 && prev.Y > current.Y {
				continue
			}

			if neighbor, ok := pm.GetAndCheck(current, current.X+1, current.Y, false, true); ok {
				maybeDot = neighbor
			}
		}
		if current.Type == "L" {
			if direction > 0 && prev.X > current.X {
				continue
			}

			if direction < 0 && prev.Y < current.Y {
				continue
			}

			if neighbor, ok := pm.GetAndCheck(current, current.X-1, current.Y, false, true); ok {
				maybeDot = neighbor
			}
		}
		if current.Type == "J" {
			if direction > 0 && prev.Y < current.Y {
				continue
			}

			if direction < 0 && prev.X < current.X {
				continue
			}

			if neighbor, ok := pm.GetAndCheck(current, current.X+1, current.Y, false, true); ok {
				maybeDot = neighbor
			}
		}
		if current.Type == "F" {
			if direction > 0 && prev.Y > current.Y {
				continue
			}

			if direction < 0 && prev.X > current.X {
				continue
			}

			if neighbor, ok := pm.GetAndCheck(current, current.X-1, current.Y, false, true); ok {
				maybeDot = neighbor
			}
		}

		if maybeDot.X < 0 || maybeDot.X >= pm.Width || maybeDot.Y < 0 || maybeDot.Y*pm.Width >= len(pm.Data) {
			continue
		}

		if maybeDot.Type != "." {
			continue
		}

		// we found a dot inside, just bfs from there and mark them
		pm.MarkInside(maybeDot)
	}

	// finally iterate on all the nodes, count the ones inside
	contained := 0
	for _, row := range pm.Network {
		for _, pipe := range row {
			if pipe.Type == "I" {
				contained++
			}
		}
	}

	return contained
}

func (pm *PipeMap) MarkInside(from Node) {
	var current Node
	queue := []Node{{"I", from.X, from.Y}}
	for {
		if len(queue) == 0 {
			break
		}

		current, queue = queue[0], queue[1:]
		if !pm.AddToNetwork(current) {
			continue
		}

		for _, pipe := range pm.Neighbors(current, false, true) {
			if pipe.X < 0 || pipe.X >= pm.Width || pipe.Y < 0 || pipe.Y*pm.Width >= len(pm.Data) {
				continue
			}
			if pipe.Type == "." {
				queue = append(queue, Node{"I", pipe.X, pipe.Y})
			}
		}
	}
}

func (pm PipeMap) Render() {
	for y := 0; y < pm.Height; y++ {
		for x := 0; x < pm.Width; x++ {
			if pipe, ok := pm.Network[x][y]; ok {
				if pipe.Type == "S" {
					fmt.Printf("\033[31m%s\033[0m", pipe.Type)
					continue
				}
				if pipe.Type == "I" {
					fmt.Printf("\033[32m%s\033[0m", pipe.Type)
					continue
				}

				fmt.Printf(pipe.Type)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	s := time.Now()
	var pipeMap = PipeMap{
		Network: map[int]map[int]Node{},
	}
	i := 0

	for line := range utils.ReadInput() {
		if pipeMap.Width == 0 {
			pipeMap.Width = len(line)
		}

		if pos := strings.Index(line, "S"); pos >= 0 {
			pipeMap.Start = Node{"S", pos, i}
			pipeMap.AddToNetwork(pipeMap.Start)
		}

		pipeMap.Data += line
		i++
	}
	pipeMap.Height = i

	furthest := pipeMap.FindFurthest()
	contained := pipeMap.FindContained()
	pipeMap.Render()

	fmt.Println("furthest", furthest)
	fmt.Println("contained", contained)
	fmt.Println("elapsed", time.Since(s))
}
