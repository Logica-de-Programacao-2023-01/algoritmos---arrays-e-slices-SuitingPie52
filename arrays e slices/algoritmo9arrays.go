package main

import "fmt"

func main() {
	conj := [6]float64{}
	var x float64
	var z int

	fmt.Println("Informe um número.")
	fmt.Scanln(&x)

	for z < 6 {
		conj[z] = x
		z++
	}

	fmt.Println("Seu conjunto é: ", conj)
}
