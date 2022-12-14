package adventofcode

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Problem2022Day03Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	totalPriorities := 0
	for {
		input, _, _ := in.ReadLine()
		if len(input) == 0 {
			break
		}
		commonItems := make(map[rune]bool)
		lhs := string(input[:len(input)/2])
		rhs := string(input[len(input)/2:])
		for _, c := range rhs {
			if commonItems[c] {
				continue
			}
			if strings.ContainsRune(lhs, c) {
				commonItems[c] = true
			}
		}
		for r := range commonItems {
			if r >= 'a' {
				totalPriorities += int(r - 'a' + 1)
			} else {
				totalPriorities += int(r-'A'+1) + 26
			}
		}
	}
	out.WriteString(strconv.Itoa(totalPriorities))
}

func Problem2022Day03Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	totalPriorities := 0
	for {
		groups := make([]string, 0)
		group1, _, _ := in.ReadLine()
		groups = append(groups, string(group1))
		group2, _, _ := in.ReadLine()
		groups = append(groups, string(group2))
		group3, _, _ := in.ReadLine()
		groups = append(groups, string(group3))
		if len(group1) == 0 {
			break
		}
		runeCountByGroup := make([]map[rune]int, len(groups))
		for i, group := range groups {
			runeCountByGroup[i] = make(map[rune]int)
			for _, c := range group {
				runeCountByGroup[i][c] += 1
			}
		}

		calculated := make(map[rune]bool)
		for _, c := range groups[0] {
			if calculated[c] {
				continue
			}
			calculated[c] = true
			if runeCountByGroup[1][c] == 0 || runeCountByGroup[2][c] == 0 {
				continue
			}
			if c >= 'a' {
				totalPriorities += int(c - 'a' + 1)
			} else {
				totalPriorities += int(c-'A'+1) + 26
			}
			break
		}
	}
	out.WriteString(strconv.Itoa(totalPriorities))
}
