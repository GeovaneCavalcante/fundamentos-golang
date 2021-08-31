package main

import (
	"fmt"
	"time"
)

func main() {
	// while
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("par")
		} else {
			fmt.Println("impar")
		}
	}

	for {
		fmt.Println("pra sempre")
		time.Sleep(time.Second)
	}
}
