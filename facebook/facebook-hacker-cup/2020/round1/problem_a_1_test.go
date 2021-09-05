package fhc_2020_round1

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemA1(t *testing.T) {
	input := `
5
2 2 2
1 2
0 0 0 100
3 3
0 0 0 100
2 2 2
10 20
0 0 0 100
3 3
0 0 0 100
5 5 3
2 4 5 9 12
0 0 0 100
4 3 6 3 2
0 0 0 100
10 3 8
9 14 15
0 1 3 53
12 7 16
5 2 1 38
50 10 17
4 9 10 26 28 59 97 100 105 106
1 0 7 832
130 12 82 487 12 30 214 104 104 527
21 81 410 605
`

	expected := `
Case #1: 120
Case #2: 200
Case #3: 9144576
Case #4: 803986060
Case #5: 271473330
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	ProblemA1(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}

func TestProblemA1_Submission(t *testing.T) {
	input, _ := os.ReadFile("./problem_a_1_input.txt")
	expected, _ := os.ReadFile("./problem_a_1_expected.txt")
	reader := strings.NewReader(string(input))
	writer := &strings.Builder{}

	ProblemA1(reader, writer)

	// os.WriteFile("./problem_a_1_expected.txt", []byte(writer.String()), 0644)
	assert.Equal(t, string(expected), writer.String())
}
