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
	sum := 0

	for open := scanner.Scan(); open; open = scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		left := strings.Split(pairs[0], "-")
		right := strings.Split(pairs[1], "-")

		if overlaps := checkForOverlap(left, right); overlaps {
			fmt.Println(left, right, overlaps)
			sum += 1
		}
	}

	fmt.Println(sum)
}

func checkForOverlap(a, b []string) bool {
	var aStart, bStart, aEnd, bEnd int
	fmt.Sscan(a[0], &aStart)
	fmt.Sscan(a[1], &aEnd)
	fmt.Sscan(b[0], &bStart)
	fmt.Sscan(b[1], &bEnd)

	return aStart >= bStart && aStart <= bEnd && aEnd <= bEnd ||
		bStart >= aStart && bStart <= aEnd && bEnd <= aEnd
}
