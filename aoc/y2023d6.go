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

// func (r race) bisect_up(low, high int, lastWon bool) (newLow, newHigh int, newWon bool) {
// 	middle := (high - low) / 2
// 	// fmt.Printf("high: %d low: %d middle: %d\n", high, low, middle)
// 	currWin := r.wins(middle)
// 	if lastWon == currWin {
// 		return low + middle, high, currWin
// 	} else {
// 		return low, middle, currWin
// 	}
// }

// func (r race) bisect_down(low, high int, lastWon bool) (newLow, newHigh int, newWon bool) {
// 	middle := (high - low) / 2
// 	// fmt.Printf("high: %d low: %d middle: %d\n", high, low, middle)
// 	currWin := r.wins(middle)
// 	if lastWon == currWin {
// 		return low, middle, currWin
// 	} else {
// 		return low + middle, high, currWin
// 	}
// if currWin {
// 	if lastWon == currWin {
// 		return low, middle, currWin
// 	} else {
// 		return middle, high, currWin
// 	}
// } else {
// 	if lastWon != currWin {
// 		return middle, high, currWin
// 	} else {
// 		return low, middle, currWin
// 	}
// }
// }

func (r race) bisect_left(low, high int, seeking bool) (newLow, newHigh int) {
	middle := (high - low) / 2
	currWin := r.wins(middle)
	fmt.Printf("high: %d low: %d middle: %d\n", high, low, middle)
	if currWin == seeking {
		return r.bisect_left(low, middle, seeking)
	}
	fmt.Printf("middle: %d High: %d\n", middle, high)
	return middle, high
}

func (r race) bisect_right(low, high int, seeking bool) (newLow, newHigh int) {
	middle := (high - low) / 2
	// fmt.Printf("high: %d low: %d middle: %d\n", high, low, middle)
	currWin := r.wins(middle)
	if currWin == seeking {
		return r.bisect_right(middle+low, high, seeking)
	}
	fmt.Printf("Low: %d Middle: %d\n", low, middle+low)
	return low, middle + low
}

func (r race) bisect(starting string, seeking bool, low, high int) (newLow, newHigh int) {
	fmt.Printf("starting: %s, seeking: %v, low: %d high: %d\n", starting, seeking, low, high)
	if high-low < 10 {
		return low, high
	}
	if starting == "left" {
		newLow, newHigh = r.bisect_left(low, high, seeking)
		fmt.Printf("back from left Low: %d High: %d\n", newLow, newHigh)
		return r.bisect("right", !seeking, newLow, newHigh)
	} else {
		newLow, newHigh = r.bisect_right(low, high, seeking)
		fmt.Printf("back from right Low: %d High: %d\n", newLow, newHigh)
		return r.bisect("left", !seeking, newLow, newHigh)
	}
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
	low, _ = r.bisect("left", true, 0, r.time)
	lowBound := r.up(low, true)

	var topBound int
	low, _ = r.bisect("right", true, 0, r.time)
	topBound = r.up(low, false)

	fmt.Printf("Low: %d Top: %d\n", lowBound, topBound)
	val := topBound - lowBound
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
