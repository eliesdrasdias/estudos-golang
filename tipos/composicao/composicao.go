package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Pode ter mais métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	// lógica para ligar o turbo
	fmt.Println("Turbo ligado")
}

func (b bmw7) fazerBaliza() {
	// lógica para fazer baliza
	fmt.Println("Baliza feita")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
