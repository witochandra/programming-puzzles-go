package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Problem2022Day01Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	mostCalories := 0
	currentCalories := 0
	for {
		input, _, err := in.ReadLine()
		if len(input) == 0 {
			currentCalories = 0
			if err == io.EOF {
				break
			}
			continue
		}
		calories, _ := strconv.Atoi(string(input))
		currentCalories += calories
		if currentCalories > mostCalories {
			mostCalories = currentCalories
		}
	}
	fmt.Fprint(out, mostCalories)
}

func Problem2022Day01Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	topX := 3
	topCalories := make([]int, topX)
	currentCalories := 0
	for {
		input, _, err := in.ReadLine()
		if len(input) == 0 {
			for i := 0; i < topX; i++ {
				if currentCalories > topCalories[i] {
					for j := i + 1; j < topX-1; j++ {
						topCalories[j] = topCalories[j-1]
					}
					topCalories[i] = currentCalories
					break
				}
			}
			currentCalories = 0
			if err == io.EOF {
				break
			}
			continue
		}
		calories, _ := strconv.Atoi(string(input))
		currentCalories += calories
	}
	fmt.Fprint(out, topCalories[0]+topCalories[1]+topCalories[2])
}
