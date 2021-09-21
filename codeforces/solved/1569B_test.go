package codeforces

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1569B(t *testing.T) {
	input := `
3
3
111
2
21
4
2122
`
	expected := `
YES
X==
=X=
==X
NO
YES
X=+-
=X==
-=X+
+=-X
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	Problem1569B(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
