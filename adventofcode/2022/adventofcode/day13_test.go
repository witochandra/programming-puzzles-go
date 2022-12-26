package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day13_Example(t *testing.T) {
	input := strings.TrimSpace(`
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day13Part1(strings.NewReader(input), writer)

		assert.Equal(t, "13", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day13Part2(strings.NewReader(input), writer)

		assert.Equal(t, "140", writer.String())
	})
}

func TestProblem2022Day13_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day13_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day13Part1(strings.NewReader(input), writer)

		assert.Equal(t, "5605", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day13Part2(strings.NewReader(input), writer)

		assert.Equal(t, "24969", writer.String())
	})
}
