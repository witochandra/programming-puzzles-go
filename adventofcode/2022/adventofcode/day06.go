package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2022Day06Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	index := 0
	buffer := make([]rune, 4)
	runes := make(map[rune]int)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		runes[r] += 1
		runes[buffer[0]] -= 1
		if runes[buffer[0]] <= 0 {
			delete(runes, buffer[0])
		}
		buffer = append(buffer[1:4], r)
		index += 1
		if len(runes) == 4 {
			break
		}
	}
	fmt.Fprint(out, index)
}

func Problem2022Day06Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	index := 0
	buffer := make([]rune, 14)
	runes := make(map[rune]int)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		runes[r] += 1
		runes[buffer[0]] -= 1
		if runes[buffer[0]] <= 0 {
			delete(runes, buffer[0])
		}
		buffer = append(buffer[1:14], r)
		index += 1
		if len(runes) == 14 {
			break
		}
	}
	fmt.Fprint(out, index)
}
