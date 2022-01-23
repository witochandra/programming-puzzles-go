package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2021Day02Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	finalX, finalY := 0, 0
	for {
		command := ""
		value := 0
		processed, _ := fmt.Fscanf(in, "%s %d\n", &command, &value)
		if processed == 0 {
			break
		}
		switch command {
		case "forward":
			finalX += value
		case "up":
			finalY -= value
		case "down":
			finalY += value
		}
	}
	fmt.Fprintln(out, finalX*finalY)
}

func Problem2021Day02Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	aim, finalX, finalY := 0, 0, 0
	for {
		command := ""
		value := 0
		processed, _ := fmt.Fscanf(in, "%s %d\n", &command, &value)
		if processed == 0 {
			break
		}
		switch command {
		case "forward":
			finalX += value
			finalY += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	fmt.Fprintln(out, finalX*finalY)
}
