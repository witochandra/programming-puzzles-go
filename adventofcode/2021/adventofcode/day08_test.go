package adventofcode

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day08(t *testing.T) {
	example := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

	fileContent, err := os.ReadFile("./day08_input.txt")
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
			want:  "26",
			eval:  Problem2021Day08Part1,
		},
		{
			name:  "part 1 puzzle",
			input: puzzle,
			want:  "264",
			eval:  Problem2021Day08Part1,
		},
		{
			name:  "part 2 example",
			input: example,
			want:  "61229",
			eval:  Problem2021Day08Part2,
		},
		{
			name:  "part 2 puzzle",
			input: puzzle,
			want:  "1063760",
			eval:  Problem2021Day08Part2,
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
