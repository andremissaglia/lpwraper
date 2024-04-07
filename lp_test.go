package lpwrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lp := NewLP()

	x := lp.AddVariable("x")
	y := lp.AddVariable("y")

	c1 := x.PlusTerm(y.Times(3)).LE(NewConstExpression(7))
	lp.AddConstraint(c1)

	c2 := x.Times(2).Plus(y.ToTerm()).LE(NewConstExpression(4))
	lp.AddConstraint(c2)

	obj := x.PlusVar(y)
	lp.SetObjFn(obj.Terms, true)

	result := lp.Solve()
	assert.InDelta(t, 3.0, result, 0.01)
	assert.InDelta(t, 1.0, x.value, 0.01)
	assert.InDelta(t, 2.0, y.value, 0.01)

}
