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

func Problem2021Day10Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	scores := make([]int, 0)
	scoreTable := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
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
		corrupted := false
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
					corrupted = true
				}
			}
			if corrupted {
				break
			}
		}
		if corrupted {
			continue
		}

		currentScore := 0
		for i := len(buffer) - 1; i >= 0; i-- {
			currentScore *= 5
			pair := pairTable[buffer[i]]
			currentScore += scoreTable[pair]
		}
		position := len(scores)
		scores = append(scores, currentScore)
		for position > 0 {
			prevScore := scores[position-1]
			if prevScore > currentScore {
				scores[position-1] = scores[position]
				scores[position] = prevScore
				position--
				continue
			}
			break
		}
	}
	fmt.Fprint(out, scores[len(scores)/2])
}
