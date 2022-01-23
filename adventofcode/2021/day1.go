package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2021Day1Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	measurementCount := 0
	before := 0
	increments := 0

	for {
		measurement := 0
		processed, _ := fmt.Fscanln(in, &measurement)
		if processed == 0 {
			break
		}
		if measurementCount > 0 {
			if before < measurement {
				increments++
			}
			before = measurement
		}
		measurementCount++
	}
	fmt.Fprintln(out, increments)
}

func Problem2021Day1Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	increments := 0
	windows := make([]int, 0)
	for {
		measurement := 0
		processed, _ := fmt.Fscanln(in, &measurement)
		if processed == 0 {
			break
		}
		windows = append(windows, measurement)
		if len(windows) < 4 {
			continue
		}
		if len(windows) > 4 {
			windows = windows[1:]
		}
		firstMeasurement := windows[0] + windows[1] + windows[2]
		secondMeasurement := windows[1] + windows[2] + windows[3]
		if firstMeasurement < secondMeasurement {
			increments++
		}
	}
	fmt.Fprintln(out, increments)
}
