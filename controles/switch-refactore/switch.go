package main

import "fmt"

func notaParaConceito(n float64) string {

	switch {
	case n >= 9 && n <= 10:
		return "a"
	case n >= 8 && n <= 9:
		return "b"
	case n >= 5 && n <= 8:
		return "c"
	case n >= 3 && n <= 5:
		return "d"
	default:
		return "E"

	}
}

func main() {
	fmt.Println(notaParaConceito(9.9))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(3.9))
}
