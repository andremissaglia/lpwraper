package lpwrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionBuilding(t *testing.T) {
	lp := NewLP()

	x := lp.AddVariable("x")
	y := lp.AddVariable("y")
	z := lp.AddVariable("z")

	exp := x.PlusVar(y).PlusTerm(z.Times(2)).MinusConst(4)
	assert.Equal(t, &Expression{
		Terms: []*Term{
			{Coefficient: 1, Variable: x},
			{Coefficient: 1, Variable: y},
			{Coefficient: 2, Variable: z},
		},
		ConstTerm: -4,
	}, exp)
	assert.Equal(t, "1.00*x + 1.00*y + 2.00*z + -4.00", exp.String())

	exp = x.ToTerm().Negative().PlusConst(3).MinusTerms(y.MinusVar(z).Terms)
	assert.Equal(t, &Expression{
		Terms: []*Term{
			{Coefficient: -1, Variable: x},
			{Coefficient: -1, Variable: y},
			{Coefficient: +1, Variable: z},
		},
		ConstTerm: 3,
	}, exp)

	exp = x.PlusConst(1)
	assert.Equal(t, &Expression{
		Terms: []*Term{
			{Coefficient: 1, Variable: x},
		},
		ConstTerm: 1,
	}, exp)

	exp = (y.MinusTerm(z.Times(2)))
	assert.Equal(t, &Expression{
		Terms: []*Term{
			{Coefficient: 1, Variable: y},
			{Coefficient: -2, Variable: z},
		},
		ConstTerm: 0,
	}, exp)

	exp = x.ToExpression()
	assert.Equal(t, &Expression{
		Terms: []*Term{
			{Coefficient: 1, Variable: x},
		},
		ConstTerm: 0,
	}, exp)

	exp = x.ToTerm().ToExpression()
	assert.Equal(t, &Expression{
		Terms: []*Term{
			{Coefficient: 1, Variable: x},
		},
		ConstTerm: 0,
	}, exp)
}
