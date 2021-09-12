package fhc_2021_round1

import (
	"bufio"
	"fmt"
	"io"
)

func ProblemA2(r io.Reader, w io.Writer) {
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

		switchIndices := make([][]int, 0)
		currentPoint := -1
		currentRune := rune('F')

		for i, c := range exercise {
			if c == 'X' {
				if currentRune == 'F' || currentRune == 'X' {
					currentPoint = i
				} else if currentRune == 'O' {
					switchIndices = append(switchIndices, []int{currentPoint, i})
					currentPoint = i
				}
				currentRune = 'X'
			} else if c == 'O' {
				if currentRune == 'F' || currentRune == 'O' {
					currentPoint = i
				} else if currentRune == 'X' {
					switchIndices = append(switchIndices, []int{currentPoint, i})
					currentPoint = i
				}
				currentRune = 'O'
			}
		}
		result := 0
		for _, indices := range switchIndices {
			leftSideLength := indices[0]
			rightSideLength := n - indices[1] - 1
			length := leftSideLength + rightSideLength + 1

			result = result + ((length*(length+1))-(leftSideLength*(leftSideLength+1))-(rightSideLength*(rightSideLength+1)))/2
			result = result % 1000000007
		}
		fmt.Fprintf(out, "Case #%d: %d\n", caseNo, result)
	}
}
