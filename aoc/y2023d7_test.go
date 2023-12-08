package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d7partInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestGetRank(t *testing.T) {

	tests := []struct {
		h            hand
		expectedRank int
	}{
		{h: hand{cards: []card("32T3K")}, expectedRank: 2},
		{h: hand{cards: []card("T55J5")}, expectedRank: 4},
		{h: hand{cards: []card("KK677")}, expectedRank: 3},
		{h: hand{cards: []card("AAAAA")}, expectedRank: 7},
		{h: hand{cards: []card("23456")}, expectedRank: 1},
		{h: hand{cards: []card("AAAAK")}, expectedRank: 6},
		{h: hand{cards: []card("22333")}, expectedRank: 5},
	}

	for _, test := range tests {
		actualRank := test.h.getRank()
		assert.Equal(t, test.expectedRank, actualRank)
	}
}

func Test_2023_Day_7_Part_1_Example(t *testing.T) {
	result := y2023d7part1(y2023d7partInput)
	assert.Equal(t, "6440", result)
}

func Test_2023_Day_7_Part_2_Example(t *testing.T) {
	result := y2023d7part2(y2023d7partInput)
	assert.Equal(t, "5905", result)
}
