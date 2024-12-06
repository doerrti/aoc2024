package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

var Day2_1 = Puzzle{
	Filepath: "inputs/day02.txt",
	Solve: func(input string) string {
		lines := strings.Split(input, "\n")

		var safeCount int = 0
		lines = lines[:len(lines)-1]
		for _, line := range lines {
			numbers := strings.Split(line, " ")

			previous := 0
			previousSign := 0
			isSafe := true

			firstNum, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			secondNum, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			difference := secondNum - firstNum
			if difference == 0 || abs(difference) > 3 {
				isSafe = false
				continue
			}
			if difference > 0 {
				previousSign = 1
			} else {
				previousSign = -1
			}
			previous = secondNum
			for _, number := range numbers[2:len(numbers)] {
				converted, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				difference = converted - previous
				if difference == 0 || (difference > 0 && previousSign == -1) || (difference < 0 && previousSign == 1) || (abs(difference) > 3) {
					isSafe = false
					break
				}
				previous = converted
				if difference < 0 {
					previousSign = -1
				} else if difference > 0 {
					previousSign = 1
				}
			}
			if isSafe {
				fmt.Println(line)
				safeCount += 1
			}
		}

		return strconv.FormatInt(int64(safeCount), 10)
	},
}
