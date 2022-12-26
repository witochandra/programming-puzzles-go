package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2022Day12Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	terrain := make([][]byte, 0)
	startPosition := []int{0, 0}
	endPosition := []int{0, 0}
	for {
		raw, _ := in.ReadBytes('\n')
		if len(raw) > 0 && raw[len(raw)-1] == '\n' {
			raw = raw[:len(raw)-1]
		}
		if len(raw) == 0 {
			break
		}
		for i, c := range raw {
			if c == 'S' {
				startPosition = []int{i, len(terrain)}
				raw[i] = 'a'
			} else if c == 'E' {
				endPosition = []int{i, len(terrain)}
				raw[i] = 'z'
			}
		}
		terrain = append(terrain, raw)
	}
	passed := make(map[string]bool)
	queue := make([][]int, 0)
	queue = append(queue, startPosition)
	steps := 0
	for {
		containsEndPosition := false
		newQueue := make([][]int, 0)
		for _, position := range queue {
			if position[0] == endPosition[0] && position[1] == endPosition[1] {
				containsEndPosition = true
				break
			}
			passed[fmt.Sprintf("%d,%d", position[0], position[1])] = true
			nextPositions := [][]int{
				{position[0] - 1, position[1]},
				{position[0] + 1, position[1]},
				{position[0], position[1] - 1},
				{position[0], position[1] + 1},
			}
			for _, nextPosition := range nextPositions {
				if nextPosition[0] < 0 || nextPosition[0] >= len(terrain[0]) || nextPosition[1] < 0 || nextPosition[1] >= len(terrain) {
					continue
				}
				if passed[fmt.Sprintf("%d,%d", nextPosition[0], nextPosition[1])] {
					continue
				}
				diff := int(terrain[nextPosition[1]][nextPosition[0]]) - int(terrain[position[1]][position[0]])
				if diff <= 1 {
					passed[fmt.Sprintf("%d,%d", nextPosition[0], nextPosition[1])] = true
					newQueue = append(newQueue, nextPosition)
				}
			}
		}
		if containsEndPosition {
			break
		}
		queue = newQueue
		steps++
		if len(queue) == 0 {
			panic("no path found")
		}
	}
	fmt.Fprint(out, steps)
}

func Problem2022Day12Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	terrain := make([][]byte, 0)
	startPositions := make([][]int, 0)
	endPosition := []int{0, 0}
	for {
		raw, _ := in.ReadBytes('\n')
		if len(raw) > 0 && raw[len(raw)-1] == '\n' {
			raw = raw[:len(raw)-1]
		}
		if len(raw) == 0 {
			break
		}
		for i, c := range raw {
			if c == 'S' || c == 'a' {
				startPositions = append(startPositions, []int{i, len(terrain)})
				raw[i] = 'a'
			} else if c == 'E' {
				endPosition = []int{i, len(terrain)}
				raw[i] = 'z'
			}
		}
		terrain = append(terrain, raw)
	}
	passed := make(map[string]bool)
	queue := make([][]int, 0)
	queue = append(queue, startPositions...)
	steps := 0
	for {
		containsEndPosition := false
		newQueue := make([][]int, 0)
		for _, position := range queue {
			if position[0] == endPosition[0] && position[1] == endPosition[1] {
				containsEndPosition = true
				break
			}
			passed[fmt.Sprintf("%d,%d", position[0], position[1])] = true
			nextPositions := [][]int{
				{position[0] - 1, position[1]},
				{position[0] + 1, position[1]},
				{position[0], position[1] - 1},
				{position[0], position[1] + 1},
			}
			for _, nextPosition := range nextPositions {
				if nextPosition[0] < 0 || nextPosition[0] >= len(terrain[0]) || nextPosition[1] < 0 || nextPosition[1] >= len(terrain) {
					continue
				}
				if passed[fmt.Sprintf("%d,%d", nextPosition[0], nextPosition[1])] {
					continue
				}
				diff := int(terrain[nextPosition[1]][nextPosition[0]]) - int(terrain[position[1]][position[0]])
				if diff <= 1 {
					passed[fmt.Sprintf("%d,%d", nextPosition[0], nextPosition[1])] = true
					newQueue = append(newQueue, nextPosition)
				}
			}
		}
		if containsEndPosition {
			break
		}
		queue = newQueue
		steps++
		if len(queue) == 0 {
			panic("no path found")
		}
	}
	fmt.Fprint(out, steps)
}
