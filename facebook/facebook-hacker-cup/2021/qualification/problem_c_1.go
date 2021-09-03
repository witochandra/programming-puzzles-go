package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Cave struct {
	Index int
	Gold  int
	Links []*Cave
}

func (c *Cave) Link(cave *Cave) {
	c.Links = append(c.Links, cave)
	cave.Links = append(cave.Links, c)
}

func (c *Cave) Unlink(cave *Cave) {
	for i, l := range c.Links {
		if l.Index == cave.Index {
			c.Links[i] = c.Links[len(c.Links)-1]
			c.Links = c.Links[:len(c.Links)-1]
			break
		}
	}
	for i, l := range cave.Links {
		if l.Index == c.Index {
			cave.Links[i] = cave.Links[len(cave.Links)-1]
			cave.Links = cave.Links[:len(cave.Links)-1]
			break
		}
	}
}

func (c *Cave) HasLink(cave *Cave) bool {
	for _, l := range c.Links {
		if l.Index == cave.Index {
			return true
		}
	}
	return false
}

func NewCave(index int, gold int) *Cave {
	return &Cave{
		Index: index,
		Gold:  gold,
		Links: make([]*Cave, 0),
	}
}

func CopyCaves(caves []*Cave) []*Cave {
	newCaves := make([]*Cave, 0)
	for i, c := range caves {
		newCaves = append(newCaves, NewCave(i, c.Gold))
	}
	for i, c := range caves {
		for _, link := range c.Links {
			newCaves[i].Links = append(newCaves[i].Links, newCaves[link.Index])
		}
	}
	return newCaves
}

func GatherGolds(caves []*Cave, pos int, gold int, maxGold int, currentMaxGold *int) {
	currentGold := gold + caves[pos].Gold
	caves[pos].Gold = 0
	if pos == 0 && currentGold > *currentMaxGold {
		*currentMaxGold = currentGold
	}
	if len(caves[0].Links) == 0 {
		return
	}
	if currentMaxGold == &maxGold {
		return
	}
	for _, nextCave := range caves[pos].Links {
		newCaves := CopyCaves(caves)
		newCaves[pos].Unlink(newCaves[nextCave.Index])
		GatherGolds(newCaves, nextCave.Index, currentGold, maxGold, currentMaxGold)
	}
}

func ProblemC1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfMines := 0
	fmt.Fscanln(in, &numberOfMines)
	for caseNo := 1; caseNo <= numberOfMines; caseNo++ {
		numberOfCaves := 0
		fmt.Fscanln(in, &numberOfCaves)

		caves := make([]*Cave, 0)
		temp, _ := in.ReadString('\n')
		totalGold := 0
		for i, goldStr := range strings.Split(strings.TrimSpace(temp), " ") {
			gold, _ := strconv.Atoi(goldStr)
			caves = append(caves, NewCave(i, gold))
			totalGold += gold
		}
		for i := 0; i < numberOfCaves-1; i++ {
			temp, _ = in.ReadString('\n')
			strs := strings.Split(strings.TrimSpace(temp), " ")
			left, _ := strconv.Atoi(strs[0])
			right, _ := strconv.Atoi(strs[1])

			caves[left-1].Link(caves[right-1])
		}
		maxGold := 0
		goldMines := make([][]*Cave, 0)
		goldMines = append(goldMines, caves)
		for tunnelOrigin := 0; tunnelOrigin < numberOfCaves-1; tunnelOrigin++ {
			for tunnelDestination := tunnelOrigin + 1; tunnelDestination < numberOfCaves; tunnelDestination++ {
				tempCaves := CopyCaves(caves)
				tempCaves[tunnelOrigin].Link(tempCaves[tunnelDestination])
				goldMines = append(goldMines, tempCaves)
			}
		}
		for _, goldMine := range goldMines {
			GatherGolds(goldMine, 0, 0, totalGold, &maxGold)
			if maxGold == totalGold {
				break
			}
		}
		fmt.Fprintf(out, "Case #%d: %d\n", caseNo, maxGold)
	}
}

// func main() {
// 	ProblemC1(os.Stdin, os.Stdout)
// }
