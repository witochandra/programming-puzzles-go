package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2022Day08Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	forest := make([][]byte, 0)
	visibilityByCoordinate := make(map[string]bool)
	for {
		raw, _ := in.ReadBytes('\n')
		if len(raw) == 0 {
			break
		}
		if raw[len(raw)-1] == '\n' {
			raw = raw[:len(raw)-1]
		}
		forest = append(forest, raw)
	}
	// mark all edges as visible
	for i := 0; i < len(forest); i++ {
		visibilityByCoordinate[fmt.Sprintf("%d,%d", 0, i)] = true
		visibilityByCoordinate[fmt.Sprintf("%d,%d", len(forest[0])-1, i)] = true
	}
	for i := 0; i < len(forest[0]); i++ {
		visibilityByCoordinate[fmt.Sprintf("%d,%d", i, 0)] = true
		visibilityByCoordinate[fmt.Sprintf("%d,%d", i, len(forest)-1)] = true
	}

	// from sides
	for y := 0; y < len(forest); y++ {
		var (
			maxHeightFromLeft  byte = '0' - 1
			maxHeightFromRight byte = '0' - 1
		)
		// from left to right
		for x := 0; x < len(forest[y]); x++ {
			if forest[y][x] > maxHeightFromLeft {
				visibilityByCoordinate[fmt.Sprintf("%d,%d", x, y)] = true
				maxHeightFromLeft = forest[y][x]
			}
		}

		// from right to left
		for x := len(forest[y]) - 1; x >= 0; x-- {
			if forest[y][x] > maxHeightFromRight {
				visibilityByCoordinate[fmt.Sprintf("%d,%d", x, y)] = true
				maxHeightFromRight = forest[y][x]
			}
		}
	}

	// from top and bottom
	for x := 0; x < len(forest[0]); x++ {
		var (
			maxHeightFromTop    byte = '0' - 1
			maxHeightFromBottom byte = '0' - 1
		)
		// from top to bottom
		for y := 0; y < len(forest); y++ {
			if forest[y][x] > maxHeightFromTop {
				visibilityByCoordinate[fmt.Sprintf("%d,%d", x, y)] = true
				maxHeightFromTop = forest[y][x]
			}
		}
		// from bottom to top
		for y := len(forest) - 1; y >= 0; y-- {
			if forest[y][x] > maxHeightFromBottom {
				visibilityByCoordinate[fmt.Sprintf("%d,%d", x, y)] = true
				maxHeightFromBottom = forest[y][x]
			}
		}
	}
	fmt.Fprint(out, len(visibilityByCoordinate))
}

func Problem2022Day08Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	forest := make([][]int, 0)
	for {
		raw, _ := in.ReadBytes('\n')
		if len(raw) == 0 {
			break
		}
		row := make([]int, 0)
		for _, r := range raw {
			if r == '\n' {
				continue
			}
			row = append(row, int(r-'0'))
		}
		forest = append(forest, row)
	}

	maxScore := 0
	for y := 1; y < len(forest)-1; y++ {
		for x := 1; x < len(forest[y])-1; x++ {
			left, right, top, bottom := 0, 0, 0, 0

			// visibilities to left
			for i := x - 1; i >= 0; i-- {
				left++
				if forest[y][x] <= forest[y][i] {
					break
				}
			}

			// visibilities to right
			for i := x + 1; i < len(forest[y]); i++ {
				right++
				if forest[y][x] <= forest[y][i] {
					break
				}
			}

			// visibilities to top
			for i := y - 1; i >= 0; i-- {
				top++
				if forest[y][x] <= forest[i][x] {
					break
				}
			}

			// visibilities to bottom
			for i := y + 1; i < len(forest); i++ {
				bottom++
				if forest[y][x] <= forest[i][x] {
					break
				}
			}

			score := left * right * top * bottom
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Fprint(out, maxScore)
}
