package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day07_Example(t *testing.T) {
	input := strings.TrimSpace(`
16,1,2,0,4,2,7,1,2,14
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "37"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day07Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "168"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day07Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}

func TestProblem2021Day07_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day07_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := string(fileContent)
	t.Run("part 1", func(t *testing.T) {
		expected := "323647"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day07Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "87640209"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day07Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}
