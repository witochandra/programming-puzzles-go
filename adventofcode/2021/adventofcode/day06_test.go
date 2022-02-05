package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day06_Example(t *testing.T) {
	input := strings.TrimSpace(`
3,4,3,1,2
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "5934"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day06Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "26984457539"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day06Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}

func TestProblem2021Day06_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day06_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		expected := "388419"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day06Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "1740449478328"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day06Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}
