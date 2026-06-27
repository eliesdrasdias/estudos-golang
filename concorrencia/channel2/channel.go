package main

import (
	"fmt"
	"time"
)

// channel é a forma de comunicação entre goroutines

// função que recebe um canal como parâmetro
func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal (escrita)

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

// funcao principal
func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo dados do canal (enviando)
	fmt.Println(a, b)
	fmt.Println(<-c)
}
