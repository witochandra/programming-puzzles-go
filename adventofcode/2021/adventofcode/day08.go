package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func Problem2021Day08Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	result := 0
	for {
		_, err := in.Peek(1)
		if err == io.EOF {
			break
		}
		input, _, _ := in.ReadLine()
		signals := strings.Split(string(input), " | ")
		outputs := strings.Split(signals[1], " ")
		for _, output := range outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				result += 1
			}
		}
	}
	fmt.Fprint(out, result)
}

func Problem2021Day08Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	result := 0
	for {
		_, err := in.Peek(1)
		if err == io.EOF {
			break
		}
		input, _, _ := in.ReadLine()
		signals := strings.Split(string(input), " | ")
		patterns := strings.Split(signals[0], " ")
		outputs := strings.Split(signals[1], " ")

		sortString := func(w string) string {
			s := strings.Split(w, "")
			sort.Strings(s)
			return strings.Join(s, "")
		}

		valueByPattern := make(map[string]int)
		patternByValue := make(map[int]string)
		fiveSegments := make([]string, 0)
		sixSegments := make([]string, 0)
		for _, pattern := range patterns {
			pattern = sortString(pattern)
			switch len(pattern) {
			case 2:
				valueByPattern[pattern] = 1
				patternByValue[1] = pattern
			case 3:
				valueByPattern[pattern] = 7
				patternByValue[7] = pattern
			case 4:
				valueByPattern[pattern] = 4
				patternByValue[4] = pattern
			case 5:
				fiveSegments = append(fiveSegments, pattern)
			case 6:
				sixSegments = append(sixSegments, pattern)
			case 7:
				valueByPattern[pattern] = 8
				patternByValue[8] = pattern
			}
		}
		calculateMatchingCount := func(lhs, rhs string) int {
			matching := 0
			for _, c := range lhs {
				for _, d := range rhs {
					if c == d {
						matching++
					}
				}
			}
			return matching
		}
		// find 9, 0, 6
		for _, segments := range sixSegments {
			if calculateMatchingCount(segments, patternByValue[4]) == 4 {
				patternByValue[9] = segments
				valueByPattern[segments] = 9
			} else if calculateMatchingCount(segments, patternByValue[1]) == 2 {
				patternByValue[0] = segments
				valueByPattern[segments] = 0
			} else {
				patternByValue[6] = segments
				valueByPattern[segments] = 6
			}
		}

		// find 5, 2, 3
		for _, segments := range fiveSegments {
			if calculateMatchingCount(segments, patternByValue[7]) == 3 {
				patternByValue[3] = segments
				valueByPattern[segments] = 3
			} else if calculateMatchingCount(segments, patternByValue[6]) == 5 {
				patternByValue[5] = segments
				valueByPattern[segments] = 5
			} else {
				patternByValue[2] = segments
				valueByPattern[segments] = 2
			}
		}
		value := 0
		for _, output := range outputs {
			value *= 10
			output = sortString(output)
			value += valueByPattern[output]
		}
		result += value
	}
	fmt.Fprint(out, result)
}
