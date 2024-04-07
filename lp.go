package lpwrapper

import (
	"fmt"

	"github.com/draffensperger/golp"
)

// LP is a linear programming model, that wraps the golp library, with some
// additional functionality to help with creating models.
type LP struct {
	variables   []*Variable
	constraints []*Constraint
	objectiveFn *Expression
	maximize    bool
	lpContext   *golp.LP
}

// NewLP creates a new LP model.
func NewLP() *LP {
	return &LP{}
}

// SetVervoseLevel sets the verbose level of the model. Defaults to golp.IMPORTANT
func (lp *LP) SetVervoseLevel(level golp.VerboseLevel) {
	lp.lpContext.SetVerboseLevel(level)
}

// AddVariable adds a new variable to the model, with a name to help identify it.
func (lp *LP) AddVariable(name string) *Variable {
	v := Variable{
		Name:  name,
		index: len(lp.variables),
		lp:    lp,
	}
	lp.variables = append(lp.variables, &v)
	return &v
}

// AddConstraintLegacy adds a new constraint to the model.
func (lp *LP) AddConstraint(constraint *Constraint) {
	if len(constraint.Expression.Terms) == 0 {
		if constraint.Expression.ConstTerm != 0 {
			panic(fmt.Sprintf("%s = %.2f", constraint.Constraint, constraint.RightHandSide))
		}
		return
	}
	lp.constraints = append(lp.constraints, constraint)
}

// SetObjFn sets the objective function of the model.
func (lp *LP) SetObjFn(terms []*Term, maximize bool) {
	lp.objectiveFn = &Expression{Terms: terms}
	lp.maximize = maximize
}

func termsToEntries(terms []*Term) []golp.Entry {
	entries := make([]golp.Entry, len(terms))
	for i, t := range terms {
		entries[i] = golp.Entry{
			Col: t.Variable.index,
			Val: t.Coefficient,
		}
	}
	return entries
}

func (lp *LP) termsToFullRow(terms []*Term) []float64 {
	row := make([]float64, len(lp.variables))
	for _, t := range terms {
		row[t.Variable.index] = t.Coefficient
	}
	return row
}

// Solve the model and return the objective function value.
func (lp *LP) Solve() float64 {
	lp.lpContext = golp.NewLP(len(lp.constraints), len(lp.variables))

	for _, c := range lp.constraints {
		err := lp.lpContext.AddConstraintSparse(termsToEntries(c.Expression.Terms), c.Constraint, c.RightHandSide)
		if err != nil {
			// if we are using the library correctly, we should never hit this
			panic(err)
		}
	}

	lp.lpContext.SetObjFn(lp.termsToFullRow(lp.objectiveFn.Terms))

	if lp.maximize {
		lp.lpContext.SetMaximize()
	}

	lp.lpContext.Solve()
	variables := lp.lpContext.Variables()
	for i, v := range lp.variables {
		v.value = variables[i]
	}
	return lp.lpContext.Objective()
}

// PrintModel prints the model to the console.
func (lp *LP) PrintModel() {
	fmt.Println("Objective function: ", lp.objectiveFn.String())
	fmt.Println()

	fmt.Println("Maximize: ", lp.maximize)
	fmt.Println()

	fmt.Println("Variables: ")
	for _, v := range lp.variables {
		fmt.Println(v.String())
	}
	fmt.Println()

	fmt.Println("Constraints: ")
	for _, c := range lp.constraints {
		fmt.Println(c.String())
	}
}
