package main

import "fmt"

func main() {
	var x, y, z, soma int
	numeros := []int{}

	fmt.Println("Informe o tamanho do seu Slice.")
	fmt.Scanln(&x)

	for {
		if len(numeros) == x {
			break
		}
		fmt.Println("Adicione um valor ao seu slice.")
		fmt.Scanln(&y)
		numeros = append(numeros, y)
	}
	fmt.Println("Seu slice é:", numeros)
	fmt.Print("Sua soma é: ")
	for {
		if z == x {
			break
		}
		soma += numeros[z]
		z++
	}
	fmt.Print(soma)
}
