package main

import "fmt"

func main() {
	var x int
	conj := [10]float64{57.5, 1, 25, 4, 67, 3.27, 7, 0, 64, 20}
	conjmaiores := []float64{}
	for len(conj) > x {
		if conj[x] > 5 {
			conjmaiores = append(conjmaiores, conj[x])
		}
		x++
	}
	fmt.Println("Seu conjunto é:", conj)
	fmt.Println("Os elementos maiores que 5 são: ", conjmaiores)
}
