package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	WIN_SCORE  int = 6
	DRAW_SCORE int = 3

	SCORE_ROCK     int = 1
	SCORE_PAPER    int = 2
	SCORE_SCISSORS int = 3

	ROCK     string = "A"
	PAPER    string = "B"
	SCISSORS string = "C"

	LOSE string = "X"
	DRAW string = "Y"
	WIN  string = "Z"
)

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sum := 0
	for open := scanner.Scan(); open; open = scanner.Scan() {
		round := scanner.Text()
		score := 0

		if round != "" {
			strategy := strings.Split(round, " ")
			opponentShape := strategy[0]
			desiredOutcome := strategy[1]

			shape := determineShapeByOutcome(opponentShape, desiredOutcome)
			score += calculateShapeScore(shape)
			score += calculateRoundOutcome([]string{opponentShape, shape})
			sum += score
		}
	}

	fmt.Println(sum)
}

func calculateShapeScore(shape string) int {
	score := 0

	switch shape {
	case ROCK, "X":
		score += SCORE_ROCK
	case PAPER, "Y":
		score += SCORE_PAPER
	case SCISSORS, "Z":
		score += SCORE_SCISSORS
	}

	return score
}

func calculateRoundOutcome(shapes []string) int {
	score := 0

	switch shapes[1] {
	case ROCK, "X":
		if shapes[0] == ROCK {
			score += DRAW_SCORE
		} else if shapes[0] == SCISSORS {
			score += WIN_SCORE
		}
	case PAPER, "Y":
		if shapes[0] == ROCK {
			score += WIN_SCORE
		} else if shapes[0] == PAPER {
			score += DRAW_SCORE
		}
	case SCISSORS, "Z":
		if shapes[0] == PAPER {
			score += WIN_SCORE
		} else if shapes[0] == SCISSORS {
			score += DRAW_SCORE
		}
	}

	return score
}

func determineShapeByOutcome(shape string, outcome string) string {
	switch outcome {
	case WIN:
		if shape == ROCK {
			return PAPER
		} else if shape == PAPER {
			return SCISSORS
		} else {
			return ROCK
		}
	case DRAW:
		return shape
	case LOSE:
		if shape == ROCK {
			return SCISSORS
		} else if shape == PAPER {
			return ROCK
		} else {
			return PAPER
		}
	default:
		return ""
	}
}
