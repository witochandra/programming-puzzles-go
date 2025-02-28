package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2023Day04Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	totalScore := 0

	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		cardStr := strings.Split(string(raw), ":")
		cardStr = strings.Split(cardStr[1], " | ")
		winningStrs := strings.Split(cardStr[0], " ")
		mineStrs := strings.Split(cardStr[1], " ")

		score := 0
		winnings := make(map[string]bool)

		for _, winning := range winningStrs {
			if winning == "" {
				continue
			}

			winnings[strings.TrimSpace(winning)] = true
		}
		for _, mine := range mineStrs {
			if mine == "" {
				continue
			}

			mine = strings.TrimSpace(mine)
			if _, ok := winnings[mine]; !ok {
				continue
			}

			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}

		totalScore += score
	}

	fmt.Fprintf(out, "%d", totalScore)
}

func Problem2023Day04Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	cards := make(map[int]int)
	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		cardStr := strings.Split(string(raw), ":")
		cardNo, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(cardStr[0], "Card ", "")))
		cards[cardNo] += 1

		cardStr = strings.Split(cardStr[1], " | ")
		winningStrs := strings.Split(cardStr[0], " ")
		mineStrs := strings.Split(cardStr[1], " ")

		score := 0
		winnings := make(map[string]bool)

		for _, winning := range winningStrs {
			if winning == "" {
				continue
			}

			winnings[strings.TrimSpace(winning)] = true
		}
		for _, mine := range mineStrs {
			if mine == "" {
				continue
			}

			mine = strings.TrimSpace(mine)
			if _, ok := winnings[mine]; !ok {
				continue
			}

			score += 1
		}
		for i := cardNo + 1; i < cardNo+score+1; i++ {
			cards[i] += cards[cardNo]
		}
	}

	totalScore := 0
	for _, v := range cards {
		totalScore += v
	}

	fmt.Fprintf(out, "%d", totalScore)
}
