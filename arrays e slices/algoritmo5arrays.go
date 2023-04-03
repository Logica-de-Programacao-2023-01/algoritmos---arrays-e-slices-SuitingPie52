package main

import "fmt"

func main() {
	var x int
	conj := [3][2]int{}

	for linha := 0; len(conj) > linha; linha++ {
		for coluna := 0; len(conj[linha]) > coluna; coluna++ {
			fmt.Println("Informe o valor da linha", linha, "e da coluna", coluna, ".")
			fmt.Scanln(&x)
			conj[linha][coluna] = x
		}
	}

	fmt.Println("Sua matriz é:", conj)
}
