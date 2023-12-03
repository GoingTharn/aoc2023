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
	for y, runes := range array {
		var lastR rune
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
	fmt.Printf("Working: %s, Incomplete: %t, lastIdx: %d\n", string(working), incomplete, lastIdx)

	// edge number, never "completes"
	if incomplete {
		lastIdx = len(array[y]) - 1
	}

	strungNum := string(working)
	outNum, err := strconv.Atoi(strungNum)
	if err != nil {
		panic(err)
	}
	return outNum, lastIdx
}

func getNumberBothDirs(x int, runeList []rune) (outNum int, lastIdx int) {
	var working []rune
	incomplete := true
	// backwards first
	for i := x; incomplete && i >= 0; i-- {
		r := runeList[i]
		if unicode.IsNumber(r) {
			working = append(working, r)
		} else {
			incomplete = false
		}
	}
	// have firstIdx of number AND first digits
	first := lo.Reverse[rune](working)

	incomplete = true
	for i := x + 1; incomplete && i < len(runeList); i++ {
		r := runeList[i]
		if unicode.IsNumber(r) {
			first = append(first, r)
		} else {
			lastIdx = i - 1
			incomplete = false
		}
	}

	if incomplete {
		lastIdx = len(runeList) - 1
	}

	strungNum := string(first)
	outNum, err := strconv.Atoi(strungNum)
	if err != nil {
		panic(err)
	}
	return outNum, lastIdx
}

func symbol(r rune) bool {
	return !unicode.IsNumber(r) && r != '.'
}

func checkRow(x int, runes []rune) []int {
	var ints []int
	if x != 0 && unicode.IsNumber(runes[x-1]) {
		num, lastIdx := getNumberBothDirs(x-1, runes)
		ints = append(ints, num)
		if lastIdx == x-1 {
			// number ends on corner, possible to have a second
			if !unicode.IsNumber(runes[x]) && x+1 < len(runes) && unicode.IsNumber(runes[x+1]) {
				second, _ := getNumberBothDirs(x+1, runes)
				ints = append(ints, second)

			}
		}
		return ints
	}
	for i := x; i <= x+1 && i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) {
			num, lastIdx := getNumberBothDirs(i, runes)
			ints = append(ints, num)
			if lastIdx >= i+1 {
				return ints
			}
		}
	}
	return ints
}

func adjNum(x, y int, array [][]rune) int {
	var working []int
	// test above
	if y != 0 {
		working = append(working, checkRow(x, array[y-1])...)
	}
	if y+1 < len(array) {
		working = append(working, checkRow(x, array[y+1])...)
	}
	working = append(working, checkRow(x, array[y])...)
	if len(working) == 2 {
		return working[0] * working[1]
	}
	return 0
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
	for y, runes := range array {
		for x, r := range runes {
			if r == '*' {
				fmt.Printf("Found * - (%d, %d)\n", x, y)
				adjProduct := adjNum(x, y, array)
				numbers = append(numbers, adjProduct)
			}
		}
	}
	sum := lo.Reduce(numbers, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return fmt.Sprint(sum)
}
