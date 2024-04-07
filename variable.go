package lpwrapper

type Variable struct {
	Name  string
	index int
	value float64

	// reference to the LP model that this variable belongs to, this helps
	// with setting bounds.
	lp *LP
}

// Times multiplies the variable by a constant, creating a term.
func (v *Variable) Times(c float64) *Term {
	return &Term{Coefficient: c, Variable: v}
}

// PlusVar adds a variable to the variable, creating an expression of type 1*v1 + 1*v2.
func (v *Variable) PlusVar(v2 *Variable) *Expression {
	return v.PlusTerm(v2.Times(1))
}

// PlusTerm adds a term to the variable, creating an expression.
func (v *Variable) PlusTerm(t *Term) *Expression {
	return v.Times(1).Plus(t)
}

// PlusConst adds a constant to the variable, creating an expression.
func (v *Variable) PlusConst(c float64) *Expression {
	return v.Times(1).PlusConst(c)
}

// MinusVar subtracts a variable from the variable, creating an expression.
func (v *Variable) MinusVar(v2 *Variable) *Expression {
	return v.PlusTerm(v2.Times(-1))
}

// MinusTerm subtracts a term from the variable, creating an expression.
func (v *Variable) MinusTerm(t *Term) *Expression {
	newTerm := &Term{
		Coefficient: -t.Coefficient,
		Variable:    t.Variable,
	}
	return v.PlusTerm(newTerm)
}

// ToTerm converts the variable to a term with a coefficient of 1.
func (v *Variable) ToTerm() *Term {
	return &Term{Coefficient: 1, Variable: v}
}

// ToExpression converts the variable to an expression with only one term.
func (v *Variable) ToExpression() *Expression {
	return &Expression{Terms: []*Term{v.Times(1)}}
}

// Value returns the value of the variable after the solution.
func (v *Variable) Value() float64 {
	return v.value
}

// String returns the name of the variable.
func (v *Variable) String() string {
	return v.Name
}

// SetMaxBound sets the maximum bound for the variable. By default it is +infinity.
func (v *Variable) SetMaxBound(bound float64) {
	v.lp.AddConstraint(v.ToExpression().LE(NewConstExpression(bound)))
}

// SetMinBound sets the minimum bound for the variable. By default it is 0.
func (v *Variable) SetMinBound(bound float64) {
	v.lp.AddConstraint(v.ToExpression().GE(NewConstExpression(bound)))
}

// SetUnbounded sets the variable to be unbounded (-inf, +inf).
func (v *Variable) SetUnbounded() {
	v.lp.lpContext.SetUnbounded(v.index)
}
