package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Printf("bom dia")
	case t.Hour() < 18:
		fmt.Printf("boa tarde")
	default:
		fmt.Printf("boa noite")
	}
}
