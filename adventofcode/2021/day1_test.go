package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day1_Example(t *testing.T) {
	input := strings.TrimSpace(`
199
200
208
210
200
207
240
269
260
263
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "7"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day1Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "5"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day1Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}

func TestProblem2021Day1_Puzzle(t *testing.T) {
	input, err := os.ReadFile("day1_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("part 1", func(t *testing.T) {
		expected := "1288"

		reader := strings.NewReader(strings.TrimSpace(string(input)))
		writer := &strings.Builder{}

		Problem2021Day1Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "1311"

		reader := strings.NewReader(strings.TrimSpace(string(input)))
		writer := &strings.Builder{}

		Problem2021Day1Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}
