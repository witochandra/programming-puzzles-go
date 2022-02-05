package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day02_Example(t *testing.T) {
	input := strings.TrimSpace(`
forward 5
down 5
forward 8
up 3
down 8
forward 2
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "150"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day02Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "900"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day02Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}

func TestProblem2021Day02_Puzzle(t *testing.T) {
	input, err := os.ReadFile("./day02_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("part 1", func(t *testing.T) {
		expected := "1499229"

		reader := strings.NewReader(strings.TrimSpace(string(input)))
		writer := &strings.Builder{}

		Problem2021Day02Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "1340836560"

		reader := strings.NewReader(strings.TrimSpace(string(input)))
		writer := &strings.Builder{}

		Problem2021Day02Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}
