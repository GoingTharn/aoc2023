package aoc

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:4:1", y2023d4part1)
	registerSolution("2023:4:2", y2023d4part2)

}

type ScratchCard struct {
	cardNumber int
	winner     []int
	candidate  []int
}

func (sc ScratchCard) MatchCount() int {
	return len(lo.Intersect[int](sc.winner, sc.candidate))
}

func NewScratchCard(cardNum int, line string) (sc ScratchCard) {
	startIdx := strings.Index(line, ":")
	pipeIdx := strings.Index(line, "|")
	winner := strings.Split(strings.TrimSpace(line[startIdx:pipeIdx]), " ")
	sc.winner = stringsToInts(winner)
	candidate := strings.Split(strings.TrimSpace(line[pipeIdx+1:]), " ")
	sc.candidate = stringsToInts(candidate)
	sc.cardNumber = cardNum + 1
	return sc
}

func Score(matchNum int) int {
	if matchNum == 0 {
		return 0
	}
	working := 1
	for i := 1; i < matchNum; i++ {
		working = working * 2
	}
	fmt.Println(working)
	return working
}

func y2023d4part2(input string) string {
	var scratchCards []ScratchCard
	running := make(map[int]int)

	for i, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		sc := NewScratchCard(i, line)
		scratchCards = append(scratchCards, sc)
	}

	for j, card := range scratchCards {
		key := j + 1
		starting, found := running[key]
		if !found {
			starting = 0
		}
		running[key] = starting + 1
		matchCount := card.MatchCount()
		fmt.Printf("%#v\n", card)
		fmt.Println(matchCount)
		for i := 1; i <= matchCount; i++ {
			fmt.Printf("TgtNum: %d, Adding: %d\n", running[key+i], running[key])
			running[key+i] += running[key]
		}
	}
	fmt.Println(running)
	count := 0
	for _, cards := range running {
		count += cards
	}
	return fmt.Sprint(count)
}

func y2023d4part1(input string) string {
	var working int = 0
	for i, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		sc := NewScratchCard(i, line)
		fmt.Printf("%#v\n", sc)
		working += Score(sc.MatchCount())
	}
	return fmt.Sprint(working)
}
