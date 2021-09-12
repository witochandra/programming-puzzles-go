package fhc_2021_round1

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemA1(t *testing.T) {
	input := `
5
1
O
3
XFO
5
FFOFF
10
FXXFXFOOXF
13
XFOFXFOFXFOFX
	`
	expected := `
Case #1: 0
Case #2: 1
Case #3: 0
Case #4: 2
Case #5: 6
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	ProblemA1(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}

func TestProblemA1_Validation(t *testing.T) {
	input, _ := os.ReadFile("./problem_a_1_validation_input.txt")
	expected, _ := os.ReadFile("./problem_a_1_validation_output.txt")
	reader := strings.NewReader(string(input))
	writer := &strings.Builder{}

	ProblemA1(reader, writer)

	// os.WriteFile("./problem_a_1_validation_output.txt", []byte(writer.String()), 0644)
	assert.Equal(t, string(expected), writer.String())
}

func TestProblemA1_Submission(t *testing.T) {
	input, _ := os.ReadFile("./problem_a_1_submission_input.txt")
	expected, _ := os.ReadFile("./problem_a_1_submission_output.txt")
	reader := strings.NewReader(string(input))
	writer := &strings.Builder{}

	ProblemA1(reader, writer)

	// os.WriteFile("./problem_a_1_submission_output.txt", []byte(writer.String()), 0644)
	assert.Equal(t, string(expected), writer.String())
}
