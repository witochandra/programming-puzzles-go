package codeforces

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1569A(t *testing.T) {
	input := `
4
1
a
6
abbaba
6
abbaba
9
babbabbaa	
`
	expected := `
-1 -1
1 2
1 2
1 2
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	Problem1569A(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
