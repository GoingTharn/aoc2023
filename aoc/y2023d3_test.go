package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d3partInput = `
467..114..
...*.....1
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
.....*....
$664.598..
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

func Test_2023_Day_3_Part_1_Example(t *testing.T) {
	result := y2023d3part1(y2023d3partInput)
	assert.Equal(t, "4361", result)
}

func Test_2023_Day_3_Part_2_Example(t *testing.T) {
	result := y2023d3part2(y2023d3partInput)
	assert.Equal(t, "still right!", result)
}
