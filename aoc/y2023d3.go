package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:3:1", y2023d3part1)
	registerSolution("2023:3:2", y2023d3part2)

}

func y2023d3part1(input string) string {
	var array [][]rune
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		lineArray := make([]rune, len(line))
		for x, r := range line {
			lineArray[x] = r
		}
		array = append(array, lineArray)
	}

	var numbers []int
	var lastR rune
	for y, runes := range array {
		for x, r := range runes {
			if unicode.IsNumber(r) && !unicode.IsNumber(lastR) {
				newNum, endX := getNumber(x, y, array)
				fmt.Printf("Num: %d, endX: %d\n", newNum, endX)
				if nearSymbol(x, endX, y, array) {
					numbers = append(numbers, newNum)
				}
			}
			lastR = r
		}
	}
	fmt.Println(numbers)
	sum := lo.Reduce(numbers, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return fmt.Sprint(sum)
}

func getNumber(x, y int, array [][]rune) (outNum int, lastIdx int) {
	var working []rune
	incomplete := true
	for i := x; incomplete && i <= len(array[y])-1; i++ {
		r := array[y][i]
		if unicode.IsNumber(r) {
			working = append(working, r)
		} else {
			lastIdx = i - 1
			incomplete = false
		}
	}
	strungNum := string(working)
	outNum, err := strconv.Atoi(strungNum)
	if err != nil {
		panic(err)
	}
	return outNum, lastIdx
}

func symbol(r rune) bool {
	return !unicode.IsNumber(r) && r != '.'
}

func nearSymbol(startX, endX, y int, array [][]rune) bool {
	// capture bounds of X
	if startX != 0 {
		startX = startX - 1
	}

	if endX != len(array[y])-1 {
		endX = endX + 1
	}
	// test edges of number
	if symbol(array[y][startX]) || symbol(array[y][endX]) {
		// found a symbol, skip checking remaining bounds
		return true
	}

	fmt.Printf("Y: %d startX: %d endX: %d\n", y, startX, endX)
	for i := startX; i <= endX; i++ {
		// line above/below. Skip bounds
		if y != 0 && symbol(array[y-1][i]) {
			return true
		} else if y != len(array)-1 && symbol(array[y+1][i]) {
			return true
		}
	}
	return false
}

func y2023d3part2(input string) string {
	return "wrong again"
}
