package main

import (
	"fmt"

	"github.com/GeovaneCavalcante/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com"),
		html.Titulo("https://www.youtube.com"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
}
