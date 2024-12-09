package aoc

import (
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
				safeCount += 1
			}
		}

		return strconv.FormatInt(int64(safeCount), 10)
	},
}

var Day2_2 = Puzzle{
	Filepath: "inputs/day02.txt",
	Solve: func(input string) string {
		lines := strings.Split(input, "\n")

		var safeCount int = 0
		lines = lines[:len(lines)-1]

		for _, line := range lines {
			levelsStrings := strings.Split(line, " ")
			levels := make([]int, 0, len(levelsStrings))

			for _, number := range levelsStrings {
				parsed, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				levels = append(levels, parsed)
			}

			mergedLevels := make([]int, 0, len(levels))

		Levels:
			for i := range levels {
				mergedLevels = mergedLevels[:0]
				mergedLevels = append(mergedLevels, levels[:i]...)
				mergedLevels = append(mergedLevels, levels[i+1:]...)
				previous := mergedLevels[0]
				previousSign := 0
				for _, level := range mergedLevels[1:] {
					if isUnsafeLevel(level, previous, previousSign) {
						continue Levels
					}
					if level-previous < 0 {
						previousSign = -1
					} else if level-previous > 0 {
						previousSign = 1
					}
					previous = level
				}
				safeCount++
				break Levels
			}
		}

		return strconv.FormatInt(int64(safeCount), 10)
	},
}

func isUnsafeLevel(current int, previous int, previousSign int) bool {
	// fmt.Printf("%v - %v = %v | %v\n", current, previous, current-previous, previousSign)
	return (current-previous == 0) || ((current-previous > 0) && previousSign < 0) || ((current-previous < 0) && previousSign > 0) || (abs(current-previous) > 3)
}
