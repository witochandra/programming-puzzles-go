package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Problem2021Day03Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	rowCount := 0
	bitsCount := make([]int, 0)
	for {
		row := ""
		processed, _ := fmt.Fscanln(in, &row)
		if len(bitsCount) == 0 {
			bitsCount = make([]int, len(row))
		}
		if processed == 0 {
			break
		}
		for i, bit := range row {
			if bit == '1' {
				bitsCount[i]++
			}
		}
		rowCount += 1
	}
	gammaRate := 0
	for i, bitCount := range bitsCount {
		if bitCount > rowCount/2 {
			gammaRate += 1 << (len(bitsCount) - i - 1)
		}
	}
	epsilonRate := (1 << len(bitsCount)) - 1 - gammaRate
	fmt.Fprintln(out, gammaRate*epsilonRate)
}

func Problem2021Day03Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	rows := make([]string, 0)
	for {
		row := ""
		processed, _ := fmt.Fscanln(in, &row)
		if processed == 0 {
			break
		}
		rows = append(rows, row)
	}
	var calculate func(rows []string, pos int, useMostCommon bool) string
	calculate = func(rows []string, pos int, useMostCommon bool) string {
		if len(rows) == 1 {
			return rows[0]
		}
		zeros := make([]string, 0)
		ones := make([]string, 0)
		for _, row := range rows {
			if row[pos] == '0' {
				zeros = append(zeros, row)
			} else {
				ones = append(ones, row)
			}
		}
		if len(ones) >= len(zeros) {
			if useMostCommon {
				return calculate(ones, pos+1, useMostCommon)
			} else {
				return calculate(zeros, pos+1, useMostCommon)
			}
		} else {
			if useMostCommon {
				return calculate(zeros, pos+1, useMostCommon)
			} else {
				return calculate(ones, pos+1, useMostCommon)
			}
		}
	}
	oxygenGeneratorRating, _ := strconv.ParseInt(calculate(rows, 0, true), 2, 64)
	co2ScrubberRating, _ := strconv.ParseInt(calculate(rows, 0, false), 2, 64)
	fmt.Fprintln(out, oxygenGeneratorRating*co2ScrubberRating)
}
