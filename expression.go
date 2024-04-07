package lpwrapper

import (
	"fmt"

	"github.com/draffensperger/golp"
)

// Expression represents a linear expression, as a sum of terms.
type Expression struct {
	Terms     []*Term
	ConstTerm float64
}

// NewConstExpression creates a new expression with a constant term.
func NewConstExpression(c float64) *Expression {
	return &Expression{ConstTerm: c}
}

// PlusTerm adds a term to the expression, modifying the same expression and returning itself.
func (e *Expression) PlusTerm(t *Term) *Expression {
	e.Terms = append(e.Terms, t)
	return e
}

// MinusTermadds a term with inverted sign to the expression, modifying the same expression and returning itself.
func (e *Expression) MinusTerm(t *Term) *Expression {
	e.Terms = append(e.Terms, &Term{Coefficient: -t.Coefficient, Variable: t.Variable})
	return e
}

// PlusTerms adds multiple terms to the expression, modifying the same expression and returning itself.
func (e *Expression) PlusTerms(terms []*Term) *Expression {
	e.Terms = append(e.Terms, terms...)
	return e
}

// MinusTerms adds multiple terms with inverted sign to the expression, modifying the same expression and returning itself.
func (e *Expression) MinusTerms(terms []*Term) *Expression {
	for _, t := range terms {
		e.Terms = append(e.Terms, &Term{Coefficient: -t.Coefficient, Variable: t.Variable})
	}
	return e
}

// PlusConst adds a constant to the expression, modifying the same expression and returning itself.
func (e *Expression) PlusConst(c float64) *Expression {
	e.ConstTerm += c
	return e
}

// MinusConst subtracts a constant from the expression, modifying the same expression and returning itself.
func (e *Expression) MinusConst(c float64) *Expression {
	e.ConstTerm -= c
	return e
}

// EQ creates a new constraint from two expressions, with the first equal to the second.
func (e *Expression) EQ(e2 *Expression) *Constraint {
	return e.createConstraint(golp.EQ, e2)
}

// LE creates a new constraint from two expressions, with the first less than or equal to the second.
func (e *Expression) LE(e2 *Expression) *Constraint {
	return e.createConstraint(golp.LE, e2)
}

// GE creates a new constraint from two expressions, with the first greater than or equal to the second.
func (e *Expression) GE(e2 *Expression) *Constraint {
	return e.createConstraint(golp.GE, e2)
}

func (e *Expression) createConstraint(constraint golp.ConstraintType, e2 *Expression) *Constraint {
	rhs := e2.ConstTerm - e.ConstTerm
	newExpression := &Expression{
		Terms:     make([]*Term, 0, len(e.Terms)+len(e2.Terms)),
		ConstTerm: 0,
	}
	newExpression.PlusTerms(e.Terms)
	newExpression.MinusTerms(e2.Terms)

	return &Constraint{Expression: newExpression, Constraint: constraint, RightHandSide: rhs}
}

// String returns a string representation of the expression.
func (e *Expression) String() string {
	s := ""
	for i, t := range e.Terms {
		if i > 0 {
			s += " + "
		}
		s += t.String()
	}
	if e.ConstTerm != 0 {
		if s != "" {
			s += " + "
		}
		s += fmt.Sprintf("%.2f", e.ConstTerm)
	}
	return s
}
