package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Problem2023Day01Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	sumOfCallibrations := 0
	for {
		input, _, err := in.ReadLine()
		if len(input) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		callibration := 0
		for i := 0; i < len(input); i++ {
			c := input[i]
			if c >= '0' && c <= '9' {
				callibration += int(c - '0')
				callibration *= 10
				break
			}
		}
		for i := len(input) - 1; i >= 0; i-- {
			c := input[i]
			if c >= '0' && c <= '9' {
				callibration += int(c - '0')
				break
			}
		}
		sumOfCallibrations += callibration
	}
	fmt.Fprint(out, sumOfCallibrations)
}

func Problem2023Day01Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	digitsFromString := func(input string) []int {
		digits := make([]int, 0)
		for i := 0; i < len(input); i++ {
			c := input[i]
			if c >= '0' && c <= '9' {
				digits = append(digits, int(c-'0'))
			} else if strings.HasPrefix(input[i:], "one") {
				digits = append(digits, 1)
			} else if strings.HasPrefix(input[i:], "two") {
				digits = append(digits, 2)
			} else if strings.HasPrefix(input[i:], "three") {
				digits = append(digits, 3)
			} else if strings.HasPrefix(input[i:], "four") {
				digits = append(digits, 4)
			} else if strings.HasPrefix(input[i:], "five") {
				digits = append(digits, 5)
			} else if strings.HasPrefix(input[i:], "six") {
				digits = append(digits, 6)
			} else if strings.HasPrefix(input[i:], "seven") {
				digits = append(digits, 7)
			} else if strings.HasPrefix(input[i:], "eight") {
				digits = append(digits, 8)
			} else if strings.HasPrefix(input[i:], "nine") {
				digits = append(digits, 9)
			}
		}
		return digits
	}

	sumOfCallibrations := 0
	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		input := string(raw)
		digits := digitsFromString(input)
		sumOfCallibrations += (digits[0] * 10) + digits[len(digits)-1]
	}
	fmt.Fprint(out, sumOfCallibrations)
}
