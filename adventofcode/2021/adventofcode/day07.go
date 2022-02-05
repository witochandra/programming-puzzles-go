package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Problem2021Day07Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	input := ""
	fmt.Fscanln(in, &input)

	positions := make([]int, 0)

	for _, s := range strings.Split(input, ",") {
		t, _ := strconv.Atoi(s)
		positions = append(positions, t)
	}
	sort.Ints(positions)
	target := positions[len(positions)/2]
	fuels := 0
	for _, pos := range positions {
		distance := 0
		if pos > target {
			distance = pos - target
		} else if pos < target {
			distance = target - pos
		}
		fuels += distance
	}
	fmt.Fprint(out, fuels)
}

func Problem2021Day07Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	input := ""
	fmt.Fscanln(in, &input)

	positions := make([]int, 0)

	for _, s := range strings.Split(input, ",") {
		t, _ := strconv.Atoi(s)
		positions = append(positions, t)
	}
	sort.Ints(positions)
	calculateFuels := func(target int) int {
		fuels := 0
		for _, pos := range positions {
			distance := 0
			if pos > target {
				distance = pos - target
			} else if pos < target {
				distance = target - pos
			}
			fuels += (distance + 1) * distance / 2
		}
		return fuels
	}
	fuelsByTarget := make(map[int]int)
	currentTarget := positions[len(positions)/2]
	for {
		fuelsByTarget[currentTarget] = calculateFuels(currentTarget)
		if _, found := fuelsByTarget[currentTarget-1]; !found {
			fuelsByTarget[currentTarget-1] = calculateFuels(currentTarget - 1)
		}
		if _, found := fuelsByTarget[currentTarget+1]; !found {
			fuelsByTarget[currentTarget+1] = calculateFuels(currentTarget + 1)
		}
		if fuelsByTarget[currentTarget] < fuelsByTarget[currentTarget-1] && fuelsByTarget[currentTarget] < fuelsByTarget[currentTarget+1] {
			break
		}
		if fuelsByTarget[currentTarget-1] < fuelsByTarget[currentTarget+1] {
			currentTarget--
		} else {
			currentTarget++
		}
	}
	fmt.Fprint(out, fuelsByTarget[currentTarget])
}
