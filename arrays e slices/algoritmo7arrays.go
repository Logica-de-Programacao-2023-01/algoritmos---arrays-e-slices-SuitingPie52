package main

import "fmt"

func main() {
	var x, z int
	conjunto := []int{2, 7, 9, 15, 25}

	fmt.Println("Informe seu número.")
	fmt.Scanln(&x)

	for {
		if x == conjunto[z] {
			fmt.Println("Seu conjunto é: ", conjunto)
			break
		} else if z <= 3 {
			z++
		} else {
			conjunto = append(conjunto, x)
			fmt.Println("Seu conjunto é: ", conjunto)
			break
		}
	}
}
