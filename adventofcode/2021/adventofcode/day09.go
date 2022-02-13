package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func Problem2021Day09Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	heightMap := make([][]int, 0)
	for {
		_, err := in.Peek(1)
		if err == io.EOF {
			break
		}
		input, _, _ := in.ReadLine()
		row := make([]int, 0)
		for _, c := range input {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		heightMap = append(heightMap, row)
	}
	sumOfRiskLevels := 0
	for y, row := range heightMap {
		for x, v := range row {
			if x > 0 && v >= heightMap[y][x-1] {
				continue
			}
			if y > 0 && v >= heightMap[y-1][x] {
				continue
			}
			if x < len(row)-1 && v >= heightMap[y][x+1] {
				continue
			}
			if y < len(heightMap)-1 && v >= heightMap[y+1][x] {
				continue
			}
			sumOfRiskLevels += v + 1
		}
	}
	fmt.Fprint(out, sumOfRiskLevels)
}

func Problem2021Day09Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	heightMap := make([][]int, 0)
	for {
		_, err := in.Peek(1)
		if err == io.EOF {
			break
		}
		input, _, _ := in.ReadLine()
		row := make([]int, 0)
		for _, c := range input {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		heightMap = append(heightMap, row)
	}

	findNeighbours := func(x, y int) [][]int {
		neighbours := make([][]int, 0)
		if x > 0 {
			neighbours = append(neighbours, []int{x - 1, y})
		}
		if x < len(heightMap[0])-1 {
			neighbours = append(neighbours, []int{x + 1, y})
		}
		if y > 0 {
			neighbours = append(neighbours, []int{x, y - 1})
		}
		if y < len(heightMap)-1 {
			neighbours = append(neighbours, []int{x, y + 1})
		}
		return neighbours
	}

	findBasinArea := func(x, y int) int {
		areas := make(map[string]bool)
		areas[fmt.Sprintf("%d,%d", x, y)] = true

		pointsToCheck := findNeighbours(x, y)
		for {
			if len(pointsToCheck) == 0 {
				break
			}
			ptcX, ptcY := pointsToCheck[0][0], pointsToCheck[0][1]
			if areas[fmt.Sprintf("%d,%d", ptcX, ptcY)] {
				pointsToCheck = pointsToCheck[1:]
				continue
			}
			if heightMap[ptcY][ptcX] == 9 {
				pointsToCheck = pointsToCheck[1:]
				continue
			}

			areas[fmt.Sprintf("%d,%d", ptcX, ptcY)] = true
			pointsToCheck = append(pointsToCheck[1:], findNeighbours(ptcX, ptcY)...)
		}
		return len(areas)
	}
	currentMin := 0
	basinAreas := make([]int, 0)
	for y, row := range heightMap {
		for x, v := range row {
			if x > 0 && v >= heightMap[y][x-1] {
				continue
			}
			if y > 0 && v >= heightMap[y-1][x] {
				continue
			}
			if x < len(row)-1 && v >= heightMap[y][x+1] {
				continue
			}
			if y < len(heightMap)-1 && v >= heightMap[y+1][x] {
				continue
			}
			area := findBasinArea(x, y)
			if len(basinAreas) < 3 || area > currentMin {
				basinAreas = append(basinAreas, area)
				if len(basinAreas) > 3 {
					sort.Ints(basinAreas)
					basinAreas = basinAreas[1:]
					currentMin = basinAreas[0]
				}
			}
		}
	}
	result := 1
	for _, v := range basinAreas {
		result *= v
	}
	fmt.Fprint(out, result)
}
