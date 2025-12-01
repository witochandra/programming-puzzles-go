package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Problem2025Day01Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	current := 50
	password := 0
	for {
		input, _, err := in.ReadLine()
		if len(input) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		steps := 1
		if input[0] == 'L' {
			steps = -1
		}
		parsed, _ := strconv.Atoi(string(input[1:]))
		steps = steps * parsed
		current = (current + steps + 100) % 100
		if current == 0 {
			password++
		}
	}
	fmt.Fprint(out, password)
}

func Problem2025Day01Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	current := 50
	password := 0
	for {
		input, _, err := in.ReadLine()
		if len(input) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		steps := 1
		if input[0] == 'L' {
			steps = -1
		}
		parsed, _ := strconv.Atoi(string(input[1:]))
		password += parsed / 100
		parsed = parsed % 100

		steps = steps * parsed
		prev := current
		current = current + steps

		if current == 0 {
			password++
		} else if current < 0 && prev != 0 {
			password++
		} else if current >= 100 && prev > 0 {
			password++
		}
		current = (current + 100) % 100
	}
	fmt.Fprint(out, password)
}
