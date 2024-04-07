package lpwrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstraint(t *testing.T) {
	lp := NewLP()

	x := lp.AddVariable("x")
	y := lp.AddVariable("y")
	z := lp.AddVariable("z")

	exp1 := x.PlusVar(y).MinusConst(4)
	exp2 := z.PlusConst(3)

	c := exp1.EQ(exp2)

	assert.Equal(t, "1.00*x + 1.00*y + -1.00*z EQ 7.00", c.String())

}
