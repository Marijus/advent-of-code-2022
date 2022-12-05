package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Marijus/advent-of-code-2022/common"
)

func main() {
	input := common.GetInput("../inputs/input.txt")

	//partOne(input)
	partTwo(input)
}

type Stack struct {
	items []string
}

func (s *Stack) insert(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) push(item string) {
	s.items = append([]string{item}, s.items...)
}

func (s *Stack) pop() string {
	result := s.items[0]
	s.items = append(s.items[:0], s.items[1:]...)
	return result
}

func partTwo(input string) {
	lines := strings.Split(input, "\n")
	stacks, lineIndex := getParsedData(input)

	for l := lineIndex + 1; l < len(lines); l++ {
		line := lines[l]

		if strings.TrimSpace(line) == "" {
			continue
		}

		move, from, to := parseInstruction(line)

		var poppedItems []string
		for i := 0; i < move; i++ {
			popped := stacks[from].pop()
			poppedItems = append([]string{popped}, poppedItems...)
		}

		for _, poppedItem := range poppedItems {
			stacks[to].push(poppedItem)
		}

		for _, stack := range stacks {
			fmt.Println(stack.items)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack.pop())
	}
}

func partOne(input string) {
	lines := strings.Split(input, "\n")
	stacks, lineIndex := getParsedData(input)

	for l := lineIndex + 1; l < len(lines); l++ {
		line := lines[l]
		if strings.TrimSpace(line) == "" {
			continue
		}

		move, from, to := parseInstruction(line)

		for i := 0; i < move; i++ {
			popped := stacks[from].pop()
			stacks[to].push(popped)
		}

		for _, stack := range stacks {
			fmt.Println(stack.items)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack.pop())
	}
}

func getParsedData(input string) (stacks []Stack, lineIndex int) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			lineIndex = i - 1
		}
	}

	indexLine := lines[lineIndex]
	for j := 0; j < len(indexLine); j++ {
		substr := indexLine[j : j+1]
		if substr != " " {
			newStack := Stack{}
			for i := 0; i < lineIndex; i++ {
				lineSubstr := lines[i][j : j+1]

				if lineSubstr != " " {
					newStack.insert(lineSubstr)
				}

			}

			stacks = append(stacks, newStack)
		}
	}

	return stacks, lineIndex
}

func parseInstruction(line string) (move, from, to int) {
	fmt.Println(line)
	items := strings.Split(line, " ")
	move, _ = strconv.Atoi(items[1])

	from, _ = strconv.Atoi(items[3])
	from--

	to, _ = strconv.Atoi(items[5])
	to--

	return
}
