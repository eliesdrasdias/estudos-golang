package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// quebra a linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é" + x) = forma errada
	xs := fmt.Sprint(x) // converte o valor de x pra string
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	// o número 2 faz com que imprima 2 casas decimáis
	// o %f significar que vai imprimir uma variável float
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
