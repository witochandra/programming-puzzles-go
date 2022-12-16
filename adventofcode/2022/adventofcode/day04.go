package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Problem2022Day04Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	overlaps := 0
	for {
		min1, max1, min2, max2 := 0, 0, 0, 0
		_, err := fmt.Fscanf(in, "%d-%d,%d-%d\n", &min1, &max1, &min2, &max2)
		if err == io.EOF {
			break
		}
		if (min1 <= min2 && max1 >= max2) ||
			(min2 <= min1 && max2 >= max1) {
			overlaps += 1
		}
	}
	out.WriteString(strconv.Itoa(overlaps))
}

func Problem2022Day04Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	overlaps := 0
	for {
		min1, max1, min2, max2 := 0, 0, 0, 0
		_, err := fmt.Fscanf(in, "%d-%d,%d-%d\n", &min1, &max1, &min2, &max2)
		if err == io.EOF {
			break
		}
		if (min1 >= min2 && min1 <= max2) ||
			(max1 >= min2 && max1 <= max2) ||
			(min2 >= min1 && min2 <= max1) ||
			(max2 >= min1 && max2 <= max1) {
			overlaps += 1
		}
	}
	out.WriteString(strconv.Itoa(overlaps))
}
