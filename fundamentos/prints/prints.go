package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("linha")

	fmt.Println("nova")
	fmt.Println("linha")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("o valor de x é " + xs)
	fmt.Println("Valor de x é", x)

	fmt.Printf("Valor de x é  %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
}