package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)
	stacks := map[int]string{}

	// configure stacks of crates
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")

		for i := 1; i < len(line); i += 4 {
			stackNumber := i / 4

			crate := string(line[i])
			if crate != " " {
				stacks[stackNumber] += crate
			}
		}

		if line == "" {
			break
		}
	}

	// rearrange stacks of crates
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		var count, from, to int
		fmt.Sscan(tokens[1], &count)
		fmt.Sscan(tokens[3], &from)
		fmt.Sscan(tokens[5], &to)

		for i := 0; i < count; i++ {
			crate := string(stacks[from-1][0])
			stacks[from-1] = stacks[from-1][1:]
			stacks[to-1] = crate + stacks[to-1]
		}
	}

	// LAST CRATE PER STACK
	for i := 0; i < len(stacks); i++ {
		fmt.Print(string(stacks[i][0]))
	}
}
