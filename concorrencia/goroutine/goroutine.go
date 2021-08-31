package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s - %d\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Geovane", "porque vc nao fala comigo", 3)
	//fale("joao", "só posso falar depois de vc", 1)
	//go fale("Maria", "ei", 500)
	//go fale("joao", "opa", 500)

	go fale("Maria", "entendi", 10)
	fale("joao", "parabnéns", 5)

}
