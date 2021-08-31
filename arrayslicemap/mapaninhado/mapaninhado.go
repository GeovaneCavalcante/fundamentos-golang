package main

import "fmt"

func main() {
	funcsPorLetras := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 123122.32,
			"guga pereira":   2311.322,
		},
		"J": {
			"Gabriela Silva": 123122.32,
			"guga pereira":   2311.322,
		},
		"P": {
			"Gabriela Silva": 123122.32,
			"guga pereira":   2311.322,
		},
	}

	fmt.Println(funcsPorLetras)
}
