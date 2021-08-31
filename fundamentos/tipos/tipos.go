package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)

	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivo).... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	// com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64

	fmt.Println("o valor maximo de int é", i1)
	fmt.Println("o tipo é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numero reis float32, float64

	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal é", reflect.TypeOf(49.99))

	// boolean

	bo := true
	fmt.Println("o tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string

	s1 := "olá mundo"
	fmt.Println("o tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println("tamanho da string é", len(s1))
	s2 := `
		olá mundo
		teste
	`
	fmt.Println(s2)
	fmt.Println("tamanho da string é", len(s2))
}
