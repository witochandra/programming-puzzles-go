package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ProblemA1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfBirthdays := 0
	fmt.Fscanln(in, &numberOfBirthdays)

	for caseNo := 1; caseNo <= numberOfBirthdays; caseNo++ {
		input := ""
		fmt.Fscanln(in, &input)

		charCount := make(map[rune]int)
		var consonants, vowels int
		for _, c := range input {
			if strings.ContainsRune("AEIOU", c) {
				vowels++
			} else {
				consonants++
			}
			charCount[c]++
		}
		var maxSameConsonants, maxSameVowels int
		for c, count := range charCount {
			isVowels := strings.ContainsRune("AEIOU", c)
			if isVowels && maxSameVowels < count {
				maxSameVowels = count
			} else if !isVowels && maxSameConsonants < count {
				maxSameConsonants = count
			}
		}
		result := func() int {
			secondsToAllVowels := consonants + (2 * (vowels - maxSameVowels))
			secondsToAllConsonants := vowels + (2 * (consonants - maxSameConsonants))
			if secondsToAllVowels < secondsToAllConsonants {
				return secondsToAllVowels
			}
			return secondsToAllConsonants
		}()
		fmt.Fprintf(out, "Case #%d: %d\n", caseNo, result)
	}
}
