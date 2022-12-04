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
	sum := 0

	for open := scanner.Scan(); open; {
		group := []string{}

		for i := 0; i < 3; i++ {
			rucksack := scanner.Text()
			group = append(group, rucksack)
			open = scanner.Scan()
		}

		sort.Strings(group)
		common := findGroupCommonItem(group)
		priority := calculatePriority(common)
		sum += priority
	}

	fmt.Println(sum)
}

func findGroupCommonItem(a []string) (common rune) {
	mid := map[rune]bool{}
	for _, char := range a[1] {
		mid[char] = true
	}

	long := map[rune]bool{}
	for _, char := range a[2] {
		long[char] = true
	}

	for _, char := range a[0] {
		if mid[char] && long[char] {
			common = char
			return
		}
	}

	return
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
