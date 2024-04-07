package main

import (
	"fmt"

	"github.com/andremissaglia/lpwrapper"
)

func main() {
	lp := lpwrapper.NewLP()

	x := lp.AddVariable("x")
	y := lp.AddVariable("y")

	c1 := x.PlusTerm(y.Times(3)).LE(lpwrapper.NewConstExpression(7))
	lp.AddConstraint(c1)

	c2 := x.Times(2).Plus(y.ToTerm()).LE(lpwrapper.NewConstExpression(4))
	lp.AddConstraint(c2)

	obj := x.PlusVar(y)
	lp.SetObjFn(obj.Terms, true)

	result := lp.Solve()

	fmt.Printf("Maximum value: %f\n", result)
	fmt.Printf("(x, y): (%.1f, %.1f)\n", x.Value(), y.Value())
	fmt.Println("")
	fmt.Println("Full model:")
	lp.PrintModel()

}
