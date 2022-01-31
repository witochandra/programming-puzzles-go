package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day05_Example(t *testing.T) {
	input := strings.TrimSpace(`
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "5"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day05Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "12"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day05Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}

func TestProblem2021Day05_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day05_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		expected := "8111"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day05Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "22088"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day05Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}
