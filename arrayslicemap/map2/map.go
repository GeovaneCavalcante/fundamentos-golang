package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"Jose":    1221.2,
		"Gabriel": 232.23,
		"Pedri":   212.2,
	}

	funcsESalarios["rafael"] = 212.2

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
