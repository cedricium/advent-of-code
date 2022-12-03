package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)
	sum := 0

	open := scanner.Scan()
	for open {
		rucksack := scanner.Text()
		left := rucksack[:len(rucksack)/2]
		right := rucksack[len(rucksack)/2:]

		common := findCommonItem(left, right)
		priority := calculatePriority(common)
		sum += priority

		open = scanner.Scan()
	}

	fmt.Println(sum)
}

func findCommonItem(a, b string) (common rune) {
	set := map[rune]bool{}

	for _, char := range a {
		set[char] = true
	}

	for _, char := range b {
		_, found := set[char]
		if found {
			common = char
			break
		}
	}

	return
}

// Given char, returns the corresponding priority integer as outlined in the
// instructions, where lowercase a-z should return 1-26 and A-Z returns 27-52.
// The int value of char is used to determine the casing.
func calculatePriority(char rune) int {
	if char > 90 { // lowercase
		return int(char) - 96
	} else { // uppercase
		return int(char) - 38
	}
}
