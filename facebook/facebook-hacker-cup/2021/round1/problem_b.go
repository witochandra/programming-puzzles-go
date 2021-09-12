package fhc_2021_round1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ProblemB(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfTowns := 0
	fmt.Fscanln(in, &numberOfTowns)

	for caseNo := 1; caseNo <= numberOfTowns; caseNo++ {
		temp, _ := in.ReadString('\n')
		strs := strings.Split(strings.TrimSpace(temp), " ")

		n, _ := strconv.Atoi(strs[0])
		m, _ := strconv.Atoi(strs[1])
		a, _ := strconv.Atoi(strs[2])
		b, _ := strconv.Atoi(strs[3])

		minIntersections := n + m - 1
		if a < minIntersections || b < minIntersections {
			fmt.Fprintf(out, "Case #%d: Impossible\n", caseNo)
			continue
		}
		town := make([][]int, n)
		for i := 0; i < n; i++ {
			town[i] = make([]int, m)
		}
		minDurations := []int{0, 0}
		flip := false
		if a < b {
			minDurations = []int{a, b}
		} else {
			minDurations = []int{b, a}
			flip = true
		}
		maxDuration := 0
		waitTime := minDurations[0] / minIntersections
		town[0][0] = waitTime + (minDurations[0] % minIntersections)
		for i := 1; i < n; i++ {
			town[i][0] = waitTime
		}
		for i := 0; i < m; i++ {
			town[n-1][i] = waitTime
		}
		rightFastestRemains := minDurations[1] - (waitTime * m)
		waitTime = rightFastestRemains / (n - 1)
		town[0][m-1] = waitTime + (rightFastestRemains % (n - 1))
		for i := 1; i < n-1; i++ {
			town[i][m-1] = waitTime
		}
		maxDuration = town[0][0]
		if town[0][m-1] > maxDuration {
			maxDuration = town[0][m-1]
		}
		fmt.Fprintf(out, "Case #%d: Possible\n", caseNo)
		for i := 0; i < n; i++ {
			row := make([]string, m)
			for j := 0; j < m; j++ {
				if town[i][j] == 0 {
					town[i][j] = maxDuration + 1
				}
				idx := j
				if flip {
					idx = m - j - 1
				}
				row[j] = strconv.Itoa(town[i][idx])
			}
			fmt.Fprintf(out, "%s\n", strings.Join(row, " "))
		}
	}
}
