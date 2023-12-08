package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

func init() {
	registerSolution("2023:7:1", y2023d7part1)
	registerSolution("2023:7:2", y2023d7part2)

}

type hand struct {
	cards []card
	bet   int
}

type card rune

func (h hand) getRank2() int {
	var working []card
	jCount := 0
	for _, r := range h.cards {
		if r == 'J' {
			jCount++
		} else {
			working = append(working, r)
		}

	}
	dupes := lo.FindDuplicates[card](working)
	dupeCounts := make([]int, 0)
	handAsString := string(working)
	for _, dupe := range dupes {
		dupeCounts = append(dupeCounts, strings.Count(handAsString, string(dupe)))
	}
	slices.Sort(dupeCounts)
	retVal := 0
	// split up the values to account for J
	if slices.Equal[int](dupeCounts, []int{}) {
		retVal = 1
	} else if slices.Equal(dupeCounts, []int{2}) {
		retVal = 3
	} else if slices.Equal(dupeCounts, []int{2, 2}) {
		retVal = 4
	} else if slices.Equal(dupeCounts, []int{3}) {
		retVal = 5
	} else if slices.Equal(dupeCounts, []int{2, 3}) {
		retVal = 6
	} else if slices.Equal(dupeCounts, []int{4}) {
		retVal = 7
	} else if slices.Equal(dupeCounts, []int{5}) {
		retVal = 9
	}
	retVal = retVal + jCount*2
	if retVal > 9 {
		retVal = 9
	}
	return retVal
}

func (h hand) getRank() int {
	dupes := lo.FindDuplicates[card](h.cards)
	dupeCounts := make([]int, 0)
	handAsString := string(h.cards)
	for _, dupe := range dupes {
		dupeCounts = append(dupeCounts, strings.Count(handAsString, string(dupe)))
	}
	slices.Sort(dupeCounts)
	if slices.Equal[int](dupeCounts, []int{}) {
		return 1
	} else if slices.Equal(dupeCounts, []int{2}) {
		return 2
	} else if slices.Equal(dupeCounts, []int{2, 2}) {
		return 3
	} else if slices.Equal(dupeCounts, []int{3}) {
		return 4
	} else if slices.Equal(dupeCounts, []int{2, 3}) {
		return 5
	} else if slices.Equal(dupeCounts, []int{4}) {
		return 6
	} else if slices.Equal(dupeCounts, []int{5}) {
		return 7
	}
	panic("not found")
}

func y2023d7part1(input string) string {
	const order = "23456789TJQKA"
	var hands []hand
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		splut := strings.Split(line, " ")
		bet, err := strconv.Atoi(splut[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand{cards: []card(splut[0]), bet: bet})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		aRank := hands[i].getRank()
		bRank := hands[j].getRank()
		if aRank < bRank {
			return true
		} else if bRank < aRank {
			return false
		} else {
			for k := 0; k < len(hands[i].cards); k++ {
				idxA := strings.IndexRune(order, rune(hands[i].cards[k]))
				idxB := strings.IndexRune(order, rune(hands[j].cards[k]))
				if idxA < idxB {
					return true
				} else if idxB < idxA {
					return false
				}
			}
			return false
		}
	})

	betSum := 0
	for i := 0; i < len(hands); i++ {
		fmt.Println(hands[i])
		betSum += (i + 1) * hands[i].bet
	}

	return fmt.Sprint(betSum)
}

func (h hand) String() string {
	return fmt.Sprintf("{cards: %s, bet: %d}", string(h.cards), h.bet)
}
func y2023d7part2(input string) string {
	const order = "J23456789TQKA"
	var hands []hand
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		splut := strings.Split(line, " ")
		bet, err := strconv.Atoi(splut[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand{cards: []card(splut[0]), bet: bet})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		aRank := hands[i].getRank2()
		bRank := hands[j].getRank2()
		if aRank < bRank {
			return true
		} else if bRank < aRank {
			return false
		} else {
			for k := 0; k < len(hands[i].cards); k++ {
				idxA := strings.IndexRune(order, rune(hands[i].cards[k]))
				idxB := strings.IndexRune(order, rune(hands[j].cards[k]))
				if idxA < idxB {
					return true
				} else if idxB < idxA {
					return false
				}
			}
			return false
		}
	})

	betSum := 0
	for i := 0; i < len(hands); i++ {
		fmt.Println(hands[i])
		betSum += (i + 1) * hands[i].bet
	}

	return fmt.Sprint(betSum)
}
