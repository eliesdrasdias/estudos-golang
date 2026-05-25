package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turbo bool
}

func main() {
	f1 := ferrari{}
	f1.nome = "F40"
	f1.velocidadeAtual = 0
	f1.turbo = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f1.nome, f1.turbo)
}
