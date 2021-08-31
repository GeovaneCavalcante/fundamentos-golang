package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("soma =>", a+b)
	fmt.Println("subtricao =>", a-b)
	fmt.Println("divisao =>", a/b)
	fmt.Println("multiplicacao =>", a*b)
	fmt.Println("módulo =>", a%b)

	fmt.Println("E =>", a&b)
	fmt.Println("ou =>", a|b)

	c := 3.0
	d := 2.0

	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("menor => ", math.Max(c, d))
	fmt.Println("exponecial => ", math.Pow(c, d))
}
