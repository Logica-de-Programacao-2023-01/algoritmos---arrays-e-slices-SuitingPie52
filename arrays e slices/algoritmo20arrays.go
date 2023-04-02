package main

import "fmt"

func main() {
	var verdadeiro, i int
	conj := [5]int{0, 10, 2, 3, 4}

	fmt.Println("Seu conjunto é:", conj)
	for len(conj) > i+1 {
		if conj[i+1] > conj[i] {
			verdadeiro = 1
		} else {
			verdadeiro = 0
			break
		}
		i++
	}
	if verdadeiro == 1 {
		fmt.Println("Seu conjento é crescente.")
	} else {
		fmt.Println("Seu conjunto não é crescente.")
	}
}
