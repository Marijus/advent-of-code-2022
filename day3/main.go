package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Marijus/advent-of-code-2022/common"
)

func main() {
	input := common.GetInput("../inputs/input.txt")

	//partOne(input)
	partTwo(input)
}

func partTwo(input string) {
	lines := strings.Split(input, "\n")

	sum := 0
	for i := 0; i < len(lines)/3; i++ {
		a := lines[i*3]
		b := lines[i*3+1]
		c := lines[i*3+2]
		fmt.Println(a, b, c)

		founnd := findCommonInLines(a, b, c)
		fmt.Println(founnd)

		if unicode.IsLower(rune(founnd[0])) {
			sum += int(founnd[0] - uint8(96))
		} else {
			sum += int(founnd[0] - uint8(38))
		}
	}

	fmt.Println(sum)
}

func findCommonInLines(a, b, c string) string {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			for z := 0; z < len(c); z++ {
				if a[i] == b[j] && b[j] == c[z] {
					return a[i : i+1]
				}
			}
		}
	}

	panic("no")
}

func partOne(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		a := line[0 : len(line)/2]
		b := line[len(line)/2:]
		c := findCommon(a, b)
		if unicode.IsLower(rune(c[0])) {
			sum += int(c[0] - uint8(96))
		} else {
			sum += int(c[0] - uint8(38))
		}
	}

	fmt.Println(sum)
}

func findCommon(a, b string) string {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				return a[i : i+1]
			}
		}
	}

	panic("no")
	return ""
}
