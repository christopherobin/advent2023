package main

import (
	"fmt"

	"github.com/christopherobin/advent2023/pkg/utils"
)

func main() {
	for line := range utils.ReadInput() {
		fmt.Println(line)
	}
}
