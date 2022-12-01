package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	max := 0
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sum := 0
	for {
		var num int
		line := scanner.Text()

		if len(line) > 0 {
			if _, err := fmt.Sscan(line, &num); err == nil {
				sum += num
			}
		} else {
			if sum > max {
				max = sum
			}
			sum = 0
		}

		if more := scanner.Scan(); more == false {
			break
		}
	}

	fmt.Println(max)
}
