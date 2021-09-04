package fhc_qualification_2021

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemA2(t *testing.T) {
	input := func() string {
		return `
7
ABC
2
BA
CA
ABC
2
AB
AC
F
0
BANANA
4
AB
AN
BA
NA
FBHC
4
FB
BF
HC
CH
FOXEN
8
NI
OE
NX
EW
OI
FE
FN
XW
CONSISTENCY
26
AB
BC
CD
DE
EF
FG
GH
HI
IJ
JK
KL
LM
MN
NO
OP
PQ
QR
RS
ST
TU
UV
VW
WX
XY
YZ
ZA
		`
	}()
	expected := func() string {
		return `
Case #1: 2
Case #2: -1
Case #3: 0
Case #4: 3
Case #5: -1
Case #6: 8
Case #7: 100
`
	}()

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	ProblemA2(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
