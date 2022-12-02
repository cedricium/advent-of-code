package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	WIN  int = 6
	DRAW int = 3

	SCORE_ROCK     int = 1
	SCORE_PAPER    int = 2
	SCORE_SCISSORS int = 3

	ROCK     string = "A"
	PAPER    string = "B"
	SCISSORS string = "C"
)

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sum := 0
	open := true
	for open {
		open = scanner.Scan()
		round := scanner.Text()
		score := 0

		if round != "" {
			shapes := strings.Split(round, " ")
			score += calculateShapeScore(shapes[1])
			score += calculateRoundOutcome(shapes)
			sum += score
		}
	}

	fmt.Println(sum)
}

func calculateShapeScore(shape string) int {
	score := 0

	switch shape {
	case "X":
		score += SCORE_ROCK
	case "Y":
		score += SCORE_PAPER
	case "Z":
		score += SCORE_SCISSORS
	}

	return score
}

func calculateRoundOutcome(shapes []string) int {
	score := 0

	switch shapes[1] {
	case "X": // ROCK
		if shapes[0] == ROCK {
			score += DRAW
		} else if shapes[0] == SCISSORS {
			score += WIN
		}
	case "Y": // PAPER
		if shapes[0] == ROCK {
			score += WIN
		} else if shapes[0] == PAPER {
			score += DRAW
		}
	case "Z": // SCISSORS
		if shapes[0] == PAPER {
			score += WIN
		} else if shapes[0] == SCISSORS {
			score += DRAW
		}
	}

	return score
}
