package fhc_2021_round1

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemB(t *testing.T) {
	input := `
6
2 2 999 999
2 3 12 11
4 3 6 6
50 50 1 1
2 2 1000 1000
2 2 3 1000
`
	expected := `
Case #1: Possible
333 333
333 333
Case #2: Possible
6 7 5
2 2 2
Case #3: Possible
1 2 1
1 2 1
1 2 1
1 1 1
Case #4: Impossible
Case #5: Possible
334 334
333 333
Case #6: Possible
1 998
1 1
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	ProblemB(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}

func TestProblemB_Validation(t *testing.T) {
	input, _ := os.ReadFile("./problem_b_validation_input.txt")
	expected, _ := os.ReadFile("./problem_b_validation_output.txt")
	reader := strings.NewReader(string(input))
	writer := &strings.Builder{}

	ProblemB(reader, writer)

	// os.WriteFile("./problem_b_validation_output.txt", []byte(writer.String()), 0644)
	assert.Equal(t, string(expected), writer.String())
}

func TestProblemB_Submission(t *testing.T) {
	input, _ := os.ReadFile("./problem_b_submission_input.txt")
	// expected, _ := os.ReadFile("./problem_b_submission_output.txt")
	reader := strings.NewReader(string(input))
	writer := &strings.Builder{}

	ProblemB(reader, writer)

	os.WriteFile("./problem_b_submission_output.txt", []byte(writer.String()), 0644)
	// assert.Equal(t, string(expected), writer.String())
}
