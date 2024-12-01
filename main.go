package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")

	var left []int = make([]int, 0, len(lines))
	var right []int = make([]int, 0, len(lines))

	for iLine, line := range lines[:len(lines)-1] {
		leftLength := 0
		for _, char := range lines[iLine] {
			if char == ' ' {
				break
			}
			leftLength += 1
		}
		leftNum, err := strconv.Atoi(line[0:leftLength])
		if err != nil {
			panic(err)
		}

		left = append(left, leftNum)
		for iChar, char := range line[leftLength:] {
			if char != ' ' {
				rightNum, err := strconv.Atoi(lines[iLine][leftLength+iChar:])
				if err != nil {
					fmt.Println(leftNum)
					fmt.Println(leftLength)
					fmt.Println(line)
					panic(err)
				}
				right = append(right, rightNum)
				break
			}
		}
	}

	slices.Sort[[]int](left)
	slices.Sort[[]int](right)

	var sum int = 0
	for i := range left {
		sum += abs(left[i] - right[i])
	}

	fmt.Println(sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
