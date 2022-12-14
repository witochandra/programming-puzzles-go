package adventofcode

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Problem2022Day02Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	scoreByHand := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	totalScore := 0
	for {
		input, _, _ := in.ReadLine()
		if len(input) == 0 {
			break
		}
		hands := strings.Split(string(input), " ")
		totalScore += scoreByHand[hands[1]]
		diff := scoreByHand[hands[1]] - scoreByHand[hands[0]]
		if diff == 1 || diff == -2 {
			totalScore += 6
		} else if diff == 0 {
			totalScore += 3
		}
	}
	out.WriteString(strconv.Itoa(totalScore))
}

func Problem2022Day02Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	scoreByHand := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	weaknessByHand := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}
	strengthByHand := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}
	totalScore := 0
	for {
		input, _, _ := in.ReadLine()
		if len(input) == 0 {
			break
		}
		hands := strings.Split(string(input), " ")
		myHand := ""
		if hands[1] == "X" {
			myHand = strengthByHand[hands[0]]
		} else if hands[1] == "Y" {
			myHand = hands[0]
			totalScore += 3
		} else {
			myHand = weaknessByHand[hands[0]]
			totalScore += 6
		}
		totalScore += scoreByHand[myHand]
	}
	out.WriteString(strconv.Itoa(totalScore))
}
