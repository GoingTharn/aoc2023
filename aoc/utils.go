package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func stringsToInts(input []string) []int {
	var out []int
	for _, cand := range input {
		if len(strings.TrimSpace(cand)) == 0 {
			continue
		}
		asInt, err := strconv.Atoi(cand)
		if err != nil {
			fmt.Println(err)
			continue
		}
		out = append(out, asInt)
	}
	return out
}

func splitThenConvert(line string, headerSplit string) []int {
	// splitThenConvert assumes a list of ints after a header value
	idx := strings.Index(line, headerSplit) + 1
	return stringsToInts(strings.Split(line[idx:], " "))
}

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
