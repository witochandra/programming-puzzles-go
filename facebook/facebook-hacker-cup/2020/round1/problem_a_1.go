package fhc_2020_round1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ProblemA1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfFloorPlans := 0
	fmt.Fscanln(in, &numberOfFloorPlans)
	for caseNo := 1; caseNo <= numberOfFloorPlans; caseNo++ {
		// consume the first line
		tmp, _ := in.ReadString('\n')
		tmps := strings.Split(tmp, " ")
		numberOfRooms, _ := strconv.Atoi(tmps[0])
		k, _ := strconv.Atoi(tmps[1])
		width, _ := strconv.Atoi(strings.TrimSpace(tmps[2]))

		// consume the second line
		tmp, _ = in.ReadString('\n')
		tmps = strings.Split(tmp, " ")
		xs := make([]int, numberOfRooms)
		for i, s := range tmps {
			xs[i], _ = strconv.Atoi(strings.TrimSpace(s))
		}

		// consume the third line
		tmp, _ = in.ReadString('\n')
		tmps = strings.Split(tmp, " ")
		ls := make([]int, 4)
		for i, s := range tmps {
			ls[i], _ = strconv.Atoi(strings.TrimSpace(s))
		}
		for i := k; i < numberOfRooms; i++ {
			xs[i] = (ls[0] * xs[i-2]) + (ls[1] * xs[i-1]) + ls[2]
			xs[i] = xs[i] % ls[3]
			xs[i] = xs[i] + 1
		}

		// consume the fourth line
		tmp, _ = in.ReadString('\n')
		tmps = strings.Split(tmp, " ")
		ys := make([]int, numberOfRooms)
		for i, s := range tmps {
			ys[i], _ = strconv.Atoi(strings.TrimSpace(s))
		}

		// consume the fifth line
		tmp, _ = in.ReadString('\n')
		tmps = strings.Split(tmp, " ")
		hs := make([]int, 4)
		for i, s := range tmps {
			hs[i], _ = strconv.Atoi(strings.TrimSpace(s))
		}
		for i := k; i < numberOfRooms; i++ {
			ys[i] = (hs[0] * ys[i-2]) + (hs[1] * ys[i-1]) + hs[2]
			ys[i] = ys[i] % hs[3]
			ys[i] = ys[i] + 1
		}

		// compute result
		result := 1
		perimeters := make([]int, numberOfRooms)
		for i := 0; i < numberOfRooms; i++ {
			if i == 0 {
				perimeters[i] = 2 * (width + ys[i])
			} else if xs[i-1]+width < xs[i] {
				perimeters[i] = perimeters[i-1] + 2*(width+ys[i])
			} else {
				previousMaxHeight := 0
				for j := i - 1; j >= 0; j-- {
					if xs[j]+width < xs[i] {
						break
					}
					if ys[j] > previousMaxHeight {
						previousMaxHeight = ys[j]
					}
				}
				perimeters[i] = perimeters[i-1] + (2 * (xs[i] - xs[i-1]))
				if ys[i] > previousMaxHeight {
					perimeters[i] = perimeters[i] + (2 * (ys[i] - previousMaxHeight))
				}
			}

			result = (result * (perimeters[i] % 1000000007)) % 1000000007
		}

		// print result
		fmt.Fprintf(out, "Case #%d: %d\n", caseNo, result)
	}
}
