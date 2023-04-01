package main

import "fmt"

func main() {
	var x, y, armazenar1 int
	conj := []int{34, 2, 3, 55, 6, 7, 9, 12}

	fmt.Println("Seu conjunto é: ", conj)
	fmt.Println("Informe 2 índices de números para trocar de posição.")
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	armazenar1 = conj[x]
	conj[x] = conj[y]
	conj[y] = armazenar1

	fmt.Println("Seu conjunto é: ", conj)
}
