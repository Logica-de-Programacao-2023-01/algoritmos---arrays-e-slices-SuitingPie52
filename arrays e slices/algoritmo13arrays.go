package main

import "fmt"

func main() {
	var x, y int
	conj := [7]int{}

	fmt.Println("Seu conjunto é: ", conj)
	fmt.Println("Informe um valor que será adicionado ao seu primeiro elemento.")
	fmt.Scanln(&x)
	fmt.Println("Informe um valor que será adicionado ao seu último elemento.")
	fmt.Scanln(&y)

	conj[0] = x
	z := len(conj) - 1
	conj[z] = y

	fmt.Println("Seu conjunto agora é: ", conj)
}
