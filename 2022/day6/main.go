package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const markerSequenceLength = 14

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	index := 0
	var buffer string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		buffer += scanner.Text()
	}

	for i := 0; i < len(buffer); i++ {
		seen := map[string]bool{}
		seen[string(buffer[i])] = true

		for j := i + 1; j < len(buffer); j++ {
			if i >= markerSequenceLength-1 && len(seen) >= markerSequenceLength {
				index = j
				break
			}

			if !seen[string(buffer[j])] {
				seen[string(buffer[j])] = true
			} else {
				break
			}
		}

		if index > 0 {
			break
		}
	}

	fmt.Println(index)
}
