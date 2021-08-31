package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "a"
	case 8, 7:
		return "b"
	case 5, 6:
		return "c"
	case 4, 3:
		return "d"
	case 2, 1, 0:
		return "e"

	default:
		return "nota invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(100))
}
