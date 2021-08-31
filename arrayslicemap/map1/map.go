package main

import "fmt"

func main() {
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[123456789] = "maria"
	aprovados[313121313] = "pedro"
	aprovados[312313231] = "ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s, %d\n", nome, cpf)
	}

	fmt.Println(aprovados[312313231])

	delete(aprovados, 312313231)

	fmt.Println(aprovados)
}
