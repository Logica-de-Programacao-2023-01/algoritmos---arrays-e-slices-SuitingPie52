package main

import "fmt"

func main() {
	var x string
	var y int
	nomes := []string{"João", "Maria", "Fábio", "Gabriel", "Cléber", "Isabelle", "Iracema", "Fábio"}

	fmt.Println("Seus nomes são:", nomes)
	fmt.Println("Digite um nome da lista para remover.")
	fmt.Scanln(&x)

	z := y + 1
	for len(nomes) > y {
		if x == nomes[y] {
			nomes = append(nomes[:y], nomes[z:]...)
		}
		y++
		z++
	}
	fmt.Println("Seus nomes são: ", nomes)
}
