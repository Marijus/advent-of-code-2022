package main

import (
	"fmt"
	"strings"

	"github.com/Marijus/advent-of-code-2022/common"
)

func main() {
	input := common.GetInput("../inputs/input.txt")

	//partOne(input)
	partTwo(input)
}

func partOne(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(strings.TrimSpace(line), " ")
		opponent := splitted[0]
		me := splitted[1]
		winnerPoints := getWinnerPoints(opponent, me)
		selectionPoints := getSelectionPoints(me)

		fmt.Println(winnerPoints, selectionPoints)
		fmt.Println(winnerPoints + selectionPoints)
		sum += winnerPoints + selectionPoints
	}

	fmt.Println(sum)
}

func partTwo(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(strings.TrimSpace(line), " ")
		opponent := splitted[0]
		me := getNewValue(opponent, splitted[1])

		winnerPoints := getWinnerPoints(opponent, me)
		selectionPoints := getSelectionPoints(me)

		fmt.Println(winnerPoints, selectionPoints)
		fmt.Println(winnerPoints + selectionPoints)
		sum += winnerPoints + selectionPoints
	}

	fmt.Println(sum)
}

func getNewValue(a, b string) string {
	// A - Rock
	// B - Papper
	// C - sciccors

	// X - Rock / loose
	// Y - paper / draw
	// Z - scissors / win

	if b == "X" {
		switch a {
		case "A":
			return "Z"
		case "B":
			return "X"
		case "C":
			return "Y"
		}
	}

	if b == "Y" {
		switch a {
		case "A":
			return "X"
		case "B":
			return "Y"
		case "C":
			return "Z"
		}
	}

	if b == "Z" {
		switch a {
		case "A":
			return "Y"
		case "B":
			return "Z"
		case "C":
			return "X"
		}
	}

	return ""
}

func getSelectionPoints(b string) int {
	if b == "Y" {
		return 2
	}

	if b == "X" {
		return 1
	}

	return 3
}

func getWinnerPoints(a, b string) int {
	// A - Rock
	// B - Papper
	// C - sciccors

	// X - Rock
	// Y - paper
	// Z - scissors
	if (a == "A" && b == "X") || (a == "B" && b == "Y") || (a == "C" && b == "Z") {
		return 3
	}

	// rock/paper +
	if a == "A" && b == "Y" {
		return 6
	}

	// rock/sciccors
	if a == "A" && b == "Z" {
		return 0
	}

	// paper/rock +
	if a == "B" && b == "X" {
		return 0
	}

	// paper/scissors
	if a == "B" && b == "Z" {
		return 6
	}

	// scissors/paper
	if a == "C" && b == "Y" {
		return 0
	}

	// scissors/rock
	if a == "C" && b == "X" {
		return 6
	}

	return 0
}
