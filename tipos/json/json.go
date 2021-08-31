package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "Notebook", 1999.0, []string{"promocao", "eletronico"}}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 produto

	jsonString := `{"id":1,"nome":"Notebook","preco":1999,"tags":["promocao","eletronico"]}`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags)
}
