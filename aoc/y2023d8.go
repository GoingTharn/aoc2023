package aoc

import (
	"fmt"
	"strings"
)

func init() {
	registerSolution("2023:8:1", y2023d8part1)
	registerSolution("2023:8:2", y2023d8part2)

}

type mapNode struct {
	name  string
	left  *mapNode
	right *mapNode
	tail  rune
}

func (mn mapNode) String() string {
	return fmt.Sprintf("{name: %s left:%v right: %v}", mn.name, mn.left, mn.right)
}

func (mn *mapNode) move(dir rune) *mapNode {
	if dir == 'L' {
		return mn.left
	} else {
		return mn.right
	}
}
func ghostWalk(mns []*mapNode, dirs string) int {
	origMapNodes := mns[:]
	fmt.Println(dirs)
	count := 0
	directionsComplete := 0
	dirRunes := []rune(dirs)
	for i := 0; ; i++ {
		if i >= len(dirRunes) {
			directionsComplete++
			fmt.Printf("Dirs complete count: %d\n", directionsComplete)
			i = 0
		}
		count++
		for j := range mns {
			newMns := mns[j].move(dirRunes[i])
			if newMns == origMapNodes[j] {
				fmt.Printf("Loop detected for node %v at index %d", newMns, i)
			}
			mns[j] = newMns
		}
		if allTerminus(mns) {
			return count
		}
	}
}

func allTerminus(mns []*mapNode) bool {
	for _, mn := range mns {
		if mn.tail != 'Z' {
			return false
		}
	}
	return true
}

func walk(mn *mapNode, dirs string, terminus string) int {
	count := 0
	dirRunes := []rune(dirs)
	for i := 0; mn.name != terminus; i++ {
		if i >= len(dirRunes) {
			i = 0
		}
		mn = mn.move(dirRunes[i])
		fmt.Printf("Node: %s\n", mn.name)
		count++
	}
	return count
}

func y2023d8part1(input string) string {
	lines := lines(input)

	nodes := make(map[string]*mapNode)
	dirs := lines[0]

	// add nodemap first pass
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		splits := strings.Split(line, " = ")
		nodeName := splits[0]
		nodes[nodeName] = &mapNode{name: nodeName}
	}

	fmt.Println(nodes)
	// enrich nodepath
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		enrichNode(line, nodes)
	}

	steps := walk(nodes["AAA"], dirs, "ZZZ")

	return fmt.Sprint(steps)
}

func y2023d8part2(input string) string {
	lines := lines(input)

	nodes := make(map[string]*mapNode)
	dirs := lines[0]

	// add nodemap first pass
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		splits := strings.Split(line, " = ")
		nodeName := splits[0]
		nodes[nodeName] = &mapNode{name: nodeName, tail: []rune(nodeName)[2]}
		if nodes[nodeName].tail == 'Z' {
			fmt.Println(nodes[nodeName])
		}
	}

	// enrich nodepath
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		enrichNode(line, nodes)
	}

	var startNodes []*mapNode
	for _, mn := range nodes {
		if mn.tail == 'A' {
			startNodes = append(startNodes, mn)
		}
	}
	fmt.Println(len(startNodes))
	// steps := ghostWalk(startNodes, dirs)
	return dirs //fmt.Sprint(steps)
}

func enrichNode(line string, nodeMap map[string]*mapNode) {
	splits := strings.Split(line, " = ")
	nodeName := splits[0]
	dirs := splits[1]
	working := strings.Split(dirs[1:len(dirs)-1], ", ")
	left, right := working[0], working[1]

	workingNode := nodeMap[nodeName]
	leftNode := nodeMap[left]
	workingNode.left = leftNode
	rightNode := nodeMap[right]
	workingNode.right = rightNode
}
