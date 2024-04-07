package lpwrapper

import (
	"fmt"

	"github.com/draffensperger/golp"
)

type Constraint struct {
	Expression    *Expression
	Constraint    golp.ConstraintType
	RightHandSide float64
}

func (e *Constraint) String() string {
	return fmt.Sprintf("%s %s %.2f", e.Expression.String(), e.Constraint, e.RightHandSide)
}
