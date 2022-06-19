package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2021Day10Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	score := 0
	scoreTable := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	pairTable := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	for {
		row := ""
		processed, _ := fmt.Fscanln(in, &row)
		if processed == 0 {
			break
		}
		buffer := make([]rune, 0)
		currentScore := 0
		for _, c := range row {
			switch c {
			case '(', '[', '{', '<':
				buffer = append(buffer, c)
			case ')', ']', '}', '>':
				if len(buffer) > 0 {
					prevOpening := buffer[len(buffer)-1]
					if pairTable[prevOpening] == c {
						buffer = buffer[:len(buffer)-1]
						break
					}
				}
				currentScore = scoreTable[c]
			}
			if currentScore > 0 {
				score += currentScore
				break
			}
		}
	}
	fmt.Fprint(out, score)
}
