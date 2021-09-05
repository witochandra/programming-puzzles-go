package fhc_qualification_2020

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func canTravel(incomingRules, outgoingRules []bool, from, to int) bool {
	distance := to - from
	if distance == 0 {
		return true
	}
	if distance == -1 || distance == 1 {
		return incomingRules[to] && outgoingRules[from]
	}
	if distance < -1 {
		return canTravel(incomingRules, outgoingRules, from, from-1) &&
			canTravel(incomingRules, outgoingRules, from-1, to)
	}
	return canTravel(incomingRules, outgoingRules, from, from+1) &&
		canTravel(incomingRules, outgoingRules, from+1, to)
}

func ProblemA(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfAirlines := 0
	fmt.Fscanln(in, &numberOfAirlines)
	for caseNumber := 1; caseNumber <= numberOfAirlines; caseNumber++ {
		numberOfCountries := 0
		incomingRulesStr := ""
		outgoingRulesStr := ""

		fmt.Fscanln(in, &numberOfCountries)
		fmt.Fscanln(in, &incomingRulesStr)
		fmt.Fscanln(in, &outgoingRulesStr)
		incomingRules := make([]bool, numberOfCountries)
		outgoingRules := make([]bool, numberOfCountries)

		for i := 0; i < numberOfCountries; i++ {
			incomingRules[i] = incomingRulesStr[i] == 'Y'
			outgoingRules[i] = outgoingRulesStr[i] == 'Y'
		}

		results := make([]string, 0)
		for i := 0; i < numberOfCountries; i++ {
			result := make([]string, 0)
			for j := 0; j < numberOfCountries; j++ {
				if canTravel(incomingRules, outgoingRules, i, j) {
					result = append(result, "Y")
				} else {
					result = append(result, "N")
				}
			}
			results = append(results, strings.Join(result, ""))
		}

		fmt.Fprintf(out, "Case #%d:\n", caseNumber)
		fmt.Fprintf(out, "%s\n", strings.Join(results, "\n"))
	}
}
