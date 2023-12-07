package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d6partInput = `Time:      7  15   30
Distance:  9  40  200
`

func TestWins(t *testing.T) {
	r := race{time: 7, distance: 9}
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{1, false},
	}

	for _, test := range tests {
		num := r.wins(test.input)
		assert.Equal(t, test.expected, num)

	}

}

func Test_2023_Day_6_Part_1_Example(t *testing.T) {
	result := y2023d6part1(y2023d6partInput)
	assert.Equal(t, "288", result)
}

func Test_2023_Day_6_Part_2_Example(t *testing.T) {
	result := y2023d6part2(y2023d6partInput)
	assert.Equal(t, "71503", result)
}
