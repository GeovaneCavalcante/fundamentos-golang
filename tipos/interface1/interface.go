package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrebome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrebome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"roberto", "silva"}
	imprimir(coisa)

	coisa = produto{"roberto", 7.9}
	imprimir(coisa)
}
