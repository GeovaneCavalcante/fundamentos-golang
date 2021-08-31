package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

func (c carro) teste() string {
	return c.nome
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f.teste())
}
