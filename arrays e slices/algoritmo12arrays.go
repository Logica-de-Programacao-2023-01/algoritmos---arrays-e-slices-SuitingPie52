package main

import "fmt"

func main() {
	var x int
	conj := [5]int{2, 3, 7, 9, 15}
	multiplos := []int{}
	for len(conj) > x {
		multiplos = append(multiplos, conj[x]*3)
		x++
	}
	fmt.Println("Seu conjunto é: ", conj)
	fmt.Println("Seus múltiplos por 3 são: ", multiplos)
}
