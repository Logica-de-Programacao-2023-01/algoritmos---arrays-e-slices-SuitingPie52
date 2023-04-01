package main

import "fmt"

func main() {
	var x, y int
	conj := [2][3]int{{2, 3, 5}, {45, 0, 20}}

	fmt.Println("Seu conjunto é: ", conj)
	fmt.Println("Indique o valor do índice da sua linha. (0-1)")
	fmt.Scanln(&x)
	fmt.Println("Indique o valor da índice da sua coluna. (0-2)")
	fmt.Scanln(&y)
	fmt.Println("Seu valor é: ", conj[x][y])
}
