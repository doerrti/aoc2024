package aoc

import (
	"os"
)

type Puzzle struct {
	Filepath string
	Input    string
	Solve    func(input string) string
}

func (puzzle *Puzzle) LoadInput() {
	f, err := os.ReadFile(puzzle.Filepath)
	if err != nil {
		panic(err)
	}
	puzzle.Input = string(f)
}
