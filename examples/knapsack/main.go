package main

import (
	"fmt"

	"github.com/andremissaglia/lpwrapper"
)

type product struct {
	name   string
	weight float64
	cost   float64
	value  float64
}

func main() {
	products := []product{
		{"apple", 5, 2, 5},
		{"orange", 5, 3, 7},
		{"spice", 1, 9, 20},
		{"pineapple", 10, 4, 8},
	}

	maxWeight := 6.0
	maxCost := 13.0

	// Create a new LP model
	lp := lpwrapper.NewLP()

	// Create a new variable for each product
	amountSelected := make([]*lpwrapper.Variable, len(products))
	for i, p := range products {
		amountSelected[i] = lp.AddVariable(p.name)
	}

	// Add constraint for the maximum weight
	weights := &lpwrapper.Expression{}
	for i, p := range products {
		weights = weights.PlusTerm(amountSelected[i].Times(p.weight))
	}
	lp.AddConstraint(weights.LE(lpwrapper.NewConstExpression(maxWeight)))

	// Add constraint for the maximum cost
	costs := &lpwrapper.Expression{}
	for i, p := range products {
		costs = costs.PlusTerm(amountSelected[i].Times(p.cost))
	}
	lp.AddConstraint(costs.LE(lpwrapper.NewConstExpression(maxCost)))

	// Add objective function
	obj := lpwrapper.Expression{}
	for i, p := range products {
		obj.PlusTerm(amountSelected[i].Times(p.value))
	}
	lp.SetObjFn(obj.Terms, true)

	// Solve the model
	result := lp.Solve()

	// Print the result
	fmt.Printf("Maximum value: %f\n", result)
	for i, p := range products {
		fmt.Printf("%s: %.1f\n", p.name, amountSelected[i].Value())
	}
}
