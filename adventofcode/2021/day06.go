package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2021Day06(r io.Reader, w io.Writer, days int) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	input := ""
	fmt.Fscanln(in, &input)
	countByTimer := make(map[int]int64)
	for _, s := range strings.Split(input, ",") {
		t, _ := strconv.Atoi(s)
		countByTimer[t]++
	}
	for day := 0; day < days; day++ {
		newBorns := countByTimer[0]
		for i := 0; i < 8; i++ {
			countByTimer[i] = countByTimer[i+1]
		}
		countByTimer[6] += newBorns
		countByTimer[8] = newBorns
	}
	var total int64
	for _, count := range countByTimer {
		total += count
	}
	fmt.Fprintln(out, total)
}

func Problem2021Day06Part1(r io.Reader, w io.Writer) {
	Problem2021Day06(r, w, 80)
}

func Problem2021Day06Part2(r io.Reader, w io.Writer) {
	Problem2021Day06(r, w, 256)
}
