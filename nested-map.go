package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstLetter := rune(name[0])
		_, ok := counts[firstLetter]

		if !ok {
			counts[firstLetter] = make(map[string]int)
		}

		counts[firstLetter][name]++
	}
	return counts
}

func main() {
	fmt.Printf("Nested Loop\n")

	name := []string{
		"ukasha",
		"urwah",
		"inshal",
		"ibad",
		"ukasha",
		"muaaz",
		"maaz",
	}

	maps_counter := getNameCounts(name)
	fmt.Println(maps_counter)
}
