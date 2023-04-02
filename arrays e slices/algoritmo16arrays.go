package main

import "fmt"

func main() {
	var i int
	conj := [10]int{0, 25, 12, 14, 3, 37, 88, 23, 9, 11}
	conj_pares := []int{}

	fmt.Println("Seu conjunto é: ", conj)
	for i < len(conj) {
		if conj[i]%2 == 0 {
			conj_pares = append(conj_pares, conj[i])
		}
		i++
	}
	fmt.Println("Os pares são: ", conj_pares)
}
