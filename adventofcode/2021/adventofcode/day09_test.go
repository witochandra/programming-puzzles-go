package adventofcode

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day09(t *testing.T) {
	example := `2199943210
3987894921
9856789892
8767896789
9899965678`

	fileContent, err := os.ReadFile("./day09_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	puzzle := string(fileContent)

	testCases := []struct {
		name  string
		input string
		want  string
		eval  func(r io.Reader, w io.Writer)
	}{
		{
			name:  "part 1 example",
			input: example,
			want:  "15",
			eval:  Problem2021Day09Part1,
		},
		{
			name:  "part 1 puzzle",
			input: puzzle,
			want:  "516",
			eval:  Problem2021Day09Part1,
		},
		{
			name:  "part 2 example",
			input: example,
			want:  "1134",
			eval:  Problem2021Day09Part2,
		},
		{
			name:  "part 2 puzzle",
			input: puzzle,
			want:  "1023660",
			eval:  Problem2021Day09Part2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			w := &strings.Builder{}
			tc.eval(r, w)
			assert.Equal(t, tc.want, w.String())
		})
	}
}
