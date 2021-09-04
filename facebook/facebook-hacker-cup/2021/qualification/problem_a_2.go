package fhc_qualification_2021

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Node struct {
	Name     rune
	Children []*Node
}

func (n *Node) AddChild(child *Node) {
	for _, c := range n.Children {
		if c.Name == child.Name {
			return
		}
	}
	n.Children = append(n.Children, child)
}

func NewNode(name rune) *Node {
	return &Node{
		Name:     name,
		Children: make([]*Node, 0),
	}
}

func DistanceTo(toName rune, skips string, nodesToCheck []*Node) int {
	nextNodeToCheckSet := make(map[rune]bool)
	nextNodesToCheck := make([]*Node, 0)
	for _, node := range nodesToCheck {
		if node.Name == toName {
			return 1
		}
		for _, gc := range node.Children {
			if strings.ContainsRune(skips, gc.Name) {
				continue
			}
			skips = skips + string(gc.Name)
			if !nextNodeToCheckSet[gc.Name] {
				nextNodeToCheckSet[gc.Name] = true
				nextNodesToCheck = append(nextNodesToCheck, gc)
			}
		}
	}
	if len(nextNodeToCheckSet) > 0 {
		distance := DistanceTo(toName, skips, nextNodesToCheck)
		if distance >= 0 {
			return 1 + distance
		}
	}

	return -1
}

func ProblemA2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfBirthdays := 0
	fmt.Fscanln(in, &numberOfBirthdays)

	for caseNo := 1; caseNo <= numberOfBirthdays; caseNo++ {
		input := ""
		replacementCount := 0
		fmt.Fscanln(in, &input)
		fmt.Fscanln(in, &replacementCount)

		nodes := make(map[rune]*Node)
		runeCount := make(map[rune]int)
		for _, v := range input {
			nodes[v] = NewNode(v)
			runeCount[v] = runeCount[v] + 1
		}
		sameRuneCount := 1
		for _, count := range runeCount {
			if sameRuneCount < count {
				sameRuneCount = count
			}
		}
		for i := 0; i < replacementCount; i++ {
			tmp := ""
			fmt.Fscanln(in, &tmp)
			runes := []rune(tmp)
			destination, found := nodes[runes[1]]
			if !found {
				destination = NewNode(runes[1])
				nodes[runes[1]] = destination
			}
			source, found := nodes[runes[0]]
			if !found {
				source = NewNode(runes[0])
				nodes[runes[0]] = source
			}
			source.AddChild(destination)
		}
		currentSteps := -1
		minSteps := len(input) - sameRuneCount
		for destination := range nodes {
			steps := 0
			finish := true
			for _, source := range input {
				if source == destination {
					continue
				}
				distance := DistanceTo(destination, string(source), nodes[source].Children)
				steps = steps + distance
				if distance < 0 {
					finish = false
					break
				}
			}
			if finish && (currentSteps == -1 || steps < currentSteps) {
				currentSteps = steps
				if currentSteps == minSteps {
					break
				}
			}
		}
		fmt.Fprintf(out, "Case #%d: %d\n", caseNo, currentSteps)
	}
}
