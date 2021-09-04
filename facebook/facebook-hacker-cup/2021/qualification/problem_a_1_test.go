package fhc_qualification_2021

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemA1(t *testing.T) {
	input := `
6
ABC
F
BANANA
FBHC
FOXEN
CONSISTENCY
	`
	expected := `
Case #1: 2
Case #2: 0
Case #3: 3
Case #4: 4
Case #5: 5
Case #6: 12
	`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	ProblemA1(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
