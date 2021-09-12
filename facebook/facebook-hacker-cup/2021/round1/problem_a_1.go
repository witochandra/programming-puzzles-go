package fhc_2021_round1

import (
	"bufio"
	"fmt"
	"io"
)

func ProblemA1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	numberOfExercises := 0
	fmt.Fscanln(in, &numberOfExercises)

	for caseNo := 1; caseNo <= numberOfExercises; caseNo++ {
		n := 0
		exercise := ""
		fmt.Fscanln(in, &n)
		fmt.Fscanln(in, &exercise)

		switches := 0
		currentHand := ""
		for _, c := range exercise {
			if c == 'X' {
				if currentHand == "right" {
					switches++
				}
				currentHand = "left"
			} else if c == 'O' {
				if currentHand == "left" {
					switches++
				}
				currentHand = "right"
			}
		}
		fmt.Fprintf(out, "Case #%d: %d\n", caseNo, switches)
	}
}
