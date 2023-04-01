package main

import "fmt"

func main() {
	var x int
	nomes := []string{"João", "Maria", "Pedro", "Gabriel", "Cléber", "Isabelle", "Iracema", "Fábio"}

	fmt.Println("Seus nomes são:", nomes)
	fmt.Println("Escolha um nome da lista para remover. (Informe o índice de 0 a 7)")
	fmt.Scanln(&x)
	y := x + 1
	if x < len(nomes) {
		nomes = append(nomes[:x], nomes[y:]...)
	}
	fmt.Println("Seus nomes são: ", nomes)
}
