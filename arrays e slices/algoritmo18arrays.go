package main

import "fmt"

func main() {
	var quantidade, y, i int
	conjPrimos := []int{}
	i = 1
	y = 1

	fmt.Println("Informe uma quantidade de números primos.")
	fmt.Scanln(&quantidade)
	for quantidade >= i {
		for y < quantidade {
			if i%y != 0 {
				if i%1 == 0 && i%i == 0 {
					conjPrimos = append(conjPrimos, i)
				}
			}
			y++
		}
		i++
	}
	fmt.Println("Os", quantidade, "primeiros números primos são", conjPrimos)
}
