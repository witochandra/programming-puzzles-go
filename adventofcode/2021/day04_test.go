package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day04_Example(t *testing.T) {
	input := strings.TrimSpace(`
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "4512"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day04Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "1924"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day04Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}

func TestProblem2021Day04_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day04_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		expected := "41503"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day04Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "3178"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day04Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}
