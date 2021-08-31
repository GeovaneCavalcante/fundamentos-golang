package main

import "fmt"

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {

	r1, r2 := troca(3, 2)
	fmt.Println(r1, r2)
}
