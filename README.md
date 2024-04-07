# lpwrapper

A simple linear programing library, built on top of golp and lpsolve

It allows easily creating models using math expressions, instead of operating the final matrix directly.

## Example
 Here is a simple example with two variables:

```go
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
```

This produces:
```
Maximum value: 3.000000
(x, y): (1.0, 2.0)

Full model:
Objective function:  1.00*x + 1.00*y

Maximize:  true

Variables: 
x
y

Constraints: 
1.00*x + 3.00*y LE 7.00
2.00*x + 1.00*y LE 4.00
```

You can find more elaborate examples in the [examples](./examples/) folder.

## Getting started

* This library requires `lpsolve` to be previously installed on your machine.
* Add this to you environment variables:

```sh
export CGO_CFLAGS="-I/usr/include/lpsolve"
export CGO_LDFLAGS="-llpsolve55 -lm -ldl"
```