package aoc

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func init() {
	registerSolution("2023:5:1", y2023d5part1)
	registerSolution("2023:5:2", y2023d5part2)

}

type Mapping struct {
	ranges   []Range
	destName string
}

type Range struct {
	sourceStart int
	destStart   int
	length      int
}

func (m Mapping) Mutate(in int) int {
	out := in
	for _, r := range m.ranges {
		out = r.Mutate(in)
		if out != in {
			return out
		}
	}
	return out
}

func (r Range) Mutate(in int) int {
	out := in
	start := r.sourceStart
	end := r.sourceStart + r.length - 1
	if start <= in && in <= end {
		out = in - start + r.destStart
	}
	return out
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
				workingMap = Mapping{destName: destination}
			} else {
				// data
				workingMap.ParseMapping(line)
			}
		}
	}

	DestToSource[destination] = workingMap
	fmt.Println(seeds)
	working := seeds[:]
	for _, step := range order {
		fmt.Println(step)
		curr := DestToSource[step]
		fmt.Println(curr)
		for i, spot := range working {
			working[i] = curr.Mutate(spot)
		}
	}
	minVal := slices.Min(working)
	fmt.Println(working)
	return fmt.Sprint(minVal)
}

func (m *Mapping) ParseMapping(line string) {
	ints := stringsToInts(strings.Split(line, " "))
	source := ints[1]
	dest := ints[0]
	length := ints[2]
	m.ranges = append(m.ranges, Range{sourceStart: source, destStart: dest, length: length})
}

func getSourceAndDest(line string) (source, destination string) {
	spaceIdx := strings.Index(line, " ")
	splits := strings.Split(line[:spaceIdx], "-to-")
	source = splits[0]
	destination = splits[1]
	return source, destination
}

type SeedRange struct {
	bounds []bound
}

type bound struct {
	low      int
	high     int
	consumed bool
}

func (b bound) NewBounds(r Range) []bound {
	var working []bound
	if r.sourceStart > b.high || (r.sourceStart+r.length-1) < b.low {
		// bound is not within source
		working = append(working, b)
		return working
	}

	// need to add to bounds
	if b.low-r.sourceStart < 0 {
		// bound lower than source
		idx := r.sourceStart - b.low
		newBound := bound{low: b.low, high: b.low + idx - 1}
		// fmt.Printf("low bound: %#v\n", newBound)
		working = append(working, newBound)

		b.low = b.low + idx
	}

	rTop := r.sourceStart + r.length - 1
	if b.high > rTop {
		newBound := bound{low: rTop, high: b.high}
		// fmt.Printf("high bound: %#v\n", newBound)
		working = append(working, newBound)
		b.high = rTop
	}

	if r.sourceStart >= r.destStart {
		diff := r.sourceStart - r.destStart
		b.low = b.low - diff
		b.high = b.high - diff
		b.consumed = true
	} else if r.sourceStart < r.destStart {
		diff := r.destStart - r.sourceStart
		b.low = b.low + diff
		b.high = b.high + diff
		b.consumed = true
	}
	working = append(working, b)
	return working
}

func (sr *SeedRange) CheckBounds(ranges []Range) []bound {

	var newBounds []bound
	for _, r := range ranges {
		var unconsumed []bound
		for _, b := range sr.bounds {
			workingBounds := b.NewBounds(r)
			for _, wb := range workingBounds {
				if !wb.consumed {
					unconsumed = append(unconsumed, wb)
				} else {
					newBounds = append(newBounds, wb)
				}
			}
		}
		sr.bounds = unconsumed
	}
	for i := range newBounds {
		newBounds[i].consumed = false
	}
	return append(newBounds, sr.bounds...)
}

func (sr *SeedRange) lowest() int {
	candidate := 99999999999999
	for _, b := range sr.bounds {
		if b.low < candidate {
			candidate = b.low
		}
	}
	return candidate
}

func y2023d5part2(input string) string {
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
					sort.Slice(workingMap.ranges, func(i, j int) bool {
						return workingMap.ranges[i].sourceStart < workingMap.ranges[j].sourceStart
					})
					DestToSource[destination] = workingMap
				}
				// new map
				_, destination = getSourceAndDest(line)
				workingMap = Mapping{destName: destination}
			} else {
				// data
				workingMap.ParseMapping(line)
			}
		}
	}

	DestToSource[destination] = workingMap

	var srs []*SeedRange
	for i := 0; i < len(seeds)-1; i = i + 2 {
		newSR := SeedRange{bounds: []bound{{low: seeds[i], high: seeds[i] + seeds[i+1] - 1}}}
		srs = append(srs, &newSR)
	}

	for _, step := range order {
		curr := DestToSource[step]
		fmt.Printf("=================step %s===============\n", step)
		fmt.Printf("Map Ranges: %#v\n", curr.ranges)
		for i, sr := range srs {
			nextBounds := sr.CheckBounds(curr.ranges)
			fmt.Printf("NextBounds for SR: %d\n%#v\n", i, nextBounds)
			sr.bounds = nextBounds
		}
	}

	lowest := 99999999999999
	for _, sr := range srs {
		candidate := sr.lowest()
		if candidate < lowest {
			lowest = candidate
		}
	}
	return fmt.Sprint(lowest)
}
