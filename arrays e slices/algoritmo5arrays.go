package main

import "fmt"

func main() {
	var a, b, c, d, e, f int

	fmt.Println("Informe o número da linha 0 coluna 0.")
	fmt.Scanln(&a)
	fmt.Println("Informe o número da linha 0 e coluna 1.")
	fmt.Scanln(&b)
	fmt.Println("Informe o número da linha 1 e coluna 0.")
	fmt.Scanln(&c)
	fmt.Println("Informe o número da linha 1 e coluna 1.")
	fmt.Scanln(&d)
	fmt.Println("Informe o número da linha 2 e coluna 0.")
	fmt.Scanln(&e)
	fmt.Println("Informe o número da linha 2 e coluna 1.")
	fmt.Scanln(&f)

	numeros := [3][2]int{{a, b}, {c, d}, {e, f}}

	fmt.Println("Sua matriz é:", numeros)
}
