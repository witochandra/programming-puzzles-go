package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2023Day02Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	maxRed, maxGreen, maxBlue := 12, 13, 14
	result := 0
	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		possible := true
		input := string(raw)
		record := strings.Split(input, ": ")
		gameID, _ := strconv.Atoi(record[0][5:])
		grabs := strings.Split(record[1], "; ")
		for _, grab := range grabs {
			red, green, blue := 0, 0, 0
			cubes := strings.Split(grab, ", ")
			for _, cube := range cubes {
				if strings.Contains(cube, "red") {
					red, _ = strconv.Atoi(strings.Split(cube, " ")[0])
				} else if strings.Contains(cube, "green") {
					green, _ = strconv.Atoi(strings.Split(cube, " ")[0])
				} else if strings.Contains(cube, "blue") {
					blue, _ = strconv.Atoi(strings.Split(cube, " ")[0])
				}
			}
			if red > maxRed || green > maxGreen || blue > maxBlue {
				possible = false
				break
			}
		}
		if possible {
			result += gameID
		}
	}
	fmt.Fprint(out, result)
}

func Problem2023Day02Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	result := 0
	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		maxRed, maxGreen, maxBlue := 0, 0, 0
		input := string(raw)
		record := strings.Split(input, ": ")
		grabs := strings.Split(record[1], "; ")
		for _, grab := range grabs {
			red, green, blue := 0, 0, 0
			cubes := strings.Split(grab, ", ")
			for _, cube := range cubes {
				if strings.Contains(cube, "red") {
					red, _ = strconv.Atoi(strings.Split(cube, " ")[0])
				} else if strings.Contains(cube, "green") {
					green, _ = strconv.Atoi(strings.Split(cube, " ")[0])
				} else if strings.Contains(cube, "blue") {
					blue, _ = strconv.Atoi(strings.Split(cube, " ")[0])
				}
			}
			if red > maxRed {
				maxRed = red
			}
			if green > maxGreen {
				maxGreen = green
			}
			if blue > maxBlue {
				maxBlue = blue
			}
		}
		result += (maxRed * maxGreen * maxBlue)
	}
	fmt.Fprint(out, result)
}
