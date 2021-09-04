package fhc_qualification_2021

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func ProblemB(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfGames := 0
	fmt.Fscanln(in, &numberOfGames)
	for caseNo := 1; caseNo <= numberOfGames; caseNo++ {
		size := 0
		fmt.Fscanln(in, &size)

		board := make([]string, size)
		boardRotated := make([]string, size)
		for i := 0; i < size; i++ {
			fmt.Fscanln(in, &board[i])
		}
		for i := 0; i < size; i++ {
			row := ""
			for j := 0; j < size; j++ {
				row = row + string(board[j][i])
			}
			boardRotated[i] = row
		}
		minStepsToWin := -1
		solutions := make(map[string]bool)
		for i := 0; i < size; i++ {
			for rotated, row := range map[bool]string{false: board[i], true: boardRotated[i]} {
				if strings.Contains(row, "O") {
					continue
				}
				steps := strings.Count(row, ".")
				if steps == 0 {
					continue
				}
				if minStepsToWin == -1 || steps < minStepsToWin {
					minStepsToWin = steps
					solutions = make(map[string]bool)
				}
				if steps == minStepsToWin {
					positions := make([]string, 0)
					for j, r := range row {
						if r != '.' {
							continue
						}
						if !rotated {
							positions = append(positions, fmt.Sprintf("%d%d", j, i))
						} else {
							positions = append(positions, fmt.Sprintf("%d%d", i, j))
						}
					}
					sort.Strings(positions)
					solutions[strings.Join(positions, "-")] = true
				}
			}
		}
		result := "Impossible"
		if len(solutions) > 0 {
			result = fmt.Sprintf("%d %d", minStepsToWin, len(solutions))
		}
		fmt.Fprintf(out, "Case #%d: %s\n", caseNo, result)
	}
}
