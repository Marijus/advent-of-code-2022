package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Marijus/advent-of-code-2022/common"
)

func main() {
	input := common.GetInput("../inputs/input.txt")

	//partOne(input)
	partTwo(input)
}

func partTwo(input string) {
	var elfs []int
	currentElf := 0
	for _, line := range strings.Split(input, "\n") {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			elfs = append(elfs, currentElf)
			currentElf = 0
			continue
		}

		count, err := strconv.Atoi(trimmedLine)
		if err != nil {
			panic(err)
		}
		currentElf += count
	}

	sort.Ints(elfs)
	fmt.Println(elfs)
	sum := 0
	for _, item := range elfs[len(elfs)-3:] {
		sum += item
	}
	fmt.Println(sum)
}

func partOne(input string) {
	maxElf := 0

	currentElf := 0
	for _, line := range strings.Split(input, "\n") {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			if currentElf > maxElf {
				maxElf = currentElf
			}

			currentElf = 0
			continue
		}

		count, err := strconv.Atoi(trimmedLine)
		if err != nil {
			panic(err)
		}
		currentElf += count
	}

	fmt.Println(maxElf)
}
