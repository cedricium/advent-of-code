package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sums := []int{}
	sum := 0
	for more := scanner.Scan(); more; more = scanner.Scan() {
		var num int
		line := scanner.Text()

		if len(line) > 0 {
			if _, err := fmt.Sscan(line, &num); err == nil {
				sum += num
			}
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	sort.Ints(sums)
	max := 0
	for _, num := range sums[len(sums)-3:] {
		max += num
	}

	fmt.Println(max)
}
