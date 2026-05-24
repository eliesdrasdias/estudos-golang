package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) total() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += float64(item.quantidade) * item.preco
	}
	return total
}

func main() {
	pedido1 := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 101, quantidade: 2, preco: 10.0},
			{produtoID: 102, quantidade: 1, preco: 20.0},
		},
	}
	fmt.Printf("Total do pedido: %.2f\n", pedido1.total())
}
