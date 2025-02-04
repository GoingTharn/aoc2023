package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d3partInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
..........
1*1.......
..........
`

func TestSymbol(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'.', false},
		{'1', false},
		{'%', true},
		{'+', true},
	}

	for _, test := range tests {
		actual := symbol(test.input)
		assert.Equal(t, test.expected, actual)
	}

}

func TestBothDirs(t *testing.T) {
	tests := []struct {
		input    string
		x        int
		expected int
		firstIdx int
		lastIdx  int
	}{
		{"123451", 2, 123451, 0, 5},
		{"123451=.", 2, 123451, 0, 5},
		{"-123451=.", 3, 123451, 0, 6},
		{"12.44", 1, 12, 0, 1},
		{"12.44", 3, 44, 3, 4},
	}

	for _, test := range tests {
		num, lastIdx := getNumberBothDirs(test.x, []rune(test.input))
		assert.Equal(t, test.expected, num)
		assert.Equal(t, test.lastIdx, lastIdx)

	}

}

func Test_2023_Day_3_Part_1_Example(t *testing.T) {
	result := y2023d3part1(y2023d3partInput)
	assert.Equal(t, "4363", result)
}

func Test_2023_Day_3_Part_2_Example(t *testing.T) {
	result := y2023d3part2(y2023d3partInput)
	assert.Equal(t, "467836", result)
}
