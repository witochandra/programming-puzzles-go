package codeforces

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1574B(t *testing.T) {
	input := `
7
2 2 1 0
1 1 1 1
1 2 3 2
1 1 4 0
1 1 4 1
1 1 4 2
1 1 4 4
`
	expected := `
YES
NO
YES
NO
YES
YES
NO
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	Problem1574B(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
