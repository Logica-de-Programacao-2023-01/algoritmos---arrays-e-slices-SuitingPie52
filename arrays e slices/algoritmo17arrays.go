package main

import "fmt"

func main() {
	var i, soma int
	conj := [10]int{3, 11, 20, 17, 7, 2, 13, 40, 75, 35}

	fmt.Println("Seu conjunto é :", conj)
	for i < len(conj) {
		if i%2 != 0 {
			soma += conj[i]
		}
		i++
	}
	fmt.Println("Sua soma é: ", soma)
}
