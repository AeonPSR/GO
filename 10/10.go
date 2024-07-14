package main

import (
	"errors"
	"fmt"
)

func diviser(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("le diviseur ne peut pas être zéro")
	}
	return a / b, nil
}

func main() {
	nombres := [][2]int{ //Slice of ints to iterate through
		{10, 2},
		{10, 0},
		{15, 3},
		{20, 4},
		{864197523, 7},
		{78, 456879}, //Goes to 0, due to INT rounding. 
	}

	for _, n := range nombres { //'_, n' -> Ignore the index, n takes the value. Itterate of the range nombres
		resultat, err := diviser(n[0], n[1]) //Get the result and if there's an error in two variable, since this weird language can do that
		if err != nil {
			fmt.Printf("Erreur: %v\n", err)
		} else {
			fmt.Printf("Résultat de %d / %d = %d\n", n[0], n[1], resultat)
		}
	}
}
