package lpwrapper

import "fmt"

// Term represents a term in a linear expression, as constant * variable.
type Term struct {
	Coefficient float64
	Variable    *Variable
}

// Plus adds two terms together, creating a new expression.
func (t *Term) Plus(t2 *Term) *Expression {
	return &Expression{Terms: []*Term{t, t2}}
}

// Negative returns the negative of the term.
func (t *Term) Negative() *Term {
	return &Term{Coefficient: -t.Coefficient, Variable: t.Variable}
}

// ToExpression converts the term to an expression with only one term.
func (t *Term) ToExpression() *Expression {
	return &Expression{Terms: []*Term{t}}
}

// PlusConst adds a constant to the term, creating a new expression.
func (t *Term) PlusConst(c float64) *Expression {
	return &Expression{Terms: []*Term{t}, ConstTerm: c}
}

// String returns a string representation of the term.
func (t *Term) String() string {
	return fmt.Sprintf("%.2f*%s", t.Coefficient, t.Variable.String())
}
