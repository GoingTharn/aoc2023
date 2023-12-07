package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:6:1", y2023d6part1)
	registerSolution("2023:6:2", y2023d6part2)

}

type race struct {
	time     int
	distance int
}

func y2023d6part1(input string) string {
	splits := strings.Split(input, "\n")
	times := splitThenConvert(splits[0], ":")
	distances := splitThenConvert(splits[1], ":")

	var races []race
	for i := 0; i < len(times); i++ {
		newRace := race{time: times[i], distance: distances[i]}
		races = append(races, newRace)
	}

	var counts []int
	for _, r := range races {
		counts = append(counts, r.getRaceCount())
	}
	fmt.Println(counts)

	product := lo.Reduce(counts, func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
	return fmt.Sprint(product)
}

func (r race) wins(s int) bool {
	timeToRace := r.time - s
	velocity := s
	distanceTravelled := timeToRace * velocity
	return distanceTravelled > r.distance
}

func (r race) getRaceCount() (count int) {
	count = 0
	for i := 0; i < r.time; i++ {
		if r.wins(i) {
			count++
		}
	}
	return count
}

func y2023d6part2(input string) string {
	splits := strings.Split(input, "\n")
	var this []rune
	for _, r := range splits[0] {
		if unicode.IsNumber(r) {
			this = append(this, r)
		}
	}
	time, _ := strconv.Atoi(string(this))
	this = []rune{}
	for _, r := range splits[1] {
		if unicode.IsNumber(r) {
			this = append(this, r)
		}
	}
	dist, _ := strconv.Atoi(string(this))
	fmt.Println(time, dist)

	r := race{time: time, distance: dist}

	var low int
	var lowBound int

	var high int
	var highBound int

	var notWinning bool = true
	for i := 1; notWinning; i = i + 10000 {
		if r.wins(i) {
			low = i
			notWinning = false
		}
	}
	lowBound = r.down(low, false)

	notWinning = true
	for i := r.time; notWinning; i = i - 10000 {
		if r.wins(i) {
			high = i
			notWinning = false
		}
	}
	highBound = r.up(high, false)

	val := highBound - lowBound - 1
	return fmt.Sprint(val)
}

func (r race) down(time int, breakOn bool) int {
	for i := time; ; i-- {
		if breakOn == r.wins(i) {
			return i
		}
	}
}

func (r race) up(time int, breakOn bool) int {
	for i := time; ; i++ {
		if breakOn == r.wins(i) {
			return i
		}
	}
}
