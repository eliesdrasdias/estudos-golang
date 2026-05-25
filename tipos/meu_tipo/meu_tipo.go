package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9, 10) {
		return "A"
	} else if n.entre(8, 8.99) {
		return "B"
	} else if n.entre(7, 7.99) {
		return "C"
	} else if n.entre(6, 6.99) {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5)) // A
	fmt.Println(notaParaConceito(8.5)) // B
	fmt.Println(notaParaConceito(7.5)) // C
	fmt.Println(notaParaConceito(6.5)) // D
	fmt.Println(notaParaConceito(5.5)) // F
}
