package adventofcode

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Problem2022Day10Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	cycle, x, signalStrength := 0, 1, 0
	incrementCycle := func() {
		cycle++
		if (cycle+20)%40 == 0 {
			signalStrength += cycle * x
		}
	}
	for {
		raw, _ := in.ReadString('\n')
		if len(raw) == 0 {
			break
		}
		input := strings.TrimSpace(raw)
		if strings.HasPrefix(input, "noop") {
			incrementCycle()
		} else if strings.HasPrefix(input, "addx ") {
			toAdd, _ := strconv.Atoi(strings.TrimPrefix(input, "addx "))
			incrementCycle()
			incrementCycle()
			x += toAdd
		}
	}
	out.WriteString(strconv.Itoa(signalStrength))
}

func Problem2022Day10Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	cycle, x := 0, 1
	crt := ""
	incrementCycle := func() {
		cycle++
		position := (cycle - 1) % 40
		if position >= x-1 && position <= x+1 {
			crt += "#"
		} else {
			crt += "."
		}
		if cycle%40 == 0 {
			crt += "\n"
		}
	}
	for {
		raw, _ := in.ReadString('\n')
		if len(raw) == 0 {
			break
		}
		input := strings.TrimSpace(raw)
		if strings.HasPrefix(input, "noop") {
			incrementCycle()
		} else if strings.HasPrefix(input, "addx ") {
			toAdd, _ := strconv.Atoi(strings.TrimPrefix(input, "addx "))
			incrementCycle()
			incrementCycle()
			x += toAdd
		}
	}
	out.WriteString(crt)
}
