package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2021Day05(r io.Reader, w io.Writer, checkDiagonal bool) {
	type Coordinate struct {
		X int
		Y int
	}

	type Line struct {
		A Coordinate
		B Coordinate
	}

	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	maxX, maxY := 0, 0
	lines := make([]Line, 0)
	for {
		_, err := in.Peek(1)
		if err != nil {
			break
		}
		line := Line{}
		fmt.Fscanf(in, "%d,%d -> %d,%d\n", &line.A.X, &line.A.Y, &line.B.X, &line.B.Y)
		if line.A.X > maxX {
			maxX = line.A.X
		} else if line.B.X > maxX {
			maxX = line.B.X
		}
		if line.A.Y > maxY {
			maxY = line.A.Y
		} else if line.B.Y > maxY {
			maxY = line.B.Y
		}
		lines = append(lines, line)
	}
	diagram := make([][]int, maxY+1)
	for i := range diagram {
		diagram[i] = make([]int, maxX+1)
	}
	for _, line := range lines {
		if !checkDiagonal && line.A.X != line.B.X && line.A.Y != line.B.Y {
			continue
		}
		currX, currY := line.A.X, line.A.Y
		diagram[currY][currX]++
		for currX != line.B.X || currY != line.B.Y {
			if currX < line.B.X {
				currX++
			} else if currX > line.B.X {
				currX--
			}
			if currY < line.B.Y {
				currY++
			} else if currY > line.B.Y {
				currY--
			}
			diagram[currY][currX]++
		}
	}
	overlaps := 0
	for _, row := range diagram {
		for _, v := range row {
			if v > 1 {
				overlaps++
			}
		}
	}
	fmt.Fprintln(out, overlaps)
}

func Problem2021Day05Part1(r io.Reader, w io.Writer) {
	Problem2021Day05(r, w, false)
}

func Problem2021Day05Part2(r io.Reader, w io.Writer) {
	Problem2021Day05(r, w, true)
}
