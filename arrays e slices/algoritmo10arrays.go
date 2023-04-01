package main

import "fmt"

func main() {
	var x int
	conj := []int{10, 37, 39, 87, 9, 94, 7, 8, 4, 11}
	menor := conj[1]
	maior := conj[1]

	fmt.Println("Seu conjunto é: ", conj)
	for x < 10 {
		if conj[x] > maior {
			maior = conj[x]
		}
		if conj[x] < menor {
			menor = conj[x]
		}
		x++
	}
	fmt.Println("Seu maior número é: ", maior)
	fmt.Println("Seu menor número é: ", menor)
}
