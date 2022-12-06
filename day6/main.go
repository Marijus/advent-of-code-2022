package main

import (
	"fmt"

	"github.com/Marijus/advent-of-code-2022/common"
)

func main() {
	input := common.GetInput("../inputs/input.txt")

	//findMarker(input, 4)
	findMarker(input, 14)
}

func findMarker(input string, length int) {
	for i := length; i < len(input); i++ {
		substr := input[i-length : i]
		if allCharsDifferent(substr) {
			fmt.Println("result: ", i)
			return
		}
	}
}

func allCharsDifferent(s string) bool {
	charsMap := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		substr := s[i : i+1]

		_, found := charsMap[substr]

		if !found {
			charsMap[substr] = true
		} else {
			return false
		}
	}

	return true
}
