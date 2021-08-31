package main

import "fmt"

func obeterValorAprovado(numero int) int {
	defer fmt.Println("fim")

	if numero > 500 {
		fmt.Println("valor alto")

		return 500
	} else {
		fmt.Println("valor baixo")

		return numero
	}
}

func main() {
	fmt.Println(obeterValorAprovado(500))
}
