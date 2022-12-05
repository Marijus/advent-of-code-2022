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

func partTwo(input string) {
	//fmt.Println(input)
	count := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, ",")
		first := strings.Split(splitted[0], "-")
		second := strings.Split(splitted[1], "-")
		a, _ := strconv.Atoi(first[0])
		b, _ := strconv.Atoi(first[1])
		c, _ := strconv.Atoi(second[0])
		d, _ := strconv.Atoi(second[1])

		for i := a; i <= b; i++ {
			if c <= i && i <= d {
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

func partOne(input string) {
	//fmt.Println(input)
	count := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, ",")
		first := strings.Split(splitted[0], "-")
		second := strings.Split(splitted[1], "-")
		a, _ := strconv.Atoi(first[0])
		b, _ := strconv.Atoi(first[1])
		c, _ := strconv.Atoi(second[0])
		d, _ := strconv.Atoi(second[1])

		if c >= a && d <= b {
			fmt.Println(a, b, c, d)
			count++
		} else if a >= c && b <= d {
			fmt.Println(a, b, c, d)
			count++
		}
		//fmt.Println(line)
	}

	fmt.Println(count)
}
