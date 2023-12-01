package aoc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d1partInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var y2023d1part2Input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
one
541
plqxtk1oneone3rkfive
threetcstlhfcqdfhjgccckcvk1hjlpmmkjmr8six
vdttwoeight7eightoneoneightvq
1oneseven
twotwo
1vttzfour9
`

func Test_2023_Day_1_Part_1_Example(t *testing.T) {
	result := y2023d1part1(y2023d1partInput)
	assert.Equal(t, "142", result)
}

func Test_2023_Day_1_Part_2_Example(t *testing.T) {
	result := y2023d1part2(y2023d1part2Input)
	assert.Equal(t, "470", result)
}

func TestGetDigits(t *testing.T) {

	tests := []struct {
		input            string
		expectedForward  string
		expectedReversed string
		same             bool
	}{
		{input: "ldrqsbjrtnlj6lqgptgplbsnhrhpsixzqgpzkxone", expectedForward: "6", expectedReversed: "1", same: false},
		{input: "8nine7", expectedForward: "8", expectedReversed: "7", same: false},
		{input: "brb9", expectedForward: "9", expectedReversed: "9", same: true},
		{input: "225eightworbm", expectedForward: "2", expectedReversed: "2", same: false},
		{input: "vfvrcknfls5q", expectedForward: "5", expectedReversed: "5", same: true},
		{input: "sevensixmvvrzhsixsixsix9", expectedForward: "7", expectedReversed: "9", same: false},
		{input: "oneight", expectedForward: "1", expectedReversed: "8", same: false},
		{input: "3oneightwo", expectedForward: "3", expectedReversed: "2", same: false},
	}

	for _, test := range tests {
		actualF, idx := getTextDigit(test.input, false)
		actualR, lidx := getTextDigit(test.input, true)
		fmt.Printf("Line: %s FirstIdx: %d, LastIdx: %d\n", test.input, idx, lidx)

		assert.Equal(t, test.expectedForward, actualF)
		assert.Equal(t, test.expectedReversed, actualR)
		assert.Equal(t, test.same, idx == lidx)
	}

}
