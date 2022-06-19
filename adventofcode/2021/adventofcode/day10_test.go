package adventofcode

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day10(t *testing.T) {
	example := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	fileContent, err := os.ReadFile("./day10_input.txt")
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
			want:  "26397",
			eval:  Problem2021Day10Part1,
		},
		{
			name:  "part 1 puzzle",
			input: puzzle,
			want:  "168417",
			eval:  Problem2021Day10Part1,
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
