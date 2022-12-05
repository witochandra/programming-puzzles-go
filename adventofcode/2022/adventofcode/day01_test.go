package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day01_Example(t *testing.T) {
	input := strings.TrimSpace(`
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "24000"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2022Day01Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "45000"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2022Day01Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}

func TestProblem2022Day01_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day01_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		expected := "73211"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2022Day01Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "213958"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2022Day01Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}
