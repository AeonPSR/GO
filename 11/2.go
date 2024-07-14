package main

import (
	"fmt"
)

func somme(tableau []int) int {
	total := 0
	for _, valeur := range tableau {
		total += valeur
	}
	return total
}

func main() {
	nombres := []int{1, 1, 1, 1, 5}
	resultat := somme(nombres)
	fmt.Printf("La somme des éléments du tableau est: %d\n", resultat)
}