package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2022Day09Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	absInt := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	headPos, tailPos := []int{0, 0}, []int{0, 0}
	travelled := make(map[string]bool)
	for {
		raw, _ := in.ReadString('\n')
		raw = strings.TrimSpace(raw)
		if len(raw) == 0 {
			break
		}
		input := strings.Split(raw, " ")
		direction := input[0]
		distance, _ := strconv.Atoi(input[1])
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				headPos[1]++
			case "D":
				headPos[1]--
			case "R":
				headPos[0]++
			case "L":
				headPos[0]--
			}
			diffX, diffY := headPos[0]-tailPos[0], headPos[1]-tailPos[1]
			if absInt(diffX) > 1 && absInt(diffY) == 0 {
				tailPos[0] += diffX / absInt(diffX)
			} else if absInt(diffY) > 1 && absInt(diffX) == 0 {
				tailPos[1] += diffY / absInt(diffY)
			} else if absInt(diffX)+absInt(diffY) > 2 {
				tailPos[0] += diffX / absInt(diffX)
				tailPos[1] += diffY / absInt(diffY)
			}
			travelled[fmt.Sprintf("%d,%d", tailPos[0], tailPos[1])] = true
		}
	}
	fmt.Fprint(out, len(travelled))
}

func Problem2022Day09Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	absInt := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	knotPositions := make([][]int, 0)
	for i := 0; i < 10; i++ {
		knotPositions = append(knotPositions, []int{0, 0})
	}
	travelled := make(map[string]bool)
	for {
		raw, _ := in.ReadString('\n')
		raw = strings.TrimSpace(raw)
		if len(raw) == 0 {
			break
		}
		input := strings.Split(raw, " ")
		direction := input[0]
		distance, _ := strconv.Atoi(input[1])
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				knotPositions[0][1]++
			case "D":
				knotPositions[0][1]--
			case "R":
				knotPositions[0][0]++
			case "L":
				knotPositions[0][0]--
			}
			for j := 1; j < len(knotPositions); j++ {
				diffX := knotPositions[j-1][0] - knotPositions[j][0]
				diffY := knotPositions[j-1][1] - knotPositions[j][1]
				if absInt(diffX) > 1 && absInt(diffY) == 0 {
					knotPositions[j][0] += diffX / absInt(diffX)
				} else if absInt(diffY) > 1 && absInt(diffX) == 0 {
					knotPositions[j][1] += diffY / absInt(diffY)
				} else if absInt(diffX)+absInt(diffY) > 2 {
					knotPositions[j][0] += diffX / absInt(diffX)
					knotPositions[j][1] += diffY / absInt(diffY)
				} else {
					break
				}
			}
			tailPos := knotPositions[len(knotPositions)-1]
			travelled[fmt.Sprintf("%d,%d", tailPos[0], tailPos[1])] = true
		}
	}
	fmt.Fprint(out, len(travelled))
}
