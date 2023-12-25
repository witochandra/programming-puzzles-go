package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Problem2023Day03Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	schematic := make([]string, 0)
	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		schematic = append(schematic, string(raw))
	}
	result := 0
	notSymbol := "0123456789."
	isPartNumber := false
	currentNumber := 0
	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			value := schematic[y][x]
			if value < '0' || value > '9' {
				if isPartNumber {
					result += currentNumber
				}
				currentNumber = 0
				isPartNumber = false
				continue
			}
			currentNumber = (currentNumber * 10) + int(value-'0')

			if !isPartNumber {
				north := y > 0 && !strings.ContainsRune(notSymbol, rune(schematic[y-1][x]))
				northEast := y > 0 && x < len(schematic[y])-1 && !strings.ContainsRune(notSymbol, rune(schematic[y-1][x+1]))
				east := x < len(schematic[y])-1 && !strings.ContainsRune(notSymbol, rune(schematic[y][x+1]))
				southEast := y < len(schematic)-1 && x < len(schematic[y])-1 && !strings.ContainsRune(notSymbol, rune(schematic[y+1][x+1]))
				south := y < len(schematic)-1 && !strings.ContainsRune(notSymbol, rune(schematic[y+1][x]))
				southWest := y < len(schematic)-1 && x > 0 && !strings.ContainsRune(notSymbol, rune(schematic[y+1][x-1]))
				west := x > 0 && !strings.ContainsRune(notSymbol, rune(schematic[y][x-1]))
				northWest := y > 0 && x > 0 && !strings.ContainsRune(notSymbol, rune(schematic[y-1][x-1]))
				isPartNumber = north || northEast || east || southEast || south || southWest || west || northWest
			}
		}
	}
	if isPartNumber {
		result += currentNumber
	}
	fmt.Fprint(out, result)
}

func Problem2023Day03Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	schematic := make([]string, 0)
	for {
		raw, _, err := in.ReadLine()
		if len(raw) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		schematic = append(schematic, string(raw))
	}
	result := 0
	currentNumber := 0
	currentGears := make(map[string]bool)
	gears := make(map[string][]int)
	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			value := schematic[y][x]
			if value < '0' || value > '9' {
				if currentNumber > 0 {
					for gear := range currentGears {
						gears[gear] = append(gears[gear], currentNumber)
					}
				}
				currentNumber = 0
				currentGears = make(map[string]bool)
				continue
			}
			currentNumber = (currentNumber * 10) + int(value-'0')

			if y > 0 && schematic[y-1][x] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x, y-1)] = true
			}
			if y > 0 && x < len(schematic[y])-1 && schematic[y-1][x+1] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x+1, y-1)] = true
			}
			if x < len(schematic[y])-1 && schematic[y][x+1] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x+1, y)] = true
			}
			if y < len(schematic)-1 && x < len(schematic[y])-1 && schematic[y+1][x+1] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x+1, y+1)] = true
			}
			if y < len(schematic)-1 && schematic[y+1][x] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x, y+1)] = true
			}
			if y < len(schematic)-1 && x > 0 && schematic[y+1][x-1] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x-1, y+1)] = true
			}
			if x > 0 && schematic[y][x-1] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x-1, y)] = true
			}
			if y > 0 && x > 0 && schematic[y-1][x-1] == '*' {
				currentGears[fmt.Sprintf("%d,%d", x-1, y-1)] = true
			}
		}
	}
	if currentNumber > 0 {
		for gear := range currentGears {
			gears[gear] = append(gears[gear], currentNumber)
		}
	}
	for _, parts := range gears {
		if len(parts) != 2 {
			continue
		}
		result += parts[0] * parts[1]
	}
	fmt.Fprint(out, result)
}
