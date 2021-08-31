package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)

	c <- 1

	fmt.Println("só depois que for lido")
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("foi lido")
	fmt.Println(<-c)
	fmt.Println("fim")
}
