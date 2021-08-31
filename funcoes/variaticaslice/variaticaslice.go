package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("lista de aprovados")

	for i, aprovado := range aprovados {
		fmt.Println(i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Joao", "maria", "geovane"}

	imprimirAprovados(aprovados...)
}
