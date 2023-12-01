package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:1:1", y2023d1part1)
	registerSolution("2023:1:2", y2023d1part2)

}

func y2023d1part1(input string) string {
	numbers := "0123456789"
	lines := strings.Split(input, "\n")
	digits := make([]int, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("Skipping empty line")
			continue
		}
		working := strings.Split(line, "")
		fmt.Print("Working: ")
		fmt.Println(working)
		justDigits := lo.Filter[string](working, func(x string, index int) bool {
			return strings.Contains(numbers, x)
		})
		first := getDigit(justDigits)
		reversed := lo.Reverse(justDigits)
		last := getDigit(reversed)
		twoDigits := strings.Join([]string{first, last}, "")
		asNum, err := strconv.Atoi(twoDigits)
		if err != nil {
			panic(err)
		}
		digits = append(digits, asNum)
	}

	sum := lo.Reduce(digits, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return fmt.Sprint(sum)
}

func getTextDigit(input string, last bool) (digit string, index int) {
	digitsAsText := map[string]string{"one": "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
	}

	var lowestIdx int
	if !last {
		lowestIdx = 9999
	} else {
		lowestIdx = -1
	}

	var first string
	for key := range digitsAsText {
		if !last {
			idx := strings.Index(input, key)
			if idx != -1 && idx < lowestIdx {
				lowestIdx = idx
				first = key
			}
		} else {
			idx := strings.LastIndex(input, key)
			if idx != -1 && idx > lowestIdx {
				lowestIdx = idx
				first = key
			}
		}
	}

	return digitsAsText[first], lowestIdx
}

func getDigit(input []string) string {
	var candidate []string = lo.Slice[string](input, 0, 1)
	return candidate[0]
}

func y2023d1part2(input string) string {
	lines := strings.Split(input, "\n")
	var digits []int

	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("Skipping empty line")
			continue
		}

		first, _ := getTextDigit(line, false)
		last, _ := getTextDigit(line, true)

		answer := strings.Join([]string{first, last}, "")

		fmt.Printf("Line: %s, Answer: %s\n", line, answer)
		asNum, err := strconv.Atoi(answer)
		if err != nil {
			panic(err)
		}
		digits = append(digits, asNum)
	}

	fmt.Println(digits)
	sum := lo.Reduce(digits, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return fmt.Sprint(sum)
}
