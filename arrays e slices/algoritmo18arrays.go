package main

import "fmt"

func main() {
	var quantidade, x, i int
	conjPrimos := []int{}
	x = 2
	i = 2

	fmt.Println("ALGORITMO INCOMPLETO")
	fmt.Println("Informe uma quantidade de números primos.")
	fmt.Scanln(&quantidade)
	for quantidade != len(conjPrimos) {
		for i%x != 0 && x < i {
			if i > 1 && i%i == 0 && i/1 == 0 {
				conjPrimos = append(conjPrimos, i)
			}
			i++
		}
		x++
	}
	fmt.Println("Os", quantidade, "primeiros números primos são", conjPrimos)
}
