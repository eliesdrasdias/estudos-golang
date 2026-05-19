package main

import "fmt"

func inc1(n int) int {
	n++
	return n
}

func inc2(n *int) int {
	*n++
	return *n
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n)
	fmt.Println(n)
}
