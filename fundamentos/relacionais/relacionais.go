package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("string: ", "banana" == "banana")
	fmt.Println("!=: ", 3 != 2)
	fmt.Println("<: ", 3 < 2)
	fmt.Println(">: ", 3 > 2)
	fmt.Println("<=: ", 3 <= 2)
	fmt.Println(">=: ", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("datas:", d1 == d2)
	fmt.Println("datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Joao"}

	fmt.Println("Pessoas", p1 == p2)
}
