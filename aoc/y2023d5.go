package aoc

import (
	"fmt"
	"slices"
	"strings"
)

func init() {
	registerSolution("2023:5:1", y2023d5part1)
	registerSolution("2023:5:2", y2023d5part2)

}

type Mapping struct {
	sourceToDest map[int]int
	destName     string
}

func y2023d5part1(input string) string {
	DestToSource := make(map[string]Mapping)

	order := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	var seeds []int
	var workingMap Mapping
	var destination string
	for i, line := range strings.Split(input, "\n") {
		// seeds
		// var source string
		if i == 0 {
			justNums := strings.Split(line, ":")[1]
			seeds = stringsToInts(strings.Split(justNums, " "))
		} else {
			if len(line) == 0 {
				continue
			}
			if strings.Contains(line, "to") {
				if destination != "" {
					DestToSource[destination] = workingMap
				}
				// new map
				_, destination = getSourceAndDest(line)
				workingMap = Mapping{destName: destination, sourceToDest: make(map[int]int)}
			} else {
				// data
				workingMap.ParseMapping(line)
			}
		}

	}
	fmt.Println(seeds)
	working := seeds[:]
	for _, step := range order {
		curr := DestToSource[step]
		fmt.Println(curr)
		for i, spot := range working {
			var candidate int
			candidate, found := curr.sourceToDest[spot]
			if !found {
				candidate = spot
			}
			working[i] = candidate
		}
	}
	minVal := slices.Min(working)
	return fmt.Sprint(minVal)
}

func (m Mapping) ParseMapping(line string) {
	ints := stringsToInts(strings.Split(line, " "))
	source := ints[1]
	dest := ints[0]
	length := ints[2]
	for i := 0; i < length; i++ {
		m.sourceToDest[source+i] = dest + i
	}
}

func getSourceAndDest(line string) (source, destination string) {
	spaceIdx := strings.Index(line, " ")
	splits := strings.Split(line[:spaceIdx], "-to-")
	source = splits[0]
	destination = splits[1]
	return source, destination
}

func y2023d5part2(input string) string {
	return "wrong again"
}
