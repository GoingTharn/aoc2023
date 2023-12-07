package aoc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d5partInput = `seeds: 79 14 55 13
seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestRangeMutate(t *testing.T) {
	testRange := Range{destStart: 2, sourceStart: 98, length: 2}
	tests := []struct {
		input    int
		expected int
	}{
		{98, 2},
		{99, 3},
		{100, 100},
	}

	for _, test := range tests {
		num := testRange.Mutate(test.input)
		assert.Equal(t, test.expected, num)

	}

}

func TestBounds(t *testing.T) {
	testRange := Range{destStart: 2, sourceStart: 98, length: 2}
	tests := []struct {
		input    bound
		expected []bound
	}{
		{input: bound{low: 2, high: 33}, expected: []bound{{low: 2, high: 33}}},
		{input: bound{low: 97, high: 99}, expected: []bound{{low: 97, high: 97, consumed: false}, {low: 2, high: 3, consumed: true}}},
		{input: bound{low: 96, high: 106}, expected: []bound{{low: 96, high: 97}, {low: 2, high: 3, consumed: true}, {low: 99, high: 106}}},
	}

	for i, test := range tests {
		fmt.Printf("Test %d\n", i)
		bounds := test.input.NewBounds(testRange)
		assert.ElementsMatch(t, test.expected, bounds)
	}

	testRange = Range{destStart: 52, sourceStart: 50, length: 48}
	tests = []struct {
		input    bound
		expected []bound
	}{
		{input: bound{low: 79, high: 92}, expected: []bound{{low: 81, high: 94, consumed: true}}},
	}

	for i, test := range tests {
		fmt.Printf("Test %d\n", i)
		bounds := test.input.NewBounds(testRange)
		assert.ElementsMatch(t, test.expected, bounds)
	}
}

func Test_2023_Day_5_Part_1_Example(t *testing.T) {
	result := y2023d5part1(y2023d5partInput)
	assert.Equal(t, "35", result)
}

func Test_2023_Day_5_Part_2_Example(t *testing.T) {
	result := y2023d5part2(y2023d5partInput)
	assert.Equal(t, "46", result)
}
