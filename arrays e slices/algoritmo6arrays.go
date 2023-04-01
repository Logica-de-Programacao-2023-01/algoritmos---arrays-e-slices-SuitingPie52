package main

import "fmt"

func main() {
	var x, z int
	conjunto := [10]int{0, 1, 4, 7, 15, 20, 25, 74, 80, 100}

	fmt.Println("Informe um número.")
	fmt.Scanln(&x)
	for {
		if x == conjunto[z] {
			fmt.Println("Seu número está contido nesse conjunto.")
			break
		} else if z >= 9 {
			fmt.Println("Seu número não está contido nesse conjunto.")
			break
		}
		z++
	}
}
