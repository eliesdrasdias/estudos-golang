package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turbo           bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
	f.velocidadeAtual += 50
}

func main() {
	carro1 := ferrari{"F8 Tributo", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"SF90 Stradale", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
